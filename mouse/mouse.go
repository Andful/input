package mouse

/*
#include <stdint.h>

#ifdef __WIN32
#cgo CFLAGS:-nostdlib
#include <Windows.h>

void Move(int32_t dx,int32_t dy) {
	POINT p;
	if (GetCursorPos(&p)) {
		SetCursorPos(p.x+dx,p.y+dy);
	}
}

void SetPosition(int32_t x,int32_t y) {
	SetCursorPos(x,y);
}

#endif
*/
import (
"C"
)

func SetPosition(x int,y int) {
	C.SetPosition(C.int32_t(x),C.int32_t(y))
}

func Move(dx int,dy int){
	C.Move(C.int32_t(dx),C.int32_t(dy))
}