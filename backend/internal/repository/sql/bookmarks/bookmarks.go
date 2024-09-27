package bookmarks

import (
	"errors"
	"fmt"
	"metroid_bookmarks/internal/repository/sql/pgerr"
	"metroid_bookmarks/internal/repository/sql/photos"
	"metroid_bookmarks/pkg/pgpool"
	"strings"
)

const bookmarksTable = "bookmarks"

var ErrZeroID = errors.New("userId must be greater than zero")

type SQL interface {
	Create(createForm *CreateBookmark) (*BookmarkPreview, error)
	Delete(id int, userID int) (*BookmarkPreview, error)
	Edit(id, userID int, editForm *EditBookmark) (*BookmarkPreview, error)
	GetAll(limit, offset, userID int, completed, orderByID *bool) ([]Bookmark, error)
	GetByID(id int) (*BookmarkPreview, error)
	Total(userID int, completed *bool) (int, error)
}

type bookmarksSQL struct {
	sql pgpool.SQL[BookmarkPreview]
}

func NewSQL(dbPool *pgpool.PgPool) SQL {
	sql := pgpool.NewSQL[BookmarkPreview](dbPool, bookmarksTable)
	return &bookmarksSQL{sql: sql}
}

func (s *bookmarksSQL) Create(createForm *CreateBookmark) (*BookmarkPreview, error) {
	entity, err := s.sql.Insert(createForm)
	if err != nil {
		err = pgerr.CreatePgError(err)
		return nil, err
	}

	return entity, nil
}

func (s *bookmarksSQL) Delete(id, userID int) (*BookmarkPreview, error) {
	entity, err := s.sql.DeleteWhere("id=$1 AND user_id=$2", id, userID)
	if err != nil {
		err = pgerr.DeletePgError(err, id)
		return nil, err
	}

	return entity, nil
}

func (s *bookmarksSQL) Edit(id, userID int, editForm *EditBookmark) (*BookmarkPreview, error) {
	entity, err := s.sql.UpdateWhere(editForm, "id=$1 AND user_id=$2", id, userID)
	if err != nil {
		err = pgerr.EditPgError(err, id)
		return nil, err
	}

	return entity, nil
}

func (s *bookmarksSQL) GetAll(limit, offset, userID int, completed, orderByID *bool) ([]Bookmark, error) {
	var (
		bookmarks  []Bookmark
		queryArray []string
	)

	args := make([]any, 0, 2)          //nolint:mnd
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

	if userID != 0 {
		whereUserID := fmt.Sprintf("b.user_id=$%d", placeHolder)
		whereArray = append(whereArray, whereUserID)
		args = append(args, userID)
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

	if orderByID != nil {
		order := "ORDER BY b.id "
		if *orderByID {
			order += "ASC"
		} else {
			order += "DESC"
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
		var (
			bookmark   Bookmark
			photoIDs   []*int32
			photoNames []*string
		)

		err = rows.Scan(
			&bookmark.ID, &bookmark.Ctime, &bookmark.Completed,
			&bookmark.Area.ID, &bookmark.Area.NameRu, &bookmark.Area.NameEn,
			&bookmark.Room.ID, &bookmark.Room.NameRu, &bookmark.Room.NameEn,
			&bookmark.Skill.ID, &bookmark.Skill.NameRu, &bookmark.Skill.NameEn,
			&photoIDs, &photoNames,
		)
		if err != nil {
			return nil, err
		}

		for i, photoID := range photoIDs {
			if photoID != nil {
				ulr := fmt.Sprintf("/%d/%d/%s", userID, bookmark.ID, *photoNames[i])
				bookmark.Photos = append(bookmark.Photos, photos.Photo{ID: int(*photoID), URL: ulr})
			}
		}

		bookmarks = append(bookmarks, bookmark)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func (s *bookmarksSQL) GetByID(id int) (*BookmarkPreview, error) {
	entity, err := s.sql.SelectOne(id)
	if err != nil {
		err = pgerr.SelectPgError(err, id)
		return nil, err
	}

	return entity, nil
}

func (s *bookmarksSQL) Total(userID int, completed *bool) (int, error) {
	var (
		count      int
		queryArray []string
	)

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

	if userID > 0 {
		whereUserID := fmt.Sprintf("b.user_id=$%d", placeHolder)
		whereArray = append(whereArray, whereUserID)
		args = append(args, userID)
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
