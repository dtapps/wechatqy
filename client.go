package wechatqy

import (
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	AppId            string
	AgentId          int
	Secret           string
	RedirectUri      string
	Key              string             // key
	ApiGormClientFun golog.ApiClientFun // 日志配置
	Debug            bool               // 日志开关
	ZapLog           *golog.ZapLog      // 日志服务
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	zapLog        *golog.ZapLog  // 日志服务
	config        struct {
		appId       string
		agentId     int
		secret      string
		redirectUri string
		key         string
	}
	log struct {
		status bool             // 状态
		client *golog.ApiClient // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.zapLog = config.ZapLog

	c.config.appId = config.AppId
	c.config.agentId = config.AgentId
	c.config.secret = config.Secret
	c.config.redirectUri = config.RedirectUri
	c.config.key = config.Key

	c.requestClient = gorequest.NewHttp()

	apiGormClient := config.ApiGormClientFun()
	if apiGormClient != nil {
		c.log.client = apiGormClient
		c.log.status = true
	}

	return c, nil
}
