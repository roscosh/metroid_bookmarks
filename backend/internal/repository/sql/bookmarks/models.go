package bookmarks

import (
	"metroid_bookmarks/internal/repository/sql/areas"
	"metroid_bookmarks/internal/repository/sql/photos"
	"metroid_bookmarks/internal/repository/sql/rooms"
	"metroid_bookmarks/internal/repository/sql/skills"
	"time"
)

type Bookmark struct {
	Id        int            `json:"id"`
	Ctime     time.Time      `json:"ctime"`
	Completed bool           `json:"completed"`
	Area      areas.Area     `json:"area"`
	Room      rooms.Room     `json:"room"`
	Skill     skills.Skill   `json:"skill"`
	Photos    []photos.Photo `json:"photos"`
}

type BookmarkPreview struct {
	Id        int       `json:"id"        db:"id"`
	UserId    int       `json:"user_id"   db:"user_id"`
	AreaId    int       `json:"area_id"   db:"area_id"`
	RoomId    int       `json:"room_id"   db:"room_id"`
	SkillId   int       `json:"skill_id"  db:"skill_id"`
	Ctime     time.Time `json:"ctime"     db:"ctime"`
	Completed bool      `json:"completed" db:"completed"`
}

type CreateBookmark struct {
	UserId  int `db:"user_id"`
	AreaId  int `db:"area_id"`
	RoomId  int `db:"room_id"`
	SkillId int `db:"skill_id"`
}

type EditBookmark struct {
	AreaId    *int  `json:"area_id"  db:"area_id"`
	RoomId    *int  `json:"room_id"  db:"room_id"`
	SkillId   *int  `json:"skill_id" db:"skill_id"`
	Completed *bool `json:"completed" db:"completed"`
}
