package corestr

import "strings"

type KeyValuePair struct {
	Key, Value string
}

func NewKeyValuePairTrimmed(key, val string) *KeyValuePair {
	return &KeyValuePair{
		Key:   strings.TrimSpace(key),
		Value: strings.TrimSpace(val),
	}
}
