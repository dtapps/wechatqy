package wechatqy

import (
	"go.dtapp.net/dorm"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

type ConfigClient struct {
	AppId       string
	AgentId     int
	Secret      string
	RedirectUri string
	Key         string           // key
	GormClient  *dorm.GormClient // 日志数据库
	LogClient   *golog.ZapLog    // 日志驱动
	LogDebug    bool             // 日志开关
}

type Client struct {
	requestClient *gorequest.App   // 请求服务
	logClient     *golog.ApiClient // 日志服务
	config        *ConfigClient    // 配置
}

func NewClient(config *ConfigClient) (*Client, error) {

	var err error
	c := &Client{config: config}

	c.requestClient = gorequest.NewHttp()

	if c.config.GormClient.Db != nil {
		c.logClient, err = golog.NewApiClient(&golog.ApiClientConfig{
			GormClient: c.config.GormClient,
			TableName:  logTable,
			LogClient:  c.config.LogClient,
			LogDebug:   c.config.LogDebug,
		})
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}
