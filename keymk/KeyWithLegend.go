package keymk

import (
	"fmt"
	"strconv"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreinterface"
)

// KeyWithLegend
//
// Chain Sequence (Root-Package-Group-State-User-Item)
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

// OutputItemsArray
//
//  Depending on Options
//  -   IsIgnoreLegendAttachments() - calls OutputWithoutLegend()
//  -   or else - calls compiles using legends
//
//  Chain may look like:
//      - root-package-group-state-user-item
//
// Ordering :
//  - Root
//  - Package
//  - Group
//  - State
//  - User
//  - Item
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

func (it *KeyWithLegend) Group(group interface{}) string {
	request := KeyLegendCompileRequest{
		GroupId: fmt.Sprintf(
			constants.SprintValueFormat,
			group),
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupString(group string) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId:   group,
	}

	return it.CompileUsingRequest(request)
}

// Item
//
// It will include the existing group. chain (root-pkg-group-item)
func (it *KeyWithLegend) Item(item interface{}) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId:   it.groupName,
		ItemId:    fmt.Sprintf(constants.SprintValueFormat, item),
	}

	return it.CompileUsingRequest(request)
}

// ItemWithoutGroup
//
// Doesn't include existing group chain (root-pkg-item)
func (it *KeyWithLegend) ItemWithoutGroup(item interface{}) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		ItemId:    fmt.Sprintf(constants.SprintValueFormat, item),
	}

	return it.CompileUsingRequest(request)
}

// ItemWithoutStateGroup
//
// Doesn't include existing group chain (root-pkg-item)
func (it *KeyWithLegend) ItemWithoutStateGroup(item interface{}) string {
	request := KeyLegendCompileRequest{
		ItemId: fmt.Sprintf(constants.SprintValueFormat, item),
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) ItemEnumByte(item coreinterface.ByteEnumNamer) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId:   it.groupName,
		ItemId:    fmt.Sprintf(constants.SprintValueFormat, item),
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) ItemString(item string) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId:   it.groupName,
		ItemId:    item,
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) ItemInt(itemId int) string {
	return it.Item(itemId)
}

func (it *KeyWithLegend) ItemUInt(itemId uint) string {
	return it.Item(itemId)
}

func (it *KeyWithLegend) GroupItemIntRange(group string, startId, endId int) []string {
	ids := make([]string, 0, (endId-startId)+constants.Capacity3)

	for i := startId; i <= endId; i++ {
		ids = append(ids, it.GroupItemString(group, strconv.Itoa(i)))
	}

	return ids
}

