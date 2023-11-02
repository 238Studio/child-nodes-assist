package _const

// 数据报常数
const (
	PortHead   byte = 0xEF
	PortEnd    byte = 0xFF
	PortMaxLen      = 0xFFFF
)

// 子节点模块常数
const (
	InitModule   byte = 0x00
	SensorModule byte = 0x01
	ReportModule byte = 0x02
	ErrorModule  byte = 0x03
)

// 特殊数据报
const (
	FailedToRev = "FailedToRev"
)

// 特殊子节点模块
const (
	SerialVerify byte = 0x0F
)
