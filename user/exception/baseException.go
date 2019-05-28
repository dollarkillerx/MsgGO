package exception

import "log"

func SimpleException(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func SimpleExceptionLog(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}