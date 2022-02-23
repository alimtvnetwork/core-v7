package coreinterface

import "gitlab.com/evatix-go/core/internal/internalinterface/internalserializer"

type Serializer interface {
	internalserializer.Serializer
}
