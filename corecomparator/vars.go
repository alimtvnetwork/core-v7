package corecomparator

var (
	CompareNames = [...]string{
		Equal:            "Equal",
		LeftGreater:      "LeftGreater",
		LeftGreaterEqual: "LeftGreaterEqual",
		LeftLess:         "LeftLess",
		LeftLessEqual:    "LeftLessEqual",
		NotEqual:         "NotEqual",
		Inconclusive:     "Inconclusive",
	}

	CompareOperatorsSymbols = [...]string{
		Equal:            "=",
		LeftGreater:      ">",
		LeftGreaterEqual: ">=",
		LeftLess:         "<",
		LeftLessEqual:    "<=",
		NotEqual:         "!=",
		Inconclusive:     "?!",
	}

	CompareOperatorsShotNames = [...]string{
		Equal:            "eq",
		LeftGreater:      "gt",
		LeftGreaterEqual: "ge",
		LeftLess:         "lt",
		LeftLessEqual:    "le",
		NotEqual:         "ne",
		Inconclusive:     "i",
	}

	SqlCompareOperators = [...]string{
		Equal:            "=",
		LeftGreater:      ">",
		LeftGreaterEqual: ">=",
		LeftLess:         "<",
		LeftLessEqual:    "<=",
		NotEqual:         "<>",
		Inconclusive:     "i",
	}

	RangesMap = map[string]Compare{
		Equal.Name():                         Equal,
		LeftGreater.Name():                   LeftGreater,
		LeftGreaterEqual.Name():              LeftGreaterEqual,
		LeftLess.Name():                      LeftLess,
		LeftLessEqual.Name():                 LeftLessEqual,
		NotEqual.Name():                      NotEqual,
		Inconclusive.Name():                  Inconclusive,
		Equal.NumberString():                 Equal,
		LeftGreater.NumberString():           LeftGreater,
		LeftGreaterEqual.NumberString():      LeftGreaterEqual,
		LeftLess.NumberString():              LeftLess,
		LeftLessEqual.NumberString():         LeftLessEqual,
		NotEqual.NumberString():              NotEqual,
		Inconclusive.NumberString():          Inconclusive,
		Equal.NumberJsonString():             Equal,
		LeftGreater.NumberJsonString():       LeftGreater,
		LeftGreaterEqual.NumberJsonString():  LeftGreaterEqual,
		LeftLess.NumberJsonString():          LeftLess,
		LeftLessEqual.NumberJsonString():     LeftLessEqual,
		NotEqual.NumberJsonString():          NotEqual,
		Inconclusive.NumberJsonString():      Inconclusive,
		Equal.OperatorSymbol():               Equal,
		LeftGreater.OperatorSymbol():         LeftGreater,
		LeftGreaterEqual.OperatorSymbol():    LeftGreaterEqual,
		LeftLess.OperatorSymbol():            LeftLess,
		LeftLessEqual.OperatorSymbol():       LeftLessEqual,
		NotEqual.OperatorSymbol():            NotEqual,
		Inconclusive.OperatorSymbol():        Inconclusive,
		Equal.OperatorShortForm():            Equal,
		LeftGreater.OperatorShortForm():      LeftGreater,
		LeftGreaterEqual.OperatorShortForm(): LeftGreaterEqual,
		LeftLess.OperatorShortForm():         LeftLess,
		LeftLessEqual.OperatorShortForm():    LeftLessEqual,
		NotEqual.OperatorShortForm():         NotEqual,
		Inconclusive.OperatorShortForm():     Inconclusive,
	}
)
