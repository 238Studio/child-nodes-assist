package _const

// 数据报常数
const (
	PortLen uint32 = 512
)

// 子节点模块常数
const (
	InitModule   uint32 = 0x00
	SensorModule uint32 = 0x01
	ReportModule uint32 = 0x02
	ErrorModule  uint32 = 0x03
)

// 特殊数据报
const (
	// FailedToRev 接收失败
	FailedToRev byte = 0
	// SuccessRev 接收成功
	SuccessRev byte = 1
	// ReadyResend 预备重发
	ReadyResend byte = 2
	// CancelResend 取消重发
	CancelResend byte = 3
)

// 特殊子节点模块
const (
	SerialVerify uint32 = 0x0F
)
