package chmodhelper

import "gitlab.com/auk-go/core/constants"

// IsPartialMatchVariableAttr
//
//  @givenVarAttr can have wildcards "*"
//   On wildcard present comparison will ignore for that segment.
//
//  Example (will consider this a match):
//   - givenVarAttr: (rwx : "r*x"),
//   - rwx         : (rwx : "r-x")
func IsPartialMatchVariableAttr(
	givenVarAttr *VarAttribute,
	rwx string,
) bool {
	r, w, x := ExpandCharRwx(rwx)

	read := givenVarAttr.isRead.ToByteCondition(
		ReadChar,
		NopChar,
		constants.WildcardChar)
	write := givenVarAttr.isWrite.ToByteCondition(
		WriteChar,
		NopChar,
		constants.WildcardChar)
	execute := givenVarAttr.isExecute.ToByteCondition(
		ExecuteChar,
		NopChar,
		constants.WildcardChar,
	)

	isRead := givenVarAttr.isRead.IsWildcardOrBool(read == r)
	isWrite := givenVarAttr.isWrite.IsWildcardOrBool(write == w)
	isExecute := givenVarAttr.isExecute.IsWildcardOrBool(execute == x)

	return isRead &&
		isWrite &&
		isExecute
}
