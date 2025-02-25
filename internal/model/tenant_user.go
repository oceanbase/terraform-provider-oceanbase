/*
 * Copyright 2025 OceanBase
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/oceanbasecloud/terraform-provider-oceanbase/obcloudsdk"
)

type TenantUser struct {
	Id             types.String        `tfsdk:"id"`
	Name           types.String        `tfsdk:"name"`
	ClusterId      types.String        `tfsdk:"cluster_id"`
	TenantId       types.String        `tfsdk:"tenant_id"`
	Password       types.String        `tfsdk:"password"`
	Description    types.String        `tfsdk:"description"`
	Type           types.String        `tfsdk:"type"`
	Status         types.String        `tfsdk:"status"`
	IsLocked       types.Bool          `tfsdk:"is_locked"`
	EncryptionType types.String        `tfsdk:"encryption_type"`
	Privileges     basetypes.ListValue `tfsdk:"privileges"`
}

type TenantUserPrivilege struct {
	WithGrant  bool   `tfsdk:"with_grant" json:"withGrant"`
	Table      string `tfsdk:"table" json:"table"`
	Role       string `tfsdk:"role" json:"role"`
	Database   string `tfsdk:"database" json:"database"`
	Privileges string `tfsdk:"privileges" json:"privileges"`
}

func (u *TenantUser) GetPrivileges() ([]TenantUserPrivilege, error) {
	return parseListValue(u.Privileges, parseObjectValue[TenantUserPrivilege])
}

func (u *TenantUser) GetPrivilegesStr() (string, error) {
	privileges, err := u.GetPrivileges()
	if err == nil {
		privilegeStr, err := json.Marshal(privileges)
		if err == nil {
			return string(privilegeStr), nil
		}
	}
	return "", err
}

func extractPrivileges(data *obcloudsdk.OcpDbUserDtoV2) []TenantUserPrivilege {
	privileges := make([]TenantUserPrivilege, 0)
	for _, db := range data.GetDatabases() {
		dbPrivileges := ""
		if db.GetRole() == DatabaseRoleOther {
			dbPrivileges = db.GetPrivileges()
		}
		privilege := TenantUserPrivilege{
			WithGrant:  db.GetWithGrant() != 0,
			Table:      db.GetTable(),
			Database:   db.GetDatabase(),
			Role:       db.GetRole(),
			Privileges: dbPrivileges,
		}
		privileges = append(privileges, privilege)
	}
	return privileges
}

func mergePrivileges(statePrivileges, currentPrivileges []TenantUserPrivilege) basetypes.ListValue {
	stateMap := make(map[string]TenantUserPrivilege)
	currentMap := make(map[string]TenantUserPrivilege)
	for _, sp := range statePrivileges {
		stateMap[sp.Database] = sp
	}
	for _, cp := range currentPrivileges {
		currentMap[cp.Database] = cp
	}

	// drop the privileges in state which are not in current privileges
	dbsDropped := make([]string, 0)
	for db, _ := range stateMap {
		if _, ok := currentMap[db]; !ok {
			dbsDropped = append(dbsDropped, db)
		}
	}
	for _, db := range dbsDropped {
		delete(stateMap, db)
	}

	// iterate over current privileges and update state privileges
	for db, currentPrivilege := range currentMap {
		statePrivilege, ok := stateMap[db]
		if ok && currentPrivilege.Role == DatabaseRoleOther && statePrivilege.Role == DatabaseRoleOther {
			currentParts := strings.Split(currentPrivilege.Privileges, ",")
			stateParts := strings.Split(statePrivilege.Privileges, ",")
			sort.Strings(currentParts)
			sort.Strings(stateParts)
			cps := strings.Join(currentParts, ",")
			sps := strings.Join(stateParts, ",")
			if cps != sps {
				stateMap[db] = currentPrivilege
			} else {
				statePrivilege.Role = currentPrivilege.Role
				statePrivilege.Table = currentPrivilege.Table
				statePrivilege.WithGrant = currentPrivilege.WithGrant
			}
		} else {
			stateMap[db] = currentPrivilege
		}
	}

	objectType := types.ObjectType{
		AttrTypes: map[string]attr.Type{"database": basetypes.StringType{}, "table": basetypes.StringType{}, "role": basetypes.StringType{}, "privileges": basetypes.StringType{}, "with_grant": basetypes.BoolType{}},
	}

	privileges := make([]attr.Value, 0)
	for _, sp := range stateMap {
		attributeTypes := map[string]attr.Type{"database": basetypes.StringType{}, "table": basetypes.StringType{}, "role": basetypes.StringType{}, "privileges": basetypes.StringType{}, "with_grant": basetypes.BoolType{}}
		attributes := map[string]attr.Value{"database": types.StringValue(sp.Database), "table": types.StringValue(sp.Table), "role": types.StringValue(sp.Role), "privileges": types.StringValue(sp.Privileges), "with_grant": types.BoolValue(sp.WithGrant)}
		privilegeObject, _ := basetypes.NewObjectValue(attributeTypes, attributes)
		privileges = append(privileges, privilegeObject)
	}
	result, _ := types.ListValue(objectType, privileges)
	return result
}

func (u *TenantUser) Refresh(data *obcloudsdk.OcpDbUserDtoV2) {
	u.Name = types.StringValue(data.GetUserName())
	u.Type = types.StringValue(data.GetUserType())
	u.Status = types.StringValue(data.GetUserStatus())
	u.IsLocked = types.BoolValue(data.GetUserStatus() == TenantUserStatusLocked)
	u.Description = types.StringValue(data.GetDescription())
	statePrivileges, err := u.GetPrivileges()
	if err == nil {
		currentPrivileges := extractPrivileges(data)
		u.Privileges = mergePrivileges(statePrivileges, currentPrivileges)
	}
}
