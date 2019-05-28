package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

/**
	用于加密解秘
	这个Sprintf有点坑，会有一个/n
 */

func NewUUID() (string,error) {
	out, err := exec.Command("uuidgen").Output()

	oot := fmt.Sprintf("%s", out)
	oot = strFilter(oot)
	return oot,err
}

// 获取没有 - 的uuid
func NewUUIDSimplicity() (string,error) {
	s, e := NewUUID()
	var u string
	for _,k :=range s {
		if k != '-' {
			u = fmt.Sprintf("%s%s",u,string(k))
		}
	}
	u = strFilter(u)
	return u,e
}

func strFilter(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}