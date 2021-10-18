// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.1.0
// - protoc             v3.9.1
// source: lorawan-stack/api/contact_info.proto

package ttnpb

import (
	gogo "github.com/TheThingsIndustries/protoc-gen-go-json/gogo"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the ContactType to JSON.
func (x ContactType) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	s.WriteEnumString(int32(x), ContactType_name)
}

// ContactType_customvalue contains custom string values that extend ContactType_value.
var ContactType_customvalue = map[string]int32{
	"OTHER":     0,
	"ABUSE":     1,
	"BILLING":   2,
	"TECHNICAL": 3,
}

// UnmarshalProtoJSON unmarshals the ContactType from JSON.
func (x *ContactType) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	v := s.ReadEnum(ContactType_value, ContactType_customvalue)
	if err := s.Err(); err != nil {
		s.SetErrorf("could not read ContactType enum: %v", err)
		return
	}
	*x = ContactType(v)
}

// MarshalProtoJSON marshals the ContactMethod to JSON.
func (x ContactMethod) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	s.WriteEnumString(int32(x), ContactMethod_name)
}

// ContactMethod_customvalue contains custom string values that extend ContactMethod_value.
var ContactMethod_customvalue = map[string]int32{
	"OTHER": 0,
	"EMAIL": 1,
	"PHONE": 2,
}

// UnmarshalProtoJSON unmarshals the ContactMethod from JSON.
func (x *ContactMethod) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	v := s.ReadEnum(ContactMethod_value, ContactMethod_customvalue)
	if err := s.Err(); err != nil {
		s.SetErrorf("could not read ContactMethod enum: %v", err)
		return
	}
	*x = ContactMethod(v)
}

// MarshalProtoJSON marshals the ContactInfo message to JSON.
func (x *ContactInfo) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.ContactType != 0 || s.HasField("contact_type") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("contact_type")
		x.ContactType.MarshalProtoJSON(s)
	}
	if x.ContactMethod != 0 || s.HasField("contact_method") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("contact_method")
		x.ContactMethod.MarshalProtoJSON(s)
	}
	if x.Value != "" || s.HasField("value") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("value")
		s.WriteString(x.Value)
	}
	if x.Public || s.HasField("public") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("public")
		s.WriteBool(x.Public)
	}
	if x.ValidatedAt != nil || s.HasField("validated_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("validated_at")
		if x.ValidatedAt == nil {
			s.WriteNil()
		} else {
			s.WriteTime(*x.ValidatedAt)
		}
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the ContactInfo message from JSON.
func (x *ContactInfo) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "contact_type", "contactType":
			s.AddField("contact_type")
			x.ContactType.UnmarshalProtoJSON(s)
		case "contact_method", "contactMethod":
			s.AddField("contact_method")
			x.ContactMethod.UnmarshalProtoJSON(s)
		case "value":
			s.AddField("value")
			x.Value = s.ReadString()
		case "public":
			s.AddField("public")
			x.Public = s.ReadBool()
		case "validated_at", "validatedAt":
			s.AddField("validated_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.ValidatedAt = v
		}
	})
}

// MarshalProtoJSON marshals the ContactInfoValidation message to JSON.
func (x *ContactInfoValidation) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Id != "" || s.HasField("id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("id")
		s.WriteString(x.Id)
	}
	if x.Token != "" || s.HasField("token") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("token")
		s.WriteString(x.Token)
	}
	if x.Entity != nil || s.HasField("entity") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("entity")
		// NOTE: EntityIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.Entity)
	}
	if len(x.ContactInfo) > 0 || s.HasField("contact_info") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("contact_info")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.ContactInfo {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("contact_info"))
		}
		s.WriteArrayEnd()
	}
	if x.CreatedAt != nil || s.HasField("created_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("created_at")
		if x.CreatedAt == nil {
			s.WriteNil()
		} else {
			s.WriteTime(*x.CreatedAt)
		}
	}
	if x.ExpiresAt != nil || s.HasField("expires_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("expires_at")
		if x.ExpiresAt == nil {
			s.WriteNil()
		} else {
			s.WriteTime(*x.ExpiresAt)
		}
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the ContactInfoValidation message from JSON.
func (x *ContactInfoValidation) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "id":
			s.AddField("id")
			x.Id = s.ReadString()
		case "token":
			s.AddField("token")
			x.Token = s.ReadString()
		case "entity":
			s.AddField("entity")
			// NOTE: EntityIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v EntityIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.Entity = &v
		case "contact_info", "contactInfo":
			s.AddField("contact_info")
			s.ReadArray(func() {
				if s.ReadNil() {
					x.ContactInfo = append(x.ContactInfo, nil)
					return
				}
				v := &ContactInfo{}
				v.UnmarshalProtoJSON(s.WithField("contact_info", false))
				if s.Err() != nil {
					return
				}
				x.ContactInfo = append(x.ContactInfo, v)
			})
		case "created_at", "createdAt":
			s.AddField("created_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.CreatedAt = v
		case "expires_at", "expiresAt":
			s.AddField("expires_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.ExpiresAt = v
		}
	})
}