package watch

import (
	"fmt"
	"net"
	"strings"
)

func isValidIPv4(ipAddress string) bool {
	testInput := net.ParseIP(ipAddress)
	return testInput.To4() != nil
}

// lookupNetworkInterfaceIPv4Address return the first valid IPv4 address to the given network interface
func lookupNetworkInterfaceIPv4Address(name string) (string, error) {
	byNameInterface, err := net.InterfaceByName(name)
	if err != nil {
		return "", fmt.Errorf("network interface not found: %s", name)
	}

	addresses, err := byNameInterface.Addrs()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve addresses for interface: %s", name)
	}

	for _, addr := range addresses {
		tokens := strings.Split(addr.String(), "/")
		if isValidIPv4(tokens[0]) {
			return tokens[0], nil
		}
	}

	return "", fmt.Errorf("no addresses found for interface: %s", name)
}
