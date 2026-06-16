package services

import (
	"fmt"
	"net"
	"strings"
)

func ValidateIPCIDR(cidr string) error {
	_, _, err := net.ParseCIDR(cidr)
	if err != nil {
		return fmt.Errorf("invalid CIDR notation: %s", cidr)
	}
	return nil
}

func ValidateConfigType(configType string) error {
	valid := []string{"terraform", "ansible", "firewall", "cloudformation"}
	for _, v := range valid {
		if strings.EqualFold(configType, v) {
			return nil
		}
	}
	return fmt.Errorf("invalid config type: %s. Must be one of: %s", configType, strings.Join(valid, ", "))
}
