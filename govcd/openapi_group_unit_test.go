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

func TestOpenApiGroupWireType(t *testing.T) {
	config := &types.OpenApiGroup{
		Name:           "group@example.test",
		ProviderType:   "OAUTH",
		OrgEntityRef:   &types.OpenApiReference{ID: "urn:vcloud:org:one"},
		RoleEntityRefs: []types.OpenApiReference{{ID: "urn:vcloud:role:one"}},
	}

	payload, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	var createPayload map[string]any
	if err := json.Unmarshal(payload, &createPayload); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if _, found := createPayload["nameInSource"]; found {
		t.Fatal("create payload must omit server-computed nameInSource")
	}
	if got := createPayload["providerType"]; got != "OAUTH" {
		t.Fatalf("providerType = %v, want OAUTH", got)
	}

	response := []byte(`{
		"id":"urn:vcloud:group:one",
		"name":"group@example.test",
		"nameInSource":"group@example.test",
		"description":"test group",
		"providerType":"OAUTH",
		"orgEntityRef":{"id":"urn:vcloud:org:one"},
		"roleEntityRefs":[{"id":"urn:vcloud:role:one"}],
		"sourceEntityRef":null
	}`)
	var group types.OpenApiGroup
	if err := json.Unmarshal(response, &group); err != nil {
		t.Fatalf("json.Unmarshal() response error = %v", err)
	}
	if group.NameInSource != config.Name || group.SourceEntityRef != nil {
		t.Fatalf("unexpected response type: %#v", group)
	}
}
