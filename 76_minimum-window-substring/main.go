package leetcode

import "math"

func minWindow(s string, t string) string {
	need, window := make(map[uint8]int, len(t)), make(map[uint8]int, len(t))
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	var left, right, valid int
	start, length := 0, math.MaxInt32
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
			}
			window[d]--
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}
