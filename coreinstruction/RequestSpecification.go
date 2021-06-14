package coreinstruction

type RequestSpecification struct {
	BaseIdentifier
	BaseTypeDotFilter
	BaseTags
	BaseIsGlobal
	BaseIsContinueOnError
	BaseIsRunAll
}

func (r RequestSpecification) ClonePtr() *RequestSpecification {
	return &RequestSpecification{
		BaseIdentifier:    BaseIdentifier{r.Id},
		BaseTypeDotFilter: BaseTypeDotFilter{
			splitDotFilters: nil,
			TypeDotFilter:   r.TypeDotFilter,
		},
		BaseTags: BaseTags{
			Tags: r.Tags,
		},
		BaseIsGlobal:          BaseIsGlobal{r.IsGlobal},
		BaseIsContinueOnError: BaseIsContinueOnError{r.IsContinueOnError},
		BaseIsRunAll:          BaseIsRunAll{r.IsRunAll},
	}
}