func (it *KeyWithLegend) UptoGroup(user string) string {
	request := KeyLegendCompileRequest{
		GroupId: it.groupName,
		UserId:  user,
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) UptoState(user string) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		UserId:    user,
		GroupId:   it.groupName,
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupIntRange(startId, endId int) []string {
	ids := make([]string, 0, (endId-startId)+constants.Capacity3)

	for i := startId; i <= endId; i++ {
		ids = append(ids, it.GroupString(strconv.Itoa(i)))
	}

	return ids
}

func (it *KeyWithLegend) GroupUIntRange(startId, endId uint) []string {
	ids := make([]string, 0, (endId-startId)+constants.Capacity3)

	for i := startId; i <= endId; i++ {
		ids = append(ids, it.Group(i))
	}

	return ids
}

func (it *KeyWithLegend) ItemIntRange(startId, endId int) []string {
	ids := make([]string, 0, (endId-startId)+constants.Capacity3)

	for i := startId; i <= endId; i++ {
		ids = append(ids, it.Item(i))
	}

	return ids
}

func (it *KeyWithLegend) ItemUIntRange(startId, endId uint) []string {
	ids := make([]string, 0, (endId-startId)+constants.Capacity3)

	for i := startId; i <= endId; i++ {
		ids = append(ids, it.Item(i))
	}

	return ids
}

func (it *KeyWithLegend) GroupUserString(group, user string) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		UserId:    user,
		GroupId:   group,
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupUser(group, user interface{}) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		UserId:    fmt.Sprintf(constants.SprintValueFormat, user),
		GroupId:   fmt.Sprintf(constants.SprintValueFormat, group),
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupUInt(group uint) string {
	request := KeyLegendCompileRequest{
		GroupId: fmt.Sprintf(constants.SprintValueFormat, group),
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupByte(group byte) string {
	request := KeyLegendCompileRequest{
		GroupId: fmt.Sprintf(constants.SprintValueFormat, group),
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupUserByte(group, user byte) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId:   fmt.Sprintf(constants.SprintValueFormat, group),
		UserId:    fmt.Sprintf(constants.SprintValueFormat, user),
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupUserItem(group, user, item interface{}) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		UserId:    fmt.Sprintf(constants.SprintValueFormat, user),
		GroupId:   fmt.Sprintf(constants.SprintValueFormat, group),
		ItemId:    fmt.Sprintf(constants.SprintValueFormat, item),
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupStateUserItem(group, state, user, item interface{}) string {
	request := KeyLegendCompileRequest{
		StateName: fmt.Sprintf(constants.SprintValueFormat, state),
		UserId:    fmt.Sprintf(constants.SprintValueFormat, user),
		GroupId:   fmt.Sprintf(constants.SprintValueFormat, group),
		ItemId:    fmt.Sprintf(constants.SprintValueFormat, item),
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) StateUserItem(state, user, item interface{}) string {
	request := KeyLegendCompileRequest{
		StateName: fmt.Sprintf(constants.SprintValueFormat, state),
		UserId:    fmt.Sprintf(constants.SprintValueFormat, user),
		GroupId:   it.groupName,
		ItemId:    fmt.Sprintf(constants.SprintValueFormat, item),
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) StateUser(state, user interface{}) string {
	request := KeyLegendCompileRequest{
		StateName: fmt.Sprintf(constants.SprintValueFormat, state),
		UserId:    fmt.Sprintf(constants.SprintValueFormat, user),
		GroupId:   it.groupName,
		ItemId:    "",
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

func (it *KeyWithLegend) GroupUserItemString(group, user, item string) string {
	request := KeyLegendCompileRequest{
		UserId:    user,
		GroupId:   group,
		ItemId:    item,
		StateName: it.stateName,
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupUserItemUint(group, user, item uint) string {
	return it.GroupUserItem(user, group, item)
}

func (it *KeyWithLegend) GroupUserItemInt(group, user, item int) string {
	return it.GroupUserItem(user, group, item)
}

func (it *KeyWithLegend) GroupItem(group, item interface{}) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId:   fmt.Sprintf(constants.SprintValueFormat, group),
		ItemId:    fmt.Sprintf(constants.SprintValueFormat, item),
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) StateItem(stateName, item interface{}) string {
	request := KeyLegendCompileRequest{
		StateName: fmt.Sprintf(constants.SprintValueFormat, stateName),
		ItemId:    fmt.Sprintf(constants.SprintValueFormat, item),
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupItemString(group, item string) string {
	request := KeyLegendCompileRequest{
		StateName: it.stateName,
		GroupId:   group,
		ItemId:    item,
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) GroupStateItemString(group, stateName, item string) string {
	request := KeyLegendCompileRequest{
		GroupId:   group,
		ItemId:    item,
		StateName: stateName,
	}

	return it.CompileUsingRequest(request)
}

func (it *KeyWithLegend) StateItemString(stateName, item string) string {
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

func (it *KeyWithLegend) CompileUsingRequest(
	request KeyLegendCompileRequest,
) string {
	finalItems := it.FinalStrings(request)

	return strings.Join(finalItems, it.option.Joiner)
}

// FinalStrings
//
// Returns compiled array from conditions using OutputItemsArray
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

	return NewKeyWithLegend(
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
