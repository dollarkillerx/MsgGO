/**
* @Author: dollarkiller
* @Date: 19-5-30 下午1:44
* @Description:
* */
package execption

func ExecptionPanic(err error)  {
	if err != nil {
		panic(err.Error())
	}
}
