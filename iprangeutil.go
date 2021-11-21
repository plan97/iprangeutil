// Package iprangeutil provides utilities for IP range operations.
package iprangeutil

import (
	"errors"
	"fmt"
)

var (
	ErrStartingIP = errors.New("starting IP error")
	ErrEndingIP   = errors.New("ending IP error")
	ErrIPFunc     = errors.New("ipFunc error")
)

// IPv4Func type allows the IP v4 address to be used in a function.
type IPv4Func func(ip0, ip1, ip2, ip3 *uint8) error

// CompareIPv4 returns true if both the input IP v4 addresses are the same.
func CompareIPv4(ip0, ip1, ip2, ip3,
	compareIP0, compareIP1, compareIP2, compareIP3 *uint8) bool {
	if (*ip0 == *compareIP0) && (*ip1 == *compareIP1) &&
		(*ip2 == *compareIP2) && (*ip3 == *compareIP3) {
		return true
	}
	return false
}

// NextIPv4 increments the input IP v4 address by 1.
// If the input IP v4 address is 255.255.255.255,
// then it would increment it to 0.0.0.0.
func NextIPv4(ip0, ip1, ip2, ip3 *uint8) {
	*ip3 += 1

	if *ip3 == 0 {
		*ip2 += 1

		if *ip2 == 0 {
			*ip1 += 1

			if *ip1 == 0 {
				*ip0 += 1
			}
		}
	}
}

// ExpandIPv4 would traverse all IP v4 addresses from start to end
// and call ipFunc for each of the traversed IP v4 address.
func ExpandIPv4(start, end string, ipFunc IPv4Func) (err error) {
	startIP := make([]uint8, 4)
	endIP := make([]uint8, 4)

	_, err = fmt.Sscanf(start, "%d.%d.%d.%d", &startIP[0], &startIP[1], &startIP[2], &startIP[3])
	if err != nil {
		return fmt.Errorf("%s: %w", ErrStartingIP, err)
	}

	_, err = fmt.Sscanf(end, "%d.%d.%d.%d", &endIP[0], &endIP[1], &endIP[2], &endIP[3])
	if err != nil {
		return fmt.Errorf("%s: %w", ErrEndingIP, err)
	}

	if ipFunc == nil {
		return
	}

	for {
		err = ipFunc(&startIP[0], &startIP[1], &startIP[2], &startIP[3])
		if err != nil {
			return fmt.Errorf("%s: %w", ErrIPFunc, err)
		}

		if CompareIPv4(&startIP[0], &startIP[1], &startIP[2], &startIP[3],
			&endIP[0], &endIP[1], &endIP[2], &endIP[3]) {
			break
		}

		NextIPv4(&startIP[0], &startIP[1], &startIP[2], &startIP[3])
	}

	return
}
