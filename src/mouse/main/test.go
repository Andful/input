package main

import (
"mouse"
"time"
)

func main() {
	time.Sleep(3000*time.Millisecond)
	mouse.SetPosition(0,0)
	time.Sleep(3000*time.Millisecond)
	mouse.Move(100,100)
}