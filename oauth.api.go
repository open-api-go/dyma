// Code generated by protoc-gen-go_api(github.com/dev-openapi/protoc-gen-go_api version=v1.0.3). DO NOT EDIT.
// source: douyin-miniapp/oauth.proto

package douyin_miniapp

import (
	context "context"
	fmt "fmt"
	io "io"
	json "encoding/json"
	bytes "bytes"
	http "net/http"
	strings "strings"
	url "net/url"
)
// Reference imports to suppress errors if they are not otherwise used.
var _ = context.Background
var _ = http.NewRequest
var _ = io.Copy
var _ = bytes.Compare
var _ = json.Marshal
var _ = strings.Compare
var _ = fmt.Errorf
var _ = url.Parse


// Client API for Oauth service

type OauthService interface {
	// GetClientToken  生成client_token https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/account-permission/client-token
	GetClientToken(ctx context.Context, in *GetClientTokenReq, opts ...Option) (*GetClientTokenRes, error)
}

type oauthService struct {
	// opts
	opts *Options
}

func NewOauthService(opts ...Option) OauthService {
	opt := newOptions(opts...)
	if len(opt.addr) <= 0 {
		opt.addr = "https://open.douyin.com"
	}
	return &oauthService {
		opts: opt,
	}
}


func (c *oauthService) GetClientToken(ctx context.Context, in *GetClientTokenReq, opts ...Option) (*GetClientTokenRes, error) {
	var res GetClientTokenRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/oauth/client_token/", opt.addr)

	// body
	var body io.Reader
	bodyForms := url.Values{} 
	if in.GetClientKey() != "" {
		bodyForms.Add("client_key", fmt.Sprintf("%v", in.GetClientKey()))
	}
	if in.GetClientSecret() != "" {
		bodyForms.Add("client_secret", fmt.Sprintf("%v", in.GetClientSecret()))
	}
	if in.GetGrantType() != "" {
		bodyForms.Add("grant_type", fmt.Sprintf("%v", in.GetGrantType()))
	}
	body = strings.NewReader(bodyForms.Encode())
	headers["Content-Type"] = "multipart/form-data"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	// header
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := opt.DoRequest(ctx, opt.client, req)
	if err != nil {
		return nil, err
	}
	err = opt.DoResponse(ctx, resp, &res)
	return &res, err 

}
