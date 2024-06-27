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
	ID        int       `json:"id"        db:"id"`
	UserID    int       `json:"user_id"   db:"user_id"`
	AreaID    int       `json:"area_id"   db:"area_id"`
	RoomID    int       `json:"room_id"   db:"room_id"`
	SkillID   int       `json:"skill_id"  db:"skill_id"`
	Ctime     time.Time `json:"ctime"     db:"ctime"`
	Completed bool      `json:"completed" db:"completed"`
}

type CreateBookmark struct {
	UserID  int `db:"user_id"`
	AreaID  int `db:"area_id"`
	RoomID  int `db:"room_id"`
	SkillID int `db:"skill_id"`
}

type EditBookmark struct {
	AreaID    *int  `json:"area_id"  db:"area_id"`
	RoomID    *int  `json:"room_id"  db:"room_id"`
	SkillID   *int  `json:"skill_id" db:"skill_id"`
	Completed *bool `json:"completed" db:"completed"`
}
