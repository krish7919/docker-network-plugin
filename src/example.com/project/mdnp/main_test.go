package main

import(
    "testing"
)

var (
    testDriverLocalScope *MyDockerNetworkPlugin
    testDriverGlobalScope *MyDockerNetworkPlugin
    err error
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

//TODO from here
//func TestCreateEndpoint(t *testing.T) {
//    req := &n.CreateEndpointRequest{
//check both scenarios of interface ebing populated and empty
//    }
//    resp, err :=  testDriverLocalScope.CreateEndpoint(req)
//}
// EndpointInfo(*InfoRequest) (*InfoResponse, error)
// Join(*JoinRequest) (*JoinResponse, error)



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

