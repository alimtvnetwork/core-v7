package internalenuminf

type HttpMethodTyper interface {
	BasicEnumer
	ByteValuePlusEqualer

	IsGetHttp() bool
	IsPutHttp() bool
	IsPostHttp() bool
	IsDeleteHttp() bool
	IsPatchHttp() bool

	IsAnyHttpMethod(methodNames ...string) bool
	IsAnyHttp() bool
	IsNotHttpMethod() bool
}
