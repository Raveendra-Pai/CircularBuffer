package cbuffer

import (
	"fmt"
	"math"
	"sync"
)

//structure to hold buffer data
type CircularBuffer[T any] struct {
	bufferSize int
	head       int
	tail       int
	buffer     []T
	mutex      *sync.Mutex //this is to synchronize multiple Go Routines to access single Circular buffer
}

//Create buffer which does boundary check for size
func CreateBuffer[T any](buffersize int) (*CircularBuffer[T], error) {

	if buffersize >= math.MaxInt {
		return nil, fmt.Errorf("Buffer size is out of the limit")
	}

	if buffersize <= 0 {
		return nil, fmt.Errorf("size should be non-zero positive number")
	}

	//Create slice with length as same as capacity
	circularBuffer := &CircularBuffer[T]{
		bufferSize: buffersize,
		head:       -1,
		tail:       -1,
		buffer:     make([]T, buffersize, buffersize),
		mutex:      new(sync.Mutex),
	}

	return circularBuffer, nil
}

func (c *CircularBuffer[T]) EnQueue(item T) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.head != -1 && c.tail != -1 &&
		(c.tail+1)%c.bufferSize == c.head%c.bufferSize { //this condition will hit if the tail index is less than head index
		return fmt.Errorf("Circular Buffer is full!!")

	} else if c.head == -1 && c.tail == -1 { //empty buffer check
		c.head, c.tail = 0, 0
	} else {
		c.tail++ //I would love to use operating overloading here, if GoLang was supporting
		//since the buffer is circular, once the tail is reached array max limit,we need to store it at beginning of the array
		if c.tail == c.bufferSize {
			c.tail = 0
		}
	}
	c.buffer[c.tail] = item
	return nil
}

func (c *CircularBuffer[T]) DeQueue() (T, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.head == -1 && c.tail == -1 { //buffer empty check
		var def T //To return default
		return def, fmt.Errorf("Circular buffer is empty")
	}

	item := c.buffer[c.head]
	//this condition will hit if and only if a single element exist in that buffer
	if c.head == c.tail {
		c.head, c.tail = -1, -1 //reset the buffer head and tail to treat it as empty buffer
	} else {
		c.head++                    //I wish Go Lang should support operating overloading, then I would do increment operator overloading to handle boundary condition mentioned in next line
		if c.head == c.bufferSize { //rotate the head to start of the array to prevent index out of range
			c.head = 0
		}
	}
	return item, nil
}

func (c *CircularBuffer[T]) PrintBuffer() {

	if c.head == -1 || c.tail == -1 {
		fmt.Println("Empty buffer")
		return
	}
	if c.tail%c.bufferSize < c.head { //this condition will hit when tail index is less than head index
		for index := c.head; index <= c.bufferSize-1; index++ { // print from head to array end since tail is in the other end
			fmt.Printf("%v\t", c.buffer[index])
		}
		for i := 0; i <= c.tail; i++ { //print from array start to tail end
			fmt.Printf("%v\t", c.buffer[i])
		}
	} else { //when head index is less than tail then its straight forward looping
		for i := c.head; i <= c.tail; i++ {
			fmt.Printf("%v\t", c.buffer[i])
		}
	}
	fmt.Println()
}
