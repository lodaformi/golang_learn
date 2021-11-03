package stringSplit

import "strings"

func SplitStr(str, sep string) (res []string) {
	ind := strings.Index(str, sep)

	for ind > -1 {
		res = append(res, str[:ind])
		str = str[(ind + len(sep)):]
		ind = strings.Index(str, sep)
	}

	res = append(res, str)
	return
}
