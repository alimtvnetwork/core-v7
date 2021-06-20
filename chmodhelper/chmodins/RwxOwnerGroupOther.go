package chmodins

import "gitlab.com/evatix-go/core/constants"

// RwxOwnerGroupOther
//
// Owner, Group, Other:
// String Index Values
//  - 0: 'r'/'*'/'-'
//  - 1: 'w'/'*'/'-'
//  - 2: 'x'/'*'/'-'
// Examples can be :
//  - "rwx" or
//  - "*wx" or
//  - "rw*" or
//  - "***"
//
// Length must be 3. Not more not less.
type RwxOwnerGroupOther struct {
	// String Index Values
	//  - 0: 'r'/'*'/'-'
	//  - 1: 'w'/'*'/'-'
	//  - 2: 'x'/'*'/'-'
	// Examples can be :
	//  - "rwx" or
	//  - "*wx" or
	//  - "rw*" or
	//  - "***"
	//
	// Length must be 3. Not more not less.
	Owner string `json:"Owner"`
	// String Index Values
	//  - 0: 'r'/'*'/'-'
	//  - 1: 'w'/'*'/'-'
	//  - 2: 'x'/'*'/'-'
	// Examples can be :
	//  - "rwx" or
	//  - "*wx" or
	//  - "rw*" or
	//  - "***"
	//
	// Length must be 3. Not more not less.
	Group string `json:"Group"`
	// String Index Values
	//  - 0: 'r'/'*'/'-'
	//  - 1: 'w'/'*'/'-'
	//  - 2: 'x'/'*'/'-'
	// Examples can be :
	//  - "rwx" or
	//  - "*wx" or
	//  - "rw*" or
	//  - "***"
	//
	// Length must be 3. Not more not less.
	Other string `json:"Other"`
}

// NewRwxOwnerGroupOther
//
// Each arg ownerRwx, groupRwx, otherRwx should have
//
// Index Values
//  - 0: 'r'/'*'/'-'
//  - 1: 'w'/'*'/'-'
//  - 2: 'x'/'*'/'-'
// Examples can be :
//  - "rwx" or
//  - "*wx" or
//  - "rw*" or
//  - "***"
//
// Length must be 3. Not more not less.
func NewRwxOwnerGroupOther(
	ownerRwx,
	groupRwx,
	otherRwx string,
) *RwxOwnerGroupOther {
	return &RwxOwnerGroupOther{
		Owner: ownerRwx,
		Group: groupRwx,
		Other: otherRwx,
	}
}

func (receiver *RwxOwnerGroupOther) IsOwner(rwx string) bool {
	return receiver.Owner == rwx
}

func (receiver *RwxOwnerGroupOther) IsGroup(rwx string) bool {
	return receiver.Group == rwx
}

func (receiver *RwxOwnerGroupOther) IsOther(rwx string) bool {
	return receiver.Other == rwx
}

func (receiver *RwxOwnerGroupOther) ExpandCharOwner() (r, w, x byte) {
	return expandCharsRwx(receiver.Owner)
}

func (receiver *RwxOwnerGroupOther) ExpandCharGroup() (r, w, x byte) {
	return expandCharsRwx(receiver.Group)
}

func (receiver *RwxOwnerGroupOther) ExpandCharOther() (r, w, x byte) {
	return expandCharsRwx(receiver.Other)
}

func (receiver *RwxOwnerGroupOther) Is(
	ownerRwx,
	groupRwx,
	otherRwx string,
) bool {
	return receiver.IsOwner(ownerRwx) &&
		receiver.IsGroup(groupRwx) &&
		receiver.IsOther(otherRwx)
}

func (receiver *RwxOwnerGroupOther) IsEqual(another *RwxOwnerGroupOther) bool {
	if another == nil && receiver == nil {
		return true
	}

	if another == nil || receiver == nil {
		return false
	}

	return receiver.Owner == another.Owner &&
		receiver.Group == another.Group &&
		receiver.Other == another.Other
}

func (receiver *RwxOwnerGroupOther) ToString(isIncludeHyphen bool) string {
	if isIncludeHyphen {
		return receiver.String()
	}

	return receiver.Owner +
		receiver.Group +
		receiver.Other
}

// String : Includes hyphen in-front
// constants.Hyphen +
//		receiver.Owner +
//		receiver.Group +
//		receiver.Other
func (receiver *RwxOwnerGroupOther) String() string {
	return constants.Hyphen +
		receiver.Owner +
		receiver.Group +
		receiver.Other
}

func (receiver *RwxOwnerGroupOther) Clone() *RwxOwnerGroupOther {
	if receiver == nil {
		return nil
	}

	return &RwxOwnerGroupOther{
		Owner: receiver.Owner,
		Group: receiver.Group,
		Other: receiver.Other,
	}
}
