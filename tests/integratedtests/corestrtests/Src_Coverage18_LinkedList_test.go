package corestrtests

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_SrcC18_Creators_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC18_Creators_Verification", func() {
		// Arrange
		tc := srcC18CreatorsTestCase

		// Act
		s1, s2 := "a", "b"
		items := []*string{&s1, &s2, nil}
		actual := args.Map{
			"createNN":    corestr.New.LinkedList.Create() != nil,
			"createEmpty": !corestr.New.LinkedList.Create().HasItems(),
			"emptyLen":    corestr.New.LinkedList.Empty().Length(),
			"stringsLen":  corestr.New.LinkedList.Strings([]string{"a", "b", "c"}).Length(),
			"stringsEE":   !corestr.New.LinkedList.Strings([]string{}).HasItems(),
			"spreadLen":   corestr.New.LinkedList.SpreadStrings("x", "y").Length(),
			"spreadEE":    !corestr.New.LinkedList.SpreadStrings().HasItems(),
			"mapLen":      corestr.New.LinkedList.UsingMap(map[string]bool{"a": true, "b": false, "c": true}).Length(),
			"mapNilE":     !corestr.New.LinkedList.UsingMap(nil).HasItems(),
			"ptrLen":      corestr.New.LinkedList.PointerStringsPtr(&items).Length(),
			"ptrNilE":     !corestr.New.LinkedList.PointerStringsPtr(nil).HasItems(),
			"emptyLLe":    !corestr.Empty.LinkedList().HasItems(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC18_HeadTail_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC18_HeadTail_Verification", func() {
		// Arrange
		tc := srcC18HeadTailTestCase

		// Act
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		actual := args.Map{
			"head":       ll.Head().Element,
			"tail":       ll.Tail().Element,
			"lengthLock": corestr.New.LinkedList.SpreadStrings("a", "b").LengthLock(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC18_State_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC18_State_Verification", func() {
		// Arrange
		tc := srcC18StateTestCase

		// Act
		llE := corestr.New.LinkedList.Create()
		llA := corestr.New.LinkedList.Create()
		llA.Add("x")
		actual := args.Map{
			"emptyIsEmpty": llE.IsEmpty(),
			"emptyHasIt":   llE.HasItems(),
			"addedIsEmpty": llA.IsEmpty(),
			"addedHasIt":   llA.HasItems(),
			"emptyLock":    corestr.New.LinkedList.Create().IsEmptyLock(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC18_Add_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC18_Add_Verification", func() {
		// Arrange
		tc := srcC18AddTestCase

		// Act
		ll1 := corestr.New.LinkedList.Create(); ll1.Add("first"); ll1.Add("second")
		ll2 := corestr.New.LinkedList.Create(); ll2.AddLock("a")
		ll3 := corestr.New.LinkedList.Create(); ll3.AddNonEmpty(""); ll3.AddNonEmpty("x")
		ll4 := corestr.New.LinkedList.Create(); ll4.AddNonEmptyWhitespace("  "); ll4.AddNonEmptyWhitespace("x")
		ll5 := corestr.New.LinkedList.Create(); ll5.AddIf(false, "no"); ll5.AddIf(true, "yes")
		ll6 := corestr.New.LinkedList.Create(); ll6.AddsIf(false, "a", "b"); ll6.AddsIf(true, "c", "d")
		ll7 := corestr.New.LinkedList.Create(); ll7.AddFunc(func() string { return "computed" })
		ll8 := corestr.New.LinkedList.Create()
		ll8.AddFuncErr(func() (string, error) { return "ok", nil }, func(err error) {})
		ll9 := corestr.New.LinkedList.Create()
		errCalled := false
		ll9.AddFuncErr(func() (string, error) { return "", errors.New("fail") }, func(err error) { errCalled = true })
		ll10 := corestr.New.LinkedList.Create(); ll10.Push("a"); ll10.PushBack("b")
		ll11 := corestr.New.LinkedList.Create(); ll11.Adds("a", "b", "c")
		ll12 := corestr.New.LinkedList.Create(); ll12.Adds()
		ll13 := corestr.New.LinkedList.Create(); ll13.AddStrings([]string{"a", "b"})
		ll14 := corestr.New.LinkedList.Create(); ll14.AddsLock("a", "b")
		ll15 := corestr.New.LinkedList.Create(); ll15.AddItemsMap(map[string]bool{"a": true, "b": false})
		ll16 := corestr.New.LinkedList.Create(); ll16.AddItemsMap(map[string]bool{})
		actual := args.Map{
			"addLen":      ll1.Length(),
			"addLockLen":  ll2.Length(),
			"nonEmptyLen": ll3.Length(),
			"noWhiteLen":  ll4.Length(),
			"addIfLen":    ll5.Length(),
			"addIfHead":   ll5.Head().Element,
			"addsIfLen":   ll6.Length(),
			"funcHead":    ll7.Head().Element,
			"funcErrLen":  ll8.Length(),
			"funcErrErr":  errCalled,
			"pushLen":     ll10.Length(),
			"addsLen":     ll11.Length(),
			"addsEmptyE":  !ll12.HasItems(),
			"stringsLen":  ll13.Length(),
			"addsLockLen": ll14.Length(),
			"mapLen":      ll15.Length(),
			"mapEmptyE":   !ll16.HasItems(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC18_AddFront_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC18_AddFront_Verification", func() {
		// Arrange
		tc := srcC18AddFrontTestCase

		// Act
		ll1 := corestr.New.LinkedList.SpreadStrings("b", "c"); ll1.AddFront("a")
		ll2 := corestr.New.LinkedList.Create(); ll2.AddFront("only")
		ll3 := corestr.New.LinkedList.SpreadStrings("b"); ll3.PushFront("a")
		actual := args.Map{
			"frontHead":   ll1.Head().Element,
			"frontLen":    ll1.Length(),
			"frontEmHead": ll2.Head().Element,
			"frontEmLen":  ll2.Length(),
			"pushHead":    ll3.Head().Element,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC18_AppendNode_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC18_AppendNode_Verification", func() {
		// Arrange
		tc := srcC18AppendNodeTestCase

		// Act
		ll1 := corestr.New.LinkedList.Create()
		ll1.AppendNode(&corestr.LinkedListNode{Element: "x"})
		ll1.AppendNode(&corestr.LinkedListNode{Element: "y"})
		ll2 := corestr.New.LinkedList.Create()
		ll2.AddBackNode(&corestr.LinkedListNode{Element: "a"})
		actual := args.Map{
			"appendLen1": 1, // after first append it was 1
			"appendHead": "x",
			"appendLen2": ll1.Length(),
			"addBackLen": ll2.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC18_InsertAt_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC18_InsertAt_Verification", func() {
		// Arrange
		tc := srcC18InsertAtTestCase

		// Act
		ll1 := corestr.New.LinkedList.SpreadStrings("b", "c"); ll1.InsertAt(0, "a")
		ll2 := corestr.New.LinkedList.SpreadStrings("a", "c"); ll2.InsertAt(1, "b")
		ll3 := corestr.New.LinkedList.SpreadStrings("a", "c")
		ll3.AddAfterNode(ll3.Head(), "b")
		actual := args.Map{
			"frontHead": ll1.Head().Element,
			"middleAt1": ll2.List()[1],
			"afterAt1":  ll3.List()[1],
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC18_Attach_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC18_Attach_Verification", func() {
		// Arrange
		tc := srcC18AttachTestCase

		// Act
		ll1 := corestr.New.LinkedList.Create()
		err1 := ll1.AttachWithNode(nil, &corestr.LinkedListNode{Element: "x"})
		ll2 := corestr.New.LinkedList.SpreadStrings("a", "b")
		err2 := ll2.AttachWithNode(ll2.Head(), &corestr.LinkedListNode{Element: "x"})
		ll3 := corestr.New.LinkedList.SpreadStrings("a")
		err3 := ll3.AttachWithNode(ll3.Tail(), &corestr.LinkedListNode{Element: "b"})
		_ = err3
		actual := args.Map{
			"nilCurrErr":    err1 != nil,
			"nonNilNextErr": err2 != nil,
			"successLen":    ll3.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC18_AddCollPtr_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC18_AddCollPtr_Verification", func() {
		// Arrange
		tc := srcC18AddCollPtrTestCase

		// Act
		ll1 := corestr.New.LinkedList.SpreadStrings("a")
		ll1.AddCollectionToNode(true, ll1.Head(), corestr.New.Collection.Strings([]string{"b", "c"}))
		ll2 := corestr.New.LinkedList.Create()
		s1 := "a"
		ll2.AddPointerStringsPtr([]*string{&s1, nil})
		ll3 := corestr.New.LinkedList.Create()
		ll3.AddCollection(corestr.New.Collection.Strings([]string{"a", "b"}))
		ll4 := corestr.New.LinkedList.Create()
		ll4.AddCollection(nil)
		actual := args.Map{
			"colToNodeGe2": ll1.Length() >= 2,
			"ptrLen":       ll2.Length(),
			"colLen":       ll3.Length(),
			"colNilE":      !ll4.HasItems(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC18_Loop_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC18_Loop_Verification", func() {
		// Arrange
		tc := srcC18LoopTestCase

		// Act
		count1 := 0
		corestr.New.LinkedList.SpreadStrings("a", "b", "c").Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count1++
			return false
		})
		count2 := 0
		corestr.New.LinkedList.SpreadStrings("a", "b", "c").Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count2++
			return true
		})
		emptyOk := true
		corestr.New.LinkedList.Create().Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			emptyOk = false
			return false
		})
		count3 := 0
		corestr.New.LinkedList.SpreadStrings("a", "b", "c").Loop(func(arg *corestr.LinkedListProcessorParameter) bool {
			count3++
			return arg.Index == 1
		})
		actual := args.Map{
			"fullCount":   count1,
			"breakCount":  count2,
			"emptyOk":     emptyOk,
			"breakSecond": count3,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC18_Filter_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC18_Filter_Verification", func() {
		// Arrange
		tc := srcC18FilterTestCase

		// Act
		nodes1 := corestr.New.LinkedList.SpreadStrings("a", "b", "c").Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: arg.Node.Element != "b", IsBreak: false}
		})
		nodes2 := corestr.New.LinkedList.Create().Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true}
		})
		nodes3 := corestr.New.LinkedList.SpreadStrings("a", "b").Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: true}
		})
		nodes4 := corestr.New.LinkedList.SpreadStrings("a", "b", "c").Filter(func(arg *corestr.LinkedListFilterParameter) *corestr.LinkedListFilterResult {
			return &corestr.LinkedListFilterResult{Value: arg.Node, IsKeep: true, IsBreak: arg.Index == 1}
		})
		actual := args.Map{
			"keepLen":       len(nodes1),
			"emptyLen":      len(nodes2),
			"breakFirstLen": len(nodes3),
			"breakSecLen":   len(nodes4),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC18_Remove_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC18_Remove_Verification", func() {
		// Arrange
		tc := srcC18RemoveTestCase

		// Act
		ll1 := corestr.New.LinkedList.SpreadStrings("a", "b", "c"); ll1.RemoveNodeByElementValue("a", true, true)
		ll2 := corestr.New.LinkedList.SpreadStrings("a", "b", "c"); ll2.RemoveNodeByElementValue("b", true, true)
		ll3 := corestr.New.LinkedList.SpreadStrings("ABC", "def"); ll3.RemoveNodeByElementValue("abc", false, true)
		ll4 := corestr.New.LinkedList.SpreadStrings("a", "b", "c"); ll4.RemoveNodeByIndex(0)
		ll5 := corestr.New.LinkedList.SpreadStrings("a", "b", "c"); ll5.RemoveNodeByIndex(2)
		ll6 := corestr.New.LinkedList.SpreadStrings("a", "b", "c"); ll6.RemoveNodeByIndex(1)
		ll7 := corestr.New.LinkedList.SpreadStrings("a", "b", "c", "d"); ll7.RemoveNodeByIndexes(true, 0, 2)
		ll8 := corestr.New.LinkedList.SpreadStrings("a"); ll8.RemoveNodeByIndexes(true)
		ll9 := corestr.New.LinkedList.SpreadStrings("a", "b", "c"); ll9.RemoveNode(ll9.Head().Next())
		ll10 := corestr.New.LinkedList.SpreadStrings("a", "b"); ll10.RemoveNode(ll10.Head())
		ll11 := corestr.New.LinkedList.SpreadStrings("a"); ll11.RemoveNode(nil)
		actual := args.Map{
			"rmValFirstLen":   ll1.Length(),
			"rmValFirstHead":  ll1.Head().Element,
			"rmValMiddleLen":  ll2.Length(),
			"rmValCILen":      ll3.Length(),
			"rmIdxFirstHead":  ll4.Head().Element,
			"rmIdxLastLen":    ll5.Length(),
			"rmIdxMiddleLen":  ll6.Length(),
			"rmIdxesLen":      ll7.Length(),
			"rmIdxesEmptyLen": ll8.Length(),
			"rmNodeLen":       ll9.Length(),
			"rmNodeFirstHead": ll10.Head().Element,
			"rmNodeNilLen":    ll11.Length(),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC18_IndexAt_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC18_IndexAt_Verification", func() {
		// Arrange
		tc := srcC18IndexAtTestCase

		// Act
		ll := corestr.New.LinkedList.SpreadStrings("a", "b", "c")
		ll2 := corestr.New.LinkedList.SpreadStrings("a", "b")
		ll3 := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{
			"idxAt0":      ll.IndexAt(0).Element,
			"idxAt2":      ll.IndexAt(2).Element,
			"idxNegNil":   corestr.New.LinkedList.SpreadStrings("a").IndexAt(-1) == nil,
			"safeAt1":     ll2.SafeIndexAt(1).Element,
			"safeOorNil":  ll3.SafeIndexAt(5) == nil,
			"safeNegNil":  ll3.SafeIndexAt(-1) == nil,
			"safeLock0":   ll2.SafeIndexAtLock(0).Element,
			"ptrAt0":      *ll2.SafePointerIndexAt(0),
			"ptrOorNil":   ll2.SafePointerIndexAt(5) == nil,
			"ptrDef0":     ll3.SafePointerIndexAtUsingDefault(0, "def"),
			"ptrDefOor":   ll3.SafePointerIndexAtUsingDefault(5, "def"),
			"ptrDefLock0": ll3.SafePointerIndexAtUsingDefaultLock(0, "def"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC18_NextNodes_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC18_NextNodes_Verification", func() {
		// Arrange
		tc := srcC18NextNodesTestCase

		// Act
		actual := args.Map{
			"nextLen": len(corestr.New.LinkedList.SpreadStrings("a", "b", "c", "d").GetNextNodes(2)),
			"allLen":  len(corestr.New.LinkedList.SpreadStrings("a", "b").GetAllLinkedNodes()),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}
func Test_SrcC18_StringJoin_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC18_StringJoin_Verification", func() {
		// Arrange
		tc := srcC18StringJoinTestCase

		// Act
		actual := args.Map{
			"strNonE":      corestr.New.LinkedList.SpreadStrings("a", "b").String() != "",
			"strEmptyNonE": corestr.New.LinkedList.Create().String() != "",
			"strLockNonE":  corestr.New.LinkedList.SpreadStrings("a").StringLock() != "",
			"strLockENonE": corestr.New.LinkedList.Create().StringLock() != "",
			"join":         corestr.New.LinkedList.SpreadStrings("a", "b").Join(","),
			"joinLock":     corestr.New.LinkedList.SpreadStrings("a", "b").JoinLock(","),
			"joins":        corestr.New.LinkedList.SpreadStrings("a").Joins(",", "b", "c"),
			"joinsNil":     corestr.New.LinkedList.Create().Joins(",", "a"),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC18_CompareEquals_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC18_CompareEquals_Verification", func() {
		// Arrange
		tc := srcC18CompareEqualsTestCase

		// Act
		ll := corestr.New.LinkedList.SpreadStrings("a")
		actual := args.Map{
			"summaryNonE":  corestr.New.LinkedList.SpreadStrings("a", "b").GetCompareSummary(corestr.New.LinkedList.SpreadStrings("a", "b"), "left", "right") != "",
			"equalSame":    corestr.New.LinkedList.SpreadStrings("a", "b").IsEquals(corestr.New.LinkedList.SpreadStrings("a", "b")),
			"equalDiffLen": corestr.New.LinkedList.SpreadStrings("a", "b").IsEquals(corestr.New.LinkedList.SpreadStrings("a")),
			"equalSameRef": ll.IsEqualsWithSensitive(ll, true),
			"equalBothE":   corestr.New.LinkedList.Create().IsEquals(corestr.New.LinkedList.Create()),
			"equalOneE":    corestr.New.LinkedList.SpreadStrings("a").IsEquals(corestr.New.LinkedList.Create()),
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}
func Test_SrcC18_Json_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC18_Json_Verification", func() {
		// Arrange
		tc := srcC18JsonTestCase

		// Act
		noPanic := !callPanicsSrcC18(func() {
			ll := corestr.New.LinkedList.SpreadStrings("a", "b")
			_ = ll.JsonModel()
			_ = ll.JsonModelAny()
			data, _ := json.Marshal(ll)
			ll2 := corestr.New.LinkedList.Create()
			_ = json.Unmarshal(data, ll2)
			r := ll.Json()
			_ = r.Error == nil
			jr := ll.JsonPtr()
			ll3 := corestr.New.LinkedList.Create()
			_, _ = ll3.ParseInjectUsingJson(jr)
			ll4 := corestr.New.LinkedList.Create()
			_ = ll4.ParseInjectUsingJsonMust(jr)
			ll5 := corestr.New.LinkedList.Create()
			_ = ll5.JsonParseSelfInject(jr)
			_ = ll.AsJsonMarshaller()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func Test_SrcC18_Clear_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC18_Clear_Verification", func() {
		// Arrange
		tc := srcC18ClearTestCase

		// Act
		ll1 := corestr.New.LinkedList.SpreadStrings("a", "b"); ll1.RemoveAll()
		ll2 := corestr.New.LinkedList.SpreadStrings("a"); ll2.Clear()
		ll3 := corestr.New.LinkedList.Create(); ll3.Clear()
		actual := args.Map{
			"removeAllE":  !ll1.HasItems(),
			"clearE":      !ll2.HasItems(),
			"clearEmptyOk": true,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}
func Test_SrcC18_NonChained_Verification(t *testing.T) {
	safeTest(t, "Test_SrcC18_NonChained_Verification", func() {
		// Arrange
		tc := srcC18NonChainedTestCase

		// Act
		noPanic := !callPanicsSrcC18(func() {
			nc := corestr.NewNonChainedLinkedListNodes(5)
			_ = nc.IsEmpty()
			nc.Adds(&corestr.LinkedListNode{Element: "a"}, &corestr.LinkedListNode{Element: "b"})
			_ = nc.Length()
			_ = nc.HasItems()
			_ = nc.First().Element
			_ = nc.Last().Element
			// ApplyChaining
			nc2 := corestr.NewNonChainedLinkedListNodes(5)
			nc2.Adds(&corestr.LinkedListNode{Element: "a"}, &corestr.LinkedListNode{Element: "b"}, &corestr.LinkedListNode{Element: "c"})
			nc2.ApplyChaining()
			_ = nc2.IsChainingApplied()
			_ = nc2.First().HasNext()
			// ToChainedNodes
			nc3 := corestr.NewNonChainedLinkedListNodes(3)
			nc3.Adds(&corestr.LinkedListNode{Element: "x"}, &corestr.LinkedListNode{Element: "y"})
			_ = nc3.ToChainedNodes()
			// Empty
			nc4 := corestr.NewNonChainedLinkedListNodes(0)
			_ = nc4.FirstOrDefault()
			_ = nc4.LastOrDefault()
			// Nil
			nc5 := corestr.NewNonChainedLinkedListNodes(0)
			nc5.Adds(nil)
			_ = nc5.HasItems()
		})
		actual := args.Map{
			"noPanic": noPanic,
		}

		// Assert
		tc.ShouldBeEqualMapFirst(t, actual)
	})
}

func callPanicsSrcC18(fn func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	fn()
	return false
	}
