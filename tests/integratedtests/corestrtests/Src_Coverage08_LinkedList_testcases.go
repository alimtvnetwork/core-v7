package corestrtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// Branch: LinkedList basic empty state
var srcC08LinkedListBasicTestCase = coretestcases.CaseV1{
	Title: "LinkedList Create returns correct state -- new empty list",
	ExpectedInput: args.Map{
		"isEmpty":      true,
		"hasItems":     false,
		"length":       0,
		"isEmptyLock":  true,
	},
}

// Branch: Add methods
var srcC08LinkedListAddTestCase = coretestcases.CaseV1{
	Title: "LinkedList Add variants return correct state -- multiple adds",
	ExpectedInput: args.Map{
		"length":      3,
		"headElement": "a",
		"tailElement": "c",
		"afterFront":  "z",
		"noPanic":     true,
	},
}

// Branch: AddStrings/AddCollection
var srcC08LinkedListAddsTestCase = coretestcases.CaseV1{
	Title: "LinkedList Adds AddStrings AddCollection execute without panic -- various",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// Branch: List methods
var srcC08LinkedListListTestCase = coretestcases.CaseV1{
	Title: "LinkedList List methods return correct -- 3 items",
	ExpectedInput: args.Map{
		"listLen": 3,
		"noPanic": true,
	},
}

// Branch: ToCollection
var srcC08LinkedListToCollectionTestCase = coretestcases.CaseV1{
	Title: "LinkedList ToCollection returns correct lengths -- items and empty",
	ExpectedInput: args.Map{
		"colLen":      2,
		"emptyColLen": 0,
	},
}

// Branch: Loop
var srcC08LinkedListLoopTestCase = coretestcases.CaseV1{
	Title: "LinkedList Loop iterates all items -- 3 items",
	ExpectedInput: args.Map{
		"count": 3,
	},
}

// Branch: Filter
var srcC08LinkedListFilterTestCase = coretestcases.CaseV1{
	Title: "LinkedList Filter returns all nodes -- keep all",
	ExpectedInput: args.Map{
		"nodesLen": 3,
	},
}

// Branch: SafeIndexAt
var srcC08LinkedListIndexAtTestCase = coretestcases.CaseV1{
	Title: "LinkedList SafeIndexAt returns correct -- index 1 of 3",
	ExpectedInput: args.Map{
		"element":     "b",
		"negOneIsNil": true,
		"outIsNil":    true,
		"noPanic":     true,
	},
}

// Branch: GetNextNodes / GetAllLinkedNodes
var srcC08LinkedListNextNodesTestCase = coretestcases.CaseV1{
	Title: "LinkedList GetNextNodes GetAllLinkedNodes return correct -- 3 items",
	ExpectedInput: args.Map{
		"nextNodesLen": 2,
		"allNodesLen":  3,
	},
}

// Branch: IsEquals
var srcC08LinkedListEqualsTestCase = coretestcases.CaseV1{
	Title: "LinkedList IsEquals IsEqualsWithSensitive return true -- same lists",
	ExpectedInput: args.Map{
		"isEquals":    true,
		"isSensitive": true,
	},
}

// Branch: RemoveNodeByIndex
var srcC08LinkedListRemoveTestCase = coretestcases.CaseV1{
	Title: "LinkedList RemoveNodeByIndex returns correct length -- remove first",
	ExpectedInput: args.Map{
		"length": 2,
	},
}

// Branch: Clear/RemoveAll
var srcC08LinkedListClearTestCase = coretestcases.CaseV1{
	Title: "LinkedList Clear RemoveAll return empty -- 2 items",
	ExpectedInput: args.Map{
		"length": 0,
	},
}

// Branch: Json/Marshal
var srcC08LinkedListJsonTestCase = coretestcases.CaseV1{
	Title: "LinkedList Json Marshal methods execute without panic -- 1 item",
	ExpectedInput: args.Map{
		"noPanic": true,
	},
}

// Branch: Joins
var srcC08LinkedListJoinsTestCase = coretestcases.CaseV1{
	Title: "LinkedList Joins returns non-empty -- 2 items plus extra",
	ExpectedInput: args.Map{
		"nonEmpty": true,
	},
}

// Branch: AppendNode
var srcC08LinkedListAppendNodeTestCase = coretestcases.CaseV1{
	Title: "LinkedList AppendNode returns correct length -- 2 nodes",
	ExpectedInput: args.Map{
		"length": 2,
	},
}

// Branch: AppendChainOfNodes
var srcC08LinkedListAppendChainTestCase = coretestcases.CaseV1{
	Title: "LinkedList AppendChainOfNodes returns correct length -- chain of 2",
	ExpectedInput: args.Map{
		"length": 2,
	},
}

// Branch: LinkedListNode methods
var srcC08LinkedListNodeTestCase = coretestcases.CaseV1{
	Title: "LinkedListNode methods return correct -- element a",
	ExpectedInput: args.Map{
		"hasNext":           false,
		"string":            "a",
		"isEqualValue":      true,
		"isSensitiveTrue":   true,
		"isSensitiveFalse":  true,
		"cloneElement":      "a",
		"noPanic":           true,
	},
}

// Branch: EndOfChain
var srcC08LinkedListEndOfChainTestCase = coretestcases.CaseV1{
	Title: "LinkedListNode EndOfChain returns correct -- chain of 2",
	ExpectedInput: args.Map{
		"endElement": "b",
		"chainLen":   2,
	},
}

// Branch: Node IsEqual
var srcC08LinkedListNodeEqualTestCase = coretestcases.CaseV1{
	Title: "LinkedListNode IsEqual IsChainEqual IsEqualSensitive return true -- same",
	ExpectedInput: args.Map{
		"isEqual":          true,
		"isChainEqual":     true,
		"isEqualSensitive": true,
	},
}
