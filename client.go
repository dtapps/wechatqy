package wechatqy

import (
	"go.dtapp.net/golog"
)

// ClientConfig 实例配置
type ClientConfig struct {
	AppId       string
	AgentId     int
	Secret      string
	RedirectUri string
}

// Client 实例
type Client struct {
	config struct {
		appId       string
		agentId     int
		secret      string
		redirectUri string
	}
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
	mongoLog struct {
		status bool            // 状态
		client *golog.ApiMongo // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {
	c := &Client{}
	c.config.appId = config.AppId
	c.config.agentId = config.AgentId
	c.config.secret = config.Secret
	c.config.redirectUri = config.RedirectUri
	return c, nil
}
