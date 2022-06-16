package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {

	for {

		var o = [3]string{"石头", "剪刀", "布"}
		fmt.Printf("石头，剪子，布（不玩请输入退出）！你出：")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "退出" {
			goto ok
		}

		if input == "石头" || input == "剪刀" || input == "布" {
			rand.Seed(time.Now().Unix())
			x := rand.Intn(3)

			switch {

			case input == "石头" && o[x] == "石头":
				fmt.Println("我出", o[x], "我们平啦！")
			case input == "石头" && o[x] == "剪刀":
				fmt.Println("我出", o[x], "我输啦！")
			case input == "石头" && o[x] == "布":
				fmt.Println("我出", o[x], "我赢啦！")
			case input == "剪刀" && o[x] == "剪刀":
				fmt.Println("我出", o[x], "我们平啦！")
			case input == "剪刀" && o[x] == "石头":
				fmt.Println("我出", o[x], "我赢啦！")
			case input == "剪刀" && o[x] == "布":
				fmt.Println("我出", o[x], "我输啦！")
			case input == "布" && o[x] == "剪刀":
				fmt.Println("我出", o[x], "我赢啦！")
			case input == "布" && o[x] == "石头":
				fmt.Println("我出", o[x], "我输啦！")
			case input == "布" && o[x] == "布":
				fmt.Println("我出", o[x], "我们平啦！")
			}
		} else {
			fmt.Printf("输入错误！！！")
			continue
		}
	}
ok:
	fmt.Println("不玩算啦！")

}
