package models

func init() {
	models["fa"] = compact(map[string]int{
		"ان ": 0,
		"ای ": 1,
		"ه ا": 2,
		" اي": 3,
		" در": 4,
		"به ": 5,
		" بر": 6,
		"در ": 7,
		"ران": 8,
		" به": 9,
		"ی ا": 10,
		"از ": 11,
		"ين ": 12,
		"می ": 13,
		" از": 14,
		"ده ": 15,
		"ست ": 16,
		"است": 17,
		" اس": 18,
		" که": 19,
		"که ": 20,
		"اير": 21,
		"ند ": 22,
		"اين": 23,
		" ها": 24,
		"يرا": 25,
		"ود ": 26,
		" را": 27,
		"های": 28,
		" خو": 29,
		"ته ": 30,
		"را ": 31,
		"رای": 32,
		"رد ": 33,
		"ن ب": 34,
		"کرد": 35,
		" و ": 36,
		" کر": 37,
		"ات ": 38,
		"برا": 39,
		"د ک": 40,
		"مان": 41,
		"ی د": 42,
		" ان": 43,
		"خوا": 44,
		"شور": 45,
		" با": 46,
		"ن ا": 47,
		" سا": 48,
		"تمی": 49,
		"ری ": 50,
		"اتم": 51,
		"ا ا": 52,
		"واه": 53,
		" ات": 54,
		" عر": 55,
		"اق ": 56,
		"ر م": 57,
		"راق": 58,
		"عرا": 59,
		"ی ب": 60,
		" تا": 61,
		" تو": 62,
		"ار ": 63,
		"ر ا": 64,
		"ن م": 65,
		"ه ب": 66,
		"ور ": 67,
		"يد ": 68,
		"ی ک": 69,
		" ام": 70,
		" دا": 71,
		" کن": 72,
		"اهد": 73,
		"هد ": 74,
		" آن": 75,
		" می": 76,
		" ني": 77,
		" گف": 78,
		"د ا": 79,
		"گفت": 80,
		" کش": 81,
		"ا ب": 82,
		"نی ": 83,
		"ها ": 84,
		"کشو": 85,
		" رو": 86,
		"ت ک": 87,
		"نيو": 88,
		"ه م": 89,
		"وی ": 90,
		"ی ت": 91,
		" شو": 92,
		"ال ": 93,
		"دار": 94,
		"مه ": 95,
		"ن ک": 96,
		"ه د": 97,
		"يه ": 98,
		" ما": 99,
		"امه": 100,
		"د ب": 101,
		"زار": 102,
		"ورا": 103,
		"گزا": 104,
		" پي": 105,
		"آن ": 106,
		"انت": 107,
		"ت ا": 108,
		"فت ": 109,
		"ه ن": 110,
		"ی خ": 111,
		"اما": 112,
		"بات": 113,
		"ما ": 114,
		"ملل": 115,
		"نام": 116,
		"ير ": 117,
		"ی م": 118,
		"ی ه": 119,
		" آم": 120,
		" ای": 121,
		" من": 122,
		"انس": 123,
		"اني": 124,
		"ت د": 125,
		"رده": 126,
		"ساز": 127,
		"ن د": 128,
		"نه ": 129,
		"ورد": 130,
		" او": 131,
		" بي": 132,
		" سو": 133,
		" شد": 134,
		"اده": 135,
		"اند": 136,
		"با ": 137,
		"ت ب": 138,
		"ر ب": 139,
		"ز ا": 140,
		"زما": 141,
		"سته": 142,
		"ن ر": 143,
		"ه س": 144,
		"وان": 145,
		"وز ": 146,
		"ی ر": 147,
		"ی س": 148,
		" هس": 149,
		"ابا": 150,
		"ام ": 151,
		"اور": 152,
		"تخا": 153,
		"خاب": 154,
		"خود": 155,
		"د د": 156,
		"دن ": 157,
		"رها": 158,
		"روز": 159,
		"رگز": 160,
		"نتخ": 161,
		"ه ش": 162,
		"ه ه": 163,
		"هست": 164,
		"يت ": 165,
		"يم ": 166,
		" دو": 167,
		" دي": 168,
		" مو": 169,
		" نو": 170,
		" هم": 171,
		" کا": 172,
		"اد ": 173,
		"اری": 174,
		"انی": 175,
		"بر ": 176,
		"بود": 177,
		"ت ه": 178,
		"ح ه": 179,
		"حال": 180,
		"رش ": 181,
		"عه ": 182,
		"لی ": 183,
		"وم ": 184,
		"ژان": 185,
		" سل": 186,
		"آمر": 187,
		"اح ": 188,
		"توس": 189,
		"داد": 190,
		"دام": 191,
		"ر د": 192,
		"ره ": 193,
		"ريک": 194,
		"زی ": 195,
		"سلا": 196,
		"شود": 197,
		"لاح": 198,
		"مري": 199,
		"نند": 200,
		"ه ع": 201,
		"يما": 202,
		"يکا": 203,
		"پيم": 204,
		"گر ": 205,
		" آژ": 206,
		" ال": 207,
		" بو": 208,
		" مق": 209,
		" مل": 210,
		" وی": 211,
		"آژا": 212,
		"ازم": 213,
		"ازی": 214,
		"بار": 215,
		"برن": 216,
		"ر آ": 217,
		"ز س": 218,
		"سعه": 219,
		"شته": 220,
		"مات": 221,
		"ن آ": 222,
		"ن پ": 223,
		"نس ": 224,
		"ه گ": 225,
		"وسع": 226,
		"يان": 227,
		"يوم": 228,
		"کا ": 229,
		"کام": 230,
		"کند": 231,
		" خا": 232,
		" سر": 233,
		"آور": 234,
		"ارد": 235,
		"اقد": 236,
		"ايم": 237,
		"ايی": 238,
		"برگ": 239,
		"ت ع": 240,
		"تن ": 241,
		"خت ": 242,
		"د و": 243,
		"ر خ": 244,
		"رک ": 245,
		"زير": 246,
		"فته": 247,
		"قدا": 248,
		"ل ت": 249,
		"مين": 250,
		"ن گ": 251,
		"ه آ": 252,
		"ه خ": 253,
		"ه ک": 254,
		"ورک": 255,
		"ويو": 256,
		"يور": 257,
		"يوي": 258,
		"يی ": 259,
		"ک ت": 260,
		"ی ش": 261,
		" اق": 262,
		" حا": 263,
		" حق": 264,
		" دس": 265,
		" شک": 266,
		" عم": 267,
		" يک": 268,
		"ا ت": 269,
		"ا د": 270,
		"ارج": 271,
		"بين": 272,
		"ت م": 273,
		"ت و": 274,
		"تاي": 275,
		"دست": 276,
		"ر ح": 277,
		"ر س": 278,
		"رنا": 279,
		"ز ب": 280,
		"شکا": 281,
		"لل ": 282,
		"م ک": 283,
		"مز ": 284,
		"ندا": 285,
		"نوا": 286,
		"و ا": 287,
		"وره": 288,
		"ون ": 289,
		"وند": 290,
		"يمز": 291,
		" آو": 292,
		" اع": 293,
		" فر": 294,
		" مت": 295,
		" نه": 296,
		" هر": 297,
		" وز": 298,
		" گز": 299})
}
