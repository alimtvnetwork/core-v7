package errcoretests

// Mock types implementing errcore interfaces for testing

type mockErrorGetter struct {
	err error
}

func (m *mockErrorGetter) Error() error {
	return m.err
}

type mockCompiledErrorGetter struct {
	err error
}

func (m *mockCompiledErrorGetter) CompiledError() error {
	return m.err
}

type mockCompiledErrorWithTracesGetter struct {
	err error
}

func (m *mockCompiledErrorWithTracesGetter) CompiledErrorWithStackTraces() error {
	return m.err
}

type mockFullStringWithTracesGetter struct {
	err error
}

func (m *mockFullStringWithTracesGetter) FullStringWithTraces() error {
	return m.err
}
