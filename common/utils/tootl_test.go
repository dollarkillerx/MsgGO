package utils

import (
	"strings"
	"testing"
)

func TestNewUUID(t *testing.T) {
	s, e := NewUUID()
	if e != nil {
		t.Error(e.Error())
	}
	t.Log(s)
}

func TestNewUUIDSimplicity(t *testing.T) {
	s, e := NewUUIDSimplicity()
	if e != nil {
		t.Error(e.Error())
	}
	t.Log(s)
}

func TestMd5String(t *testing.T) {
	s := Md5Encode("123")
	s = strings.ToUpper(s)
	if s != "202CB962AC59075B964B07152D234B70" {
		t.Error("非正常", s)
	}
}

func TestGetCurrentTime(t *testing.T) {
	time := TimeGetNowTimeStr()
	t.Log(time)
}

func TestGetTimeToString(t *testing.T) {
	s, e := TimeGetTimeToString("1558057058")
	if e != nil {
		t.Error(e.Error())
	}
	t.Log(s)
}

func TestGetTimeStringToTime(t *testing.T) {
	s, e := TimeGetStringToTime("2019-05-17")
	if e != nil {
		t.Error(e.Error())
	}
	t.Log(s)
}

func TestGetSession(t *testing.T) {
	session := SessionGenerate("dollarkiller", 6*60*60)
	t.Log(session)
	bool := SessionCheck(session)
	t.Log(SessionMap.Load(session))
	t.Log(bool)

	node, e := SessionGetData(session)
	if e != nil {
		t.Fatal(e.Error())
	}
	t.Log(node)

	SessionDel(session)
	bool = SessionCheck(session)
	t.Log(bool)
}

func TestGenRsaKey(t *testing.T) {
	e, priKey, pubKey := GenRsaKey(1024)
	if e == nil {
		t.Log(priKey)
		t.Log(pubKey)
	}

	data := "1231231245asdasd你好"
	s, e := RsaEncryptSimple(data, pubKey)
	if e != nil {
		t.Fatal(e.Error())
	}
	simple, e := RsaDecryptSimple(s, priKey)
	if e != nil {
		t.Fatal(e.Error())
	}
	if strings.EqualFold(data, simple) {
		t.Log("OK")
	}
	t.Logf(data)
	t.Logf(s)
	t.Logf(simple)
}

func TestRsaSign(t *testing.T) {
	e, priKey, pubKey := GenRsaKey(1024)
	if e == nil {
		t.Log(priKey)
		t.Log(pubKey)
	}
	data := "1a2sd1as23d你好"
	s, e := RsaSignSimple(data, priKey)
	if e != nil {
		t.Log(e.Error())
	}
	t.Log("签名: ", s)
	e = RsaSignVerSimple(data, s, pubKey)
	if e != nil {
		t.Log("验证失败")
	}
	t.Log("验证成功")
}

func TestRsaDecryptSimple(t *testing.T) {
	e, priKey, pubKey := GenRsaKey(1024)
	if e == nil {
		t.Log(len(priKey))
		t.Log(len(pubKey))
	}
	t.Logf(priKey)
	t.Logf(pubKey)
}