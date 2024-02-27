package function

import util "kensoft.tech/go-lib"

func ToUpper(s string) string {
	return util.ToUpper(s)
}

func ProcessStr(str string, conf byte) string {
	return util.ProcessStr(str, conf)
}

func RandString(n int) string {
	return util.RandString(n)
}

func RandSeq(n int) string {
	return util.RandSeq(n)
}
