package resource

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

type NetworkType string

const (
	NetworkBridge NetworkType = "bridge"
	NetworkNAT    NetworkType = "nat"
	NetworkHost   NetworkType = "host"
)

type NetworkStatus string

const (
	StateDown       NetworkStatus = "down"
	StateUp         NetworkStatus = "up"
	StateConnecting NetworkStatus = "connecting"
	StateConnected  NetworkStatus = "connected"
)

type IPAddress struct {
	Address netip.Addr   `db:"addr" json:"addr"`
	Prefix  netip.Prefix `db:"prefix" json:"prefix"`
}

type NetworkInterface struct {
	Name string `db:"name" json:"name"`

	InterfaceType NetInterfaceType `db:"type" json:"type"`
	State         NetworkStatus    `db:"state" json:"state"`

	MACAddress string `db:"mac" json:"mac"`
	MTU        int    `db:"mtu" json:"mtu"`

	IPv4 []IPAddress `db:"ipv4" json:"ipv4"`
	IPv6 []IPAddress `db:"ipv6" json:"ipv6"`

	GatewayIPv4 string `db:"gateway_ipv4 "json:"gateway_ipv4"`
	GatewayIPv6 string `db:"gateway_ipv6 "json:"gateway_ipv6"`

	DNS []string `db:"dns" json:"dns"`

	Connected bool `db:"connected" json:"connected"`
}

type IPMethod string

const (
	IPMethodDHCP  IPMethod = "dhcp"
	IPMethodState IPMethod = "static"
)

type NetworkConfig NetworkInterface

type Network Resource[NetworkConfig, NetworkStatus]
