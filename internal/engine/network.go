package engine

import (
	"net/netip"
)

type NetInterfaceType string

const (
	InterfaceEthernet NetInterfaceType = "ethernet"
	InterfaceWiFi     NetInterfaceType = "wifi"
	InterfaceLoopback NetInterfaceType = "loopback"
	InterfaceVPN      NetInterfaceType = "vpn"
	InterfaceBridge   NetInterfaceType = "bridge"
)

type InterfaceState string

const (
	StateDown       InterfaceState = "down"
	StateUp         InterfaceState = "up"
	StateConnecting InterfaceState = "connecting"
	StateConnected  InterfaceState = "connected"
)

type IPAddress struct {
	Address netip.Addr   `json:"addr"`
	Prefix  netip.Prefix `json:"prefix"`
}

type NetworkInterface struct {
	ID int `json:"id"`

	Name        string `json:"name"`
	DisplayName string `json:"display_name"`

	InterfaceType NetInterfaceType `json:"type"`
	State         InterfaceState   `json:"state"`

	MACAddress string `json:"mac"`
	MTU        int    `json:"mtu"`

	IPv4 []IPAddress `json:"ipv4"`
	IPv6 []IPAddress `json:"ipv6"`

	GatewayIPv4 string `json:"gateway_ipv4"`
	GatewayIPv6 string `json:"gateway_ipv6"`

	DNS []string `json:"dns"`

	Connected bool `json:"connected"`
}

type IPMethod string

const (
	IPMethodDHCP  IPMethod = "dhcp"
	IPMethodState IPMethod = "static"
)
