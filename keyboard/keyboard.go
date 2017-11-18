package keyboard

/*
#ifdef __WIN32
#cgo CFLAGS:-nostdlib
#include <stdint.h>
#include <Windows.h>
void PressButton(WORD key,byte value) {
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
*/
import(
	"C"
)

func SetKey(keyId uint16, value bool) {
	var arg C.byte

	if value {
		arg = 1
	} else {
		arg = 0
	}
	C.PressButton(C.WORD(keyId),arg);
}