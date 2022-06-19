package main

import "fmt"

func FindSubstring(input_str string) string {
	var stack []byte
	var i, lenInput = 0, len(input_str)
	cnt := make([]int, 26)
	visit := make([]bool, 26)

	for ; i < lenInput; i++ {
		cnt[input_str[i]-'a']++
	}

	for i = 0; i < lenInput; i++ {
		cnt[input_str[i]-'a']--

		if !visit[input_str[i]-'a'] {
			n := len(stack) - 1
			for len(stack) > 0 && stack[n] < input_str[i] && cnt[stack[n]-'a'] != 0 {
				visit[stack[n]-'a'] = false
				stack = stack[:n]
				n = len(stack) - 1
			}

			stack = append(stack, input_str[i])
			visit[input_str[i]-'a'] = true
		}

	}
	return string(stack)
}
func main() {
	fmt.Println("Output: yzxyz -> ", FindSubstring("yzxyz"))
}