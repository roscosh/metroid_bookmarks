package areas

type Area struct {
	ID     int    `db:"id"      json:"id"`
	NameEn string `db:"name_en" json:"name_en"`
	NameRu string `db:"name_ru" json:"name_ru"`
}

type CreateArea struct {
	NameEn string `binding:"required" db:"name_en" json:"name_en"`
	NameRu string `binding:"required" db:"name_ru" json:"name_ru"`
}

type EditArea struct {
	NameEn *string `db:"name_en" json:"name_en"`
	NameRu *string `db:"name_ru" json:"name_ru"`
}
