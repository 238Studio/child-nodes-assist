package _const

// 数据报常数
const (
	PortHead   byte   = 0xEF
	PortEnd    byte   = 0xFF
	PortMaxLen uint32 = 0xFFFF
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
	FailedToRev = "FailedToRev"
)

// 特殊子节点模块
const (
	SerialVerify uint32 = 0x0F
)
