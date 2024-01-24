package json

import jsoniter "github.com/json-iterator/go"

// Marshal 将数据转换为json格式
// 传入参数：数据
// 传出参数：json格式的数据,错误信息
func Marshal(v interface{}) ([]byte, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(v)
}

// Unmarshal 将json格式的数据转换为数据
// 传入参数：json格式的数据,接收数据的结构体指针
// 传出参数：错误信息
func Unmarshal(data []byte, v interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(data, v)
}
