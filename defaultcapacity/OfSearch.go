package defaultcapacity

import "gitlab.com/evatix-go/core/constants"

func OfSearch(length int) int {
	if length <= constants.Capacity3 {
		return length
	}

	if length > constants.Capacity3 && length <= constants.N20 {
		return length / constants.Capacity3
	}

	defaultCapacity := length

	if length >= constants.ArbitraryCapacity1000 {
		defaultCapacity = constants.ArbitraryCapacity100
	} else if length > constants.ArbitraryCapacity250 {
		defaultCapacity = length / constants.N20
	} else if length >= constants.ArbitraryCapacity100 {
		defaultCapacity = length / constants.N10
	} else {
		defaultCapacity = length / constants.N5
	}

	return defaultCapacity
}
