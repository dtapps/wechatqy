package wechatqy

import (
	"context"
	"fmt"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
	"go.opentelemetry.io/otel/codes"
	"net/http"
)

type CgiBinUserGetUserInfoResponse struct {
	Errcode        int    `json:"errcode"`         // 返回码
	Errmsg         string `json:"errmsg"`          // 对返回码的文本描述内容
	UserId         string `json:"userId"`          // 成员UserID。若需要获得用户详情信息
	OpenId         string `json:"OpenId"`          // 非企业成员的标识，对当前企业唯一。不超过64字节
	DeviceId       string `json:"DeviceId"`        // 手机设备号(由企业微信在安装时随机生成，删除重装会改变，升级不受影响)
	ExternalUserid string `json:"external_userid"` // 外部联系人id，当且仅当用户是企业的客户，且跟进人在应用的可见范围内时返回。如果是第三方应用调用，针对同一个客户，同一个服务商不同应用获取到的id相同
}

type CgiBinUserGetUserInfoResult struct {
	Result CgiBinUserGetUserInfoResponse // 结果
	Body   []byte                        // 内容
	Http   gorequest.Response            // 请求
}

func newCgiBinUserGetUserInfoResult(result CgiBinUserGetUserInfoResponse, body []byte, http gorequest.Response) *CgiBinUserGetUserInfoResult {
	return &CgiBinUserGetUserInfoResult{Result: result, Body: body, Http: http}
}

// CgiBinUserGetUserInfo 获取访问用户身份
// https://open.work.weixin.qq.com/api/doc/90000/90135/91023
func (c *Client) CgiBinUserGetUserInfo(ctx context.Context, accessToken, code string, notMustParams ...gorequest.Params) (*CgiBinUserGetUserInfoResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "cgi-bin/user/getuserinfo")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	request, err := c.request(ctx, fmt.Sprintf("cgi-bin/user/getuserinfo?access_token=%s&code=%s", accessToken, code), params, http.MethodGet)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
		return newCgiBinUserGetUserInfoResult(CgiBinUserGetUserInfoResponse{}, request.ResponseBody, request), err
	}

	// 定义
	var response CgiBinUserGetUserInfoResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
	}
	return newCgiBinUserGetUserInfoResult(response, request.ResponseBody, request), err
}
