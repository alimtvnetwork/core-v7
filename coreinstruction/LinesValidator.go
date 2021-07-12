package coreinstruction

type LinesValidator struct {
	LineNumber                    int    `json:"LineNumber,omitempty"`
	LineText                      string `json:"LineText,omitempty"`
	IsTrimCompare                 bool
	IsCompareRegardlessWhitespace bool
}
