package photos

type Photo struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

type PhotoPreview struct {
	Id         int    `json:"id"          db:"id"`
	BookmarkId int    `json:"bookmark_id" db:"bookmark_id"`
	Name       string `json:"name"        db:"name"`
}

type CreatePhoto struct {
	BookmarkId int    `db:"bookmark_id"`
	Name       string `db:"name"`
}
