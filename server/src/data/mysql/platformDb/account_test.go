package platformDb

import (
	"testing"
)

func Test_FindOne(t *testing.T) {
	InitPlatFromDb()
	query := &Account{ID: 1}
	res, err := query.FindOne()
	t.Logf("Res => %v", res)
	t.Logf("Err => %v", err)
}
