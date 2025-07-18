package main

import "fmt"

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	for _, name := range users {
		var sum int
		for _, c := range name {
			switch c {
			case 'e', 'E':
				sum += 1
			case 'i', 'I':
				sum += 2
			case 'o', 'O':
				sum += 3
			case 'u', 'U':
				sum += 4
			}
		}
		distribution[name] = sum
		coins -= sum
	}
	return coins
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	fmt.Println(distribution)
}
