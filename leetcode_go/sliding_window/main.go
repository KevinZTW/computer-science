package main

import "fmt"

func main() {
	fmt.Println(longestSubstring("abcabcbb"))
	fmt.Println(longestKSubstring("abcabcbb", 2))
}

func longestSubstring(s string) int {
	longest := 0
	for i := 0; i < len(s); i++ {
		wl := 0
		words := make(map[byte]byte, len(s))
		for k := i; k < len(s)-1; {
			if _, ok := words[s[k]]; ok != true {

				words[s[k]] = s[k]

				wl += 1
				k++
			} else {
				break
			}
		}
		if wl > longest {
			longest = wl
		}
	}
	return longest
}

func longestKSubstring(s string, rep int) int {
	longest := 0
	for i := 0; i < len(s); i++ {
		wl := 0
		words := make(map[byte]int, len(s))
		for k := i; k < len(s)-1; {
			if i, _ := words[s[k]]; i < rep {
				words[s[k]] += 1
				wl += 1
				k++
			} else {
				break
			}
		}
		if wl > longest {
			longest = wl
		}
	}
	return longest
}
