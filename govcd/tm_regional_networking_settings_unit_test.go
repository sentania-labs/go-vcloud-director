//go:build unit || ALL

// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package govcd

import (
	"encoding/json"
	"testing"

	"github.com/vmware/go-vcloud-director/v3/types/v56"
)

func TestTmRegionalNetworkingAviSettingWireType(t *testing.T) {
	wireValue := []byte(`{
		"status":null,
		"active":true,
		"serviceEngineGroupMode":"TENANT_MANAGED",
		"serviceEngineQuota":60,
		"applicationLimit":null,
		"serviceEngineGroupRefs":null
	}`)

	var setting types.TmRegionalNetworkingAviSetting
	if err := json.Unmarshal(wireValue, &setting); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if !setting.Active || setting.ServiceEngineGroupMode != "TENANT_MANAGED" || setting.ServiceEngineQuota != 60 {
		t.Fatalf("unexpected Avi setting: %#v", setting)
	}
	if setting.Status != nil || setting.ApplicationLimit != nil || setting.ServiceEngineGroupRefs != nil {
		t.Fatalf("nullable fields were not preserved: %#v", setting)
	}

	payload, err := json.Marshal(&setting)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}
	var roundTrip map[string]any
	if err := json.Unmarshal(payload, &roundTrip); err != nil {
		t.Fatalf("json.Unmarshal() round trip error = %v", err)
	}
	if value, found := roundTrip["applicationLimit"]; !found || value != nil {
		t.Fatalf("applicationLimit = %v, found = %v, want explicit null", value, found)
	}
	if value, found := roundTrip["serviceEngineGroupRefs"]; !found || value != nil {
		t.Fatalf("serviceEngineGroupRefs = %v, found = %v, want explicit null", value, found)
	}
}

func TestTmRegionalNetworkingAviSettingInactivePayload(t *testing.T) {
	payload, err := json.Marshal(&types.TmRegionalNetworkingAviSetting{Active: false})
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}
	var config map[string]any
	if err := json.Unmarshal(payload, &config); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if _, found := config["serviceEngineGroupMode"]; found {
		t.Fatal("inactive payload must omit serviceEngineGroupMode")
	}
	if _, found := config["serviceEngineQuota"]; found {
		t.Fatal("inactive payload must omit serviceEngineQuota")
	}
}
