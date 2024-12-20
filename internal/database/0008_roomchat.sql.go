// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 0008_roomchat.sql

package database

import (
	"context"
)

const addMemberToRoomChat = `-- name: AddMemberToRoomChat :exec
insert into room_members (user_id, room_id)
values (?,?)
`

type AddMemberToRoomChatParams struct {
	UserID uint64
	RoomID int64
}

func (q *Queries) AddMemberToRoomChat(ctx context.Context, arg AddMemberToRoomChatParams) error {
	_, err := q.db.ExecContext(ctx, addMemberToRoomChat, arg.UserID, arg.RoomID)
	return err
}

const createRoomChat = `-- name: CreateRoomChat :exec
insert into room_chats(name, is_group, admin_id, avatar_url)
values (?,?,?,?)
`

type CreateRoomChatParams struct {
	Name      string
	IsGroup   bool
	AdminID   uint64
	AvatarUrl string
}

func (q *Queries) CreateRoomChat(ctx context.Context, arg CreateRoomChatParams) error {
	_, err := q.db.ExecContext(ctx, createRoomChat,
		arg.Name,
		arg.IsGroup,
		arg.AdminID,
		arg.AvatarUrl,
	)
	return err
}

const deleteMemberFromRoomChat = `-- name: DeleteMemberFromRoomChat :exec
delete from room_members
where user_id = ? and room_id = ?
`

type DeleteMemberFromRoomChatParams struct {
	UserID uint64
	RoomID int64
}

func (q *Queries) DeleteMemberFromRoomChat(ctx context.Context, arg DeleteMemberFromRoomChatParams) error {
	_, err := q.db.ExecContext(ctx, deleteMemberFromRoomChat, arg.UserID, arg.RoomID)
	return err
}
