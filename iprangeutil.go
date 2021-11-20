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

func nextIPv4(ip0, ip1, ip2, ip3,
	endIP0, endIP1, endIP2, endIP3 *uint8) bool {
	if (*ip0 == *endIP0) && (*ip1 == *endIP1) &&
		(*ip2 == *endIP2) && (*ip3 == *endIP3) {
		return true
	}

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

	return false
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

		if nextIPv4(&startIP[0], &startIP[1], &startIP[2], &startIP[3],
			&endIP[0], &endIP[1], &endIP[2], &endIP[3]) {
			break
		}
	}

	return
}
