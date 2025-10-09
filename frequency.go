package main

func Frequency(s string) map[string]int {

	frequency := make(map[string]int)

	for i := 0; i < len(s); i++ {
		ch := string(s[i])
		frequency[ch] += 1
	}

	return frequency
}
