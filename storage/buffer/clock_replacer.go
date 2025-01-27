// this code is from https://github.com/brunocalza/go-bustub
// there is license and copyright notice in licenses/go-bustub dir

package buffer

import (
	"fmt"
	"sync"
)

// FrameID is the type for frame id
type FrameID uint32

/**
 * ClockReplacer implements the clock replacement policy, which approximates the Least Recently Used policy.
 */
type ClockReplacer struct {
	cList     *circularList
	clockHand **node
	mutex     *sync.Mutex
}

// Victim removes the victim frame as defined by the replacement policy
func (c *ClockReplacer) Victim() *FrameID {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.cList.size == 0 {
		fmt.Println("Victim: page which can be cache out is not exist!")
		//panic("Victim: page which can be cache out is not exist!")
		return nil
	}

	var victimFrameID *FrameID
	currentNode := (*c.clockHand)
	for {
		//if currentNode.value.(bool) {
		if currentNode.value {
			currentNode.value = false
			c.clockHand = &currentNode.next
		} else {
			//frameID := currentNode.key.(FrameID)
			frameID := currentNode.key
			victimFrameID = &frameID

			c.clockHand = &currentNode.next

			c.cList.remove(currentNode.key)
			return victimFrameID
		}
	}
}

// Unpin unpins a frame, indicating that it can now be victimized
func (c *ClockReplacer) Unpin(id FrameID) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if !c.cList.hasKey(id) {
		c.cList.insert(id, true)
		if c.cList.size == 1 {
			c.clockHand = &c.cList.head
		}
	}
}

// Pin pins a frame, indicating that it should not be victimized until it is unpinned
func (c *ClockReplacer) Pin(id FrameID) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	//node := c.cList.find(id)
	//if node == nil {
	node, ok := c.cList.supportMap[id]
	if !ok {
		return
	}

	if (*c.clockHand) == node {
		c.clockHand = &(*c.clockHand).next
	}
	c.cList.remove(id)
	delete(c.cList.supportMap, id)
}

// Size returns the size of the clock
func (c *ClockReplacer) Size() uint32 {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.cList.size
}

// NewClockReplacer instantiates a new clock replacer
func NewClockReplacer(poolSize uint32) *ClockReplacer {
	cList := newCircularList(poolSize)
	return &ClockReplacer{cList, &cList.head, new(sync.Mutex)}
}
