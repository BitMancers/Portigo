package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generate random MAC address (standard format)
func RandomMacAddress() string {
	octets := make([]string, 6)
	for i := 0; i < 6; i++ {
		octets[i] = fmt.Sprintf("%02X", rand.Intn(256))
	}
	return strings.Join(octets, ":")
}

// Generate locally administered MAC address (bit 1 of first octet set)
func RandomLocalMacAddress() string {
	// First octet: set bit 1 (locally administered), clear bit 0 (unicast)
	firstOctet := (rand.Intn(128) & 0xFE) | 0x02
	octets := []string{fmt.Sprintf("%02X", firstOctet)}

	for i := 1; i < 6; i++ {
		octets = append(octets, fmt.Sprintf("%02X", rand.Intn(256)))
	}
	return strings.Join(octets, ":")
}

// Generate MAC with specific vendor prefix (OUI)
func RandomMacWithVendor(vendorPrefix string) string {
	vendor := strings.Split(vendorPrefix, ":")
	octets := vendor

	for i := len(vendor); i < 6; i++ {
		octets = append(octets, fmt.Sprintf("%02X", rand.Intn(256)))
	}
	return strings.Join(octets, ":")
}

// Generate MAC with different separator
func RandomMacFormatted(separator string) string {
	octets := make([]string, 6)
	for i := 0; i < 6; i++ {
		octets[i] = fmt.Sprintf("%02X", rand.Intn(256))
	}
	return strings.Join(octets, separator)
}

// Validate MAC address format
func IsValidMacAddress(mac string) bool {
	patterns := []string{
		`^([0-9A-Fa-f]{2}:){5}[0-9A-Fa-f]{2}$`, // 00:1A:2B:3C:4D:5E
		`^([0-9A-Fa-f]{2}-){5}[0-9A-Fa-f]{2}$`, // 00-1A-2B-3C-4D-5E
		`^[0-9A-Fa-f]{12}$`,                    // 001A2B3C4D5E
	}

	for _, pattern := range patterns {
		matched, _ := regexp.MatchString(pattern, mac)
		if matched {
			return true
		}
	}
	return false
}

// Generate multiple MAC addresses
func GenerateMacAddresses(count int) []string {
	macs := make([]string, count)
	for i := 0; i < count; i++ {
		macs[i] = RandomMacAddress()
	}
	return macs
}
