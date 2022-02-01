package coreinterface

type StandardExecutor interface {
	SimpleExecutor
	BasicEnumerGetter
	ByteTypeEnumGetter
	StandardLoggerGetter
	RawPayloadsGetter
	MustExecutor
}

type StandardLoggerGetter interface {
	StandardLogger() StandardLogger
}
