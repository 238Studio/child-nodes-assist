package _const

// 数据报常数
const (
	PortLen uint32 = 512
)

// 子节点模块常数
const (
	InitModule   uint32 = 0x00
	// 初始化模块 用于处理下位机初始化相关的讯息
	SensorModule uint32 = 0x01
	// 传感器模块 用于处理下位机呈递的传感器讯息
	ReportModule uint32 = 0x02
	// 日志报告模块 用于处理下位机呈递的日志
	ErrorModule  uint32 = 0x03
	// 错误报告模块 用于处理下位机呈递的错误
	FeedbackModule uint32 = 0x0f
	// 反馈模块 用于处理下位机的数据帧传递反馈要求 例如重传 并且下位机的同名模块用于接收传来的数据帧
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
