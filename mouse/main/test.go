package main

import (
"mouse"
"time"
)

func main() {
	time.Sleep(3000*time.Millisecond)
	mouse.SetPosition(110,110)
	time.Sleep(100*time.Millisecond)
	mouse.SetPosition(100,100)
	time.Sleep(100*time.Millisecond)
	mouse.SetPosition(90,90)
	time.Sleep(100*time.Millisecond)
	mouse.SetPosition(80,80)
	time.Sleep(100*time.Millisecond)
	mouse.SetPosition(70,70)
	time.Sleep(100*time.Millisecond)
	mouse.SetPosition(60,60)
	time.Sleep(100*time.Millisecond)
	mouse.SetPosition(50,50)
	time.Sleep(100*time.Millisecond)
	mouse.SetPosition(40,40)
	time.Sleep(100*time.Millisecond)
	mouse.SetPosition(30,30)
	time.Sleep(100*time.Millisecond)
	mouse.SetPosition(20,20)
	time.Sleep(100*time.Millisecond)
	mouse.SetPosition(10,10)
	time.Sleep(100*time.Millisecond)
	mouse.SetPosition(0,0)
	time.Sleep(3000*time.Millisecond)
	mouse.Move(100,100)
}