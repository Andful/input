package main

import (
"keyboard"
"time"
)

func main() {
	time.Sleep(7000*time.Millisecond)

	keyboard.SetKey(65,true)
	keyboard.SetKey(65,false)
}