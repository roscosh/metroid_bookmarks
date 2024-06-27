package bookmarks

import (
	"metroid_bookmarks/internal/repository/sql/areas"
	"metroid_bookmarks/internal/repository/sql/photos"
	"metroid_bookmarks/internal/repository/sql/rooms"
	"metroid_bookmarks/internal/repository/sql/skills"
	"time"
)

type Bookmark struct {
	ID        int            `json:"id"`
	Ctime     time.Time      `json:"ctime"`
	Completed bool           `json:"completed"`
	Area      areas.Area     `json:"area"`
	Room      rooms.Room     `json:"room"`
	Skill     skills.Skill   `json:"skill"`
	Photos    []photos.Photo `json:"photos"`
}

type BookmarkPreview struct {
	ID        int       `db:"id"        json:"id"`
	UserID    int       `db:"user_id"   json:"user_id"`
	AreaID    int       `db:"area_id"   json:"area_id"`
	RoomID    int       `db:"room_id"   json:"room_id"`
	SkillID   int       `db:"skill_id"  json:"skill_id"`
	Ctime     time.Time `db:"ctime"     json:"ctime"`
	Completed bool      `db:"completed" json:"completed"`
}

type CreateBookmark struct {
	UserID  int `db:"user_id"`
	AreaID  int `db:"area_id"`
	RoomID  int `db:"room_id"`
	SkillID int `db:"skill_id"`
}

type EditBookmark struct {
	AreaID    *int  `db:"area_id"   json:"area_id"`
	RoomID    *int  `db:"room_id"   json:"room_id"`
	SkillID   *int  `db:"skill_id"  json:"skill_id"`
	Completed *bool `db:"completed" json:"completed"`
}
