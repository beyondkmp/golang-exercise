//一行一行读取文件，并进行处理
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("test.txt")
	fmt.Printf("the type of f is %T", f)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
