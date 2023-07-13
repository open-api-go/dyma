// Code generated by protoc-gen-go_api(github.com/dev-openapi/protoc-gen-go_api version=v1.0.3). DO NOT EDIT.
// source: douyin-miniapp/customer.proto

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
	multipart "mime/multipart"
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
var _ = multipart.ErrMessageTooLarge


// Client API for Customer service

type CustomerService interface {
	// GetCustomer  获取官方平台客服链接 https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/server/im/customer-service-url
	GetCustomer(ctx context.Context, in *GetCustomerReq, opts ...Option) (*GetCustomerRes, error)
}

type customerService struct {
	// opts
	opts *Options
}

func NewCustomerService(opts ...Option) CustomerService {
	opt := newOptions(opts...)
	if len(opt.addr) <= 0 {
		opt.addr = "https://developer.toutiao.com"
	}
	return &customerService {
		opts: opt,
	}
}


func (c *customerService) GetCustomer(ctx context.Context, in *GetCustomerReq, opts ...Option) (*GetCustomerRes, error) {
	var res GetCustomerRes
	// options
	opt := buildOptions(c.opts, opts...)
	headers := make(map[string]string)
	// route
	rawURL := fmt.Sprintf("%s/api/apps/chat/customer_service_url", opt.addr)

	// body
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return nil, err
	}
	
	params := req.URL.Query()
	if in.GetAppid() != "" {
		params.Add("appid", fmt.Sprintf("%v", in.GetAppid()))
	}
	if in.GetImType() != "" {
		params.Add("im_type", fmt.Sprintf("%v", in.GetImType()))
	}
	if in.GetOpenid() != "" {
		params.Add("openid", fmt.Sprintf("%v", in.GetOpenid()))
	}
	if in.GetOrderId() != "" {
		params.Add("order_id", fmt.Sprintf("%v", in.GetOrderId()))
	}
	if in.GetScene() != "" {
		params.Add("scene", fmt.Sprintf("%v", in.GetScene()))
	}
	if in.GetType() != "" {
		params.Add("type", fmt.Sprintf("%v", in.GetType()))
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
