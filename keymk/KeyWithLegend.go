package keymk

import (
	"fmt"
	"strconv"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coreinterface/enuminf"
)

// KeyWithLegend
//
// Chain Sequence
//
//	(Root-Package-Group-State-User-item) aka. LegendChainSample
//	fixed chain (Root-Package) and variable chain (Group-State-User-item)
//
//	Chain example LegendChainSample
//
//	Depending on Options
//	-   IsIgnoreLegendAttachments() calls or invokes -> OutputWithoutLegend()
//	-   or else - calls compiles using legends
//
//	Chain may look like (Fixed chain "{root}-{package}"):
//	 - root-package-group-state-user-item (LegendChainSample)
//	 - Fixed chain "{root}-{package}" -- rest will depend on value given on parameter.
//	    - Given request to item will only print
//	        - "{root}-{package}-{item}"
//	    - Given request to state, item will only print
//	        - "{root}-{package}-{state}-{item}"
//	    - Given request to group will only print
//	        - "{root}-{package}-{group}"
//
// Ordering :
//   - Root
//   - Package
//   - Group
//   - State
//   - User
//   - ItemWithoutUser
//
// Example:
//   - On any value empty in request will be
//     ignored if Option.IsSkipEmptyEntry
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
type KeyWithLegend struct {
	option                *Option
	LegendName            LegendName
	isAttachLegendNames   bool
	rootName, packageName string
	stateName, groupName  string
}

func (it *KeyWithLegend) IsIgnoreLegendAttachments() bool {
	return !it.isAttachLegendNames
}

func (it *KeyWithLegend) RootName() string {
	return it.rootName
}

func (it *KeyWithLegend) PackageName() string {
	return it.packageName
}

func (it *KeyWithLegend) GroupName() string {
	return it.groupName
}

func (it *KeyWithLegend) StateName() string {
	return it.stateName
}

// OutputItemsArray
//
//	Chain example LegendChainSample
//
//	Depending on Options
//	-   IsIgnoreLegendAttachments() calls or invokes -> OutputWithoutLegend()
//	-   or else - calls compiles using legends
//
//	Chain may look like:
//	    - root-package-group-state-user-item (LegendChainSample)
//
// Ordering :
//   - Root
//   - Package
//   - Group
//   - State
//   - User
//   - ItemWithoutUser
//
// Example:
//   - On any value empty in request will be
//     ignored if Option.IsSkipEmptyEntry
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) OutputItemsArray(request KeyLegendCompileRequest) []string {
	if it.IsIgnoreLegendAttachments() {
		return it.OutputWithoutLegend(request)
	}

	slice := make(
		[]string,
		0,
		constants.ArbitraryCapacity14)

	isAddRegardless := it.
		option.
		IsAddEntryRegardlessOfEmptiness()

	slice = it.appendLegendNameValue(
		isAddRegardless,
		slice,
		it.LegendName.Root,
		it.rootName)

	slice = it.appendLegendNameValue(
		isAddRegardless,
		slice,
		it.LegendName.Package,
		it.packageName)

	slice = it.appendLegendNameValue(
		isAddRegardless,
		slice,
		it.LegendName.Group,
		request.GroupId)

	slice = it.appendLegendNameValue(
		isAddRegardless,
		slice,
		it.LegendName.State,
		request.StateName)

	slice = it.appendLegendNameValue(
		isAddRegardless,
		slice,
		it.LegendName.User,
		request.UserId)

	slice = it.appendLegendNameValue(
		isAddRegardless,
		slice,
		it.LegendName.Item,
		request.ItemId)

	return slice
}

func (it *KeyWithLegend) appendLegendNameValue(
	isAddRegardless bool,
	list []string,
	legendName,
	valueId string,
) []string {
	if isAddRegardless || valueId != "" {
		return append(
			list,
			legendName,
			valueId)
	}

	return list
}

