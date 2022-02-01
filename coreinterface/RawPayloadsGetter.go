package coreinterface

type RawPayloadsGetter interface {
	RawPayloads() (payloads []byte, err error)
	RawPayloadsMust() (payloads []byte)
}
