/**
* @Author: dollarkiller
* @Date: 19-5-29 下午3:37
* @Description:
* */
package auth

import (
	"MsgGO/user/dbops/defsModel"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	user := &defsModel.ImUser{Name: "dollarkiller"}
	s, e := GenerateToken(user)
	if e != nil {
		t.Fatal(e.Error())
	}
	t.Log(s)
}
