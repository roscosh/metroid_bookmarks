package rooms

type Room struct {
	ID     int    `db:"id"      json:"id"`
	NameEn string `db:"name_en" json:"name_en"`
	NameRu string `db:"name_ru" json:"name_ru"`
}

type CreateRoom struct {
	NameEn string `binding:"required" db:"name_en" json:"name_en"`
	NameRu string `binding:"required" db:"name_ru" json:"name_ru"`
}

type EditRoom struct {
	NameEn *string `db:"name_en" json:"name_en"`
	NameRu *string `db:"name_ru" json:"name_ru"`
}
