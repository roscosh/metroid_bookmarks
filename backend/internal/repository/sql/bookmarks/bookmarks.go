package bookmarks

import (
	"errors"
	"fmt"
	"metroid_bookmarks/internal/repository/sql/photos"
	"metroid_bookmarks/pkg/pgpool"
	"strings"
)

const bookmarksTable = "bookmarks"

var ErrZeroID = errors.New("userId must be greater than zero")

type SQL struct {
	sql pgpool.SQL[BookmarkPreview]
}

func NewSQL(dbPool *pgpool.DbPool) *SQL {
	sql := pgpool.NewSQL[BookmarkPreview](dbPool, bookmarksTable)
	return &SQL{sql: sql}
}

func (s *SQL) Create(createForm *CreateBookmark) (*BookmarkPreview, error) {
	return s.sql.Insert(createForm)
}

func (s *SQL) Delete(id, userId int) (*BookmarkPreview, error) {
	return s.sql.DeleteWhere("id=$1 AND user_id=$2", id, userId)
}

func (s *SQL) Edit(id, userId int, editForm *EditBookmark) (*BookmarkPreview, error) {
	return s.sql.UpdateWhere(editForm, "id=$1 AND user_id=$2", id, userId)
}

func (s *SQL) GetAll(limit, offset, userId int, completed, orderById *bool) ([]Bookmark, error) {
	var bookmarks []Bookmark
	var queryArray []string
	args := make([]any, 0, 3)          //nolint:mnd
	whereArray := make([]string, 0, 3) //nolint:mnd
	placeHolder := 1
	baseQuery := `
        SELECT
			b.id, b.ctime, b.completed,
            a.id, a.name_ru, a.name_en,
            r.id, r.name_ru, r.name_en,
            s.id, s.name_ru, s.name_en,
			array_agg(p.id) AS photo_ids,
            array_agg(p.name) AS photo_names
        FROM bookmarks b
        JOIN areas a ON b.area_id = a.id
        JOIN rooms r ON b.room_id = r.id
        JOIN skills s ON b.skill_id = s.id
		LEFT JOIN photos p ON b.id = p.bookmark_id
	`
	queryArray = append(queryArray, baseQuery)

	if userId != 0 {
		whereUserId := fmt.Sprintf("b.user_id=$%d", placeHolder)
		whereArray = append(whereArray, whereUserId)
		args = append(args, userId)
		placeHolder++
	} else {
		return nil, ErrZeroID
	}
	if completed != nil {
		whereCompleted := fmt.Sprintf("b.completed=$%d", placeHolder)
		whereArray = append(whereArray, whereCompleted)
		args = append(args, *completed)
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

	rows, err := s.sql.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bookmark Bookmark
		var photoIds []*int32
		var photoNames []*string
		err = rows.Scan(
			&bookmark.Id, &bookmark.Ctime, &bookmark.Completed,
			&bookmark.Area.Id, &bookmark.Area.NameRu, &bookmark.Area.NameEn,
			&bookmark.Room.Id, &bookmark.Room.NameRu, &bookmark.Room.NameEn,
			&bookmark.Skill.Id, &bookmark.Skill.NameRu, &bookmark.Skill.NameEn,
			&photoIds, &photoNames,
		)
		if err != nil {
			return nil, err
		}
		for i, photoId := range photoIds {
			if photoId != nil {
				ulr := fmt.Sprintf("/%d/%d/%s", userId, bookmark.Id, *photoNames[i])
				bookmark.Photos = append(bookmark.Photos, photos.Photo{Id: int(*photoId), Url: ulr})
			}
		}
		bookmarks = append(bookmarks, bookmark)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func (s *SQL) GetByID(id int) (*BookmarkPreview, error) {
	return s.sql.SelectOne(id)
}

func (s *SQL) Total(userId int, completed *bool) (int, error) {
	var count int
	var queryArray []string
	args := make([]any, 0, 3)          //nolint:mnd
	whereArray := make([]string, 0, 3) //nolint:mnd
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
	}
	if completed != nil {
		whereCompleted := fmt.Sprintf("b.completed=$%d", placeHolder)
		whereArray = append(whereArray, whereCompleted)
		args = append(args, *completed)
	}
	if whereArray != nil {
		where := "WHERE " + strings.Join(whereArray, " AND ")
		queryArray = append(queryArray, where)
	}

	query := strings.Join(queryArray, " ")
	err := s.sql.QueryRow(query, args...).Scan(&count)
	return count, err
}
