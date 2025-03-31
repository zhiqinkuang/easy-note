package service

import (
	"context"
	"github.com/zhiqinkuang/easy-note/cmd/note/dal/db"
	"github.com/zhiqinkuang/easy-note/kitex_gen/demonote"
)

type UpdateNoteService struct {
	ctx context.Context
}

// NewUpdateNoteService new UpdateNoteService
func NewUpdateNoteService(ctx context.Context) *UpdateNoteService {
	return &UpdateNoteService{ctx: ctx}
}

// UpdateNote update note info
func (s *UpdateNoteService) UpdateNote(req *demonote.UpdateNoteRequest) error {
	return db.UpdateNote(s.ctx, req.NoteId, req.UserId, req.Title, req.Cotent)
}