// Group
//
//	Returns up to state.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	State will be used from creation.
//
// Example:
//   - "{root}-{package}-{group}-{state}"
//
// Missing:
//   - "{user}-{item}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) Group(group interface{}) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId: fmt.Sprintf(
			constants.SprintValueFormat,
			group),
	}

	return it.CompileUsingRequest(request)
}

// GroupString
//
//	Returns up to state.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	State will be used from creation.
//
// Example:
//   - "{root}-{package}-{group}-{state}"
//
// Missing:
//   - "{user}-{item}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) GroupString(group string) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId:   group,
	}

	return it.CompileUsingRequest(request)
}

// UpToGroup
//
//	Returns up to group without states
//	Chain sample KeyWithLegend, LegendChainSample
//
//	State will be used from creation.
//
// Example:
//   - "{root}-{package}-{group}"
//
// Missing:
//   - "{state}-{user}-{item}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) UpToGroup(group interface{}) string {
	request := KeyLegendCompileRequest{
		GroupId: fmt.Sprintf(constants.SprintValueFormat, group),
	}

	return it.CompileUsingRequest(request)
}

// UpToGroupString
//
//	Returns up to group without states
//	Chain sample KeyWithLegend, LegendChainSample
//
//	State will be used from creation.
//
// Example:
//   - "{root}-{package}-{group}"
//
// Missing:
//   - "{state}-{user}-{item}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) UpToGroupString(group string) string {
	request := KeyLegendCompileRequest{
		GroupId: group,
	}

	return it.CompileUsingRequest(request)
}

// ItemWithoutUser
//
//	Returns up to item without user.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	State, Group will be used from creation.
//
// Example:
//   - "{root}-{package}-{state}-{group}-{item}"
//
// Missing:
//   - "{user}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) ItemWithoutUser(item interface{}) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId:   it.groupName,
		ItemId:    fmt.Sprintf(constants.SprintValueFormat, item),
	}

	return it.CompileUsingRequest(request)
}

// ItemWithoutUserGroup
//
//	Returns up to item without user, group.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	State will be used from creation.
//
// Example:
//   - "{root}-{package}-{state}-{item}"
//
// Missing:
//   - "{group}-{user}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) ItemWithoutUserGroup(item interface{}) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		ItemId:    fmt.Sprintf(constants.SprintValueFormat, item),
	}

	return it.CompileUsingRequest(request)
}

// ItemWithoutUserStateGroup
//
//	Returns up to item without user, group, state.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	Nothing will be used from creation.
//
// Example:
//   - "{root}-{package}-{item}"
//
// Missing:
//   - "{group}-{state}-{user}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) ItemWithoutUserStateGroup(item interface{}) string {
	request := KeyLegendCompileRequest{
		ItemId: fmt.Sprintf(constants.SprintValueFormat, item),
	}

	return it.CompileUsingRequest(request)
}

// ItemEnumByte
//
//	Returns up to item without user.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	Group, State will be used from creation.
//
// Example:
//   - "{root}-{package}-{state}-{group}-{item}"
//
// Missing:
//   - "{user}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) ItemEnumByte(item enuminf.ByteEnumNamer) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId:   it.groupName,
		ItemId: fmt.Sprintf(
			constants.SprintValueFormat,
			item),
	}

	return it.CompileUsingRequest(request)
}

// Item
//
//	Returns up to item without user.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	State, Group will be used from creation.
//
// Example:
//   - "{root}-{package}-{state}-{group}-{item}"
//
// Missing:
//   - "{user}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) Item(item interface{}) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId:   it.groupName,
		ItemId: fmt.Sprintf(
			constants.SprintValueFormat,
			item),
	}

	return it.CompileUsingRequest(request)
}

// ItemString
//
//	Returns up to item without user, group, state.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	State, Group will be used from creation.
//
// Example:
//   - "{root}-{package}-{state}-{group}-{item}"
//
// Missing:
//   - "{user}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) ItemString(item string) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId:   it.groupName,
		ItemId:    item,
	}

	return it.CompileUsingRequest(request)
}

