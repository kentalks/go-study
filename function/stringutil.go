package function

import (
	"kensoft.tech/go-lib/util"
)

func ToUpper(s string) string {
	return util.ToUpper(s)
}

func RandString(n int) string {
	return util.RandString(n)
}

func RandSeq(n int) string {
	return util.RandSeq(n)
}
