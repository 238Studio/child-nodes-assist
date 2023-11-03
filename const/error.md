# 错误常量

## 错误等级

| 错误等级                    | 值 | 描述                            |
|-------------------------|---|-------------------------------|
| `NotException`          | 0 | 没有错误                          |
| `TrivialException`      | 1 | 轻微错误。直接传入数据库。                 |
| `CommonException`       | 2 | 一般错误。需要进行特殊处理。                |
| `FatalException`        | 3 | 严重错误。被`recover`或者入参检测所捕获的错误。  |
| `ExtremeFatalException` | 4 | 极其严重的错误。此类错误一般无法捕获，会直接导致程序崩溃。 |

## 错误类型

| 错误类型                        | 值  | 描述            |
|-----------------------------|----|---------------|
| `NoError`                   | 0  | 没有错误          |
| `JsonMarshalError`          | 1  | JSON序列化错误     |
| `JsonUnmarshalError`        | 2  | JSON反序列化错误    |
| `HttpNewRequestError`       | 3  | HTTP请求错误      |
| `HttpDoError`               | 4  | HTTP请求错误      |
| `MapKeyNotExistError`       | 5  | Map键不存在       |
| `CreateTableError`          | 6  | 创建表错误         |
| `AddDataError`              | 7  | 添加数据错误        |
| `DeleteDataError`           | 8  | 删除数据错误        |
| `UpdateDataError`           | 9  | 更新数据错误        |
| `GetDataError`              | 10 | 获取数据错误        |
| `CreateConfigTableError`    | 11 | 创建配置文件错误      |
| `DeleteConfigTableError`    | 12 | 删除配置文件错误      |
| `ConfigModuleNotExist`      | 13 | 配置模块不存在       |
| `ReadConfigError`           | 14 | 读取配置文件错误      |
| `WriteConfigError`          | 15 | 写入配置文件错误      |
| `NoDeviceError`             | 16 | 不存在的下位机       |
| `DeviceNoTargetModuleError` | 17 | 指向的下位机没有该功能模块 |
| `PortNotConnectedError`     | 18 | COM口没有连接      |
| `FailedToSendToDevice`      | 19 | 发送数据到下位机失败    |
| `FailedToFlushPort`         | 20 | 刷新COM口失败      |
| `FailedToSendToDeviceError` | 21 | 发送数据到下位机失败    |
| `NilPointer`                | 22 | 空指针           |