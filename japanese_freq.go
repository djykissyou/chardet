package chardet

var freq_dad = [...]float64{
	7.6455537, // の
	7.3060647, // い
	5.9703762, // た
	5.4956975, // と
	4.8033590, // か
	4.7351322, // て
	4.7191237, // し
	4.5240123, // な
	4.4300475, // に
	4.2768832, // っ
	4.1119433, // う
	4.1045409, // は
	3.7992355, // が
	3.6168427, // で
	3.2427086, // る
	3.2243530, // を
	3.2052752, // ん
	2.8824367, // こ
	2.8408707, // ら
	2.8164166, // そ
	2.8162761, // あ
	2.6973156, // れ
	2.5002783, // も
	2.4854935, // ま
	0.4605158, // ン
	0.3328489, // ロ
	0.3328288, // ス
	0.2113406, // ト
	0.2085521, // ル
	0.2036774, // レ
}

var freq_sub = [...]map[int]float64{
	{
		0x03: 7.4947066,
		0x04: 9.2921303,
		0x07: 7.9909336,
		0x0b: 13.792685,
		0x0c: 5.4202539,
		0x0d: 19.181225,
		0x11: 10.096167,
		0x1e: 11.329273,
		0x1f: 8.5729743,
		0x20: 6.8296505,
	},
	{
		0x00: 4.1923378,
		0x02: 18.564749,
		0x04: 5.5367018,
		0x05: 10.588869,
		0x07: 4.3001199,
		0x09: 9.1027378,
		0x0a: 16.947046,
		0x0e: 18.195765,
		0x17: 8.3380670,
		0x21: 4.2336057,
	},
	{
		0x00: 22.533146,
		0x01: 7.7188602,
		0x03: 11.327864,
		0x04: 6.4426613,
		0x06: 9.5425387,
		0x0c: 11.476766,
		0x10: 12.749957,
		0x12: 6.4358930,
		0x1f: 6.1117671,
		0x22: 5.6605476,
	},
	{
		0x01: 26.535088,
		0x06: 11.034846,
		0x08: 7.3456278,
		0x0b: 9.7803564,
		0x0c: 5.9236428,
		0x0d: 4.8685937,
		0x0f: 7.2100774,
		0x11: 11.400055,
		0x16: 6.7930653,
		0x23: 9.1086476,
	},
	{
		0x06: 9.3496615,
		0x07: 5.4074887,
		0x08: 6.5506534,
		0x09: 16.620974,
		0x12: 39.455906,
		0x15: 4.1981485,
		0x16: 4.3035931,
		0x22: 5.1587850,
		0x24: 4.6744671,
		0x25: 4.2803226,
	},
	{
		0x01: 54.476628,
		0x06: 3.5903560,
		0x0b: 3.0424364,
		0x0e: 4.1179822,
		0x14: 3.2968834,
		0x16: 7.6771177,
		0x23: 7.8777094,
		0x26: 6.5196182,
		0x27: 5.7953029,
		0x28: 3.6059662,
	},
	{
		0x01: 7.6134032,
		0x02: 26.295394,
		0x04: 9.2246230,
		0x05: 32.655093,
		0x07: 3.3436109,
		0x15: 1.7047977,
		0x17: 5.2547599,
		0x26: 2.8703226,
		0x29: 7.1671086,
		0x2a: 3.8708865,
	},
	{
		0x01: 28.409887,
		0x02: 4.3736679,
		0x04: 16.503354,
		0x08: 7.1788233,
		0x09: 9.8845609,
		0x0c: 5.9407499,
		0x0e: 5.6410660,
		0x10: 10.189251,
		0x12: 5.7433447,
		0x26: 6.1352940,
	},
	{
		0x01: 7.3896294,
		0x04: 10.007508,
		0x06: 9.6358109,
		0x07: 20.735024,
		0x0b: 18.076530,
		0x14: 6.5933119,
		0x16: 12.430922,
		0x1f: 4.0492806,
		0x27: 5.3539120,
		0x2b: 5.7280705,
	},
	{
		0x02: 41.003322,
		0x03: 6.8645696,
		0x04: 1.7479089,
		0x05: 39.242906,
		0x06: 3.3917637,
		0x0b: 1.3552047,
		0x23: 2.0240525,
		0x26: 1.2591548,
		0x2c: 1.8699724,
		0x2d: 1.2411454,
	},
	{
		0x00: 8.3504489,
		0x01: 8.4257575,
		0x03: 11.836160,
		0x04: 8.9587857,
		0x06: 6.8061342,
		0x07: 17.120474,
		0x08: 17.930286,
		0x0d: 9.4458463,
		0x11: 6.8159146,
		0x1e: 4.3101931,
	},
	{
		0x01: 15.981607,
		0x07: 22.204617,
		0x09: 5.1806630,
		0x11: 6.5228198,
		0x13: 11.851058,
		0x14: 15.660049,
		0x17: 6.9422438,
		0x2e: 5.5286296,
		0x2f: 5.4338708,
		0x30: 4.6944419,
	},
	{
		0x01: 14.058542,
		0x06: 5.0073039,
		0x07: 8.0619811,
		0x09: 10.319706,
		0x11: 4.6337901,
		0x12: 14.518962,
		0x13: 5.3438361,
		0x14: 23.967752,
		0x27: 4.9314917,
		0x2b: 9.1566354,
	},
	{
		0x01: 5.0523654,
		0x06: 9.8527405,
		0x07: 1.4156669,
		0x08: 1.2666494,
		0x0b: 10.553458,
		0x14: 18.260513,
		0x16: 9.8694840,
		0x23: 4.3566710,
		0x31: 37.652052,
		0x32: 1.7203995,
	},
	{
		0x00: 21.021511,
		0x03: 33.575661,
		0x04: 6.0399343,
		0x0d: 2.8145695,
		0x10: 9.7744361,
		0x11: 7.9856097,
		0x16: 2.8046109,
		0x1e: 2.7647762,
		0x1f: 9.6661355,
		0x20: 3.5527561,
	},
	{
		0x03: 6.911655,
		0x04: 14.436984,
		0x06: 22.178904,
		0x0b: 5.9418093,
		0x14: 6.8538974,
		0x16: 6.6349000,
		0x27: 10.199023,
		0x2b: 12.319207,
		0x31: 6.9838519,
		0x33: 7.5397685,
	},
	{
		0x00: 7.3216357,
		0x03: 7.4039433,
		0x04: 5.0525646,
		0x07: 12.911070,
		0x08: 4.0555202,
		0x0b: 3.8422687,
		0x0c: 4.7373639,
		0x0d: 34.754387,
		0x1e: 15.723559,
		0x2e: 4.1976879,
	},
	{
		0x00: 16.785606,
		0x03: 34.049565,
		0x04: 2.5163271,
		0x08: 4.9686432,
		0x0a: 7.3958739,
		0x0d: 3.8510445,
		0x10: 6.0585615,
		0x11: 5.3899053,
		0x15: 7.2929372,
		0x34: 11.691536,
	},
	{
		0x00: 5.2212539,
		0x01: 11.033522,
		0x04: 6.5990325,
		0x06: 17.621261,
		0x07: 15.883981,
		0x09: 7.4234411,
		0x10: 4.9652732,
		0x15: 19.766982,
		0x26: 6.3750494,
		0x30: 5.1102035,
	},
	{
		0x00: 25.513354,
		0x01: 0.8516226,
		0x06: 3.6407059,
		0x0a: 21.840396,
		0x10: 4.3756047,
		0x11: 7.5532552,
		0x12: 1.1994901,
		0x15: 32.395448,
		0x34: 0.8662131,
		0x35: 1.7639109,
	},
	{
		0x00: 13.132114,
		0x01: 6.400206,
		0x02: 7.1264486,
		0x03: 4.4305949,
		0x07: 6.3559104,
		0x09: 13.727530,
		0x0e: 29.995364,
		0x14: 6.8318311,
		0x22: 8.2987381,
		0x36: 3.7012619,
	},
	{
		0x02: 18.893974,
		0x04: 6.6687834,
		0x05: 15.179119,
		0x07: 6.6939651,
		0x08: 7.8523255,
		0x0b: 13.271875,
		0x0c: 7.0454147,
		0x0d: 8.5519401,
		0x0e: 9.0665236,
		0x35: 6.7760795,
	},
	{
		0x00: 26.180321,
		0x01: 7.1145214,
		0x03: 4.3469354,
		0x06: 10.727655,
		0x07: 12.858287,
		0x09: 8.1407260,
		0x0a: 12.811726,
		0x10: 6.0287188,
		0x13: 5.0993612,
		0x14: 6.6917475,
	},
	{
		0x02: 8.3241807,
		0x06: 17.567878,
		0x09: 9.7736967,
		0x0d: 11.389237,
		0x17: 5.7418712,
		0x1e: 5.2144262,
		0x22: 5.9589793,
		0x24: 9.0316978,
		0x31: 14.620316,
		0x37: 12.377717,
	},
	{
		0x00: 9.8279912,
		0x03: 7.8881564,
		0x0d: 3.5765704,
		0x0f: 6.6530272,
		0x1a: 49.390013,
		0x1b: 4.2661211,
		0x38: 5.3042358,
		0x39: 5.0769114,
		0x3a: 4.0387967,
		0x3b: 3.9781769,
	},
	{
		0x00: 11.262873,
		0x03: 2.6221605,
		0x08: 3.9222847,
		0x0b: 13.424878,
		0x0c: 8.4654152,
		0x0f: 2.1693083,
		0x18: 4.1560149,
		0x1d: 43.970492,
		0x3c: 6.4787086,
		0x3d: 3.5278650,
	},
	{
		0x00: 20.560014,
		0x03: 3.2972784,
		0x08: 7.1441033,
		0x0b: 25.191905,
		0x0c: 11.243894,
		0x0f: 7.4842987,
		0x1b: 10.764131,
		0x39: 5.6001396,
		0x3e: 5.1203768,
		0x3f: 3.5938590,
	},
	{
		0x00: 16.983205,
		0x08: 8.3189452,
		0x0c: 4.9913671,
		0x0f: 10.940198,
		0x18: 5.3523780,
		0x19: 5.5564276,
		0x1c: 14.927013,
		0x3c: 5.8389578,
		0x40: 15.319416,
		0x41: 11.772092,
	},
	{
		0x00: 26.930964,
		0x08: 8.1681476,
		0x0b: 8.1852358,
		0x0c: 9.0567327,
		0x0d: 5.3144224,
		0x0f: 15.311005,
		0x39: 6.3397129,
		0x3c: 9.1250854,
		0x42: 6.2030075,
		0x43: 5.3656869,
	},
	{
		0x18: 70.461538,
		0x1a: 1.8241758,
		0x3c: 13.923077,
		0x3d: 3.8351648,
		0x44: 2.3736264,
		0x45: 1.8901099,
		0x46: 1.6263736,
		0x47: 1.4945055,
		0x48: 1.3076923,
		0x49: 1.2637363,
	},
}