// ItemInt
//
//	Returns up to item without user.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	State, Group will be used from creation.
//
// Example:
//   - "{root}-{package}-{state}-{group}-{item}"
//
// Missing:
//   - "{user}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) ItemInt(itemId int) string {
	return it.Item(itemId)
}

// ItemUInt
//
//	Returns up to item without user.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	State, Group will be used from creation.
//
// Example:
//   - "{root}-{package}-{state}-{group}-{item}"
//
// Missing:
//   - "{user}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) ItemUInt(itemId uint) string {
	return it.Item(itemId)
}

// GroupItemIntRange
//
//	Returns up to item without user.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	State, Group will be used from creation.
//
// Example:
//   - "{root}-{package}-{state}-{group}-{item}"
//
// Missing:
//   - "{user}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) GroupItemIntRange(group string, startId, endId int) []string {
	ids := make([]string, 0, (endId-startId)+constants.Capacity3)

	for i := startId; i <= endId; i++ {
		ids = append(ids, it.GroupItemString(group, strconv.Itoa(i)))
	}

	return ids
}

// UserStringWithoutState
//
//	Returns up to user without state.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	Group will be used from creation.
//
// Example:
//   - "{root}-{package}-{group}--{user}"
//
// Missing:
//   - "{state}-{item}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) UserStringWithoutState(user string) string {
	request := KeyLegendCompileRequest{
		GroupId: it.groupName,
		UserId:  user,
	}

	return it.CompileUsingRequest(request)
}

// UpToState
//
//	Returns up to state.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	State, Group will be used from creation.
//
// Example:
//   - "{root}-{package}-{group}-{state}"
//
// Missing:
//   - "{user}-{item}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) UpToState(
	user string,
) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		UserId:    user,
		GroupId:   it.groupName,
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupIntRange(
	startId, endId int,
) []string {
	ids := make([]string, 0, (endId-startId)+constants.Capacity3)

	for i := startId; i <= endId; i++ {
		ids = append(ids, it.GroupString(strconv.Itoa(i)))
	}

	return ids
}

func (it *KeyWithLegend) GroupUIntRange(
	startId, endId uint,
) []string {
	ids := make([]string, 0, (endId-startId)+constants.Capacity3)

	for i := startId; i <= endId; i++ {
		ids = append(ids, it.Group(i))
	}

	return ids
}

func (it *KeyWithLegend) ItemIntRange(
	startId, endId int,
) []string {
	ids := make([]string, 0, (endId-startId)+constants.Capacity3)

	for i := startId; i <= endId; i++ {
		ids = append(ids, it.ItemInt(i))
	}

	return ids
}

func (it *KeyWithLegend) ItemUIntRange(
	startId, endId uint,
) []string {
	ids := make([]string, 0, (endId-startId)+constants.Capacity3)

	for i := startId; i <= endId; i++ {
		ids = append(ids, it.ItemInt(int(i)))
	}

	return ids
}

