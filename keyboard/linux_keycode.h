int keycode[] = 
{
-1,
-1,
-1,
223,
-1,
-1,
-1,
-1,
158,
15,
-1,
-1,
-1,
28,
-1,
-1,
-1,
-1,
-1,
119,
-1,
122,
-1,
-1,
-1,
123,
-1,
1,
-1,
-1,
-1,
-1,
57,
-1,
-1,
107,
102,
105,
103,
106,
108,
-1,
210,
-1,
-1,
110,
111,
138,
11,
2,
3,
4,
5,
6,
7,
8,
9,
10,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
30,
48,
46,
32,
18,
33,
34,
35,
23,
36,
37,
38,
50,
49,
24,
25,
16,
19,
31,
20,
22,
47,
17,
45,
21,
44,
-1,
-1,
-1,
-1,
142,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
59,
60,
61,
62,
63,
64,
65,
66,
67,
68,
87,
88,
183,
184,
185,
186,
187,
188,
189,
190,
191,
192,
193,
194,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
69,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
-1,
207};

int get_keycode(int keyc) {
	if (keyc < sizeof(keycode)/sizeof(keycode[0])) {
		return keycode[keyc];
	} else {
		return -1;
	}
}