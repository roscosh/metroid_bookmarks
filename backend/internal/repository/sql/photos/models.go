package photos

type Photo struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

type PhotoPreview struct {
	ID         int    `db:"id"          json:"id"`
	BookmarkID int    `db:"bookmark_id" json:"bookmark_id"`
	Name       string `db:"name"        json:"name"`
}

type CreatePhoto struct {
	BookmarkID int    `db:"bookmark_id"`
	Name       string `db:"name"`
}
