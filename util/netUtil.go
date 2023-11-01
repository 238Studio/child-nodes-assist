package util

// SendWebsocketNetMessage 包装了网络讯息发送类
// 传入：需要发送的消息，核心的目标模块，是否是原始类型
// 传出：无
func SendWebsocketNetMessage(data *[]byte, targetModuleID string, isBytes bool) *map[string][]interface{} {
	r := make(map[string][]interface{})
	r[targetModuleID] = make([]interface{}, 2)
	r[targetModuleID][0] = isBytes
	r[targetModuleID][1] = *data
	return &r
}
