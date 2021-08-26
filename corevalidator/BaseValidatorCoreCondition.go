package corevalidator

type BaseValidatorCoreCondition struct {
	ValidatorCoreCondition *ValidatorCoreCondition `json:"ValidatorCoreCondition,omitempty"`
}

func (it *BaseValidatorCoreCondition) ValidatorCoreConditionDefault() ValidatorCoreCondition {
	if it.ValidatorCoreCondition != nil {
		return *it.ValidatorCoreCondition
	}

	it.ValidatorCoreCondition = &ValidatorCoreCondition{}

	return *it.ValidatorCoreCondition
}
