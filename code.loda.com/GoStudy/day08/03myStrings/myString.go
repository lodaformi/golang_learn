package mystrings

import (
	"strings"
)

func MyStringSplit(str string, sep string) (res []string) {
	ind := strings.Index(str, sep) //在str中找第一次出现的sep，找到就返回其索引，找不到返回-1
	for ind > -1 {
		tmp := str[:ind]
		//当出现连续的分隔符时，tmp有可能为空，空字符串不写入切片中，但要更新字符串
		if strings.Compare(tmp, "") != 0 {
			res = append(res, tmp) //找到将部分结果存放到res中
		}
		//更新str，如果sep是单个字符ind+1即可，但可能出现sep是多个字符的情况，所以要用len函数就求出sep的长度
		str = str[ind+len(sep):]
		ind = strings.Index(str, sep) //继续找，更新ind，用于下一次for循环判断
	}
	res = append(res, str)
	return
}
