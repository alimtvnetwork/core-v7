package regconsts

const (
	RegExForEachWordsWithDollarSymbol       = "\\$(\\w+)(\\d*)"  // Selects a full word that starts with a "$" symbol
	EachWordsWithinPercentSymbol            = "\\%(\\w+)(\\d*)%" // Selects a full word that is within two "%" symbol
	NginxBlockStart                         = "\\{\\s*(\\#(\\s)*.*)"
	NginxBlockEnding                        = "\\}\\s*(\\#(\\s)*.*)"
	NginxEmptyBlockOnlyBraces               = "\\{\\}"
	NginxEmptyBlockWithSpaces               = "\\{\\s*\\}"
	NginxEmptyBlockWithCommentAndSpacesOnly = "\\{\\s*\\#+\\s*[aA0-zZ9]*\\s*\\n*\\}"
	// Reference : https://regexr.com/5jgh0
	NginxEmptyBlock = "(\\{\\}|\\{\\s*\\}|\\{\\s*\\#+\\s*[aA0-zZ9]*.*[aA0-zZ9]*\\.*\\s*\\n*\\})"
)
