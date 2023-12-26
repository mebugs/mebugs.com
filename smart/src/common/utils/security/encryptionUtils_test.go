package security

import "testing"

func TestAESDecrypt(t *testing.T) {
	deStr, err := AESDecrypt("VEmHXzPPPKzYZtY2w0tYHM3OGMA7jPZn50BermJ37Jc=", "KEY_SiteOL_Stone")
	t.Log(deStr)
	t.Log(err)
}

func TestAESEncrypt(t *testing.T) {
	enStr, err := AESEncrypt("admin123", "mSZFrzZLYCFf1tta")
	t.Log(enStr)
	t.Log(err)
}
