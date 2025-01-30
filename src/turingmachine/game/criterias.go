package game

const (
	// The total number of criteria cards
	NumberOfCriterias = 48
)

var Criterias = [NumberOfCriterias]*Criteria{
	newCriteria(1, "△ compared to 1", []*Law{
		laws[1],
		laws[16],
	}),
	newCriteria(2, "△ compared to 3", []*Law{
		laws[25],
		laws[3],
		laws[18],
	}),
	newCriteria(3, "□ compared to 3", []*Law{
		laws[28],
		laws[8],
		laws[21],
	}),
	newCriteria(4, "□ compared to 4", []*Law{
		laws[29],
		laws[9],
		laws[138],
	}),
	newCriteria(5, "△ is even or odd", []*Law{
		laws[34],
		laws[37],
	}),
	newCriteria(6, "□ is even or odd", []*Law{
		laws[35],
		laws[38],
	}),
	newCriteria(7, "○ is even or odd", []*Law{
		laws[36],
		laws[39],
	}),
	newCriteria(8, "number of 1s in the code", []*Law{
		laws[40],
		laws[41],
		laws[42],
		// Skipping "three 1s"
	}),
	newCriteria(9, "number of 3s in the code", []*Law{
		laws[46],
		laws[47],
		laws[48],
		// Skipping "three 3s"
	}),
	newCriteria(10, "number of 4s in the code", []*Law{
		laws[49],
		laws[50],
		laws[51],
		// Skipping "three 4s"
	}),
	newCriteria(11, "△ number compared to the □ number", []*Law{
		laws[139],
		laws[89],
		laws[92],
	}),
	newCriteria(12, "□ number compared to the ○ number", []*Law{
		laws[140],
		laws[90],
		laws[93],
	}),
	newCriteria(13, "□ number compared to the ○ number", []*Law{
		laws[141],
		laws[91],
		laws[95],
	}),
	newCriteria(14, "which colour's number is smaller than the others'", []*Law{
		laws[116],
		laws[117],
		laws[118],
	}),
	newCriteria(15, "which colour's number is larger than the others'", []*Law{
		laws[113],
		laws[114],
		laws[115],
	}),
	newCriteria(16, "number of even numbers compared to the number of odd numbers", []*Law{
		laws[131],
		laws[132],
	}),
	newCriteria(17, "how many even numbers are there in the code", []*Law{
		laws[85],
		laws[86],
		laws[87],
		laws[88],
	}),
	newCriteria(18, "sum of all the numbers is even or odd", []*Law{
		laws[55],
		laws[56],
	}),
	newCriteria(19, "sum of △ and □ compared to 6", []*Law{
		laws[137],
		laws[100],
		laws[136],
	}),
	newCriteria(20, "a number repeats itself in the code", []*Law{
		laws[119],
		laws[120],
		laws[121],
	}),
	newCriteria(21, "there is a number present exactly twice", []*Law{
		laws[81],
		laws[82],
	}),
	newCriteria(22, "ascending order, descending order, or no order", []*Law{
		laws[133],
		laws[134],
		laws[135],
	}),
	newCriteria(23, "sum of all numbers compared to 6", []*Law{
		laws[74],
		laws[60],
		laws[67],
	}),
	newCriteria(24, "sequence of ascending numbers", []*Law{
		laws[83],
		laws[84],
		// Skipping "numbers in ascending order"
	}),
	newCriteria(25, "sequence of ascending or descending numbers", []*Law{
		laws[122],
		laws[123],
		laws[124],
	}),
	newCriteria(26, "a specific colour is less than 3", []*Law{
		laws[25],
		laws[28],
		laws[31],
	}),
	newCriteria(27, "a specific colour is less than 4", []*Law{
		laws[26],
		laws[29],
		laws[32],
	}),
	newCriteria(28, "a specific colour is equal to 1", []*Law{
		laws[1],
		laws[6],
		laws[11],
	}),
	newCriteria(29, "a specific colour is equal to 3", []*Law{
		laws[3],
		laws[8],
		laws[13],
	}),
	newCriteria(30, "a specific colour is equal to 4", []*Law{
		laws[4],
		laws[9],
		laws[14],
	}),
	newCriteria(31, "a specific colour is greater than 1", []*Law{
		laws[16],
		laws[19],
		laws[22],
	}),
	newCriteria(32, "a specific colour is greater than 3", []*Law{
		laws[18],
		laws[21],
		laws[24],
	}),
	newCriteria(33, "a specific colour is even or odd", []*Law{
		laws[34],
		laws[35],
		laws[36],
		laws[37],
		laws[38],
		laws[39],
	}),
	newCriteria(34, "which colour has the smallest number (or is tied for the smallest number)", []*Law{
		laws[128],
		laws[129],
		laws[130],
	}),
	newCriteria(35, "which colour has the largest number (or is tied for the largest number)", []*Law{
		laws[125],
		laws[126],
		laws[127],
	}),
	newCriteria(36, "sum of all the numbers is a multiple of 3 or 4 or 5", []*Law{
		laws[57],
		laws[58],
		laws[59],
	}),
	newCriteria(37, "sum of 2 specific colours is equal to 4", []*Law{
		laws[98],
		laws[103],
		laws[108],
	}),
	newCriteria(38, "sum of 2 specific colours is equal to 6", []*Law{
		laws[100],
		laws[105],
		laws[110],
	}),
	newCriteria(39, "number of one specific colour compared to 1", []*Law{
		laws[1],
		laws[6],
		laws[11],
		laws[16],
		laws[19],
		laws[22],
	}),
	newCriteria(40, "number of one specific colour compared to 3", []*Law{
		laws[25],
		laws[28],
		laws[31],
		laws[3],
		laws[8],
		laws[13],
		laws[18],
		laws[21],
		laws[24],
	}),
	newCriteria(41, "number of one specific colour compared to 4", []*Law{
		laws[26],
		laws[29],
		laws[32],
		laws[4],
		laws[9],
		laws[14],
		laws[142],
		laws[138],
		laws[143],
	}),
	newCriteria(42, "which colour is the smallest or the largest", []*Law{
		laws[116],
		laws[117],
		laws[118],
		laws[113],
		laws[114],
		laws[115],
	}),
	newCriteria(43, "△ number compared to teh number of another specific colour", []*Law{
		laws[139],
		laws[89],
		laws[92],
		laws[140],
		laws[90],
		laws[93],
	}),
	newCriteria(44, "□ number compared to teh number of another specific colour", []*Law{
		laws[144],
		laws[89],
		laws[94],
		laws[141],
		laws[91],
		laws[95],
	}),
	newCriteria(45, "how many 1s OR how many 3s there are in the code", []*Law{
		laws[40],
		laws[41],
		laws[42],
		laws[46],
		laws[47],
		laws[48],
	}),
	newCriteria(46, "how many 3s OR how many 4s there are in the code", []*Law{
		laws[46],
		laws[47],
		laws[48],
		laws[49],
		laws[50],
		laws[51],
	}),
	newCriteria(47, "how many 1s OR how many 4s there are in the code", []*Law{
		laws[40],
		laws[41],
		laws[42],
		laws[49],
		laws[50],
		laws[51],
	}),
	newCriteria(48, "one specific colour compared to another specific colour", []*Law{
		laws[139],
		laws[140],
		laws[141],
		laws[89],
		laws[90],
		laws[91],
		laws[92],
		laws[93],
		laws[95],
	}),
}