func (it *KeyWithLegend) GroupUserString(
	group, user string,
) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		UserId:    user,
		GroupId:   group,
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupUser(
	group, user interface{},
) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		UserId:    fmt.Sprintf(constants.SprintValueFormat, user),
		GroupId:   fmt.Sprintf(constants.SprintValueFormat, group),
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupUInt(
	group uint,
) string {
	request := KeyLegendCompileRequest{
		GroupId: fmt.Sprintf(constants.SprintValueFormat, group),
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupByte(
	group byte,
) string {
	request := KeyLegendCompileRequest{
		GroupId: fmt.Sprintf(constants.SprintValueFormat, group),
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupUserByte(
	group, user byte,
) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId:   fmt.Sprintf(constants.SprintValueFormat, group),
		UserId:    fmt.Sprintf(constants.SprintValueFormat, user),
	}

	return it.CompileUsingRequest(request)
}

// GroupUserItem
//
//	Returns up to item.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	State will be used from creation.
//
// Example:
//   - "{root}-{package}-{group}-{state}-{user}-{item}"
//
// Missing:
//   - Nothing
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) GroupUserItem(
	group, user, item interface{},
) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		UserId:    fmt.Sprintf(constants.SprintValueFormat, user),
		GroupId:   fmt.Sprintf(constants.SprintValueFormat, group),
		ItemId:    fmt.Sprintf(constants.SprintValueFormat, item),
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupStateUserItem(
	group, state, user, item interface{},
) string {
	request := KeyLegendCompileRequest{
		StateName: fmt.Sprintf(constants.SprintValueFormat, state),
		UserId:    fmt.Sprintf(constants.SprintValueFormat, user),
		GroupId:   fmt.Sprintf(constants.SprintValueFormat, group),
		ItemId:    fmt.Sprintf(constants.SprintValueFormat, item),
	}

	return it.CompileUsingRequest(request)
}

// StateUserItem
//
//	Returns up to item.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	Group will be used from creation.
//
// Example:
//   - "{root}-{package}-{group}-{state}-{user}-{item}"
//
// Missing:
//   - Nothing
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) StateUserItem(
	state, user, item interface{},
) string {
	request := KeyLegendCompileRequest{
		StateName: fmt.Sprintf(constants.SprintValueFormat, state),
		UserId:    fmt.Sprintf(constants.SprintValueFormat, user),
		GroupId:   it.groupName,
		ItemId:    fmt.Sprintf(constants.SprintValueFormat, item),
	}

	return it.CompileUsingRequest(request)
}

// StateUser
//
//	Returns up to user.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	Group will be used from creation.
//
// Example:
//   - "{root}-{package}-{group}-{state}-{user}"
//
// Missing:
//   - "{item}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) StateUser(
	state, user interface{},
) string {
	request := KeyLegendCompileRequest{
		StateName: fmt.Sprintf(constants.SprintValueFormat, state),
		UserId:    fmt.Sprintf(constants.SprintValueFormat, user),
		GroupId:   it.groupName,
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupStateUserItemString(
	group, state, user, item string,
) string {
	request := KeyLegendCompileRequest{
		UserId:    user,
		GroupId:   group,
		ItemId:    item,
		StateName: state,
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupUserItemString(
	group, user, item string,
) string {
	request := KeyLegendCompileRequest{
		UserId:    user,
		GroupId:   group,
		ItemId:    item,
		StateName: it.stateName,
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupUserItemUint(
	group, user, item uint,
) string {
	return it.GroupUserItem(user, group, item)
}

func (it *KeyWithLegend) GroupUserItemInt(
	group, user, item int,
) string {
	return it.GroupUserItem(user, group, item)
}

func (it *KeyWithLegend) GroupItem(
	group, item interface{},
) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId:   fmt.Sprintf(constants.SprintValueFormat, group),
		ItemId:    fmt.Sprintf(constants.SprintValueFormat, item),
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) StateItem(
	stateName, item interface{},
) string {
	request := KeyLegendCompileRequest{
		StateName: fmt.Sprintf(constants.SprintValueFormat, stateName),
		ItemId:    fmt.Sprintf(constants.SprintValueFormat, item),
	}

	return it.CompileUsingRequest(request)
}

// GroupItemString
//
//	Returns up to item without user.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	State, Group will be used from creation.
//
// Example:
//   - "{root}-{package}-{state}-{group}-{item}"
//
// Missing:
//   - "{user}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) GroupItemString(
	group, item string,
) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId:   group,
		ItemId:    item,
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupStateItemString(
	group, stateName, item string,
) string {
	request := KeyLegendCompileRequest{
		GroupId:   group,
		ItemId:    item,
		StateName: stateName,
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) StateItemString(
	stateName, item string,
) string {
	request := KeyLegendCompileRequest{
		GroupId:   it.groupName,
		ItemId:    item,
		StateName: stateName,
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) Compile(
	itemId string,
) string {
	return it.ItemString(itemId)
}

// CompileDefault
//
//	Returns up to item without user, group, state.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	Group, State will be used from creation.
//
// Example:
//   - "{root}-{package}-{state}-{group}"
//
// Missing:
//   - "{user}-{item}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) CompileDefault() string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId:   it.groupName,
	}

	return it.CompileUsingRequest(request)
}

// CompileUsingJoiner
//
//	Returns up to item without user, group, state.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	State, Group will be used from creation.
//
// Example:
//   - "{root}-{package}-{state}-{group}"
//
// Missing:
//   - "{user}-{item}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) CompileUsingJoiner(
	joiner string,
) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId:   it.groupName,
	}

	finalItems := it.FinalStrings(request)

	return strings.Join(finalItems, joiner)
}

// CompileStrings
//
//	Returns up to item without user, group, state.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	State, Group will be used from creation.
//
// Example:
//   - "{root}-{package}-{state}-{group}"
//
// Missing:
//   - "{user}-{item}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) CompileStrings() []string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId:   it.groupName,
	}

	return it.FinalStrings(request)
}

