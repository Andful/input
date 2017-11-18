package keyboard

/*
#include <stdint.h>
#include <Windows.h>
void PressButton(WORD key,byte value) {
	INPUT ip;

	ip.type = INPUT_KEYBOARD;
	ip.ki.wScan = 0; // hardware scan code for key
	ip.ki.time = 0;
	ip.ki.dwExtraInfo = 0;

	ip.ki.wVk = key; // virtual-key code for the "a" key
	if (value) {
		ip.ki.dwFlags = 0; // 0 for key press
	} else {
		ip.ki.dwFlags = KEYEVENTF_KEYUP;
	}
	
	SendInput(1, &ip, sizeof(INPUT));
}
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