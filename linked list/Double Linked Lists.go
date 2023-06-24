package main

import "fmt"

type Dllist struct {
	head *Node
	tail *Node
}

type Node struct {
	value    any
	previous *Node
	next     *Node
}

//末尾添加元素
func (d *Dllist) Append(value any) {
	newNode := &Node{
		value: value,
	}

	//判断是否是第一次添加
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		//若不是第一个元素
		//节点指向前一个数据的指针
		newNode.previous = d.tail
		//节点指向后的指针(注意先执行，这是为了更改前一个节点的next指向)
		d.tail.next = newNode
		//将末尾节点换成新节点指针
		d.tail = newNode
	}

	// d.InsertNode(value, -1)  //可以直接调用下面的中间节点插入
}

//中间节点插入数据
func (d *Dllist) InsertNode(value any, position int) {
	newNode := &Node{
		value: value,
	}
	length := d.Len()
	//判断是否是第一次添加
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}

	//头部插入插入数据
	if position == 0 {
		newNode.next = d.head
		//还需要以前首个节点的指针指向(现在这个节点成了第二个)
		d.head.previous = newNode
		d.head = newNode
		//不需要修改双链的末尾指针
	} else if position > length || position < 0 { //添加到末尾，之后可以优化-1才是最后一个位置，-2倒数第二位置
		newNode.previous = d.tail
		d.tail.next = newNode
		d.tail = newNode
	} else { //添加到中间
		current := d.head
		for i := 0; i < position-1; i++ {
			//挨个遍历各个节点获取挨个节点指针
			current = current.next
		}
		//将新节点加入到链表
		newNode.previous = current
		newNode.next = current.next
		//修改新节点后一个节点的previous指向
		current.next.previous = newNode
		//修改新节点前一个节点的next指向
		current.next = newNode

	}
}

//待完成
func (d *Dllist) RemoveNode(position int) {

}

//打印所有数据
func (d *Dllist) PrintValue() {
	current := d.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

//获取链表长度
func (d *Dllist) Len() int {
	current := d.head
	var len int
	for current != nil {
		len++
		current = current.next
	}

	return len
}

func main() {
	doublelinklist := &Dllist{}
	doublelinklist.Append("a")
	doublelinklist.Append("b")
	doublelinklist.Append("c")
	doublelinklist.Append("d")
	doublelinklist.Append("e")

	doublelinklist.InsertNode("我", 1)
	doublelinklist.PrintValue()
	fmt.Println(doublelinklist.Len())
}
