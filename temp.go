package main

import (
	"unicode"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	low, hei := 0, len(s)-1
	for low < hei {
		if !((s[low] >= 'A' && s[low] <= 'Z') ||
			(s[low] >= 'a' && s[low] <= 'z') ||
			(s[low] >= '0' && s[low] <= '9')) {
			low++
		} else if !((s[hei] >= 'A' && s[hei] <= 'Z') ||
			(s[hei] >= 'a' && s[hei] <= 'z') ||
			(s[hei] >= '0' && s[hei] <= '9')) {
			hei--
		} else {
			if s[low] == s[hei] || ignoreCases(s[low], s[hei]) {
				low++
				hei--
			} else {
				return false
			}
		}
	}
	return true
}

func ignoreCases(a, b byte) bool {
	return unicode.ToLower(rune(a)) == unicode.ToLower(rune(b))
}

func singleNumber(nums []int) int {
	result := 0
	for _, val := range nums {
		result = result ^ val
	}
	return result
}

func convertToTitle(n int) string {
	/*  An three bits Excel sheet title can be represented as:
		a * 26^2 + b * 26 + c
	*/
	title := ""
	for n != 0 {
		last_bit := n % 26
		if last_bit == 0 {
			title = "Z" + title
		} else {
			title = string(byte('A' + last_bit - 1))
		}

		n = (n - n % 26) / 10
	}
	return title
}

func isHappy(n int) bool {
	mapping := make(map[int]int)
	for {
		n = ss(n)
		if _, ok := mapping[n]; ok {
			return false
		}
		if n == 1 {
			return true
		}
		mapping[n] = 0
	}
	return false
}

func ss(n int) int {
	result := 0
	for n != 0 {
		result += squares(n % 10)
		n /= 10
	}
	return result
}

func squares(a int) int {
	return a * a
}


func countPrimes(n int) int {
	result := 0
	flagArray := make(map[int]bool, n)
	for i := 2; i < n; i++ {
		if flagArray[i] == false {
			// is prime
			result++

			// remove it's multiples

			for j := 2; j * i < n; j++ {
				flagArray[j*i] = true
			}
		}
	}
	return result
}

func main() {
	countPrimes(10)
}
