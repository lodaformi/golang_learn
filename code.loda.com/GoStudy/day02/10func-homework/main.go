package main

import (
	"fmt"
)

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
	//需对每个字符做判断
	for _, v := range users {
		//初始化map，如果不先初始化，则Sarah的值不会显示
		//即值为0的数据不会显示，
		distribution[v] = 0
		for _, c := range v {
			switch c {
			case 'e', 'E':
				distribution[v] += 1
				coins -= 1
			case 'i', 'I':
				distribution[v] += 2
				coins -= 2
			case 'o', 'O':
				distribution[v] += 3
				coins -= 3
			case 'u', 'U':
				distribution[v] += 4
				coins -= 4
			}
		}
	}
	// fmt.Println(distribution)

	//使用strings.Contains只能得出字符串中是否包含某个字母，不能处理多次包含同一个字母的情况
	//对字符串进行判断，是否包含对应的字符
	//如果符合条件，改变对应的值，同时coins减去相应的值
	// for _, v := range users {
	// 	if strings.Contains(v, "e") || strings.Contains(v, "E") {
	// 		// fmt.Println("e", v)
	// 		distribution[v] += 1
	// 		coins -= 1
	// 	}
	// 	if strings.Contains(v, "i") || strings.Contains(v, "I") {
	// 		// fmt.Println("i", v)
	// 		distribution[v] += 2
	// 		coins -= 2
	// 	}
	// 	if strings.Contains(v, "o") || strings.Contains(v, "O") {
	// 		// fmt.Println("o", v)
	// 		distribution[v] += 3
	// 		coins -= 3
	// 	}
	// 	if strings.Contains(v, "u") || strings.Contains(v, "U") {
	// 		// fmt.Println("u", v)
	// 		distribution[v] += 4
	// 		coins -= 4
	// 	}
	// }
	// fmt.Println(distribution)

	return coins
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	for user, value := range distribution {
		fmt.Println(user, value)
	}
}
