package coretests

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
)

// SimpleGherkins
//
// https://www.guru99.com/gherkin-test-cucumber.html
// Feature:  Login functionality of social networking site Facebook.
// Given:  I am a facebook user.
// When: I enter username as username.
// And I enter the password as the password
// Then I should be redirected to the home page of facebook
// Given -> When -> Then
type SimpleGherkins struct {
	Feature,
	Given,
	When,
	Then,
	Expect,
	Actual string
}

func (it *SimpleGherkins) ToString(testIndex int) string {
	return errcore.GherkinsString(
		testIndex,
		it.Feature,
		it.Given,
		it.When,
		it.Then)
}

func (it *SimpleGherkins) String() string {
	return it.ToString(constants.Zero)
}

func (it *SimpleGherkins) GetWithExpectation(
	testIndex int,
) string {
	return errcore.GherkinsStringWithExpectation(
		testIndex,
		it.Feature,
		it.Given,
		it.When,
		it.Then,
		it.Actual,
		it.Expect)
}

func (it *SimpleGherkins) GetMessageConditional(
	isExpectation bool,
	testIndex int,
) string {
	if isExpectation {
		return it.GetWithExpectation(testIndex)
	}

	return it.ToString(testIndex)
}
