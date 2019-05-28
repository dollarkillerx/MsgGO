/**
* @Author: dollarkiller
* @Date: 19-5-28 下午10:12
* @Description:
* */
package main

import (
	"MsgGO/user/utils"
	"fmt"
)

func main()  {
	encode := utils.Md5Encode("123456c6c1b557d1e3478cbe529c1155c00330")
	fmt.Println(encode)
}
