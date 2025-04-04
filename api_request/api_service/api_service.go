// Code generated by hertz generator.

package api_service

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/protocol"
	demoapi "github.com/zhiqinkuang/easy-note/hertz_gen/demoapi"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
)

type Client interface {
	CreateUser(context context.Context, req *demoapi.CreateUserRequest, reqOpt ...config.RequestOption) (resp *demoapi.CreateUserResponse, rawResponse *protocol.Response, err error)

	CheckUser(context context.Context, req *demoapi.CheckUserRequest, reqOpt ...config.RequestOption) (resp *demoapi.CheckUserResponse, rawResponse *protocol.Response, err error)

	CreateNote(context context.Context, req *demoapi.CreateNoteRequest, reqOpt ...config.RequestOption) (resp *demoapi.CreateNoteResponse, rawResponse *protocol.Response, err error)

	QueryNote(context context.Context, req *demoapi.QueryNoteRequest, reqOpt ...config.RequestOption) (resp *demoapi.QueryNoteResponse, rawResponse *protocol.Response, err error)

	UpdateNote(context context.Context, req *demoapi.UpdateNoteRequest, reqOpt ...config.RequestOption) (resp *demoapi.UpdateNoteResponse, rawResponse *protocol.Response, err error)

	DeleteNote(context context.Context, req *demoapi.DeleteNoteRequest, reqOpt ...config.RequestOption) (resp *demoapi.DeleteNoteResponse, rawResponse *protocol.Response, err error)
}

type ApiServiceClient struct {
	client *cli
}

func NewApiServiceClient(hostUrl string, ops ...Option) (Client, error) {
	opts := getOptions(append(ops, withHostUrl(hostUrl))...)
	cli, err := newClient(opts)
	if err != nil {
		return nil, err
	}
	return &ApiServiceClient{
		client: cli,
	}, nil
}

func (s *ApiServiceClient) CreateUser(context context.Context, req *demoapi.CreateUserRequest, reqOpt ...config.RequestOption) (resp *demoapi.CreateUserResponse, rawResponse *protocol.Response, err error) {
	httpResp := &demoapi.CreateUserResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{}).
		addHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("POST", "/v1/user/register")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) CheckUser(context context.Context, req *demoapi.CheckUserRequest, reqOpt ...config.RequestOption) (resp *demoapi.CheckUserResponse, rawResponse *protocol.Response, err error) {
	httpResp := &demoapi.CheckUserResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{}).
		addHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("POST", "/v1/user/login")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) CreateNote(context context.Context, req *demoapi.CreateNoteRequest, reqOpt ...config.RequestOption) (resp *demoapi.CreateNoteResponse, rawResponse *protocol.Response, err error) {
	httpResp := &demoapi.CreateNoteResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{}).
		addHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("POST", "/v1/note")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) QueryNote(context context.Context, req *demoapi.QueryNoteRequest, reqOpt ...config.RequestOption) (resp *demoapi.QueryNoteResponse, rawResponse *protocol.Response, err error) {
	httpResp := &demoapi.QueryNoteResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{
			"user_id":    req.GetUserID(),
			"search_key": req.GetSearchKey(),
			"offset":     req.GetOffset(),
			"limit":      req.GetLimit(),
		}).
		setPathParams(map[string]string{}).
		addHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("GET", "/v1/note/query")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) UpdateNote(context context.Context, req *demoapi.UpdateNoteRequest, reqOpt ...config.RequestOption) (resp *demoapi.UpdateNoteResponse, rawResponse *protocol.Response, err error) {
	httpResp := &demoapi.UpdateNoteResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{
			"note_id": fmt.Sprint(req.GetNoteID()),
		}).
		addHeaders(map[string]string{}).
		setFormParams(map[string]string{
			"title":   req.GetTitle(),
			"content": req.GetContent(),
		}).
		setFormFileParams(map[string]string{}).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("PUT", "/v1/note/:note_id")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

func (s *ApiServiceClient) DeleteNote(context context.Context, req *demoapi.DeleteNoteRequest, reqOpt ...config.RequestOption) (resp *demoapi.DeleteNoteResponse, rawResponse *protocol.Response, err error) {
	httpResp := &demoapi.DeleteNoteResponse{}
	ret, err := s.client.r().
		setContext(context).
		setQueryParams(map[string]interface{}{}).
		setPathParams(map[string]string{
			"note_id": fmt.Sprint(req.GetNoteID()),
		}).
		addHeaders(map[string]string{}).
		setFormParams(map[string]string{}).
		setFormFileParams(map[string]string{}).
		setBodyParam(req).
		setRequestOption(reqOpt...).
		setResult(httpResp).
		execute("DELETE", "/v1/note/:note_id")
	if err != nil {
		return nil, nil, err
	}

	resp = httpResp
	rawResponse = ret.rawResponse
	return resp, rawResponse, nil
}

var defaultClient, _ = NewApiServiceClient("http://127.0.0.1:8080")

func ConfigDefaultClient(ops ...Option) (err error) {
	defaultClient, err = NewApiServiceClient("http://127.0.0.1:8080", ops...)
	return
}

func CreateUser(context context.Context, req *demoapi.CreateUserRequest, reqOpt ...config.RequestOption) (resp *demoapi.CreateUserResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.CreateUser(context, req, reqOpt...)
}

func CheckUser(context context.Context, req *demoapi.CheckUserRequest, reqOpt ...config.RequestOption) (resp *demoapi.CheckUserResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.CheckUser(context, req, reqOpt...)
}

func CreateNote(context context.Context, req *demoapi.CreateNoteRequest, reqOpt ...config.RequestOption) (resp *demoapi.CreateNoteResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.CreateNote(context, req, reqOpt...)
}

func QueryNote(context context.Context, req *demoapi.QueryNoteRequest, reqOpt ...config.RequestOption) (resp *demoapi.QueryNoteResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.QueryNote(context, req, reqOpt...)
}

func UpdateNote(context context.Context, req *demoapi.UpdateNoteRequest, reqOpt ...config.RequestOption) (resp *demoapi.UpdateNoteResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.UpdateNote(context, req, reqOpt...)
}

func DeleteNote(context context.Context, req *demoapi.DeleteNoteRequest, reqOpt ...config.RequestOption) (resp *demoapi.DeleteNoteResponse, rawResponse *protocol.Response, err error) {
	return defaultClient.DeleteNote(context, req, reqOpt...)
}
