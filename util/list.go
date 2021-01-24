package util

type Node struct {
	next   *Node
	entity interface{}
	prior  *Node
}

type LinkedList struct {
	first  *Node
	last   *Node
	length int
}

func (this *LinkedList) Add(node interface{}) {
	if this.first == nil {
		this.first = &Node{entity: node}
		this.last = this.first
	} else {
		this.last.next = &Node{entity: node, prior: this.last}
		this.last = this.last.next
	}
	this.length += 1
}
func (this *LinkedList) Length() int {
	return this.length
}
func (this *LinkedList) Get(index int) interface{} {
	if index >= this.length || index < 0 {
		return nil
	}
	ret := this.first
	for i := 1; i < index; i++ {
		ret = ret.next
	}
	return ret
}
func (this *LinkedList) remove(node *Node) {
	node.prior.next = node.prior
	node.next.prior = node.next

}
func (this *LinkedList) Remove(entity interface{}) {
	ret := this.first
	for i := 0; i < this.length; i++ {
		if entity == ret.entity {
			this.remove(ret)
			return
		} else {
			ret = ret.next
		}
	}
}
