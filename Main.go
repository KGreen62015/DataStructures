package main

import (
	"fmt"
)

//////////List//////////
var (
	array [100]string
	Length = 0
	queueLength = 0
)
func add(value string) {
	array[Length] = value
	Length++
}
func get(position int)string{
	value := array[position]
	return value
}
func clearList(){
	Length = 0
}

//////////Stack//////////
func push(value string)  {
	array[Length] = value
	Length++
}

func postPop() {
	Length--
}

func pop(position int)string  {
	defer postPop()
	return array[Length]
}

func clearStack()  {
	Length = 0
}

//////////Stack//////////
func enqueue(value string)  {
	if Length>99 {

	}
	array[Length] = value
	Length++
}
func dequeue()  {
	
}
func clearQueue()  {

}
func main() {
	add("Red")
	fmt.Println(get(Length))
	clearList()
	push("Blue")
	fmt.Println(pop(Length))
	clearStack()
}