// Code generated by hertz generator.

package demoapi

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/zhiqinkuang/easy-note/cmd/api/mw"
	"github.com/zhiqinkuang/easy-note/cmd/api/rpc"
	"github.com/zhiqinkuang/easy-note/kitex_gen/demonote"
	"github.com/zhiqinkuang/easy-note/kitex_gen/demouser"
	"github.com/zhiqinkuang/easy-note/pkg/consts"
	"github.com/zhiqinkuang/easy-note/pkg/errno"

	demoapi "github.com/zhiqinkuang/easy-note/hertz_gen/demoapi"
)

func CreateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demoapi.CreateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	err = rpc.CreateUser(ctx, &demouser.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

// CheckUser .
// @router /v1/user/login [POST]
func CheckUser(ctx context.Context, c *app.RequestContext) {
	mw.JwtMiddleware.LoginHandler(ctx, c)
}

// CreateNote .
// @router /v1/note [POST]
func CreateNote(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demoapi.CreateNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	err = rpc.CreateNote(ctx, &demonote.CreateNoteRequest{
		Title:   req.Title,
		Content: req.Content,
		UserId:  v.(*demoapi.User).UserID,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

// QueryNote .
// @router /v1/note/query [GET]
func QueryNote(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demoapi.QueryNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	notes, total, err := rpc.QueryNotes(ctx, &demonote.QueryNoteRequest{
		UserId:    v.(*demoapi.User).UserID,
		SearchKey: req.SearchKey,
		Offset:    req.Offset,
		Limit:     req.Limit,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, utils.H{
		consts.Total: total,
		consts.Notes: notes,
	})
}

// UpdateNote .
// @router /v1/note/:note_id [PUT]
func UpdateNote(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demoapi.UpdateNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	err = rpc.UpdateNote(ctx, &demonote.UpdateNoteRequest{
		NoteId: req.NoteID,
		UserId: v.(*demoapi.User).UserID,
		Title:  req.Title,
		Cotent: req.Content,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

// DeleteNote .
// @router /v1/note/:note_id [DELETE]
func DeleteNote(ctx context.Context, c *app.RequestContext) {
	var err error
	var req demoapi.DeleteNoteRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	v, _ := c.Get(consts.IdentityKey)
	err = rpc.DeleteNote(ctx, &demonote.DeleteNoteRequest{
		NoteId: req.NoteID,
		UserId: v.(*demoapi.User).UserID,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
