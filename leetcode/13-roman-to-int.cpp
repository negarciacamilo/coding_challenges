class Solution {
public:
	int romanToInt(string s) {
		map<char, int> numbers = {
			{'I', 1},
			{'V', 5},
			{'X', 10},
			{'L', 50},
			{'C', 100},
			{'D', 500},
			{'M', 1000},
		};

		int result = 0, previousNumber = 0;
		for(int i = s.size() - 1; i >= 0; i--) {
			int v = numbers[s[i]];
			if (numbers.count(s[i]) != 0) {
				if (previousNumber > v) {
					result -= v;
				} else {
					result += v;
				}
				previousNumber = v;
			}
		}
		return result;
	}
};

