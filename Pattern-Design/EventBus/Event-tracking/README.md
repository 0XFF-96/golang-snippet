
### From juno 
1. http://jupiterconsole.douyu.com/workspace
2. 灵感来自于 douyu 开源的配置中心管理相关的代码。 

### 亮点
1、设计了一套 Dashboard 操作事件流记录机制
2、定义 Who 对 What 进行了 Operation 的相关操作。
记录此状态变更，以及相关事件状态。以便进行审查。 
3、

```db
type AppEvent struct {
	ID            int    `gorm:"primary_key,not null;AUTO_INCREMENT" json:"id"` // 数据id
	AppName       string `gorm:"not null;index:idx_app_name" json:"app_name"`   // 应用名称
	Aid           int    `gorm:"not null;index:idx_aid" json:"aid"`             // 应用id
	ZoneCode      string `gorm:"not null;index:idx_zone_code" json:"zone_code"` // 环境
	Env           string `gorm:"not null;index:idx_env" json:"env"`             // 环境
	HostName      string `gorm:"not null;" json:"host_name"`
	UserName      string `gorm:"not null;" json:"user_name"`                     // 用户名
	UID           int    `gorm:"not null;" json:"uid"`                           // 用户id
	Operation     string `gorm:"not null; index:idx_operation" json:"operation"` // 操作
	CreateTime    int64  `gorm:"" json:"create_time"`                            // 事件发生时间
	Source        string `gorm:"not null;index:idx_source" json:"source"`        // 事件来源
	Metadata      string `gorm:"not null;type:text" json:"metadata"`             // 事件内容
	OperationName string `gorm:"-" json:"operation_name"`
	SourceName    string `gorm:"-" json:"source_name"`
}
```