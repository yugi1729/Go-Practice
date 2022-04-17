package main

import "fmt"

func top(arr []rune) rune{
	return arr[len(arr)-1]
}

func pop(arr []rune) []rune{
	return arr[:len(arr)-2]
}

func main(){
	var brackets string
	fmt.Scanf("%x",brackets)
	stack := make([]rune,len(brackets))
	for _, c := range brackets{
		switch c { 
		case '{':
			stack = append(stack, c)
		case '[':
			stack = append(stack, c)
		case '(':
			stack = append(stack, c)
		case '}':
			if top(stack) == '{'{
				stack = pop(stack)
			}
		case ']':
			if top(stack) == '['{
				stack = pop(stack)
			}
		case ')':
			if top(stack) == '('{
				stack = pop(stack)
			}
		}
	}
	if len(brackets) == 0{
		fmt.Println("Balanced Brackets")
	} else {
		fmt.Println("Not Balanced")
	}
}