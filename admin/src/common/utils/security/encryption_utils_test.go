package security

import "testing"

func TestAESDecrypt(t *testing.T) {
	deStr, err := AESDecrypt("J0eU4dO/7UWe06eNdckWVw==", "KEY_SiteOL_Stone")
	t.Log(deStr)
	t.Log(err)
}

func TestAESEncrypt(t *testing.T) {
	enStr, err := AESEncrypt("123456", "KEY_SiteOL_Stone")
	t.Log(enStr)
	t.Log(err)
	TestAESDecrypt(t)
}
