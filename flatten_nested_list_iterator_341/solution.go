package flatten_nested_list_iterator_341

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedInteger struct {
	value any
	isInt bool
}

func NewNestedInteger(value any) *NestedInteger {
	var isInt bool
	if _, ok := value.(int); ok {
		isInt = true
	}
	return &NestedInteger{
		value: value,
		isInt: isInt,
	}
}

func (this NestedInteger) IsInteger() bool {
	return this.isInt
}
func (this NestedInteger) GetInteger() int {
	v := this.value.(int)
	return v
}
func (this NestedInteger) GetList() []*NestedInteger {
	v := this.value.([]*NestedInteger)
	return v
}

type ListIndex struct {
	Idx  int
	List []*NestedInteger
}

type Stack struct {
	items []*ListIndex
}

func (s *Stack) Push(v *ListIndex) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() *ListIndex {
	if len(s.items) <= 0 {
		return nil
	}

	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

func (s *Stack) Peek() *ListIndex {
	if len(s.items) <= 0 {
		return nil
	}

	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

type NestedIterator struct {
	stack *Stack
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	stack := new(Stack)
	stack.Push(&ListIndex{
		Idx:  0,
		List: nestedList,
	})

	return &NestedIterator{
		stack: stack,
	}
}

func (this *NestedIterator) Next() int {
	current := this.stack.Peek()
	for !current.List[current.Idx].IsInteger() {
		next := &ListIndex{
			Idx:  0,
			List: current.List[current.Idx].GetList(),
		}
		this.stack.Push(next)

		current.Idx++
		current = next
	}

	value := current.List[current.Idx].GetInteger()
	current.Idx++

	if current.Idx == len(current.List) {
		this.stack.Pop()
	}

	return value
}

func (this *NestedIterator) HasNext() bool {
	return !this.stack.IsEmpty() && this.stack.Peek().Idx != len(this.stack.Peek().List)
}
