package coretests

type SimpleTestCaseWrapperContractsBinder interface {
	SimpleTestCaseWrapper
	AsSimpleTestCaseWrapper() SimpleTestCaseWrapper
}
