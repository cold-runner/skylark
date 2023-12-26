package mysql

import (
	"context"
	"github.com/cold-runner/skylark/biz/infrastructure/store/orm_gen"
	"github.com/google/uuid"
)

func (m *mysqlIns) CreateComment(c context.Context, comment *orm_gen.Comment) error {
	return m.Query.Comment.WithContext(c).Create(comment)
}

func (m *mysqlIns) DeleteComment(c context.Context, commentId string) error {
	uid, _ := uuid.Parse(commentId)
	tar := &orm_gen.Comment{ID: uid}
	if _, err := m.Query.Comment.WithContext(c).Delete(tar); err != nil {
		return err
	}
	return nil
}

func (m *mysqlIns) UpdateComment(c context.Context, commentId, content string) error {
	uid, _ := uuid.Parse(commentId)
	t := m.Query.Comment
	if _, err := t.WithContext(c).Where(t.ID.Eq(uid)).Update(t.Content, content); err != nil {
		return err
	}
	return nil
}

func (m *mysqlIns) QueryComment(c context.Context, postId string, idx, limit int) ([]*orm_gen.Comment, error) {
	t := m.Query.Comment
	comments, err := t.WithContext(c).Where(t.PostID.Eq(postId)).Offset(idx).Limit(limit).Find()
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (m *mysqlIns) QueryReplyComment(c context.Context, postId, parentId string, idx, limit int) ([]*orm_gen.Comment, error) {
	t := m.Query.Comment
	comments, err := t.WithContext(c).Where(t.PostID.Eq(postId), t.ParentID.Eq(parentId)).Offset(idx).Limit(limit).Find()
	if err != nil {
		return nil, err
	}
	return comments, nil
}
