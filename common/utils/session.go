/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-9
* Time: 下午12:08
* */
package utils

import (
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	SessionMap sync.Map
)

type SessionNode struct {
	Name           string
	CreationTime   int64 // 创建时间
	ExpirationTime int64 // 过期时间
}

// 获得session
func SessionGenerate(name string, times int64) string { // name,过期时间
	timeNano := time.Now().UnixNano()
	time := time.Now().Unix()
	outtime := time + times
	intn := rand.Intn(100000)
	encode := Md5Encode(strconv.FormatInt(timeNano, 10) + strconv.Itoa(intn))
	node := &SessionNode{
		Name:           name,
		CreationTime:   time,
		ExpirationTime: outtime,
	}

	SessionMap.Store(encode, node)
	return encode
}

// 获取session数据
func SessionGetData(sessionId string) (*SessionNode, error) {
	if sessionId == "" || len(sessionId) == 0 {
		return nil, errors.New("not data")
	}
	value, ok := SessionMap.Load(sessionId)
	if ok != true {
		return nil, errors.New("not data")
	}

	node := value.(*SessionNode)
	nowTime := time.Now().Unix()
	if nowTime >= node.CreationTime && nowTime < node.ExpirationTime {
		return node, nil
	}
	// 删除过期的session
	SessionMap.Delete(sessionId)
	return nil, errors.New("not data")
}

// 验证session
func SessionCheck(sessionId string) bool {
	if sessionId == "" || len(sessionId) == 0 {
		return false
	}
	value, ok := SessionMap.Load(sessionId)
	if ok != true {
		return false
	}

	node := value.(*SessionNode)
	nowTime := time.Now().Unix()
	if nowTime >= node.CreationTime && nowTime < node.ExpirationTime {
		return true
	}
	// 删除过期的session
	SessionMap.Delete(sessionId)
	return false
}

// 删除session
func SessionDel(sessionId string) {
	SessionMap.Delete(sessionId)
}
