package main

import (
	n "github.com/docker/go-plugins-helpers/network"
	"testing"
)

var (
	testDriverLocalScope  *MyDockerNetworkPlugin
	testDriverGlobalScope *MyDockerNetworkPlugin
	err                   error
)

func init() {
	// TODO: err checking
	testDriverLocalScope, err = NewMyDockerNetworkPlugin("local")
	// TODO(Krish): test this
	// testDriverGlobalScope = NewMyDockerNetworkPlugin("global")
}

func TestGetCapabilities(t *testing.T) {
	resp, err := testDriverLocalScope.GetCapabilities()
	if err != nil {
		t.Error(err)
	}
	if resp.Scope != "local" {
		t.Errorf("Expected: local, Got: %s", resp.Scope)
	}
}

func TestCreateEndpointEmptyInterface(t *testing.T) {
	options := make(map[string]interface{})

	// case where interface is blank
	req := &n.CreateEndpointRequest{
		NetworkID:  "1234567890abcdef",
		EndpointID: "fedcba9876543210",
		Interface:  nil,
		Options:    options,
	}
	resp, err := testDriverLocalScope.CreateEndpoint(req)
	if err != nil {
		t.Error(err)
	}
	if resp.Interface.Address != "1.1.1.1/24" {
		t.Errorf("Expected 1.1.1.1/24, Got: %s", resp.Interface.Address)
	}
	if resp.Interface.MacAddress != "00:00:00:00:00:aa" {
		t.Errorf("Expected 00:00:00:00:00:aa, Got: %s", resp.Interface.MacAddress)
	}
}

func TestCreateEndpointNonEmptyInterface(t *testing.T) {
	intf := &n.EndpointInterface{
		Address: "1.1.1.1",
		//AddressIPv6 - optional
		MacAddress: "00:00:00:00:00:aa",
	}
	options := make(map[string]interface{})

	// case where interface is pre-populated by daemon
	req := &n.CreateEndpointRequest{
		NetworkID:  "1234567890abcdef",
		EndpointID: "fedcba9876543210",
		Interface:  intf,
		Options:    options,
	}
	resp, err := testDriverLocalScope.CreateEndpoint(req)
	if err != nil {
		t.Error(err)
	}
	if resp.Interface.Address != "" {
		t.Errorf("Expected empty, Got: %s", resp.Interface.Address)
	}
	if resp.Interface.MacAddress != "" {
		t.Errorf("Expected empty, Got: %s", resp.Interface.MacAddress)
	}
}

func TestEndpointInfo(t *testing.T) {
	req := &n.InfoRequest{
		NetworkID:  "1234567890abcdef",
		EndpointID: "fedcba0987654321",
	}
	resp, err := testDriverLocalScope.EndpointInfo(req)
	if err != nil {
		t.Error(err)
	}
	if len(resp.Value) != 0 {
		t.Errorf("Expected empty map, Got: %s", resp.Value)
	}
}

func TestJoin(t *testing.T) {
	req := &n.JoinRequest{
		NetworkID:  "1234567890abcdef",
		EndpointID: "fedcba0987654321",
		SandboxKey: "aabbccddeeff",
		Options:    make(map[string]interface{}),
	}
	resp, err := testDriverLocalScope.Join(req)
	if err != nil {
		t.Error(err)
	}
	if resp.InterfaceName.SrcName != "veth0" {
		t.Errorf("Expected: veth0, Got: %s", resp.InterfaceName.SrcName)
	}
	if resp.InterfaceName.DstPrefix != "krish-mdnp" {
		t.Errorf("Expected: krish-mdnp, Got: %s", resp.InterfaceName.DstPrefix)
	}
	if resp.Gateway != "" {
		t.Errorf("Expected empty, Got: %s", resp.Gateway)
	}
	if resp.GatewayIPv6 != "" {
		t.Errorf("Expected empty, Got: %s", resp.GatewayIPv6)
	}
	if resp.StaticRoutes != nil {
		t.Errorf("Expected nil, Got: %+v", resp.StaticRoutes)
	}
	if resp.DisableGatewayService != false {
		t.Errorf("Expected false, Got: %s", resp.DisableGatewayService)
	}
}

// TODO(Krish): noop methods - test later
//func TestCreateNetwork(t *testing.T) {
//    req := &n.CreateNetworkRequest{
//    }
//}
// CreateNetwork
// DeleteNetwork(*DeleteNetworkRequest) error
// DeleteEndpoint(*DeleteEndpointRequest) error
// Leave(*LeaveRequest) error
// DiscoverNew(*DiscoveryNotification) error
// DiscoverDelete(*DiscoveryNotification) error
// ProgramExternalConnectivity(*ProgramExternalConnectivityRequest) error
// RevokeExternalConnectivity(*RevokeExternalConnectivityRequest) error
