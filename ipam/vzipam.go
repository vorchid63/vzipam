package vzipam

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
        ipamapi "github.com/docker/go-plugins-helpers/api"
)

//======================================================================
// implementation of the ipam.Driver interface
//======================================================================


// GetCapabilities find out if mojaisolator requires mac address or not
func (isolator *vzipam.Ipam) GetCapabilities() (*vzipam.CapabilitiesResponse, error) {
	log.Errorf("VLU-GetCapabilities: mojaisolator")
	r := vzipam.CapabilitiesResponse{false}
	return &r, nil
}

// GetDefaultAddressSpaces get mojaisolator's default address space
func (isolator *vzipam.Ipam) GetDefaultAddressSpaces() (*vzipam.AddressSpacesResponse, error) {
     	log.Errorf("VLU-GetDefaultAddressSpaces:")	
	r := vzipam.AddressSpacesResponse{"mojaLocal", "mojaGlobal"}
     	log.Errorf("VLU-GetDefaultAddressSpaces: vzipam %#v error=%#v", r, nil)	
	return &r, nil
}

// RequestPool request an address pool from mojaisolator
func (isolator *vzipam.Ipam) RequestPool(*vzipam.RequestPoolRequest) (*vzipam.RequestPoolResponse, error) {
     	log.Errorf("VLU-RequestPoolRequest:")	
	r := vzipam.RequestPoolResponse{"poolId", "pool", map[string]string{netlabel.Prefix: "10.2.3.0"}}
     	log.Errorf("VLU-RequestPoolRequest: vzipam %#v error=%#v", r, nil)	
	return &r, nil
}

// ReleasePool relase an address pool from mojaisolator
func (isolator *vzipam.Ipam) ReleasePool(*vzipam.ReleasePoolRequest) error {
     	log.Errorf("VLU-ReleasePoolRequest:")	
	return nil
}

// RequestAddress request an address from mojaisolator
func (isolator *vzipam.Ipam) RequestAddress(*vzipam.RequestAddressRequest) (*vzipam.RequestAddressResponse, error) {
     	log.Errorf("VLU-RequestAddress:")	
	r := vzipam.RequestAddressResponse{"10.2.3.x", map[string]string{netlabel.DriverPrefix: "10.2.3.5"}}
     	log.Errorf("VLU-RequestAddress: vzipam %#v error=%#v", r, nil)	
	return &r, nil
}

// ReleaseAddress release an address from mojaisolator
func (isolator *vzipam.Ipam) ReleaseAddress(*vzipam.ReleaseAddressRequest) error {
     	log.Errorf("VLU-ReleaseAddress:")	
	return nil
}

