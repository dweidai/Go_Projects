package skiplist

import(
	"math/rand"
	"errors"
	"fmt"
)
const defaultmax int = 20
const defaultprob float32 = 0.5

type node struct{
	key int
	value interface{}
	forward	[]*node
	level int
}

func newNode(newkey int, newvalue interface{}, newlevel int, maxlevel int) *node{
	empty := make([]*node, maxlevel)
	for i := 0; i < maxlevel; i++ {
		empty[i] = nil
	}
	toReturn := new(node)
	toReturn.key = newkey
	toReturn.value = newvalue
	toReturn.forward = empty
	toReturn.level = newlevel
	return toReturn
}

type skiplist struct{
	head *node
	level int
}

func newSkiplist() *skiplist{
	toReturn := new(skiplist)
	toReturn.head = newNode(0, "head", 1, defaultmax)
	toReturn.level = 1
	return toReturn
}

func (s *skiplist) randomLevel() int{
	level := 1
	random:= rand.Intn(2)
	for random < 1 && level > defaultmax{
		level ++
		random = rand.Intn(2)
	}
	return level
}

func(s *skiplist) search(skey int)(interface{}, error){
	current := s.head
	for i:=s.level-1; i>=0; i--{
		for current.forward[i] != nil && current.forward[i].key < skey{
			current = current.forward[i]
		}
	}
	current = current.forward[0]
	if current != nil && current.key == skey{
		//fmt.Println("reached")
		return current.value, nil
	}
	return nil, errors.New("not found")
}

func (s *skiplist) insert(skey int, value interface{}){
	current := s.head
	updateList := make([]*node, defaultmax)
	for i:= s.head.level-1; i>=0; i--{
		for current.forward[i] != nil && current.forward[i].key < skey{
			current = current.forward[i]
		}
		updateList[i] = current
	}
	current = current.forward[0]
	if current != nil && current.key == skey{
		current.value = value
	} else{
		newlevel := s.randomLevel()
		if newlevel > s.level{
			for i := s.level+1; i <= newlevel; i++{
				updateList[i-1] = s.head
			}
			s.level = newlevel
			s.head.level = newlevel
		}
		newnode := newNode(skey, value, newlevel, defaultmax)
		for i:=0; i <= newlevel-1; i++{
			newnode.forward[i] = updateList[i].forward[i]
			updateList[i].forward[i] = newnode
		}
	}

}

func (s *skiplist) delete(skey int) error{
	current := s.head
	updateList := make([]*node, defaultmax)
	for i:= s.head.level-1; i>=0; i--{
		for current.forward[i] != nil && current.forward[i].key < skey{
			current = current.forward[i]
		}
		updateList[i] = current
	}
	current = current.forward[0]
	if current.key == skey {
		for i := 0; i <= current.level-1; i++ {
			if updateList[i].forward[i] != nil && updateList[i].forward[i].key != current.key {
				break
			}
			updateList[i].forward[i] = current.forward[i]
		}

		for current.level > 1 && s.head.forward[current.level] == nil {
			current.level--
		}
		current = nil
		return nil
	}
	return errors.New("Not found")
}	

func printlist(s *skiplist){
	fmt.Printf("\nhead->")
	currentNode := s.head
	for {
		fmt.Printf("[key:%d, val:%v]->", currentNode.key, currentNode.value)
		if currentNode.forward[0] == nil {
			break
		}
		currentNode = currentNode.forward[0]
	}
	fmt.Printf("nil\n")
}