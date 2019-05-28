/**
* @Author: dollarkiller
* @Date: 19-5-28 下午10:22
* @Description:
* */
package utils

import "strings"

func StrFilter(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}