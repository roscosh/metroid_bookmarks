package photos

type Photo struct {
	URL string `json:"url"`
	ID  int    `json:"id"`
}

type PhotoPreview struct {
	Name       string `db:"name"        json:"name"`
	ID         int    `db:"id"          json:"id"`
	BookmarkID int    `db:"bookmark_id" json:"bookmark_id"`
}

type CreatePhoto struct {
	Name       string `db:"name"`
	BookmarkID int    `db:"bookmark_id"`
}
