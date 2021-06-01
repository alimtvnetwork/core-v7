package defaultcapacity

import "math"

// Predictive Result must be positive possibleLen * multiplier + additionalCap.
//
// Less than zero yields additionalCap
func Predictive(
	possibleLen int,
	multiplier float64,
	additionalCap int,
) int {
	if possibleLen <= 0 {
		return additionalCap
	}

	multipliedVal := float64(possibleLen) * multiplier
	multipliedInt := int(math.Ceil(multipliedVal))
	finalCap := multipliedInt + additionalCap

	return finalCap
}
