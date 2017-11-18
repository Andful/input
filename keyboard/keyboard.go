package keyboard

/*
#include <stdint.h>
#ifdef __WIN32
#cgo CFLAGS:-nostdlib
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
*/
import(
	"C"
)

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