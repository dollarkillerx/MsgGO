/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-12
* Time: 上午11:53
* */
package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 如果文件夹不存在就会创建
func DirPing(path string) error {
	b, e := PathExists(path)
	if e != nil {
		return e
	}
	if !b {
		e := os.Mkdir(path, 00777)
		if e != nil {
			return e
		}
	}
	return nil
}

// 获取文件后缀
func FileGetPostfix(filename string) (string, error) {
	split := strings.Split(filename, ".")
	if len(split) == 0 {
		return "", errors.New("File Get Postfix Error")
	}
	return split[len(split)-1], nil
}

// 获得随机文件名 传入postfilx后缀
func FileGetRandomName(postfilx string) string {
	nano := time.Now().UnixNano()
	intn := rand.Intn(10000)
	intn = rand.Intn(10000)
	intn += intn
	formatInt := strconv.FormatInt(nano, 10)
	itoa := strconv.Itoa(intn)
	encode := Md5Encode(formatInt + itoa + postfilx)
	name := encode + "." + postfilx
	return name
}

// 获得文件sha1
func FileGetSha1(file *os.File) string {
	hash := sha1.New()
	io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil))
}

// 获取文件MD5
func FileGetMD5(file *os.File) string {
	_md5 := md5.New()
	io.Copy(_md5, file)
	return hex.EncodeToString(_md5.Sum(nil))
}

// 获取文件大小
func FielGetSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}
