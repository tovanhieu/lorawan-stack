// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
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

export default [
  'error:pkg/account/session:no_user_id_password_match',
  'error:pkg/deviceclaimingserver/gateways:gateway_not_authorized',
  'error:pkg/identityserver/blacklist:blacklisted_id',
  'error:pkg/identityserver/gormstore:account_type',
  'error:pkg/identityserver/store:entity_quota',
  'error:pkg/identityserver:application_has_devices',
  'error:pkg/identityserver:application_needs_collaborator',
  'error:pkg/identityserver:external_user',
  'error:pkg/identityserver:gateway_needs_collaborator',
  'error:pkg/identityserver:organization_needs_collaborator',
  'error:pkg/identityserver:password_contains_user_id',
  'error:pkg/identityserver:password_equals_old',
]