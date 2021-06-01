package corestr

type NonChainedLinkedListNodes struct {
	items             *[]*LinkedListNode
	isChainingApplied bool
}

func (receiver *NonChainedLinkedListNodes) IsChainingApplied() bool {
	return receiver.isChainingApplied
}

func (receiver *NonChainedLinkedListNodes) Items() *[]*LinkedListNode {
	return receiver.items
}

func NewNonChainedLinkedListNodes(
	capacity int,
) *NonChainedLinkedListNodes {
	items := make([]*LinkedListNode, 0, capacity)

	return &NonChainedLinkedListNodes{
		items: &items,
	}
}

func (receiver *NonChainedLinkedListNodes) Length() int {
	if receiver.items == nil {
		return 0
	}

	return len(*receiver.items)
}

func (receiver *NonChainedLinkedListNodes) IsEmpty() bool {
	return receiver.items == nil || len(*receiver.items) == 0
}

func (receiver *NonChainedLinkedListNodes) Adds(
	nodes ...*LinkedListNode,
) *NonChainedLinkedListNodes {
	if nodes == nil {
		return receiver
	}

	for i := range nodes {
		*receiver.items = append(
			*receiver.items,
			nodes[i])
	}

	return receiver
}

func (receiver *NonChainedLinkedListNodes) HasItems() bool {
	return !receiver.IsEmpty()
}

func (receiver *NonChainedLinkedListNodes) First() *LinkedListNode {
	return (*receiver.items)[0]
}

func (receiver *NonChainedLinkedListNodes) FirstOrDefault() *LinkedListNode {
	if receiver.IsEmpty() {
		return nil
	}

	return (*receiver.items)[0]
}

func (receiver *NonChainedLinkedListNodes) Last() *LinkedListNode {
	return (*receiver.items)[receiver.Length()-1]
}

func (receiver *NonChainedLinkedListNodes) LastOrDefault() *LinkedListNode {
	if receiver.IsEmpty() {
		return nil
	}

	return (*receiver.items)[receiver.Length()-1]
}

// ApplyChaining Warning Mutates data inside.
func (receiver *NonChainedLinkedListNodes) ApplyChaining() *NonChainedLinkedListNodes {
	length := receiver.Length()
	if length == 0 {
		return receiver
	}

	receiver.isChainingApplied = true
	for i, node := range *receiver.items {
		if i+1 >= length {
			break
		}

		nextNode := (*receiver.items)[i+1]
		node.next = nextNode
	}

	if receiver.HasItems() {
		receiver.Last().next = nil
	}

	return receiver
}

func (receiver *NonChainedLinkedListNodes) ToChainedNodes() *[]*LinkedListNode {
	length := receiver.Length()
	list := make([]*LinkedListNode, length)

	if length == 0 {
		return &list
	}

	for i, node := range *receiver.items {
		if i+1 >= length {
			break
		}

		curNode := node.Clone()
		list = append(list, curNode)
		nextNode := (*receiver.items)[i+1]
		nextNodeClone := nextNode.Clone()
		curNode.next = nextNodeClone
		list = append(list, nextNodeClone)
	}

	return &list
}
