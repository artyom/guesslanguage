package models

func init() {
	models["ur"] = compact(map[string]int{
		"یں ": 0,
		" کی": 1,
		"کے ": 2,
		" کے": 3,
		"نے ": 4,
		" کہ": 5,
		"ے ک": 6,
		"کی ": 7,
		"میں": 8,
		" می": 9,
		"ہے ": 10,
		"وں ": 11,
		"کہ ": 12,
		" ہے": 13,
		"ان ": 14,
		"ہیں": 15,
		"ور ": 16,
		" کو": 17,
		"یا ": 18,
		" ان": 19,
		" نے": 20,
		"سے ": 21,
		" سے": 22,
		" کر": 23,
		"ستا": 24,
		" او": 25,
		"اور": 26,
		"تان": 27,
		"ر ک": 28,
		"ی ک": 29,
		" اس": 30,
		"ے ا": 31,
		" پا": 32,
		" ہو": 33,
		" پر": 34,
		"رف ": 35,
		" کا": 36,
		"ا ک": 37,
		"ی ا": 38,
		" ہی": 39,
		"در ": 40,
		"کو ": 41,
		" ای": 42,
		"ں ک": 43,
		" مش": 44,
		" مل": 45,
		"ات ": 46,
		"صدر": 47,
		"اکس": 48,
		"شرف": 49,
		"مشر": 50,
		"پاک": 51,
		"کست": 52,
		"ی م": 53,
		" دی": 54,
		" صد": 55,
		" یہ": 56,
		"ا ہ": 57,
		"ن ک": 58,
		"وال": 59,
		"یہ ": 60,
		"ے و": 61,
		" بھ": 62,
		" دو": 63,
		"اس ": 64,
		"ر ا": 65,
		"نہی": 66,
		"کا ": 67,
		"ے س": 68,
		"ئی ": 69,
		"ہ ا": 70,
		"یت ": 71,
		"ے ہ": 72,
		"ت ک": 73,
		" سا": 74,
		"لے ": 75,
		"ہا ": 76,
		"ے ب": 77,
		" وا": 78,
		"ار ": 79,
		"نی ": 80,
		"کہا": 81,
		"ی ہ": 82,
		"ے م": 83,
		" سی": 84,
		" لی": 85,
		"انہ": 86,
		"انی": 87,
		"ر م": 88,
		"ر پ": 89,
		"ریت": 90,
		"ن م": 91,
		"ھا ": 92,
		"یر ": 93,
		" جا": 94,
		" جن": 95,
		"ئے ": 96,
		"پر ": 97,
		"ں ن": 98,
		"ہ ک": 99,
		"ی و": 100,
		"ے د": 101,
		" تو": 102,
		" تھ": 103,
		" گی": 104,
		"ایک": 105,
		"ل ک": 106,
		"نا ": 107,
		"کر ": 108,
		"ں م": 109,
		"یک ": 110,
		" با": 111,
		"ا ت": 112,
		"دی ": 113,
		"ن س": 114,
		"کیا": 115,
		"یوں": 116,
		"ے ج": 117,
		"ال ": 118,
		"تو ": 119,
		"ں ا": 120,
		"ے پ": 121,
		" چا": 122,
		"ام ": 123,
		"بھی": 124,
		"تی ": 125,
		"تے ": 126,
		"دوس": 127,
		"س ک": 128,
		"ملک": 129,
		"ن ا": 130,
		"ہور": 131,
		"یے ": 132,
		" مو": 133,
		" وک": 134,
		"ائی": 135,
		"ارت": 136,
		"الے": 137,
		"بھا": 138,
		"ردی": 139,
		"ری ": 140,
		"وہ ": 141,
		"ویز": 142,
		"ں د": 143,
		"ھی ": 144,
		"ی س": 145,
		" رہ": 146,
		" من": 147,
		" نہ": 148,
		" ور": 149,
		" وہ": 150,
		" ہن": 151,
		"ا ا": 152,
		"است": 153,
		"ت ا": 154,
		"ت پ": 155,
		"د ک": 156,
		"ز م": 157,
		"ند ": 158,
		"ورد": 159,
		"وکل": 160,
		"گی ": 161,
		"گیا": 162,
		"ہ پ": 163,
		"یز ": 164,
		"ے ت": 165,
		" اع": 166,
		" اپ": 167,
		" جس": 168,
		" جم": 169,
		" جو": 170,
		" سر": 171,
		"اپن": 172,
		"اکث": 173,
		"تھا": 174,
		"ثری": 175,
		"دیا": 176,
		"ر د": 177,
		"رت ": 178,
		"روی": 179,
		"سی ": 180,
		"ملا": 181,
		"ندو": 182,
		"وست": 183,
		"پرو": 184,
		"چاہ": 185,
		"کثر": 186,
		"کلا": 187,
		"ہ ہ": 188,
		"ہند": 189,
		"ہو ": 190,
		"ے ل": 191,
		" اک": 192,
		" دا": 193,
		" سن": 194,
		" وز": 195,
		" پی": 196,
		"ا چ": 197,
		"اء ": 198,
		"اتھ": 199,
		"اقا": 200,
		"اہ ": 201,
		"تھ ": 202,
		"دو ": 203,
		"ر ب": 204,
		"روا": 205,
		"رے ": 206,
		"سات": 207,
		"ف ک": 208,
		"قات": 209,
		"لا ": 210,
		"لاء": 211,
		"م م": 212,
		"م ک": 213,
		"من ": 214,
		"نوں": 215,
		"و ا": 216,
		"کرن": 217,
		"ں ہ": 218,
		"ھار": 219,
		"ہوئ": 220,
		"ہی ": 221,
		"یش ": 222,
		" ام": 223,
		" لا": 224,
		" مس": 225,
		" پو": 226,
		" پہ": 227,
		"انے": 228,
		"ت م": 229,
		"ت ہ": 230,
		"ج ک": 231,
		"دون": 232,
		"زیر": 233,
		"س س": 234,
		"ش ک": 235,
		"ف ن": 236,
		"ل ہ": 237,
		"لاق": 238,
		"لی ": 239,
		"وری": 240,
		"وزی": 241,
		"ونو": 242,
		"کھن": 243,
		"گا ": 244,
		"ں س": 245,
		"ں گ": 246,
		"ھنے": 247,
		"ھے ": 248,
		"ہ ب": 249,
		"ہ ج": 250,
		"ہر ": 251,
		"ی آ": 252,
		"ی پ": 253,
		" حا": 254,
		" وف": 255,
		" گا": 256,
		"ا ج": 257,
		"ا گ": 258,
		"اد ": 259,
		"ادی": 260,
		"اعظ": 261,
		"اہت": 262,
		"جس ": 263,
		"جمہ": 264,
		"جو ": 265,
		"ر س": 266,
		"ر ہ": 267,
		"رنے": 268,
		"س م": 269,
		"سا ": 270,
		"سند": 271,
		"سنگ": 272,
		"ظم ": 273,
		"عظم": 274,
		"ل م": 275,
		"لیے": 276,
		"مل ": 277,
		"موہ": 278,
		"مہو": 279,
		"نگھ": 280,
		"و ص": 281,
		"ورٹ": 282,
		"وہن": 283,
		"کن ": 284,
		"گھ ": 285,
		"گے ": 286,
		"ں ج": 287,
		"ں و": 288,
		"ں ی": 289,
		"ہ د": 290,
		"ہن ": 291,
		"ہوں": 292,
		"ے ح": 293,
		"ے گ": 294,
		"ے ی": 295,
		" اگ": 296,
		" بع": 297,
		" رو": 298,
		" شا": 299})
}
