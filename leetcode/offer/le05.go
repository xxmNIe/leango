package main

func replaceSpace(s string) string {
	ls := len(s)
	if ls == 0 {
		return ""
	}
	ans := []byte{}
	adi := "%20"
	for i := range s {
		if s[i] == ' ' {
			for i := range adi {
				ans = append(ans, adi[i])
			}
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}
