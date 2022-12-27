package index

import (
	"encoding/json"
	// "fmt"
)

type Node struct {
	Item IndexItem
	Next *Node
}
type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Insert(item IndexItem) {
	newNode := Node{Item: item, Next: nil}
	if list.Head == nil {
		list.Head = &newNode
	} else {
		tempHead := list.Head // to avoid the lost of head
		if tempHead.Item.Key == item.Key {
			tempHead.Item.Value = tempHead.Item.Value + "; " + item.Value
			return
		}
		for tempHead.Next != nil {
			tempHead = tempHead.Next
			if tempHead.Item.Key == item.Key {
				tempHead.Item.Value = tempHead.Item.Value + "; " + item.Value
				return
			}
		}
		tempHead.Next = &newNode
	}

}

func (list *LinkedList) Search(key string) string {
	if list.Head == nil {
		return ""
	} else {
		tempHead := list.Head // to avoid the lost of head
		if tempHead.Item.Key == key {
			return tempHead.Item.Value
		}
		for tempHead.Next != nil {
			tempHead = tempHead.Next
			if tempHead.Item.Key == key {
				return tempHead.Item.Value
			}
		}
	}

	return ""
}

func (list *LinkedList) Display() string {
	fullList := ""
	if list.Head == nil {
		return fullList
	} else {
		tempHead := list.Head // to avoid the lost of head
		for tempHead.Next != nil {
			ser, _ := json.Marshal(tempHead.Item)
			str := string(ser)
			fullList += "->" + str
			tempHead = tempHead.Next
		}
		ser, _ := json.Marshal(tempHead.Item)
		str := string(ser)
		fullList += "->" + str
	}

	return fullList
}