// Strings
//
//	Returns up to item without user, group, state.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	State, Group will be used from creation.
//
// Example:
//   - "{root}-{package}-{state}-{group}"
//
// Missing:
//   - "{user}-{item}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) Strings() []string {
	return it.CompileStrings()
}

// CompileItemUsingJoiner
//
//	Returns up to item without user, group, state.
//	Chain sample KeyWithLegend, LegendChainSample
//
//	State, Group will be used from creation.
//
// Example:
//   - "{root}-{package}-{state}-{group}"
//
// Missing:
//   - "{user}-{item}"
//
// Conditions apply:
//   - if Option.IsSkipEmptyEntry then empty input wil be ignored.
func (it *KeyWithLegend) CompileItemUsingJoiner(
	joiner, item string,
) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId:   it.groupName,
		ItemId:    item,
	}

	finalItems := it.FinalStrings(request)

	return strings.Join(finalItems, joiner)
}

// CompileUsingRequest
//
// Compiles using FinalStrings
func (it *KeyWithLegend) CompileUsingRequest(
	request KeyLegendCompileRequest,
) string {
	finalItems := it.FinalStrings(request)

	return strings.Join(finalItems, it.option.Joiner)
}

// FinalStrings
//
//	Returns compiled array from
//	conditions using OutputItemsArray
//
// Conditions:
//   - When request given
func (it *KeyWithLegend) FinalStrings(
	request KeyLegendCompileRequest,
) []string {
	array := it.OutputItemsArray(request)

	if it.option.IsUseBrackets {
		return it.addBrackets(array)
	}

	return array
}

func (it *KeyWithLegend) addBrackets(inputItems []string) []string {
	for i, item := range inputItems {
		inputItems[i] = it.option.StartBracket + item + it.option.EndBracket
	}

	return inputItems
}

func (it *KeyWithLegend) OutputWithoutLegend(request KeyLegendCompileRequest) []string {
	slice := make([]string, 0, constants.Capacity6)

	slice = append(slice, it.rootName)
	slice = append(slice, it.packageName)

	isAddRegardless := it.
		option.
		IsAddEntryRegardlessOfEmptiness()

	if isAddRegardless || request.GroupId != "" {
		slice = append(slice, request.GroupId)
	}

	if isAddRegardless || request.StateName != "" {
		slice = append(slice, request.StateName)
	}

	if isAddRegardless || request.UserId != "" {
		slice = append(slice, request.UserId)
	}

	if isAddRegardless || request.ItemId != "" {
		slice = append(slice, request.ItemId)
	}

	return slice
}

func (it *KeyWithLegend) CloneUsing(groupName string) *KeyWithLegend {
	if it == nil {
		return nil
	}

	return NewKeyWithLegend.All(
		it.option.ClonePtr(),
		it.LegendName,
		it.isAttachLegendNames,
		it.rootName,
		it.packageName,
		groupName,
		it.stateName)
}

func (it *KeyWithLegend) Clone() *KeyWithLegend {
	return it.CloneUsing(it.groupName)
}
