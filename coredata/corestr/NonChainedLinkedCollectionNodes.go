package corestr

type NonChainedLinkedCollectionNodes struct {
	items             *[]*LinkedCollectionNode
	isChainingApplied bool
}

func NewNonChainedLinkedCollectionNodes(
	capacity int,
) *NonChainedLinkedCollectionNodes {
	items := make([]*LinkedCollectionNode, 0, capacity)

	return &NonChainedLinkedCollectionNodes{
		items: &items,
	}
}

func (receiver *NonChainedLinkedCollectionNodes) IsChainingApplied() bool {
	return receiver.isChainingApplied
}

func (receiver *NonChainedLinkedCollectionNodes) Items() *[]*LinkedCollectionNode {
	return receiver.items
}

func (receiver *NonChainedLinkedCollectionNodes) Length() int {
	if receiver.items == nil {
		return 0
	}

	return len(*receiver.items)
}

func (receiver *NonChainedLinkedCollectionNodes) IsEmpty() bool {
	return receiver.items == nil || len(*receiver.items) == 0
}

func (receiver *NonChainedLinkedCollectionNodes) Adds(
	nodes ...*LinkedCollectionNode,
) *NonChainedLinkedCollectionNodes {
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

func (receiver *NonChainedLinkedCollectionNodes) HasItems() bool {
	return !receiver.IsEmpty()
}

func (receiver *NonChainedLinkedCollectionNodes) First() *LinkedCollectionNode {
	return (*receiver.items)[0]
}

func (receiver *NonChainedLinkedCollectionNodes) FirstOrDefault() *LinkedCollectionNode {
	if receiver.IsEmpty() {
		return nil
	}

	return (*receiver.items)[0]
}

func (receiver *NonChainedLinkedCollectionNodes) Last() *LinkedCollectionNode {
	return (*receiver.items)[receiver.Length()-1]
}

func (receiver *NonChainedLinkedCollectionNodes) LastOrDefault() *LinkedCollectionNode {
	if receiver.IsEmpty() {
		return nil
	}

	return (*receiver.items)[receiver.Length()-1]
}

// ApplyChaining Warning Mutates data inside.
func (receiver *NonChainedLinkedCollectionNodes) ApplyChaining() *NonChainedLinkedCollectionNodes {
	length := receiver.Length()
	if length == 0 || receiver.isChainingApplied {
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

func (receiver *NonChainedLinkedCollectionNodes) ToChainedNodes() *[]*LinkedCollectionNode {
	length := receiver.Length()
	list := make([]*LinkedCollectionNode, length)

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
