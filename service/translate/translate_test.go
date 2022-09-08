package translate_test

import (
	"os"
	"strings"
	"testing"
	"wiki-knogen-zero/service/translate"
)

func TestTransater(t *testing.T) {
	t.Log("Appid", os.Getenv("TestBaiduTranslateAppid"))
	t.Log("Secret", os.Getenv("TestBaiduTranslateSecret"))
	t.Log("pgconn", os.Getenv("TestPgConnUri"))

	baiduTranslate := translate.NewTranslate(
		os.Getenv("TestBaiduTranslateAppid"),
		os.Getenv("TestBaiduTranslateSecret"),
		strings.ReplaceAll(os.Getenv("TestPgConnUri"), "\"", ""),
	)

	// ret := baiduTranslate.TranslateOne("good night \n god like")
	// t.Log(ret)

	// ret = baiduTranslate.TranslateOne("河南人")
	// t.Log(ret)

	retS := baiduTranslate.TranslateSliceString([]string{"good night \n god like", "good basic", "tomcat"})
	t.Log(retS)

	retS = baiduTranslate.TranslateSliceString([]string{"tomcat"})
	t.Log(retS)

	retS = baiduTranslate.TranslateSliceString([]string{"york", "golang 2021", "tomcat", "facebook", "龙猫"})
	t.Log(retS)
}
