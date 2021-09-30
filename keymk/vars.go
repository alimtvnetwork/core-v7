package keymk

import "gitlab.com/evatix-go/core/constants"

var (
	PipeJoinerOption = &Option{
		Joiner:           DefaultJoiner,
		IsSkipEmptyEntry: true,
		IsUseBrackets:    true,
		StartBracket:     constants.Pipe,
		EndBracket:       constants.Pipe,
	}

	PipeCurlyJoinerOption = &Option{
		Joiner:           DefaultJoiner,
		IsSkipEmptyEntry: true,
		IsUseBrackets:    true,
		StartBracket:     constants.Pipe + constants.CurlyStart,
		EndBracket:       constants.CurlyEnd + constants.Pipe,
	}

	PipeSquareJoinerOption = &Option{
		Joiner:           DefaultJoiner,
		IsSkipEmptyEntry: true,
		IsUseBrackets:    true,
		StartBracket:     constants.Pipe + constants.SquareStart,
		EndBracket:       constants.SquareEnd + constants.Pipe,
	}

	CurlyBraceJoinerOption = &Option{
		Joiner:           DefaultJoiner,
		IsSkipEmptyEntry: true,
		IsUseBrackets:    true,
		StartBracket:     constants.CurlyStart,
		EndBracket:       constants.CurlyEnd,
	}

	BracketJoinerOption = &Option{
		Joiner:           DefaultJoiner,
		IsSkipEmptyEntry: true,
		IsUseBrackets:    true,
		StartBracket:     constants.SquareStart,
		EndBracket:       constants.SquareEnd,
	}

	ParentThesisJoinerOption = &Option{
		Joiner:           DefaultJoiner,
		IsSkipEmptyEntry: true,
		IsUseBrackets:    true,
		StartBracket:     constants.ParenthesisStart,
		EndBracket:       constants.ParenthesisEnd,
	}

	JoinerOption = &Option{
		Joiner:           constants.Hyphen,
		IsSkipEmptyEntry: true,
		IsUseBrackets:    false,
		StartBracket:     constants.EmptyString,
		EndBracket:       constants.EmptyString,
	}

	FullLegends = LegendName{
		Root:    "root",
		Package: "package",
		Group:   "group",
		User:    "user",
		Item:    "item",
	}

	FullCategoryLegends = LegendName{
		Root:    "root",
		Package: "package",
		Group:   "category",
		User:    "user",
		Item:    "item",
	}

	FullEventLegends = LegendName{
		Root:    "root",
		Package: "package",
		Group:   "event",
		User:    "user",
		Item:    "item",
	}

	ShortLegends = LegendName{
		Root:    "r",
		Package: "p",
		Group:   "g",
		User:    "u",
		Item:    "i",
	}

	ShortEventLegends = LegendName{
		Root:    "r",
		Group:   "e",
		Package: "p",
		Item:    "i",
	}
)
