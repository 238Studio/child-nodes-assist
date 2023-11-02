package _const

// Exceptions 错误等级的定义
const (
	NotException          = iota // 无错误
	TrivialException             // 轻微错误
	CommonException              // 一般错误
	FatalException               // 严重错误
	ExtremeFatalException        // 极其严重错误
)

// ErrorType 错误类型的定义
const (
	NoError = iota // 无错误

	JsonMarshalError   // Json序列化错误
	JsonUnmarshalError // Json反序列化错误

	HttpNewRequestError // Http请求错误
	HttpDoError         // Http请求错误

	MapKeyNotExistError // Map中的key不存在

	CreateTableError // 创建表单错误
	AddDataError     // 添加数据错误
	DeleteDataError  // 删除数据错误
	UpdateDataError  // 更新数据错误
	GetDataError     // 获取数据错误

	CreateConfigTableError // 配置文件创建错误
	DeleteConfigTableError // 配置文件删除错误
	ConfigModuleNotExist   // 配置模块不存在
	ReadConfigError        // 配置文件读取错误
	WriteConfigError       // 配置文件写入错误

	NoDeviceError             // 不存在下位机错误
	DeviceNoTargetModuleError // 指向下位机的功能模块不存在错误

	PortNotConnectedError // COM口没有连接错误

	FailedToSendToDevice // 发送讯息给下位机失败错误
	FailedToFlushPort    //刷新串口失败错误

	FailedToSendToDeviceError
)
