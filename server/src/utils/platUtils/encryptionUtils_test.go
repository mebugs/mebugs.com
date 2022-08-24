package platUtils

import "testing"

func TestDecrypt(t *testing.T) {
	deStr, err := Decrypt("VEmHXzPPPKzYZtY2w0tYHM3OGMA7jPZn50BermJ37Jc=", "KEY_SiteOL_Stone")
	t.Log(deStr)
	t.Log(err)
}

func TestEncrypt(t *testing.T) {
	enStr, err := Encrypt("AdminAdminAdmin1", "KEY_SiteOL_Stone")
	t.Log(enStr)
	t.Log(err)
}
