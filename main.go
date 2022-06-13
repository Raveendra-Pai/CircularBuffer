package main

import (
	"fmt"
	cbuffer "main/circularbuffer"
)

func main() {
	//creating a circular buffer to store heterogenious data
	cbuffer, err := cbuffer.CreateBuffer[interface{}](5)

	if err != nil {
		fmt.Println("Faile to create circular buffer, Error: ", err)
		return
	}

	/*Its strongly recommonded to check error returned on each EnQueue and DeQueue.
	For simplicity I have ignored return values of EnQueue and DeQueue calls.
	*/
	cbuffer.EnQueue(10)
	cbuffer.EnQueue("Hello")
	cbuffer.EnQueue(true)
	cbuffer.EnQueue(false)
	cbuffer.EnQueue('a')
	cbuffer.PrintBuffer()
	cbuffer.DeQueue()
	cbuffer.PrintBuffer()
	cbuffer.DeQueue()
	cbuffer.PrintBuffer()
	cbuffer.DeQueue()
	cbuffer.PrintBuffer()
	cbuffer.DeQueue()
	cbuffer.PrintBuffer()
	cbuffer.EnQueue('a')
	cbuffer.PrintBuffer()
	cbuffer.EnQueue("asmi")
	cbuffer.PrintBuffer()
	cbuffer.EnQueue(true)
	cbuffer.PrintBuffer()
	cbuffer.DeQueue()
	cbuffer.DeQueue()
	cbuffer.DeQueue()
	cbuffer.DeQueue()
	cbuffer.PrintBuffer()
	cbuffer.DeQueue()
	cbuffer.PrintBuffer()
}
