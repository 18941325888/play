package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	x := rand.Intn(100) + 1

	var o int
	z := 1
	y := 100

	fmt.Printf("猜数字%d-%d，请输入你猜的数", z, y)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	o, _ = strconv.Atoi(input)

	for {
		if o < z || o > y {
			fmt.Printf("输入错误请重新输入，现在%d-%d,请继续输入你猜的数", z, y)
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			o, _ = strconv.Atoi(input)
			continue
		}
		switch {
		case o == x:
			goto ok
		case o < x:
			z = o
			fmt.Printf("你猜的数值比正确的小啦，现在%d-%d,请继续输入你猜的数", z, y)
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			o, _ = strconv.Atoi(input)
		case o > x:
			y = o
			fmt.Printf("你猜的数值比正确的大啦，现在%d-%d,请继续输入你猜的数", z, y)
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			o, _ = strconv.Atoi(input)
		}

	}
ok:
	fmt.Println("你猜对啦！")

}
