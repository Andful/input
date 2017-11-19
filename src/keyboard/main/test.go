package main

import (
"keyboard"
)

func main() {
	//time.Sleep(7000*time.Millisecond)

	keyboard.SetKey(66,true)
	keyboard.SetKey(66,false)
}