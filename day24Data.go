package main

func compiledProgram(x int, y int, outputs []int) int {
	memory := [312]int{}
	changed := true
	for i := 0; i < 312; i++ {
		memory[i] = -1
	}
	memory[0] = (x >> 0) & 1
	memory[45] = (y >> 0) & 1
	memory[1] = (x >> 1) & 1
	memory[46] = (y >> 1) & 1
	memory[2] = (x >> 2) & 1
	memory[47] = (y >> 2) & 1
	memory[3] = (x >> 3) & 1
	memory[48] = (y >> 3) & 1
	memory[4] = (x >> 4) & 1
	memory[49] = (y >> 4) & 1
	memory[5] = (x >> 5) & 1
	memory[50] = (y >> 5) & 1
	memory[6] = (x >> 6) & 1
	memory[51] = (y >> 6) & 1
	memory[7] = (x >> 7) & 1
	memory[52] = (y >> 7) & 1
	memory[8] = (x >> 8) & 1
	memory[53] = (y >> 8) & 1
	memory[9] = (x >> 9) & 1
	memory[54] = (y >> 9) & 1
	memory[10] = (x >> 10) & 1
	memory[55] = (y >> 10) & 1
	memory[11] = (x >> 11) & 1
	memory[56] = (y >> 11) & 1
	memory[12] = (x >> 12) & 1
	memory[57] = (y >> 12) & 1
	memory[13] = (x >> 13) & 1
	memory[58] = (y >> 13) & 1
	memory[14] = (x >> 14) & 1
	memory[59] = (y >> 14) & 1
	memory[15] = (x >> 15) & 1
	memory[60] = (y >> 15) & 1
	memory[16] = (x >> 16) & 1
	memory[61] = (y >> 16) & 1
	memory[17] = (x >> 17) & 1
	memory[62] = (y >> 17) & 1
	memory[18] = (x >> 18) & 1
	memory[63] = (y >> 18) & 1
	memory[19] = (x >> 19) & 1
	memory[64] = (y >> 19) & 1
	memory[20] = (x >> 20) & 1
	memory[65] = (y >> 20) & 1
	memory[21] = (x >> 21) & 1
	memory[66] = (y >> 21) & 1
	memory[22] = (x >> 22) & 1
	memory[67] = (y >> 22) & 1
	memory[23] = (x >> 23) & 1
	memory[68] = (y >> 23) & 1
	memory[24] = (x >> 24) & 1
	memory[69] = (y >> 24) & 1
	memory[25] = (x >> 25) & 1
	memory[70] = (y >> 25) & 1
	memory[26] = (x >> 26) & 1
	memory[71] = (y >> 26) & 1
	memory[27] = (x >> 27) & 1
	memory[72] = (y >> 27) & 1
	memory[28] = (x >> 28) & 1
	memory[73] = (y >> 28) & 1
	memory[29] = (x >> 29) & 1
	memory[74] = (y >> 29) & 1
	memory[30] = (x >> 30) & 1
	memory[75] = (y >> 30) & 1
	memory[31] = (x >> 31) & 1
	memory[76] = (y >> 31) & 1
	memory[32] = (x >> 32) & 1
	memory[77] = (y >> 32) & 1
	memory[33] = (x >> 33) & 1
	memory[78] = (y >> 33) & 1
	memory[34] = (x >> 34) & 1
	memory[79] = (y >> 34) & 1
	memory[35] = (x >> 35) & 1
	memory[80] = (y >> 35) & 1
	memory[36] = (x >> 36) & 1
	memory[81] = (y >> 36) & 1
	memory[37] = (x >> 37) & 1
	memory[82] = (y >> 37) & 1
	memory[38] = (x >> 38) & 1
	memory[83] = (y >> 38) & 1
	memory[39] = (x >> 39) & 1
	memory[84] = (y >> 39) & 1
	memory[40] = (x >> 40) & 1
	memory[85] = (y >> 40) & 1
	memory[41] = (x >> 41) & 1
	memory[86] = (y >> 41) & 1
	memory[42] = (x >> 42) & 1
	memory[87] = (y >> 42) & 1
	memory[43] = (x >> 43) & 1
	memory[88] = (y >> 43) & 1
	memory[44] = (x >> 44) & 1
	memory[89] = (y >> 44) & 1
	for i := 0; i < 8 && changed; i++ {
		changed = false
		if memory[14] != -1 && memory[59] != -1 && memory[outputs[90]] == -1 {
			changed = true
			memory[outputs[90]] = memory[14] ^ memory[59]
		}
		if memory[39] != -1 && memory[84] != -1 && memory[outputs[91]] == -1 {
			changed = true
			memory[outputs[91]] = memory[39] ^ memory[84]
		}
		if memory[1] != -1 && memory[46] != -1 && memory[outputs[92]] == -1 {
			changed = true
			memory[outputs[92]] = memory[1] & memory[46]
		}
		if memory[47] != -1 && memory[2] != -1 && memory[outputs[93]] == -1 {
			changed = true
			memory[outputs[93]] = memory[47] & memory[2]
		}
		if memory[43] != -1 && memory[88] != -1 && memory[outputs[94]] == -1 {
			changed = true
			memory[outputs[94]] = memory[43] ^ memory[88]
		}
		if memory[25] != -1 && memory[70] != -1 && memory[outputs[95]] == -1 {
			changed = true
			memory[outputs[95]] = memory[25] & memory[70]
		}
		if memory[69] != -1 && memory[24] != -1 && memory[outputs[96]] == -1 {
			changed = true
			memory[outputs[96]] = memory[69] ^ memory[24]
		}
		if memory[36] != -1 && memory[81] != -1 && memory[outputs[97]] == -1 {
			changed = true
			memory[outputs[97]] = memory[36] & memory[81]
		}
		if memory[38] != -1 && memory[83] != -1 && memory[outputs[98]] == -1 {
			changed = true
			memory[outputs[98]] = memory[38] & memory[83]
		}
		if memory[32] != -1 && memory[77] != -1 && memory[outputs[99]] == -1 {
			changed = true
			memory[outputs[99]] = memory[32] & memory[77]
		}
		if memory[66] != -1 && memory[21] != -1 && memory[outputs[100]] == -1 {
			changed = true
			memory[outputs[100]] = memory[66] ^ memory[21]
		}
		if memory[30] != -1 && memory[75] != -1 && memory[outputs[101]] == -1 {
			changed = true
			memory[outputs[101]] = memory[30] & memory[75]
		}
		if memory[67] != -1 && memory[22] != -1 && memory[outputs[102]] == -1 {
			changed = true
			memory[outputs[102]] = memory[67] ^ memory[22]
		}
		if memory[68] != -1 && memory[23] != -1 && memory[outputs[103]] == -1 {
			changed = true
			memory[outputs[103]] = memory[68] & memory[23]
		}
		if memory[33] != -1 && memory[78] != -1 && memory[outputs[104]] == -1 {
			changed = true
			memory[outputs[104]] = memory[33] & memory[78]
		}
		if memory[48] != -1 && memory[3] != -1 && memory[outputs[105]] == -1 {
			changed = true
			memory[outputs[105]] = memory[48] & memory[3]
		}
		if memory[54] != -1 && memory[9] != -1 && memory[outputs[106]] == -1 {
			changed = true
			memory[outputs[106]] = memory[54] & memory[9]
		}
		if memory[34] != -1 && memory[79] != -1 && memory[outputs[107]] == -1 {
			changed = true
			memory[outputs[107]] = memory[34] ^ memory[79]
		}
		if memory[10] != -1 && memory[55] != -1 && memory[outputs[108]] == -1 {
			changed = true
			memory[outputs[108]] = memory[10] ^ memory[55]
		}
		if memory[42] != -1 && memory[87] != -1 && memory[outputs[109]] == -1 {
			changed = true
			memory[outputs[109]] = memory[42] ^ memory[87]
		}
		if memory[24] != -1 && memory[69] != -1 && memory[outputs[110]] == -1 {
			changed = true
			memory[outputs[110]] = memory[24] & memory[69]
		}
		if memory[53] != -1 && memory[8] != -1 && memory[outputs[111]] == -1 {
			changed = true
			memory[outputs[111]] = memory[53] ^ memory[8]
		}
		if memory[27] != -1 && memory[72] != -1 && memory[outputs[112]] == -1 {
			changed = true
			memory[outputs[112]] = memory[27] ^ memory[72]
		}
		if memory[88] != -1 && memory[43] != -1 && memory[outputs[113]] == -1 {
			changed = true
			memory[outputs[113]] = memory[88] & memory[43]
		}
		if memory[61] != -1 && memory[16] != -1 && memory[outputs[114]] == -1 {
			changed = true
			memory[outputs[114]] = memory[61] ^ memory[16]
		}
		if memory[14] != -1 && memory[59] != -1 && memory[outputs[115]] == -1 {
			changed = true
			memory[outputs[115]] = memory[14] & memory[59]
		}
		if memory[18] != -1 && memory[63] != -1 && memory[outputs[116]] == -1 {
			changed = true
			memory[outputs[116]] = memory[18] ^ memory[63]
		}
		if memory[52] != -1 && memory[7] != -1 && memory[outputs[117]] == -1 {
			changed = true
			memory[outputs[117]] = memory[52] & memory[7]
		}
		if memory[29] != -1 && memory[74] != -1 && memory[outputs[118]] == -1 {
			changed = true
			memory[outputs[118]] = memory[29] ^ memory[74]
		}
		if memory[44] != -1 && memory[89] != -1 && memory[outputs[119]] == -1 {
			changed = true
			memory[outputs[119]] = memory[44] & memory[89]
		}
		if memory[64] != -1 && memory[19] != -1 && memory[outputs[120]] == -1 {
			changed = true
			memory[outputs[120]] = memory[64] & memory[19]
		}
		if memory[82] != -1 && memory[37] != -1 && memory[outputs[121]] == -1 {
			changed = true
			memory[outputs[121]] = memory[82] ^ memory[37]
		}
		if memory[31] != -1 && memory[76] != -1 && memory[outputs[122]] == -1 {
			changed = true
			memory[outputs[122]] = memory[31] & memory[76]
		}
		if memory[71] != -1 && memory[26] != -1 && memory[outputs[123]] == -1 {
			changed = true
			memory[outputs[123]] = memory[71] & memory[26]
		}
		if memory[85] != -1 && memory[40] != -1 && memory[outputs[124]] == -1 {
			changed = true
			memory[outputs[124]] = memory[85] ^ memory[40]
		}
		if memory[9] != -1 && memory[54] != -1 && memory[outputs[125]] == -1 {
			changed = true
			memory[outputs[125]] = memory[9] ^ memory[54]
		}
		if memory[80] != -1 && memory[35] != -1 && memory[outputs[126]] == -1 {
			changed = true
			memory[outputs[126]] = memory[80] & memory[35]
		}
		if memory[4] != -1 && memory[49] != -1 && memory[outputs[127]] == -1 {
			changed = true
			memory[outputs[127]] = memory[4] ^ memory[49]
		}
		if memory[22] != -1 && memory[67] != -1 && memory[outputs[128]] == -1 {
			changed = true
			memory[outputs[128]] = memory[22] & memory[67]
		}
		if memory[16] != -1 && memory[61] != -1 && memory[outputs[129]] == -1 {
			changed = true
			memory[outputs[129]] = memory[16] & memory[61]
		}
		if memory[23] != -1 && memory[68] != -1 && memory[outputs[130]] == -1 {
			changed = true
			memory[outputs[130]] = memory[23] ^ memory[68]
		}
		if memory[50] != -1 && memory[5] != -1 && memory[outputs[131]] == -1 {
			changed = true
			memory[outputs[131]] = memory[50] ^ memory[5]
		}
		if memory[29] != -1 && memory[74] != -1 && memory[outputs[132]] == -1 {
			changed = true
			memory[outputs[132]] = memory[29] & memory[74]
		}
		if memory[26] != -1 && memory[71] != -1 && memory[outputs[133]] == -1 {
			changed = true
			memory[outputs[133]] = memory[26] ^ memory[71]
		}
		if memory[19] != -1 && memory[64] != -1 && memory[outputs[134]] == -1 {
			changed = true
			memory[outputs[134]] = memory[19] ^ memory[64]
		}
		if memory[20] != -1 && memory[65] != -1 && memory[outputs[135]] == -1 {
			changed = true
			memory[outputs[135]] = memory[20] & memory[65]
		}
		if memory[30] != -1 && memory[75] != -1 && memory[outputs[136]] == -1 {
			changed = true
			memory[outputs[136]] = memory[30] ^ memory[75]
		}
		if memory[32] != -1 && memory[77] != -1 && memory[outputs[137]] == -1 {
			changed = true
			memory[outputs[137]] = memory[32] ^ memory[77]
		}
		if memory[56] != -1 && memory[11] != -1 && memory[outputs[138]] == -1 {
			changed = true
			memory[outputs[138]] = memory[56] ^ memory[11]
		}
		if memory[62] != -1 && memory[17] != -1 && memory[outputs[139]] == -1 {
			changed = true
			memory[outputs[139]] = memory[62] ^ memory[17]
		}
		if memory[20] != -1 && memory[65] != -1 && memory[outputs[140]] == -1 {
			changed = true
			memory[outputs[140]] = memory[20] ^ memory[65]
		}
		if memory[52] != -1 && memory[7] != -1 && memory[outputs[141]] == -1 {
			changed = true
			memory[outputs[141]] = memory[52] ^ memory[7]
		}
		if memory[11] != -1 && memory[56] != -1 && memory[outputs[142]] == -1 {
			changed = true
			memory[outputs[142]] = memory[11] & memory[56]
		}
		if memory[45] != -1 && memory[0] != -1 && memory[outputs[143]] == -1 {
			changed = true
			memory[outputs[143]] = memory[45] & memory[0]
		}
		if memory[12] != -1 && memory[57] != -1 && memory[outputs[144]] == -1 {
			changed = true
			memory[outputs[144]] = memory[12] & memory[57]
		}
		if memory[76] != -1 && memory[31] != -1 && memory[outputs[145]] == -1 {
			changed = true
			memory[outputs[145]] = memory[76] ^ memory[31]
		}
		if memory[45] != -1 && memory[0] != -1 && memory[outputs[146]] == -1 {
			changed = true
			memory[outputs[146]] = memory[45] ^ memory[0]
		}
		if memory[85] != -1 && memory[40] != -1 && memory[outputs[147]] == -1 {
			changed = true
			memory[outputs[147]] = memory[85] & memory[40]
		}
		if memory[51] != -1 && memory[6] != -1 && memory[outputs[148]] == -1 {
			changed = true
			memory[outputs[148]] = memory[51] & memory[6]
		}
		if memory[17] != -1 && memory[62] != -1 && memory[outputs[149]] == -1 {
			changed = true
			memory[outputs[149]] = memory[17] & memory[62]
		}
		if memory[58] != -1 && memory[13] != -1 && memory[outputs[150]] == -1 {
			changed = true
			memory[outputs[150]] = memory[58] & memory[13]
		}
		if memory[80] != -1 && memory[35] != -1 && memory[outputs[151]] == -1 {
			changed = true
			memory[outputs[151]] = memory[80] ^ memory[35]
		}
		if memory[10] != -1 && memory[55] != -1 && memory[outputs[152]] == -1 {
			changed = true
			memory[outputs[152]] = memory[10] & memory[55]
		}
		if memory[1] != -1 && memory[46] != -1 && memory[outputs[153]] == -1 {
			changed = true
			memory[outputs[153]] = memory[1] ^ memory[46]
		}
		if memory[73] != -1 && memory[28] != -1 && memory[outputs[154]] == -1 {
			changed = true
			memory[outputs[154]] = memory[73] ^ memory[28]
		}
		if memory[73] != -1 && memory[28] != -1 && memory[outputs[155]] == -1 {
			changed = true
			memory[outputs[155]] = memory[73] & memory[28]
		}
		if memory[153] != -1 && memory[143] != -1 && memory[outputs[156]] == -1 {
			changed = true
			memory[outputs[156]] = memory[153] ^ memory[143]
		}
		if memory[72] != -1 && memory[27] != -1 && memory[outputs[157]] == -1 {
			changed = true
			memory[outputs[157]] = memory[72] & memory[27]
		}
		if memory[8] != -1 && memory[53] != -1 && memory[outputs[158]] == -1 {
			changed = true
			memory[outputs[158]] = memory[8] & memory[53]
		}
		if memory[82] != -1 && memory[37] != -1 && memory[outputs[159]] == -1 {
			changed = true
			memory[outputs[159]] = memory[82] & memory[37]
		}
		if memory[86] != -1 && memory[41] != -1 && memory[outputs[160]] == -1 {
			changed = true
			memory[outputs[160]] = memory[86] ^ memory[41]
		}
		if memory[66] != -1 && memory[21] != -1 && memory[outputs[161]] == -1 {
			changed = true
			memory[outputs[161]] = memory[66] & memory[21]
		}
		if memory[36] != -1 && memory[81] != -1 && memory[outputs[162]] == -1 {
			changed = true
			memory[outputs[162]] = memory[36] ^ memory[81]
		}
		if memory[15] != -1 && memory[60] != -1 && memory[outputs[163]] == -1 {
			changed = true
			memory[outputs[163]] = memory[15] ^ memory[60]
		}
		if memory[38] != -1 && memory[83] != -1 && memory[outputs[164]] == -1 {
			changed = true
			memory[outputs[164]] = memory[38] ^ memory[83]
		}
		if memory[5] != -1 && memory[50] != -1 && memory[outputs[165]] == -1 {
			changed = true
			memory[outputs[165]] = memory[5] & memory[50]
		}
		if memory[79] != -1 && memory[34] != -1 && memory[outputs[166]] == -1 {
			changed = true
			memory[outputs[166]] = memory[79] & memory[34]
		}
		if memory[13] != -1 && memory[58] != -1 && memory[outputs[167]] == -1 {
			changed = true
			memory[outputs[167]] = memory[13] ^ memory[58]
		}
		if memory[84] != -1 && memory[39] != -1 && memory[outputs[168]] == -1 {
			changed = true
			memory[outputs[168]] = memory[84] & memory[39]
		}
		if memory[18] != -1 && memory[63] != -1 && memory[outputs[169]] == -1 {
			changed = true
			memory[outputs[169]] = memory[18] & memory[63]
		}
		if memory[4] != -1 && memory[49] != -1 && memory[outputs[170]] == -1 {
			changed = true
			memory[outputs[170]] = memory[4] & memory[49]
		}
		if memory[42] != -1 && memory[87] != -1 && memory[outputs[171]] == -1 {
			changed = true
			memory[outputs[171]] = memory[42] & memory[87]
		}
		if memory[3] != -1 && memory[48] != -1 && memory[outputs[172]] == -1 {
			changed = true
			memory[outputs[172]] = memory[3] ^ memory[48]
		}
		if memory[25] != -1 && memory[70] != -1 && memory[outputs[173]] == -1 {
			changed = true
			memory[outputs[173]] = memory[25] ^ memory[70]
		}
		if memory[44] != -1 && memory[89] != -1 && memory[outputs[174]] == -1 {
			changed = true
			memory[outputs[174]] = memory[44] ^ memory[89]
		}
		if memory[2] != -1 && memory[47] != -1 && memory[outputs[175]] == -1 {
			changed = true
			memory[outputs[175]] = memory[2] ^ memory[47]
		}
		if memory[12] != -1 && memory[57] != -1 && memory[outputs[176]] == -1 {
			changed = true
			memory[outputs[176]] = memory[12] ^ memory[57]
		}
		if memory[41] != -1 && memory[86] != -1 && memory[outputs[177]] == -1 {
			changed = true
			memory[outputs[177]] = memory[41] & memory[86]
		}
		if memory[60] != -1 && memory[15] != -1 && memory[outputs[178]] == -1 {
			changed = true
			memory[outputs[178]] = memory[60] & memory[15]
		}
		if memory[6] != -1 && memory[51] != -1 && memory[outputs[179]] == -1 {
			changed = true
			memory[outputs[179]] = memory[6] ^ memory[51]
		}
		if memory[78] != -1 && memory[33] != -1 && memory[outputs[180]] == -1 {
			changed = true
			memory[outputs[180]] = memory[78] ^ memory[33]
		}
		if memory[153] != -1 && memory[143] != -1 && memory[outputs[181]] == -1 {
			changed = true
			memory[outputs[181]] = memory[153] & memory[143]
		}
		if memory[181] != -1 && memory[92] != -1 && memory[outputs[182]] == -1 {
			changed = true
			memory[outputs[182]] = memory[181] | memory[92]
		}
		if memory[182] != -1 && memory[175] != -1 && memory[outputs[183]] == -1 {
			changed = true
			memory[outputs[183]] = memory[182] & memory[175]
		}
		if memory[175] != -1 && memory[182] != -1 && memory[outputs[184]] == -1 {
			changed = true
			memory[outputs[184]] = memory[175] ^ memory[182]
		}
		if memory[93] != -1 && memory[183] != -1 && memory[outputs[185]] == -1 {
			changed = true
			memory[outputs[185]] = memory[93] | memory[183]
		}
		if memory[185] != -1 && memory[172] != -1 && memory[outputs[186]] == -1 {
			changed = true
			memory[outputs[186]] = memory[185] & memory[172]
		}
		if memory[185] != -1 && memory[172] != -1 && memory[outputs[187]] == -1 {
			changed = true
			memory[outputs[187]] = memory[185] ^ memory[172]
		}
		if memory[105] != -1 && memory[186] != -1 && memory[outputs[188]] == -1 {
			changed = true
			memory[outputs[188]] = memory[105] | memory[186]
		}
		if memory[127] != -1 && memory[188] != -1 && memory[outputs[189]] == -1 {
			changed = true
			memory[outputs[189]] = memory[127] & memory[188]
		}
		if memory[127] != -1 && memory[188] != -1 && memory[outputs[190]] == -1 {
			changed = true
			memory[outputs[190]] = memory[127] ^ memory[188]
		}
		if memory[170] != -1 && memory[189] != -1 && memory[outputs[191]] == -1 {
			changed = true
			memory[outputs[191]] = memory[170] | memory[189]
		}
		if memory[191] != -1 && memory[131] != -1 && memory[outputs[192]] == -1 {
			changed = true
			memory[outputs[192]] = memory[191] ^ memory[131]
		}
		if memory[131] != -1 && memory[191] != -1 && memory[outputs[193]] == -1 {
			changed = true
			memory[outputs[193]] = memory[131] & memory[191]
		}
		if memory[165] != -1 && memory[193] != -1 && memory[outputs[194]] == -1 {
			changed = true
			memory[outputs[194]] = memory[165] | memory[193]
		}
		if memory[194] != -1 && memory[179] != -1 && memory[outputs[195]] == -1 {
			changed = true
			memory[outputs[195]] = memory[194] & memory[179]
		}
		if memory[179] != -1 && memory[194] != -1 && memory[outputs[196]] == -1 {
			changed = true
			memory[outputs[196]] = memory[179] ^ memory[194]
		}
		if memory[195] != -1 && memory[148] != -1 && memory[outputs[197]] == -1 {
			changed = true
			memory[outputs[197]] = memory[195] | memory[148]
		}
		if memory[197] != -1 && memory[141] != -1 && memory[outputs[198]] == -1 {
			changed = true
			memory[outputs[198]] = memory[197] ^ memory[141]
		}
		if memory[141] != -1 && memory[197] != -1 && memory[outputs[199]] == -1 {
			changed = true
			memory[outputs[199]] = memory[141] & memory[197]
		}
		if memory[199] != -1 && memory[117] != -1 && memory[outputs[200]] == -1 {
			changed = true
			memory[outputs[200]] = memory[199] | memory[117]
		}
		if memory[200] != -1 && memory[111] != -1 && memory[outputs[201]] == -1 {
			changed = true
			memory[outputs[201]] = memory[200] ^ memory[111]
		}
		if memory[111] != -1 && memory[200] != -1 && memory[outputs[202]] == -1 {
			changed = true
			memory[outputs[202]] = memory[111] & memory[200]
		}
		if memory[158] != -1 && memory[202] != -1 && memory[outputs[203]] == -1 {
			changed = true
			memory[outputs[203]] = memory[158] | memory[202]
		}
		if memory[125] != -1 && memory[203] != -1 && memory[outputs[204]] == -1 {
			changed = true
			memory[outputs[204]] = memory[125] & memory[203]
		}
		if memory[203] != -1 && memory[125] != -1 && memory[outputs[205]] == -1 {
			changed = true
			memory[outputs[205]] = memory[203] ^ memory[125]
		}
		if memory[106] != -1 && memory[204] != -1 && memory[outputs[206]] == -1 {
			changed = true
			memory[outputs[206]] = memory[106] | memory[204]
		}
		if memory[206] != -1 && memory[108] != -1 && memory[outputs[207]] == -1 {
			changed = true
			memory[outputs[207]] = memory[206] ^ memory[108]
		}
		if memory[108] != -1 && memory[206] != -1 && memory[outputs[208]] == -1 {
			changed = true
			memory[outputs[208]] = memory[108] & memory[206]
		}
		if memory[208] != -1 && memory[152] != -1 && memory[outputs[209]] == -1 {
			changed = true
			memory[outputs[209]] = memory[208] | memory[152]
		}
		if memory[209] != -1 && memory[138] != -1 && memory[outputs[210]] == -1 {
			changed = true
			memory[outputs[210]] = memory[209] ^ memory[138]
		}
		if memory[138] != -1 && memory[209] != -1 && memory[outputs[211]] == -1 {
			changed = true
			memory[outputs[211]] = memory[138] & memory[209]
		}
		if memory[211] != -1 && memory[142] != -1 && memory[outputs[212]] == -1 {
			changed = true
			memory[outputs[212]] = memory[211] | memory[142]
		}
		if memory[176] != -1 && memory[212] != -1 && memory[outputs[213]] == -1 {
			changed = true
			memory[outputs[213]] = memory[176] ^ memory[212]
		}
		if memory[212] != -1 && memory[176] != -1 && memory[outputs[214]] == -1 {
			changed = true
			memory[outputs[214]] = memory[212] & memory[176]
		}
		if memory[214] != -1 && memory[144] != -1 && memory[outputs[215]] == -1 {
			changed = true
			memory[outputs[215]] = memory[214] | memory[144]
		}
		if memory[215] != -1 && memory[167] != -1 && memory[outputs[216]] == -1 {
			changed = true
			memory[outputs[216]] = memory[215] ^ memory[167]
		}
		if memory[215] != -1 && memory[167] != -1 && memory[outputs[217]] == -1 {
			changed = true
			memory[outputs[217]] = memory[215] & memory[167]
		}
		if memory[217] != -1 && memory[150] != -1 && memory[outputs[218]] == -1 {
			changed = true
			memory[outputs[218]] = memory[217] | memory[150]
		}
		if memory[90] != -1 && memory[218] != -1 && memory[outputs[219]] == -1 {
			changed = true
			memory[outputs[219]] = memory[90] & memory[218]
		}
		if memory[219] != -1 && memory[115] != -1 && memory[outputs[220]] == -1 {
			changed = true
			memory[outputs[220]] = memory[219] | memory[115]
		}
		if memory[90] != -1 && memory[218] != -1 && memory[outputs[221]] == -1 {
			changed = true
			memory[outputs[221]] = memory[90] ^ memory[218]
		}
		if memory[220] != -1 && memory[163] != -1 && memory[outputs[222]] == -1 {
			changed = true
			memory[outputs[222]] = memory[220] ^ memory[163]
		}
		if memory[114] != -1 && memory[222] != -1 && memory[outputs[223]] == -1 {
			changed = true
			memory[outputs[223]] = memory[114] & memory[222]
		}
		if memory[163] != -1 && memory[220] != -1 && memory[outputs[224]] == -1 {
			changed = true
			memory[outputs[224]] = memory[163] & memory[220]
		}
		if memory[223] != -1 && memory[129] != -1 && memory[outputs[225]] == -1 {
			changed = true
			memory[outputs[225]] = memory[223] | memory[129]
		}
		if memory[114] != -1 && memory[222] != -1 && memory[outputs[226]] == -1 {
			changed = true
			memory[outputs[226]] = memory[114] ^ memory[222]
		}
		if memory[225] != -1 && memory[139] != -1 && memory[outputs[227]] == -1 {
			changed = true
			memory[outputs[227]] = memory[225] & memory[139]
		}
		if memory[139] != -1 && memory[225] != -1 && memory[outputs[228]] == -1 {
			changed = true
			memory[outputs[228]] = memory[139] ^ memory[225]
		}
		if memory[224] != -1 && memory[178] != -1 && memory[outputs[229]] == -1 {
			changed = true
			memory[outputs[229]] = memory[224] | memory[178]
		}
		if memory[227] != -1 && memory[149] != -1 && memory[outputs[230]] == -1 {
			changed = true
			memory[outputs[230]] = memory[227] | memory[149]
		}
		if memory[230] != -1 && memory[116] != -1 && memory[outputs[231]] == -1 {
			changed = true
			memory[outputs[231]] = memory[230] ^ memory[116]
		}
		if memory[230] != -1 && memory[116] != -1 && memory[outputs[232]] == -1 {
			changed = true
			memory[outputs[232]] = memory[230] & memory[116]
		}
		if memory[169] != -1 && memory[232] != -1 && memory[outputs[233]] == -1 {
			changed = true
			memory[outputs[233]] = memory[169] | memory[232]
		}
		if memory[233] != -1 && memory[134] != -1 && memory[outputs[234]] == -1 {
			changed = true
			memory[outputs[234]] = memory[233] ^ memory[134]
		}
		if memory[233] != -1 && memory[134] != -1 && memory[outputs[235]] == -1 {
			changed = true
			memory[outputs[235]] = memory[233] & memory[134]
		}
		if memory[235] != -1 && memory[120] != -1 && memory[outputs[236]] == -1 {
			changed = true
			memory[outputs[236]] = memory[235] | memory[120]
		}
		if memory[236] != -1 && memory[140] != -1 && memory[outputs[237]] == -1 {
			changed = true
			memory[outputs[237]] = memory[236] & memory[140]
		}
		if memory[236] != -1 && memory[140] != -1 && memory[outputs[238]] == -1 {
			changed = true
			memory[outputs[238]] = memory[236] ^ memory[140]
		}
		if memory[237] != -1 && memory[238] != -1 && memory[outputs[239]] == -1 {
			changed = true
			memory[outputs[239]] = memory[237] | memory[238]
		}
		if memory[239] != -1 && memory[100] != -1 && memory[outputs[240]] == -1 {
			changed = true
			memory[outputs[240]] = memory[239] ^ memory[100]
		}
		if memory[239] != -1 && memory[100] != -1 && memory[outputs[241]] == -1 {
			changed = true
			memory[outputs[241]] = memory[239] & memory[100]
		}
		if memory[241] != -1 && memory[161] != -1 && memory[outputs[242]] == -1 {
			changed = true
			memory[outputs[242]] = memory[241] | memory[161]
		}
		if memory[102] != -1 && memory[242] != -1 && memory[outputs[243]] == -1 {
			changed = true
			memory[outputs[243]] = memory[102] & memory[242]
		}
		if memory[242] != -1 && memory[102] != -1 && memory[outputs[244]] == -1 {
			changed = true
			memory[outputs[244]] = memory[242] ^ memory[102]
		}
		if memory[243] != -1 && memory[128] != -1 && memory[outputs[245]] == -1 {
			changed = true
			memory[outputs[245]] = memory[243] | memory[128]
		}
		if memory[130] != -1 && memory[245] != -1 && memory[outputs[246]] == -1 {
			changed = true
			memory[outputs[246]] = memory[130] & memory[245]
		}
		if memory[245] != -1 && memory[130] != -1 && memory[outputs[247]] == -1 {
			changed = true
			memory[outputs[247]] = memory[245] ^ memory[130]
		}
		if memory[246] != -1 && memory[103] != -1 && memory[outputs[248]] == -1 {
			changed = true
			memory[outputs[248]] = memory[246] | memory[103]
		}
		if memory[96] != -1 && memory[248] != -1 && memory[outputs[249]] == -1 {
			changed = true
			memory[outputs[249]] = memory[96] ^ memory[248]
		}
		if memory[96] != -1 && memory[248] != -1 && memory[outputs[250]] == -1 {
			changed = true
			memory[outputs[250]] = memory[96] & memory[248]
		}
		if memory[250] != -1 && memory[110] != -1 && memory[outputs[251]] == -1 {
			changed = true
			memory[outputs[251]] = memory[250] | memory[110]
		}
		if memory[173] != -1 && memory[251] != -1 && memory[outputs[252]] == -1 {
			changed = true
			memory[outputs[252]] = memory[173] ^ memory[251]
		}
		if memory[251] != -1 && memory[173] != -1 && memory[outputs[253]] == -1 {
			changed = true
			memory[outputs[253]] = memory[251] & memory[173]
		}
		if memory[253] != -1 && memory[95] != -1 && memory[outputs[254]] == -1 {
			changed = true
			memory[outputs[254]] = memory[253] | memory[95]
		}
		if memory[133] != -1 && memory[254] != -1 && memory[outputs[255]] == -1 {
			changed = true
			memory[outputs[255]] = memory[133] & memory[254]
		}
		if memory[255] != -1 && memory[123] != -1 && memory[outputs[256]] == -1 {
			changed = true
			memory[outputs[256]] = memory[255] | memory[123]
		}
		if memory[157] != -1 && memory[256] != -1 && memory[outputs[257]] == -1 {
			changed = true
			memory[outputs[257]] = memory[157] & memory[256]
		}
		if memory[254] != -1 && memory[133] != -1 && memory[outputs[258]] == -1 {
			changed = true
			memory[outputs[258]] = memory[254] ^ memory[133]
		}
		if memory[256] != -1 && memory[157] != -1 && memory[outputs[259]] == -1 {
			changed = true
			memory[outputs[259]] = memory[256] ^ memory[157]
		}
		if memory[112] != -1 && memory[257] != -1 && memory[outputs[260]] == -1 {
			changed = true
			memory[outputs[260]] = memory[112] | memory[257]
		}
		if memory[260] != -1 && memory[154] != -1 && memory[outputs[261]] == -1 {
			changed = true
			memory[outputs[261]] = memory[260] ^ memory[154]
		}
		if memory[260] != -1 && memory[154] != -1 && memory[outputs[262]] == -1 {
			changed = true
			memory[outputs[262]] = memory[260] & memory[154]
		}
		if memory[262] != -1 && memory[155] != -1 && memory[outputs[263]] == -1 {
			changed = true
			memory[outputs[263]] = memory[262] | memory[155]
		}
		if memory[118] != -1 && memory[263] != -1 && memory[outputs[264]] == -1 {
			changed = true
			memory[outputs[264]] = memory[118] ^ memory[263]
		}
		if memory[263] != -1 && memory[118] != -1 && memory[outputs[265]] == -1 {
			changed = true
			memory[outputs[265]] = memory[263] & memory[118]
		}
		if memory[265] != -1 && memory[132] != -1 && memory[outputs[266]] == -1 {
			changed = true
			memory[outputs[266]] = memory[265] | memory[132]
		}
		if memory[266] != -1 && memory[136] != -1 && memory[outputs[267]] == -1 {
			changed = true
			memory[outputs[267]] = memory[266] ^ memory[136]
		}
		if memory[266] != -1 && memory[136] != -1 && memory[outputs[268]] == -1 {
			changed = true
			memory[outputs[268]] = memory[266] & memory[136]
		}
		if memory[268] != -1 && memory[101] != -1 && memory[outputs[269]] == -1 {
			changed = true
			memory[outputs[269]] = memory[268] | memory[101]
		}
		if memory[269] != -1 && memory[145] != -1 && memory[outputs[270]] == -1 {
			changed = true
			memory[outputs[270]] = memory[269] ^ memory[145]
		}
		if memory[269] != -1 && memory[145] != -1 && memory[outputs[271]] == -1 {
			changed = true
			memory[outputs[271]] = memory[269] & memory[145]
		}
		if memory[122] != -1 && memory[271] != -1 && memory[outputs[272]] == -1 {
			changed = true
			memory[outputs[272]] = memory[122] | memory[271]
		}
		if memory[137] != -1 && memory[272] != -1 && memory[outputs[273]] == -1 {
			changed = true
			memory[outputs[273]] = memory[137] & memory[272]
		}
		if memory[273] != -1 && memory[99] != -1 && memory[outputs[274]] == -1 {
			changed = true
			memory[outputs[274]] = memory[273] | memory[99]
		}
		if memory[272] != -1 && memory[137] != -1 && memory[outputs[275]] == -1 {
			changed = true
			memory[outputs[275]] = memory[272] ^ memory[137]
		}
		if memory[274] != -1 && memory[180] != -1 && memory[outputs[276]] == -1 {
			changed = true
			memory[outputs[276]] = memory[274] ^ memory[180]
		}
		if memory[180] != -1 && memory[274] != -1 && memory[outputs[277]] == -1 {
			changed = true
			memory[outputs[277]] = memory[180] & memory[274]
		}
		if memory[104] != -1 && memory[277] != -1 && memory[outputs[278]] == -1 {
			changed = true
			memory[outputs[278]] = memory[104] | memory[277]
		}
		if memory[107] != -1 && memory[278] != -1 && memory[outputs[279]] == -1 {
			changed = true
			memory[outputs[279]] = memory[107] & memory[278]
		}
		if memory[107] != -1 && memory[278] != -1 && memory[outputs[280]] == -1 {
			changed = true
			memory[outputs[280]] = memory[107] ^ memory[278]
		}
		if memory[279] != -1 && memory[166] != -1 && memory[outputs[281]] == -1 {
			changed = true
			memory[outputs[281]] = memory[279] | memory[166]
		}
		if memory[151] != -1 && memory[281] != -1 && memory[outputs[282]] == -1 {
			changed = true
			memory[outputs[282]] = memory[151] ^ memory[281]
		}
		if memory[281] != -1 && memory[151] != -1 && memory[outputs[283]] == -1 {
			changed = true
			memory[outputs[283]] = memory[281] & memory[151]
		}
		if memory[126] != -1 && memory[283] != -1 && memory[outputs[284]] == -1 {
			changed = true
			memory[outputs[284]] = memory[126] | memory[283]
		}
		if memory[284] != -1 && memory[162] != -1 && memory[outputs[285]] == -1 {
			changed = true
			memory[outputs[285]] = memory[284] & memory[162]
		}
		if memory[284] != -1 && memory[162] != -1 && memory[outputs[286]] == -1 {
			changed = true
			memory[outputs[286]] = memory[284] ^ memory[162]
		}
		if memory[97] != -1 && memory[285] != -1 && memory[outputs[287]] == -1 {
			changed = true
			memory[outputs[287]] = memory[97] | memory[285]
		}
		if memory[287] != -1 && memory[121] != -1 && memory[outputs[288]] == -1 {
			changed = true
			memory[outputs[288]] = memory[287] ^ memory[121]
		}
		if memory[121] != -1 && memory[287] != -1 && memory[outputs[289]] == -1 {
			changed = true
			memory[outputs[289]] = memory[121] & memory[287]
		}
		if memory[288] != -1 && memory[159] != -1 && memory[outputs[290]] == -1 {
			changed = true
			memory[outputs[290]] = memory[288] | memory[159]
		}
		if memory[290] != -1 && memory[164] != -1 && memory[outputs[291]] == -1 {
			changed = true
			memory[outputs[291]] = memory[290] & memory[164]
		}
		if memory[164] != -1 && memory[290] != -1 && memory[outputs[292]] == -1 {
			changed = true
			memory[outputs[292]] = memory[164] ^ memory[290]
		}
		if memory[98] != -1 && memory[291] != -1 && memory[outputs[293]] == -1 {
			changed = true
			memory[outputs[293]] = memory[98] | memory[291]
		}
		if memory[293] != -1 && memory[91] != -1 && memory[outputs[294]] == -1 {
			changed = true
			memory[outputs[294]] = memory[293] & memory[91]
		}
		if memory[294] != -1 && memory[168] != -1 && memory[outputs[295]] == -1 {
			changed = true
			memory[outputs[295]] = memory[294] | memory[168]
		}
		if memory[293] != -1 && memory[91] != -1 && memory[outputs[296]] == -1 {
			changed = true
			memory[outputs[296]] = memory[293] ^ memory[91]
		}
		if memory[295] != -1 && memory[124] != -1 && memory[outputs[297]] == -1 {
			changed = true
			memory[outputs[297]] = memory[295] & memory[124]
		}
		if memory[295] != -1 && memory[124] != -1 && memory[outputs[298]] == -1 {
			changed = true
			memory[outputs[298]] = memory[295] ^ memory[124]
		}
		if memory[147] != -1 && memory[297] != -1 && memory[outputs[299]] == -1 {
			changed = true
			memory[outputs[299]] = memory[147] | memory[297]
		}
		if memory[299] != -1 && memory[160] != -1 && memory[outputs[300]] == -1 {
			changed = true
			memory[outputs[300]] = memory[299] ^ memory[160]
		}
		if memory[299] != -1 && memory[160] != -1 && memory[outputs[301]] == -1 {
			changed = true
			memory[outputs[301]] = memory[299] & memory[160]
		}
		if memory[301] != -1 && memory[177] != -1 && memory[outputs[302]] == -1 {
			changed = true
			memory[outputs[302]] = memory[301] | memory[177]
		}
		if memory[109] != -1 && memory[302] != -1 && memory[outputs[303]] == -1 {
			changed = true
			memory[outputs[303]] = memory[109] & memory[302]
		}
		if memory[303] != -1 && memory[171] != -1 && memory[outputs[304]] == -1 {
			changed = true
			memory[outputs[304]] = memory[303] | memory[171]
		}
		if memory[304] != -1 && memory[94] != -1 && memory[outputs[305]] == -1 {
			changed = true
			memory[outputs[305]] = memory[304] & memory[94]
		}
		if memory[109] != -1 && memory[302] != -1 && memory[outputs[306]] == -1 {
			changed = true
			memory[outputs[306]] = memory[109] ^ memory[302]
		}
		if memory[304] != -1 && memory[94] != -1 && memory[outputs[307]] == -1 {
			changed = true
			memory[outputs[307]] = memory[304] ^ memory[94]
		}
		if memory[305] != -1 && memory[113] != -1 && memory[outputs[308]] == -1 {
			changed = true
			memory[outputs[308]] = memory[305] | memory[113]
		}
		if memory[308] != -1 && memory[174] != -1 && memory[outputs[309]] == -1 {
			changed = true
			memory[outputs[309]] = memory[308] ^ memory[174]
		}
		if memory[308] != -1 && memory[174] != -1 && memory[outputs[310]] == -1 {
			changed = true
			memory[outputs[310]] = memory[308] & memory[174]
		}
		if memory[310] != -1 && memory[119] != -1 && memory[outputs[311]] == -1 {
			changed = true
			memory[outputs[311]] = memory[310] | memory[119]
		}
	}
	if changed {
		panic("failed to assign all values")
	}
	for i := 0; i < len(memory); i++ {
		if memory[i] == -1 {
			return -1
		}
	}
	z := 0
	z |= memory[146] << 0
	z |= memory[156] << 1
	z |= memory[184] << 2
	z |= memory[187] << 3
	z |= memory[190] << 4
	z |= memory[192] << 5
	z |= memory[196] << 6
	z |= memory[198] << 7
	z |= memory[201] << 8
	z |= memory[205] << 9
	z |= memory[207] << 10
	z |= memory[210] << 11
	z |= memory[213] << 12
	z |= memory[216] << 13
	z |= memory[221] << 14
	z |= memory[229] << 15
	z |= memory[226] << 16
	z |= memory[228] << 17
	z |= memory[231] << 18
	z |= memory[234] << 19
	z |= memory[135] << 20
	z |= memory[240] << 21
	z |= memory[244] << 22
	z |= memory[247] << 23
	z |= memory[249] << 24
	z |= memory[252] << 25
	z |= memory[258] << 26
	z |= memory[259] << 27
	z |= memory[261] << 28
	z |= memory[264] << 29
	z |= memory[267] << 30
	z |= memory[270] << 31
	z |= memory[275] << 32
	z |= memory[276] << 33
	z |= memory[280] << 34
	z |= memory[282] << 35
	z |= memory[286] << 36
	z |= memory[289] << 37
	z |= memory[292] << 38
	z |= memory[296] << 39
	z |= memory[298] << 40
	z |= memory[300] << 41
	z |= memory[306] << 42
	z |= memory[307] << 43
	z |= memory[309] << 44
	z |= memory[311] << 45
	return z
}

