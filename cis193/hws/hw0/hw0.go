package main

import "fmt"

func main() {
	fmt.Println("Hello, à¤¦à¥à¤¨à¤¿à¤¯à¤¾!")
	for i := 1; i < 30; i++ {
		fmt.Println(Fizzbuzz(i))
	}
	fmt.Println(IsPalindrome("racecar"))
	fmt.Println(IsPalindrome("airplane"))
}

func Fizzbuzz(n int) string {
	var result = ""
	if n%3 == 0 {
		result += "fizz"
	}
	if n%5 == 0 {
		result += "buzz"
	}
	if len(result) == 0 {
		result = fmt.Sprint(n)
	}
	return result
}

func IsPrime(n int) bool {
	if n <= 2 {
		return true
	} else {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
	}
	return true
}

func IsPalindrome(s string) bool {
	var revString string = ""
	for i := 0; i < len(s); i++ {
		revString = s[i:i+1] + revString
	}
	return s == revString
}
