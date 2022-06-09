package main

import (
	"fmt"
	cbuffer "main/circularbuffer"
)

func main() {
	cbuffer, err := cbuffer.CreateBuffer[int](5)

	if err != nil {
		fmt.Println("Faile to create circular buffer, Error: ", err)
		return
	}

	/*Its strongly recommonded to check error returned on each EnQueue and DeQueue.
	For simplicity I have ignored return values of EnQueue and DeQueue calls.
	*/
	cbuffer.EnQueue(10)
	cbuffer.EnQueue(20)
	cbuffer.EnQueue(30)
	cbuffer.EnQueue(40)
	cbuffer.EnQueue(50)
	cbuffer.PrintBuffer()
	cbuffer.DeQueue()
	cbuffer.PrintBuffer()
	cbuffer.DeQueue()
	cbuffer.PrintBuffer()
	cbuffer.DeQueue()
	cbuffer.PrintBuffer()
	cbuffer.DeQueue()
	cbuffer.PrintBuffer()
	cbuffer.EnQueue(60)
	cbuffer.PrintBuffer()
	cbuffer.EnQueue(20)
	cbuffer.PrintBuffer()
	cbuffer.EnQueue(30)
	cbuffer.PrintBuffer()
	cbuffer.DeQueue()
	cbuffer.DeQueue()
	cbuffer.DeQueue()
	cbuffer.DeQueue()
	cbuffer.PrintBuffer()
	cbuffer.DeQueue()
	cbuffer.PrintBuffer()
}
