// Code generated by protoc-gen-go_api(github.com/dev-openapi/protoc-gen-go_api version=v1.0.3). DO NOT EDIT.
// source: douyin-miniapp/datacache.proto

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


// Client API for Datacache service

type DatacacheService interface {
	// SetUserStorage  设置用户缓存 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/data-caching/set-user-storage
	SetUserStorage(ctx context.Context, in *SetUserStorageReq, opts ...Option) (*SetUserStorageRes, error)
	// RemoveUserStorage  清除用户缓存 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/data-caching/remove-user-storage
	RemoveUserStorage(ctx context.Context, in *RemoveUserStorageReq, opts ...Option) (*RemoveUserStorageRes, error)
}

type datacacheService struct {
	// opts
	opts *Options
}

func NewDatacacheService(opts ...Option) DatacacheService {
	opt := newOptions(opts...)
	if len(opt.addr) <= 0 {
		opt.addr = "https://developer.toutiao.com"
	}
	return &datacacheService {
		opts: opt,
	}
}


func (c *datacacheService) SetUserStorage(ctx context.Context, in *SetUserStorageReq, opts ...Option) (*SetUserStorageRes, error) {
	var res SetUserStorageRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/api/apps/set_user_storage", opt.addr)

	// body
	var body io.Reader
	bs, err := json.Marshal(in.GetBody())
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(bs)
	headers["Content-Type"] = "application/json"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	params := req.URL.Query()
	if in.GetAccessToken() != "" {
		params.Add("access_token", fmt.Sprintf("%v", in.GetAccessToken()))
	}
	if in.GetOpenid() != "" {
		params.Add("openid", fmt.Sprintf("%v", in.GetOpenid()))
	}
	if in.GetSigMethod() != "" {
		params.Add("sig_method", fmt.Sprintf("%v", in.GetSigMethod()))
	}
	if in.GetSignature() != "" {
		params.Add("signature", fmt.Sprintf("%v", in.GetSignature()))
	}	
	req.URL.RawQuery = params.Encode()
	
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

func (c *datacacheService) RemoveUserStorage(ctx context.Context, in *RemoveUserStorageReq, opts ...Option) (*RemoveUserStorageRes, error) {
	var res RemoveUserStorageRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/api/apps/remove_user_storage", opt.addr)

	// body
	var body io.Reader
	bs, err := json.Marshal(in.GetBody())
	if err != nil {
		return nil, err
	}
	body = bytes.NewReader(bs)
	headers["Content-Type"] = "application/json"

	req, err := http.NewRequest("POST", rawURL, body)
	if err != nil {
		return nil, err
	}
	
	params := req.URL.Query()
	if in.GetAccessToken() != "" {
		params.Add("access_token", fmt.Sprintf("%v", in.GetAccessToken()))
	}
	if in.GetOpenid() != "" {
		params.Add("openid", fmt.Sprintf("%v", in.GetOpenid()))
	}
	if in.GetSigMethod() != "" {
		params.Add("sig_method", fmt.Sprintf("%v", in.GetSigMethod()))
	}
	if in.GetSignature() != "" {
		params.Add("signature", fmt.Sprintf("%v", in.GetSignature()))
	}	
	req.URL.RawQuery = params.Encode()
	
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
