// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gatewayserver_test

import (
	"context"
	"encoding/base64"
	"net"
	"regexp"
	"testing"
	"time"

	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/pkg/component"
	"go.thethings.network/lorawan-stack/pkg/config"
	"go.thethings.network/lorawan-stack/pkg/gatewayserver"
	"go.thethings.network/lorawan-stack/pkg/gatewayserver/udp"
	"go.thethings.network/lorawan-stack/pkg/log"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/types"
	"go.thethings.network/lorawan-stack/pkg/util/test"
	"go.thethings.network/lorawan-stack/pkg/util/test/assertions/should"
)

const testUDPAddress = "127.0.0.1:8332"

func testPushData(eui types.EUI64, ns *GsNsServer) func(t *testing.T) {
	return func(t *testing.T) {
		a := assertions.New(t)

		p := udp.Packet{
			GatewayEUI:      &eui,
			ProtocolVersion: udp.Version1,
			Token:           [2]byte{0x11, 0x00},
			Data: &udp.Data{
				RxPacket: []*udp.RxPacket{
					{
						Freq: 868.0,
						Chan: 2,
						Modu: "LORA",
						DatR: udp.DataRate{DataRate: types.DataRate{LoRa: "SF10BW125"}},
						CodR: "4/7",
						Data: "QCkuASaAAAAByFaF53Iu+vzmwQ==",
						Tmst: 1003503,
					},
				},
			},
			PacketType: udp.PushData,
		}
		p.Data.RxPacket[0].Size = uint16(base64.StdEncoding.DecodedLen(len(p.Data.RxPacket[0].Data)))

		packetBytes, err := p.MarshalBinary()
		if !a.So(err, should.BeNil) {
			t.FailNow()
		}

		conn, err := net.Dial("udp", testUDPAddress)
		if !a.So(err, should.BeNil) {
			t.FailNow()
		}

		_, err = conn.Write(packetBytes)
		a.So(err, should.BeNil)

		_, err = conn.Read(make([]byte, 2)) // Receive ACK
		a.So(err, should.BeNil)

		select {
		case msg := <-ns.messageReceived:
			if msg != "HandleUplink" {
				t.Fatal("Expected Gateway Server to call HandleUplink on the Network Server, instead received", msg)
			}
		case <-time.After(nsReceptionTimeout):
			t.Fatal("The Gateway Server never called the Network Server's HandleUplink to handle the PUSH_DATA.")
		}

		conn.Close()
	}
}

func testPullData(gatewayEUI types.EUI64, ns *GsNsServer, conn net.Conn) func(t *testing.T) {
	return func(t *testing.T) {
		a := assertions.New(t)

		p := udp.Packet{
			GatewayEUI:      &gatewayEUI,
			ProtocolVersion: udp.Version1,
			Token:           [2]byte{0x11, 0x00},
			PacketType:      udp.PullData,
		}

		packetBytes, err := p.MarshalBinary()
		if !a.So(err, should.BeNil) {
			t.FailNow()
		}

		_, err = conn.Write(packetBytes)
		a.So(err, should.BeNil)

		_, err = conn.Read(make([]byte, 2)) // Receive ACK
		a.So(err, should.BeNil)

		select {
		case msg := <-ns.messageReceived:
			if msg != "StartServingGateway" {
				t.Fatal("Expected Gateway Server to call StartServingGateway on the Network Server, instead received", msg)
			}
		case <-time.After(nsReceptionTimeout):
			t.Fatal("The Gateway Server never called the Network Server's StartServingGateway to handle the PULL_DATA.")
		}
	}
}

func testDownlink(registeredGatewayID string, gs *gatewayserver.GatewayServer, conn net.Conn) func(t *testing.T) {
	downlink := ttnpb.NewPopulatedDownlinkMessage(test.Randy, false)
	downlink.TxMetadata.GatewayIdentifiers = ttnpb.GatewayIdentifiers{GatewayID: registeredGatewayID}
	downlink.TxMetadata.Timestamp = 3003503000
	downlink.Settings.Frequency = 868300000
	downlink.Settings.SpreadingFactor = 7
	downlink.Settings.Bandwidth = 125000
	downlink.Settings.CodingRate = "4/5"

	base64payloadRegexp := regexp.MustCompile(regexp.QuoteMeta(base64.StdEncoding.EncodeToString(downlink.RawPayload)))

	return func(t *testing.T) {
		a := assertions.New(t)

		_, err := gs.ScheduleDownlink(test.Context(), downlink)
		if !a.So(err, should.BeNil) {
			t.FailNow()
		}

		buffer := make([]byte, 2048)

		_, err = conn.Read(buffer)
		if !a.So(err, should.BeNil) {
			t.FailNow()
		}

		found := base64payloadRegexp.Match(buffer)
		a.So(found, should.BeTrue)
	}
}

func TestUDP(t *testing.T) {
	a := assertions.New(t)

	logger := test.GetLogger(t)
	ctx := log.NewContext(test.Context(), logger)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	store, err := test.NewFrequencyPlansStore()
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}
	defer store.Destroy()

	registeredGatewayID := "registered-gateway"
	registeredGatewayEUI := types.EUI64{0xAA, 0xEE, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	registeredGatewayIdentifiers := ttnpb.GatewayIdentifiers{GatewayID: registeredGatewayID, EUI: &registeredGatewayEUI}
	_, isAddr := StartMockIsGatewayServer(ctx, []ttnpb.Gateway{
		{
			GatewayIdentifiers: registeredGatewayIdentifiers,
			FrequencyPlanID:    "EU_863_870",
			DisableTxDelay:     true,
		},
	})

	ns, nsAddr := StartMockGsNsServer(ctx)

	c := component.MustNew(logger, &component.Config{
		ServiceBase: config.ServiceBase{
			Cluster: config.Cluster{
				Name:           "test-gateway-server",
				IdentityServer: isAddr,
				NetworkServer:  nsAddr,
			},
			FrequencyPlans: config.FrequencyPlans{
				StoreDirectory: store.Directory(),
			},
		},
	})
	gs, err := gatewayserver.New(c, gatewayserver.Config{
		UDPAddress: testUDPAddress,
	})
	if !a.So(err, should.BeNil) {
		t.Fatal("Gateway Server could not be initialized:", err)
	}

	err = gs.Start()
	if !a.So(err, should.BeNil) {
		t.Fatal("Gateway Server could not start:", err)
	}

	gsStart := time.Now()
	for gs.GetPeer(ttnpb.PeerInfo_IDENTITY_SERVER, []string{}, nil) == nil || gs.GetPeer(ttnpb.PeerInfo_NETWORK_SERVER, []string{}, nil) == nil {
		if time.Since(gsStart) > nsReceptionTimeout {
			t.Fatal("Identity Server and Network Server were not initialized in time by the Gateway Server - timeout")
		}
		time.Sleep(2 * time.Millisecond)
	}

	conn, err := net.Dial("udp", testUDPAddress)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	t.Run("PullData", testPullData(registeredGatewayEUI, &ns, conn))
	t.Run("PushData", testPushData(registeredGatewayEUI, &ns))
	t.Run("Downlink", testDownlink(registeredGatewayID, gs, conn))
}
