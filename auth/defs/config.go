/**
* @Author: dollarkiller
* @Date: 19-5-30 下午1:40
* @Description:
* */
package defs

type Config struct {
	Host string `json:"host"`
	DriverName string `json:"driverName"`
	Dsn string `json:"dsn"`
	Test string `json:"Test"`
	PrivateKeyPath string `json:"private_key_path"`
	PublicKeyPath string `json:"public_key_path"`
}
