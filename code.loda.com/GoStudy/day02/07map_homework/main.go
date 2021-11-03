package main

import (
	"fmt"
	"strings"
	"unicode"
)

//统计一个字符串中每个单词出现的次数
//比如how do you do -> how:1 do:2 you:1
func count_word1() {
	str := "how  do   you do"
	m_total := make(map[string]int, 20)

	// 方法一：太麻烦，对字符串的操作不熟悉
	char_count := 0
	for index, v := range str {
		if unicode.IsSpace(v) {
			// fmt.Println("empty char", char_count)
			s_word := []string{}
			//遍历原始字符串
			for i := index - char_count; i < index; i++ {
				//将字符拼接成单词写入到string中
				s_word = append(s_word, string(str[i]))
			}
			word := fmt.Sprintf(strings.Join(s_word, ""))
			// fmt.Println(word)
			//将string写入到map中
			if m_total[word] == 0 {
				m_total[word] = 1
			} else {
				m_total[word]++
			}
			char_count = 0
		} else if index == len(str)-1 { //需对末尾不是空格的情况，最后一个字符做判读
			fmt.Println(char_count)
			s_word := []string{}
			//遍历原始字符串
			for i := index - char_count; i <= index; i++ {
				//将字符拼接成单词写入到string中
				s_word = append(s_word, string(str[i]))
			}
			word := fmt.Sprintf(strings.Join(s_word, ""))
			fmt.Println(word)
			//将string写入到map中
			if m_total[word] == 0 {
				m_total[word] = 1
			} else {
				m_total[word]++
			}
		} else {
			// fmt.Println("char++")
			char_count++
		}
	}
	fmt.Println(m_total)
}

func count_word2() {
	str := "how do you do "
	m_total := make(map[string]int, 20)

	//方法二：
	//将字符串按照空格分割，存放到切片中
	str2 := strings.Split(str, " ") //不能处理多个空格的情况，最后一个字符是空格也无法处理
	for _, v := range str2 {
		fmt.Println(v)
		_, ok := m_total[v]
		if !ok {
			m_total[v] = 1
		} else {
			m_total[v]++
		}
	}

	for key, value := range m_total {
		fmt.Println(key, value)
	}
	// fmt.Println()
	// //只遍历key
	// for key := range m_total {
	// 	fmt.Println(key)
	// }
	// fmt.Println()
	// //只遍历值
	// for _, v := range m_total {
	// 	fmt.Println(v)
	// }
}

//判断是否是回文字符串
func plalindrome(str string) {
	//slice的长度必须是0
	//如果长度不是0，则创建出来的slice中的元素是零值，后续append就是直接向后追加，零值不会被覆盖
	//考虑到非ASCII的存在，比如汉字，要用rune类型来存储
	r_slice := make([]rune, 0, len(str))
	for _, v := range str {
		r_slice = append(r_slice, v)
	}
	// fmt.Println(r_slice)

	//由于Go没有逗号表达式，而++和--是语句而不是表达式，如果想在for中执行多个变量,需要使用平行赋值
	for i, j := 0, len(r_slice)-1; i < j; i, j = i+1, j-1 {
		if r_slice[i] != r_slice[j] {
			fmt.Println("str is not a plalindrome string")
			return
		}
	}

	// for i := 0; i < len(r_slice)/2; i++ {
	// 	if r_slice[i] != r_slice[len(r_slice)-1-i] {
	// 		fmt.Println("str is not a plalindrome string")
	// 		return
	// 	}
	// }

	fmt.Println("str is a plalindrome string")
	// if i == j || i > j {
	// 	fmt.Println("str is a plalindrome string")
	// }
}

func main() {
	count_word1()
	// count_word2()
	plalindrome("abcba")
	plalindrome("上海自来水来自海上")

}
