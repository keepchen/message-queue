package queue

import (
	"errors"
	"fmt"
)

type doubleLinkedList struct {
	val  interface{}
	prev *doubleLinkedList
	next *doubleLinkedList
}

//lPush 向队列头部添加节点
func (l *doubleLinkedList) lPush(val interface{}) (*doubleLinkedList, error) {
	newNode := &doubleLinkedList{val: val}
	if l == nil {
		l = newNode
	} else {
		if l.prev != nil {
			//移动节点到头部
			l = l.movePointerToHead()
		}
		//新节点作为新的头部
		newNode.next = l
		l.prev = newNode
		l = newNode
	}

	return l, nil
}

//lPop 从队列头部返回节点值
func (l *doubleLinkedList) lPop() (*doubleLinkedList, interface{}, error) {
	var val interface{}
	if l == nil {
		return l, val, errors.New("[lPop] empty list")
	}
	//仅有一个节点
	if l.prev == nil && l.next == nil {
		val = l.val
		l = nil
	} else {
		if l.prev != nil {
			//移动节点到头部
			l = l.movePointerToHead()
		}
		val = l.val
		next := l.next
		//移除当前节点
		next.prev = nil
		l = next
	}

	return l, val, nil
}

//rPush 向队列尾部追加节点
func (l *doubleLinkedList) rPush(val interface{}) (*doubleLinkedList, error) {
	newNode := &doubleLinkedList{val: val}
	if l == nil {
		l = newNode
	} else {
		if l.next != nil {
			//将指针移动到队列尾部
			l = l.movePointerToTail()
		}
		newNode.prev = l
		l.next = newNode
		l = newNode
	}

	return l, nil
}

//rPop 从队列尾部弹出节点值
func (l *doubleLinkedList) rPop() (*doubleLinkedList, interface{}, error) {
	var val interface{}
	if l == nil {
		return l, nil, errors.New("[rPop] empty list")
	}

	//当前仅有一个节点
	if l.prev == nil && l.next == nil {
		val = l.val
		l = nil
	} else {
		if l.next != nil {
			//将指针移动到队列尾部
			l = l.movePointerToTail()
		}
		val = l.val
		prev := l.prev
		//移除当前节点
		prev.next = nil
		l = prev
	}

	return l, val, nil
}

//displayQueue 打印队列信息
func (l *doubleLinkedList) displayQueue() {
	if l == nil {
		fmt.Printf("[displayQueue] db queue is empty\n")
		return
	}
	//将指针移动到队列头部
	l = l.movePointerToHead()
	for {
		if l == nil {
			break
		}
		fmt.Printf("item---> %#v\n", l.val)
		l = l.next
	}
}

//移动指针到队列头部
func (l *doubleLinkedList) movePointerToHead() *doubleLinkedList {
	if l != nil {
		for {
			if l.prev == nil {
				break
			}
			l = l.prev
		}
	}

	return l
}

//移动指针到队列尾部
func (l *doubleLinkedList) movePointerToTail() *doubleLinkedList {
	if l != nil {
		for {
			if l.next == nil {
				break
			}
			l = l.next
		}
	}

	return l
}
