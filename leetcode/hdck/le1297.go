package main

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	maps := make(map[string]int)
	for i:=0;i<len(s)-minSize;i++ {
		str :=s[i:i+minSize]
		count :=0
		for j:=0;j<len(str);j++ {
			if 1<< int(str[j] - 'a') & count ==0 {
				count = count | 1<< int(str[j] - 'a')
			}
		}
		sum :=0
		for j :=0;j<32;j++ {
			if (1 <<j) &count != 0 {
				sum++
			}
		}
		if sum <=maxLetters {
			_,ok := maps[str]
			if !ok {
				maps[str] = 1
			}else {
				maps[str] = maps[str]+1
			}
		}
	}
	max :=0
	for _,v :=range maps {
		if v > max {
			max = v
		}
	}
	return max
}

