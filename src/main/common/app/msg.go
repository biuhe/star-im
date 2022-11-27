package app

// MessageMap 返回值常量对应的消息内容，消息集合：{消息码，消息内容}
var MessageMap = map[int]string{
	SUCCESS: "成功",
	ERROR:   "失败",
}

// GetMsg 根据代码获取返回信息
func GetMsg(code int) string {
	msg, ok := MessageMap[code]
	if ok {
		return msg
	}

	return MessageMap[ERROR]
}
