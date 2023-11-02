package util

// SendWebsocketNetMessage 包装了网络讯息发送类
// 传入：需要发送的消息，核心的目标模块，是否是原始类型
// 传出：无
func SendWebsocketNetMessage(data *[]byte, targetModuleID string, isBytes bool) *[3]interface{} {
	r := [3]interface{}{}
	// 核心的目标模块ID
	r[0] = targetModuleID
	// 是否是原始二进制数据
	r[1] = isBytes
	// 数据
	r[2] = *data
	return &r
}
