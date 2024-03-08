package wechatqy

import (
	"context"
	"fmt"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"net/http"
)

type WebhookSendResponse struct {
	Errcode   int64  `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

type WebhookSendResult struct {
	Result WebhookSendResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newWebhookSendResult(result WebhookSendResponse, body []byte, http gorequest.Response) *WebhookSendResult {
	return &WebhookSendResult{Result: result, Body: body, Http: http}
}

// WebhookSend 发送消息
// https://developer.work.weixin.qq.com/document/path/90372
func (c *Client) WebhookSend(ctx context.Context, key string, Type string, notMustParams ...gorequest.Params) (*WebhookSendResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+fmt.Sprintf("/cgi-bin/webhook/send?key=%s&type=%s", key, Type), params, http.MethodPost)
	if err != nil {
		return newWebhookSendResult(WebhookSendResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WebhookSendResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWebhookSendResult(response, request.ResponseBody, request), err
}
