package corevalidator

type BaseValidatorCoreCondition struct {
	ValidatorCoreCondition *Condition `json:"Condition,omitempty"`
}

func (it *BaseValidatorCoreCondition) ValidatorCoreConditionDefault() Condition {
	if it.ValidatorCoreCondition != nil {
		return *it.ValidatorCoreCondition
	}

	it.ValidatorCoreCondition = &Condition{}

	return *it.ValidatorCoreCondition
}
