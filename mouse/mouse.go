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
    Display *display = XOpenDisplay(NULL);

    if(display == NULL) {
        exit(EXIT_FAILURE);
    }

    XWarpPointer(display, None, None, 0, 0, 0, 0, dx, dy);

    XCloseDisplay(display);
}

void SetPosition(int32_t x,int32_t y) {
	Display *display = XOpenDisplay(NULL);

    if(display == NULL) {
        exit(EXIT_FAILURE);
    }

    Window root_window = XRootWindow(display, 0);

    XWarpPointer(display, None, root_window, 0, 0, 0, 0, x, y);

    XCloseDisplay(display);
}

void SetButton(int button,byte value){
    Display *display = XOpenDisplay(NULL);

    XEvent event;
    
    if(display == NULL)
    {
        exit(EXIT_FAILURE);
    }
    
    memset(&event, 0x00, sizeof(event));
    
    event.xbutton.button = button;
    event.xbutton.same_screen = True;

    if(value) {
        event.type = ButtonPress;    
        XQueryPointer(display, RootWindow(display, DefaultScreen(display)), &event.xbutton.root, &event.xbutton.window, &event.xbutton.x_root, &event.xbutton.y_root, &event.xbutton.x, &event.xbutton.y, &event.xbutton.state);
        
        event.xbutton.subwindow = event.xbutton.window;
        
        while(event.xbutton.subwindow)
        {
            event.xbutton.window = event.xbutton.subwindow;
            
            XQueryPointer(display, event.xbutton.window, &event.xbutton.root, &event.xbutton.subwindow, &event.xbutton.x_root, &event.xbutton.y_root, &event.xbutton.x, &event.xbutton.y, &event.xbutton.state);
        }
        
        if(XSendEvent(display, PointerWindow, True, 0xfff, &event) == 0) fprintf(stderr, "Errore nell'invio dell'evento !!!\n");
        
    } else {
        event.type = ButtonRelease;
        event.xbutton.state = 0x100;
        
        if(XSendEvent(display, PointerWindow, True, 0xfff, &event) == 0) fprintf(stderr, "Errore nell'invio dell'evento !!!\n");

    }
    
    XCloseDisplay(display);
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

func SetButton(button int, value bool){
    C.SetButton(button,boolToByte(value))
}

func boolToByte(value bool) C.uint8_t {
    if(value) {
        return 1
    } else {
        return 0
    }
}
