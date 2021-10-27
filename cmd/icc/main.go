package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(errors.New("引数の個数が正しくありません"))
	}
	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".global main\n")
	fmt.Printf("main:\n")

	p := os.Args[1]
	fmt.Printf("  mov rax, %c\n", p[0])

	for i := 1; i < len(p); i++ {
		if p[i] == '+' {
			i++
			fmt.Printf("  add rax, %c\n", p[i])
			continue
		}
		if p[i] == '-' {
			i++
			fmt.Printf("  sub rax, %c\n", p[i])
			continue
		}

		fmt.Println(errors.New(fmt.Sprintf("予期しない文字です: '%c'\n", p[i])))
		return
	}
	fmt.Printf("  ret\n")
}
