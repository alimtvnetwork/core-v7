package coremath

func MinFloat32(v1, v2 float32) float32 {
	if v1 > v2 {
		return v2
	}

	return v1
}
