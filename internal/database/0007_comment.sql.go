// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 0007_comment.sql

package database

import (
	"context"
	"database/sql"
)

const createComment = `-- name: CreateComment :exec
INSERT INTO Comment (
    post_id, 
    user_id, 
    comment_content, 
    comment_left, 
    comment_right, 
    comment_parent, 
    isDeleted
) 
VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreateCommentParams struct {
	PostID         uint64
	UserID         uint64
	CommentContent sql.NullString
	CommentLeft    sql.NullInt32
	CommentRight   sql.NullInt32
	CommentParent  sql.NullInt32
	Isdeleted      sql.NullBool
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) error {
	_, err := q.db.ExecContext(ctx, createComment,
		arg.PostID,
		arg.UserID,
		arg.CommentContent,
		arg.CommentLeft,
		arg.CommentRight,
		arg.CommentParent,
		arg.Isdeleted,
	)
	return err
}

const deleteCommentsInRange = `-- name: DeleteCommentsInRange :exec
delete from Comment
where post_id = ?
and comment_left >= ?
and comment_right <= ?
`

type DeleteCommentsInRangeParams struct {
	PostID       uint64
	CommentLeft  sql.NullInt32
	CommentRight sql.NullInt32
}

func (q *Queries) DeleteCommentsInRange(ctx context.Context, arg DeleteCommentsInRangeParams) error {
	_, err := q.db.ExecContext(ctx, deleteCommentsInRange, arg.PostID, arg.CommentLeft, arg.CommentRight)
	return err
}

const getCommentByID = `-- name: GetCommentByID :one
SELECT id, post_id, user_id, comment_content, comment_left, comment_right, comment_parent, isdeleted, created_at, updated_at 
FROM Comment
WHERE id = ?
`

func (q *Queries) GetCommentByID(ctx context.Context, id int32) (Comment, error) {
	row := q.db.QueryRowContext(ctx, getCommentByID, id)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.UserID,
		&i.CommentContent,
		&i.CommentLeft,
		&i.CommentRight,
		&i.CommentParent,
		&i.Isdeleted,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCommentByParentID = `-- name: GetCommentByParentID :many
select c.id, c.post_id, c.user_id, c.comment_content, c.comment_left, c.comment_right, c.comment_parent, c.isdeleted, c.created_at, c.updated_at from Comment c 
where c.post_id = ?
and c.comment_left > ( select sub.comment_left from Comment sub where sub.id = ?)
and c.comment_right < (select sub.comment_right from Comment sub where sub.id = ?)
and c.isDeleted = false
order by c.comment_left
limit 10 
offset 0
`

type GetCommentByParentIDParams struct {
	PostID uint64
	ID     int32
	ID_2   int32
}

func (q *Queries) GetCommentByParentID(ctx context.Context, arg GetCommentByParentIDParams) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, getCommentByParentID, arg.PostID, arg.ID, arg.ID_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.CommentContent,
			&i.CommentLeft,
			&i.CommentRight,
			&i.CommentParent,
			&i.Isdeleted,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMaxRightComment = `-- name: GetMaxRightComment :one
select comment_right from Comment 
where post_id = ?
order by comment_right desc
limit 1
`

func (q *Queries) GetMaxRightComment(ctx context.Context, postID uint64) (sql.NullInt32, error) {
	row := q.db.QueryRowContext(ctx, getMaxRightComment, postID)
	var comment_right sql.NullInt32
	err := row.Scan(&comment_right)
	return comment_right, err
}

const updateCommentLeft = `-- name: UpdateCommentLeft :exec
update Comment 
set comment_left = comment_left - ?
where post_id = ?
and comment_left >= ?
`

type UpdateCommentLeftParams struct {
	CommentLeft   sql.NullInt32
	PostID        uint64
	CommentLeft_2 sql.NullInt32
}

func (q *Queries) UpdateCommentLeft(ctx context.Context, arg UpdateCommentLeftParams) error {
	_, err := q.db.ExecContext(ctx, updateCommentLeft, arg.CommentLeft, arg.PostID, arg.CommentLeft_2)
	return err
}

const updateCommentLeftCreate = `-- name: UpdateCommentLeftCreate :exec
update Comment
set comment_left = comment_left + 2
where post_id = ? and comment_left > $2
`

func (q *Queries) UpdateCommentLeftCreate(ctx context.Context, postID uint64) error {
	_, err := q.db.ExecContext(ctx, updateCommentLeftCreate, postID)
	return err
}

const updateCommentRight = `-- name: UpdateCommentRight :exec
update Comment
set comment_right = comment_right - ?
where post_id = ?
and comment_right >= ?
`

type UpdateCommentRightParams struct {
	CommentRight   sql.NullInt32
	PostID         uint64
	CommentRight_2 sql.NullInt32
}

func (q *Queries) UpdateCommentRight(ctx context.Context, arg UpdateCommentRightParams) error {
	_, err := q.db.ExecContext(ctx, updateCommentRight, arg.CommentRight, arg.PostID, arg.CommentRight_2)
	return err
}

const updateCommentRightCreate = `-- name: UpdateCommentRightCreate :exec
update Comment
set comment_right = comment_right + 2
where post_id = ? and comment_right >= $2
`

func (q *Queries) UpdateCommentRightCreate(ctx context.Context, postID uint64) error {
	_, err := q.db.ExecContext(ctx, updateCommentRightCreate, postID)
	return err
}
