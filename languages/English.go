package languages

var ErrorMsgEn = map[int]string{
	200: "success",
	500 : "fail",
	400 : "params error.",
}


//获取errorMsg
func GetErrorMsg(code int) (string) {
	return ErrorMsgEn[code]
}
