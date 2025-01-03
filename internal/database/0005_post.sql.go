// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 0005_post.sql

package database

import (
	"context"
	"database/sql"
	"encoding/json"
)

const createPost = `-- name: CreatePost :exec
INSERT INTO post (title, image_paths, user_id, created_at, updated_at,user_nickname)
VALUES (?, ?, ?, NOW(), NOW(),?)
`

type CreatePostParams struct {
	Title        string
	ImagePaths   json.RawMessage
	UserID       uint64
	UserNickname string
}

// Create a new post
func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) error {
	_, err := q.db.ExecContext(ctx, createPost,
		arg.Title,
		arg.ImagePaths,
		arg.UserID,
		arg.UserNickname,
	)
	return err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM post
WHERE id = ?
`

// Delete a post
func (q *Queries) DeletePost(ctx context.Context, id uint64) error {
	_, err := q.db.ExecContext(ctx, deletePost, id)
	return err
}

const getAllpost = `-- name: GetAllpost :many
SELECT id, title, image_paths, user_nickname, created_at, updated_at
FROM post
order by created_at desc
`

type GetAllpostRow struct {
	ID           uint64
	Title        string
	ImagePaths   json.RawMessage
	UserNickname string
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
}

// Get all post
func (q *Queries) GetAllpost(ctx context.Context) ([]GetAllpostRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllpost)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllpostRow
	for rows.Next() {
		var i GetAllpostRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.ImagePaths,
			&i.UserNickname,
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

const getPostById = `-- name: GetPostById :one
SELECT id, title, image_paths, user_nickname, created_at, updated_at
FROM post
WHERE id = ?
`

type GetPostByIdRow struct {
	ID           uint64
	Title        string
	ImagePaths   json.RawMessage
	UserNickname string
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
}

// Get post by ID
func (q *Queries) GetPostById(ctx context.Context, id uint64) (GetPostByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getPostById, id)
	var i GetPostByIdRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ImagePaths,
		&i.UserNickname,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getpostByUserId = `-- name: GetpostByUserId :many
SELECT id, title, image_paths, user_id, created_at, updated_at
FROM post
WHERE user_id = ?
`

type GetpostByUserIdRow struct {
	ID         uint64
	Title      string
	ImagePaths json.RawMessage
	UserID     uint64
	CreatedAt  sql.NullTime
	UpdatedAt  sql.NullTime
}

// Get post by user ID
func (q *Queries) GetpostByUserId(ctx context.Context, userID uint64) ([]GetpostByUserIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getpostByUserId, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetpostByUserIdRow
	for rows.Next() {
		var i GetpostByUserIdRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.ImagePaths,
			&i.UserID,
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

const updatePost = `-- name: UpdatePost :exec
UPDATE post
SET title = ?, image_paths = ?, updated_at = NOW()
WHERE id = ?
`

type UpdatePostParams struct {
	Title      string
	ImagePaths json.RawMessage
	ID         uint64
}

// Update a post
func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) error {
	_, err := q.db.ExecContext(ctx, updatePost, arg.Title, arg.ImagePaths, arg.ID)
	return err
}
