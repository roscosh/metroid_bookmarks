package photos

type Photo struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

type PhotoPreview struct {
	ID         int    `json:"id"          db:"id"`
	BookmarkID int    `json:"bookmark_id" db:"bookmark_id"`
	Name       string `json:"name"        db:"name"`
}

type CreatePhoto struct {
	BookmarkID int    `db:"bookmark_id"`
	Name       string `db:"name"`
}
