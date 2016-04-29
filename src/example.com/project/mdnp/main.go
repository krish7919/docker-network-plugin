// main package contains the implementation of my docker network plugin, a.k.a.
// mdnp
package main

import (
	n "github.com/docker/go-plugins-helpers/network"
	"log"
)

var (
	PLUGIN_NAME = "mdnp"
)

// MyDockerNetworkPlugin implements the Driver interface
type MyDockerNetworkPlugin struct {
	scope string
}

// --- Plugin handlers start here ---
func (self *MyDockerNetworkPlugin) GetCapabilities() (*n.CapabilitiesResponse,
	error) {
	log.Printf("Received GetCapabilities req")
	capabilities := &n.CapabilitiesResponse{
		Scope: self.scope,
	}
	return capabilities, nil
}

func (self *MyDockerNetworkPlugin) CreateNetwork(req *n.CreateNetworkRequest) error {
	log.Printf("Received CreateNetwork req:\n%+v\n", req)
	// used for plumbing, do nothing API for now
	return nil
}

func (self *MyDockerNetworkPlugin) DeleteNetwork(req *n.DeleteNetworkRequest) error {
	log.Printf("Received DeleteNetwork req:\n%+v\n", req)
	// used for plumbing, do nothing API for now
	return nil
}

func (self *MyDockerNetworkPlugin) CreateEndpoint(req *n.CreateEndpointRequest) (*n.CreateEndpointResponse, error) {
	log.Printf("Received CreateEndpoint req:\n%+v\n", req)

	// If the remote process was supplied a non-empty value in Interface, it
	// must respond with an empty Interface value. LibNetwork will treat it as
	// an error if it supplies a non-empty value and receives a non-empty value
	// back, and roll back the operation.
	intfInfo := new(n.EndpointInterface)

	if req.Interface == nil {
		// case never hit in docker v1.11.0, but in tests
		intfInfo.Address = "1.1.1.1/24"
		// AddressIPv6 - optional
		intfInfo.MacAddress = "00:00:00:00:00:aa"
	}
	resp := &n.CreateEndpointResponse{
		Interface: intfInfo,
	}
	return resp, nil
}

func (self *MyDockerNetworkPlugin) DeleteEndpoint(req *n.DeleteEndpointRequest) error {
	log.Printf("Received DeleteEndpoint req:\n%+v\n", req)
	// used for plumbing, do nothing API for now
	return nil
}

func (self *MyDockerNetworkPlugin) EndpointInfo(req *n.InfoRequest) (*n.InfoResponse, error) {
	log.Printf("Received EndpointOperInfo req:\n%+v\n", req)
	// return empty map for now - the value of the Value field is an arbitrary (possibly empty) map
	value := make(map[string]string)
	resp := &n.InfoResponse{
		Value: value,
	}
	return resp, nil
}

func (self *MyDockerNetworkPlugin) Join(req *n.JoinRequest) (*n.JoinResponse, error) {
	log.Printf("Received Join req:\n%+v\n", req)
	resp := &n.JoinResponse{
		InterfaceName: n.InterfaceName{
			// SrcName is the name of the OS level interface that the remote
			// process created
			SrcName:   "veth0",
			DstPrefix: "krish-mdnp",
		},
		// my local veth IP
		//Gateway: "172.16.0.205",
		// GatewayIPv6 - optional
		//StaticRoutes: []*n.StaticRoute{
		//	{
		//		Destination: "8.8.8.8/8",
		//		RouteType:   1,
		//		//RouteType:   0,
		//		//NextHop:     "172.16.0.205",
		//	},
		//},
		// confusion on what DisableGatewayService does - TODO google this!
		DisableGatewayService: false,
	}
	return resp, nil
}

func (self *MyDockerNetworkPlugin) Leave(req *n.LeaveRequest) error {
	log.Printf("Received Leave req:\n%+v\n", req)
	// used for plumbing, do nothing API for now
	return nil
}

func (self *MyDockerNetworkPlugin) DiscoverNew(req *n.DiscoveryNotification) error {
	log.Printf("Received DiscoverNew req:\n%+v\n", req)
	if req.DiscoveryType == 1 {
		// Node Discovery
		log.Printf("....which is of type NodeDiscovery\n")
	}
	// used for plumbing, do nothing API for now
	return nil
}

func (self *MyDockerNetworkPlugin) DiscoverDelete(req *n.DiscoveryNotification) error {
	log.Printf("Received DiscoverDelete req:\n%+v\n", req)
	if req.DiscoveryType == 1 {
		// Node Discovery
		log.Printf("....which is of type NodeDiscovery\n")
	}
	// used for plumbing, do nothing API for now
	return nil
}

func (self *MyDockerNetworkPlugin) ProgramExternalConnectivity(req *n.ProgramExternalConnectivityRequest) error {
	log.Printf("Received ProgramExternalConnectivity req:\n%+v\n", req)
	// no documentation?!
	// used for plumbing, do nothing API for now
	return nil
}

func (self *MyDockerNetworkPlugin) RevokeExternalConnectivity(req *n.RevokeExternalConnectivityRequest) error {
	log.Printf("Received RevokeExternalConnectivity req:\n%+v\n", req)
	// no documentation?!
	// used for plumbing, do nothing API for now
	return nil
}

// --- Plugin handlers end here ---

// return the default prod config for ad_service
func NewMyDockerNetworkPlugin(scope string) (*MyDockerNetworkPlugin, error) {
	mdnp := &MyDockerNetworkPlugin{
		scope: scope, // TODO(Krish): local vs global?
	}
	return mdnp, nil
}

func main() {
	driver, err := NewMyDockerNetworkPlugin("local")
	if err != nil {
		log.Fatalf("ERROR: %s init failed!", PLUGIN_NAME)
	}
	requestHandler := n.NewHandler(driver)
	requestHandler.ServeTCP(PLUGIN_NAME, ":2804")
}
