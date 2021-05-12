package main

import "fmt"

func main() {
	fmt.Println(longestSubstring(" ac"))
	fmt.Println(longestKSubstring("abcabcbb", 2))
	fmt.Println(longestPalinSubstring(" cc ") + "<- answer")
}

//3. Longest Substring Without Repeating Characters
func longestSubstring(s string) int {
	longest := 0
	for i := 0; i < len(s); i++ {
		wl := 0
		words := make(map[byte]byte, len(s))
		for k := i; k < len(s); {
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
	if len(s) == 0 {
		return 0
	}

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

//0005. Longest Palindromic Substring
func longestPalinSubstring(s string) string {

	lf, lr, longest := 0, 0, 0

	//if the result string lens in even
	for i := 0; i < len(s); i++ {
		left, right := i, i

		for left > 0 && right < len(s)-1 {
			if s[left-1] == s[right+1] {
				left--
				right++
				if right-left > longest {
					longest = right - left + 1
					lf, lr = left, right
				}
			} else {
				break
			}
		}
	}

	//if the result string is odd
	for i := 0; i < len(s)-1; i++ {

		left := i
		right := i
		if s[i+1] == s[i] {

			right = i + 1
			if right-left+1 > longest {
				longest = right - left + 1
				lf, lr = left, right
			}
		} else {
			continue
		}

		for left > 0 && right < len(s)-1 && s[left-1] == s[right+1] {
			left--
			right++
			if right-left+1 > longest {
				longest = right - left + 1
				lf, lr = left, right
			}
		}
	}

	return s[lf : lr+1]
}
