package sql

import (
	"fmt"
	"strings"
	"time"
)

const bookmarksTable = "bookmarks"

type Bookmark struct {
	Id        int       `json:"id"`
	Ctime     time.Time `json:"ctime"`
	Completed bool      `json:"completed"`
	Area      Area      `json:"area"`
	Room      Room      `json:"room"`
	Skill     Skill     `json:"skill"`
	Photos    []Photo   `json:"photos"`
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

type BookmarksSQL struct {
	*baseSQL
}

func NewBookmarksSQL(baseSQL *baseSQL) *BookmarksSQL {
	return &BookmarksSQL{baseSQL: baseSQL}
}

func (s *BookmarksSQL) Create(createForm *CreateBookmark) (*BookmarkPreview, error) {
	return insert[BookmarkPreview](s.baseSQL, bookmarksTable, *createForm)
}

func (s *BookmarksSQL) Delete(id int, userId int) (*BookmarkPreview, error) {
	return deleteWhere[BookmarkPreview](s.baseSQL, bookmarksTable, "id=$1 AND user_id=$2", id, userId)
}

func (s *BookmarksSQL) Edit(id int, userId int, editForm *EditBookmark) (*BookmarkPreview, error) {
	return updateWhere[BookmarkPreview](s.baseSQL, bookmarksTable, *editForm, "id=$1 AND user_id=$2", id, userId)
}

func (s *BookmarksSQL) GetAll(limit, offset, userId int, completed *bool, orderById *bool) ([]Bookmark, error) {
	var bookmarks []Bookmark
	var queryArray []string
	var args = make([]any, 0, 3)
	var whereArray = make([]string, 0, 3)
	placeHolder := 1
	baseQuery := `
        SELECT
			b.id, b.ctime, b.completed,
            a.id, a.name_ru, a.name_en,
            r.id, r.name_ru, r.name_en,
            s.id, s.name_ru, s.name_en,
			array_agg(p.id) AS photo_ids,
            array_agg(p.path) AS photo_paths
        FROM bookmarks b
        JOIN areas a ON b.area_id = a.id
        JOIN rooms r ON b.room_id = r.id
        JOIN skills s ON b.skill_id = s.id
		LEFT JOIN photos p ON b.id = p.bookmark_id
	`
	queryArray = append(queryArray, baseQuery)

	if userId > 0 {
		whereUserId := fmt.Sprintf("b.user_id=$%d", placeHolder)
		whereArray = append(whereArray, whereUserId)
		args = append(args, userId)
		placeHolder++
	}
	if completed != nil {
		whereCompleted := fmt.Sprintf("b.completed=$%d", placeHolder)
		whereArray = append(whereArray, whereCompleted)
		args = append(args, *completed)
		placeHolder++
	}
	if whereArray != nil {
		where := "WHERE " + strings.Join(whereArray, " AND ")
		queryArray = append(queryArray, where)
	}
	groupBy := `GROUP BY b.id, b.ctime, b.completed, a.id, a.name_ru, a.name_en, r.id, r.name_ru, r.name_en,  s.id, s.name_ru, s.name_en`
	queryArray = append(queryArray, groupBy)

	if orderById != nil {
		var order string
		if *orderById {
			order = fmt.Sprintf("ORDER BY b.id %s", "ASC")

		} else {
			order = fmt.Sprintf("ORDER BY b.id %s", "DESC")
		}
		queryArray = append(queryArray, order)
	}
	if limit > 0 {
		limitQuery := fmt.Sprintf("LIMIT %d", limit)
		queryArray = append(queryArray, limitQuery)
	}
	if offset > 0 {
		offsetQuery := fmt.Sprintf("OFFSET %d", offset)
		queryArray = append(queryArray, offsetQuery)
	}

	query := strings.Join(queryArray, " ")

	rows, err := s.baseSQL.query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bookmark Bookmark
		var photoIds []*int32
		var photoPaths []*string
		err = rows.Scan(
			&bookmark.Id, &bookmark.Ctime, &bookmark.Completed,
			&bookmark.Area.Id, &bookmark.Area.NameRu, &bookmark.Area.NameEn,
			&bookmark.Room.Id, &bookmark.Room.NameRu, &bookmark.Room.NameEn,
			&bookmark.Skill.Id, &bookmark.Skill.NameRu, &bookmark.Skill.NameEn,
			&photoIds, &photoPaths,
		)
		if err != nil {
			return nil, err
		}
		for i, _ := range photoIds {
			if photoIds[i] != nil {
				bookmark.Photos = append(bookmark.Photos, Photo{int(*photoIds[i]), *photoPaths[i]})
			}
		}
		bookmarks = append(bookmarks, bookmark)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func (s *BookmarksSQL) Total(userId int, completed *bool) (int, error) {
	var count int
	var queryArray []string
	var args = make([]any, 0, 3)
	var whereArray = make([]string, 0, 3)
	placeHolder := 1
	baseQuery := `
        SELECT COUNT(*)
        FROM bookmarks b
        JOIN areas a ON b.area_id = a.id
        JOIN rooms r ON b.room_id = r.id
        JOIN skills s ON b.skill_id = s.id
	`
	queryArray = append(queryArray, baseQuery)

	if userId > 0 {
		whereUserId := fmt.Sprintf("b.user_id=$%d", placeHolder)
		whereArray = append(whereArray, whereUserId)
		args = append(args, userId)
		placeHolder++
	}
	if completed != nil {
		whereCompleted := fmt.Sprintf("b.completed=$%d", placeHolder)
		whereArray = append(whereArray, whereCompleted)
		args = append(args, *completed)
		placeHolder++
	}
	if whereArray != nil {
		where := "WHERE " + strings.Join(whereArray, " AND ")
		queryArray = append(queryArray, where)
	}

	query := strings.Join(queryArray, " ")
	err := s.queryRow(query, args...).Scan(&count)
	return count, err
}
