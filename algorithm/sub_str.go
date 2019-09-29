/*
	Print the longest sub-string occuring twice or more in the given string.
	eg. banana     --> ana
		abcsdfnabc --> abc
*/

package main

import "fmt"

func longestSubString(s string) string {
	hashmap := make(map[byte]int)
	res := ""
	for i := 0; i < len(s); i++ {
		if _, ok := hashmap[s[i]]; ok {
			counter := i+1
			for j := hashmap[s[i]]+1; counter < len(s) && s[j] == s[counter]; j, counter = j+1, counter+1 {
			}

			if counter - i > len(res) {
				res = s[i:counter]
			}
		} else {
			hashmap[s[i]] = i
		}
	}
	return res
}

func main() {
	s := ""
	fmt.Scan(&s)

	fmt.Println(longestSubString(s))
}
