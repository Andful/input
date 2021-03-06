package keyboard

/*
#include <stdint.h>
#define OS_DOES_NOT_SUPPORT_BUTTON -1
#define GENERAL_ERROR -2

#ifdef __WIN32
#cgo windows CFLAGS:-nostdlib
#include <Windows.h>

void SetKey(uint16_t key, uint8_t value) {
	INPUT ip;

	ip.type = INPUT_KEYBOARD;
	ip.ki.wScan = 0;
	ip.ki.time = 0;
	ip.ki.dwExtraInfo = 0;

	ip.ki.wVk = key;
	if (value) {
		ip.ki.dwFlags = 0;
	} else {
		ip.ki.dwFlags = KEYEVENTF_KEYUP;
	}
	
	SendInput(1, &ip, sizeof(INPUT));
}
#endif

#ifdef __linux__
#cgo linux LDFLAGS: -lX11 -lXtst
#include <stdlib.h>

#include <X11/Xlib.h>
#include <X11/keysym.h>
#include <X11/extensions/XTest.h>

#include "linux_keycode.h"

int adjust(int keycode) {
	return keycode + 8;
}

void SetKey(uint16_t key, uint8_t value) {
	Display *display;
	display = XOpenDisplay(NULL);

	if(display == NULL) {
	    exit(EXIT_FAILURE);
	}

	XTestFakeKeyEvent(display,adjust(get_keycode(key)), value, 0);
	XCloseDisplay(display);
}
#endif
*/
import "C"

func SetKey(keyId uint16, value bool) {
	C.SetKey(C.uint16_t(keyId),boolToByte(value));
}

func boolToByte(value bool) C.uint8_t {
	if(value) {
		return 1
	} else {
		return 0
	}
}
