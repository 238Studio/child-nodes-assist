package _const

// 数据报常数
const (
	PortLen uint32 = 512
)

// 子节点模块常数
const (
	// InitModule 初始化模块 用于处理下位机初始化相关的讯息
	InitModule uint32 = 0x00
	// SensorModule 传感器模块 用于处理下位机呈递的传感器讯息
	SensorModule uint32 = 0x01
	// ReportModule 日志报告模块 用于处理下位机呈递的日志
	ReportModule uint32 = 0x02
	// ErrorModule 错误报告模块 用于处理下位机呈递的错误
	ErrorModule uint32 = 0x03
	// FeedbackModule 反馈模块 用于处理下位机的数据帧传递反馈要求 例如重传 并且下位机的同名模块用于接收传来的数据帧
	FeedbackModule uint32 = 0x0f
)

// 特殊数据报
const (
	// InitData 初始化数据
	InitData string = "InitData"
	// ReSendData 重新发送的数据
	ReSendData string = "ReSendData"
	// WrongOddVariation 错误的校验
	WrongOddVariation string = "WrongOddVariation"
)