var xyIndexesToCheckFirst = [][]int{
	{0, 15},
	{1, 15},
	{2, 15},
	{3, 15},
	{4, 15},
	{5, 15},
	{6, 15},
	{7, 15},
	{8, 15},
	{9, 15},
	{10, 15},
	{11, 15},
	{12, 15},
	{13, 15},
	{14, 15},
	{15, 15},
	{16, 15},
	{17, 15},
	{18, 15},
	{19, 15},
	{20, 15},
	{21, 15},
	{22, 15},
	{23, 15},
	{24, 15},
	{25, 15},
	{26, 15},
	{27, 15},
	{28, 15},
	{29, 15},
	{30, 15},
	{31, 15},
	{32, 15},
	{33, 15},
	{34, 15},
	{35, 15},
	{36, 15},
	{37, 15},
	{38, 15},
	{39, 15},
	{40, 15},
	{41, 15},
	{42, 15},
	{43, 15},
	{44, 15},
	{15, 0},
	{27, 0},
	{37, 0},
	{15, 1},
	{27, 1},
	{37, 1},
	{15, 2},
	{27, 2},
	{37, 2},
	{15, 3},
	{27, 3},
	{37, 3},
	{15, 4},
	{27, 4},
	{37, 4},
	{15, 5},
	{27, 5},
	{37, 5},
	{15, 6},
	{27, 6},
	{37, 6},
	{15, 7},
	{27, 7},
	{37, 7},
	{15, 8},
	{27, 8},
	{37, 8},
	{15, 9},
	{27, 9},
	{37, 9},
	{15, 10},
	{27, 10},
	{37, 10},
	{15, 11},
	{27, 11},
	{37, 11},
	{15, 12},
	{27, 12},
	{37, 12},
	{15, 13},
	{27, 13},
	{37, 13},
	{14, 14},
	{15, 14},
	{27, 14},
	{37, 14},
	{15, 16},
	{27, 16},
	{37, 16},
	{15, 17},
	{27, 17},
	{37, 17},
	{15, 18},
	{27, 18},
	{37, 18},
	{15, 19},
	{27, 19},
	{37, 19},
	{15, 20},
	{27, 20},
	{37, 20},
	{15, 21},
	{27, 21},
	{37, 21},
	{15, 22},
	{27, 22},
	{37, 22},
	{15, 23},
	{27, 23},
	{37, 23},
	{15, 24},
	{27, 24},
	{37, 24},
	{15, 25},
	{27, 25},
	{37, 25},
	{15, 26},
	{27, 26},
	{37, 26},
	{0, 27},
	{1, 27},
	{2, 27},
	{3, 27},
	{4, 27},
	{5, 27},
	{6, 27},
	{7, 27},
	{8, 27},
	{9, 27},
	{10, 27},
	{11, 27},
	{12, 27},
	{13, 27},
	{14, 27},
	{15, 27},
	{16, 27},
	{17, 27},
	{18, 27},
	{19, 27},
	{20, 27},
	{21, 27},
	{22, 27},
	{23, 27},
	{24, 27},
	{25, 27},
	{26, 27},
	{27, 27},
	{28, 27},
	{29, 27},
	{30, 27},
	{31, 27},
	{32, 27},
	{33, 27},
	{34, 27},
	{35, 27},
	{36, 27},
	{37, 27},
	{38, 27},
	{39, 27},
	{40, 27},
	{41, 27},
	{42, 27},
	{43, 27},
	{44, 27},
	{15, 28},
	{27, 28},
	{37, 28},
	{15, 29},
	{27, 29},
	{37, 29},
	{15, 30},
	{27, 30},
	{37, 30},
	{15, 31},
	{27, 31},
	{37, 31},
	{15, 32},
	{27, 32},
	{37, 32},
	{15, 33},
	{27, 33},
	{37, 33},
	{15, 34},
	{27, 34},
	{37, 34},
	{15, 35},
	{27, 35},
	{37, 35},
	{15, 36},
	{27, 36},
	{36, 36},
	{37, 36},
	{0, 37},
	{1, 37},
	{2, 37},
	{3, 37},
	{4, 37},
	{5, 37},
	{6, 37},
	{7, 37},
	{8, 37},
	{9, 37},
	{10, 37},
	{11, 37},
	{12, 37},
	{13, 37},
	{14, 37},
	{15, 37},
	{16, 37},
	{17, 37},
	{18, 37},
	{19, 37},
	{20, 37},
	{21, 37},
	{22, 37},
	{23, 37},
	{24, 37},
	{25, 37},
	{26, 37},
	{27, 37},
	{28, 37},
	{29, 37},
	{30, 37},
	{31, 37},
	{32, 37},
	{33, 37},
	{34, 37},
	{35, 37},
	{36, 37},
	{38, 37},
	{39, 37},
	{40, 37},
	{41, 37},
	{42, 37},
	{43, 37},
	{44, 37},
	{15, 38},
	{27, 38},
	{37, 38},
	{15, 39},
	{27, 39},
	{37, 39},
	{15, 40},
	{27, 40},
	{37, 40},
	{15, 41},
	{27, 41},
	{37, 41},
	{15, 42},
	{27, 42},
	{37, 42},
	{15, 43},
	{27, 43},
	{37, 43},
	{15, 44},
	{27, 44},
	{37, 44},
	{44, 44},
}
