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

#ifdef __linux__
#cgo LDFLAGS: -lX11

#include <stdlib.h>

#include <X11/Xlib.h>
#include <X11/Xutil.h>

void Move(int32_t dx,int32_t dy) {
    Display *displayMain = XOpenDisplay(NULL);

    if(displayMain == NULL) {
        exit(EXIT_FAILURE);
    }

    XWarpPointer(displayMain, None, None, 0, 0, 0, 0, dx, dy);

    XCloseDisplay(displayMain);
}

void SetPosition(int32_t x,int32_t y) {
	Display *displayMain = XOpenDisplay(NULL);

    if(displayMain == NULL) {
        exit(EXIT_FAILURE);
    }

    Window root_window = XRootWindow(displayMain, 0);

    XWarpPointer(displayMain, None, root_window, 0, 0, 0, 0, x, y);

    XCloseDisplay(displayMain);
}
#endif
*/
import "C"

func SetPosition(x int,y int) {
	C.SetPosition(C.int32_t(x),C.int32_t(y))
}

func Move(dx int,dy int){
	C.Move(C.int32_t(dx),C.int32_t(dy))
}
