package sql

const areasTable = "areas"

type Area struct {
	Id     int    `json:"id"      db:"id"`
	NameEn string `json:"name_en" db:"name_en"`
	NameRu string `json:"name_ru" db:"name_ru"`
}

type CreateArea struct {
	NameEn string `json:"name_en" db:"name_en" binding:"required"`
	NameRu string `json:"name_ru" db:"name_ru" binding:"required"`
}

type EditArea struct {
	NameEn *string `json:"name_en" db:"name_en"`
	NameRu *string `json:"name_ru" db:"name_ru"`
}

type AreasSQL struct {
	*baseSQL
}

func NewAreasSQL(baseSQL *baseSQL) *AreasSQL {
	return &AreasSQL{baseSQL: baseSQL}
}

func (s *AreasSQL) GetByID(id int) (*Area, error) {
	return selectById[Area](s.baseSQL, areasTable, id)
}

func (s *AreasSQL) GetAll() ([]Area, error) {
	return selectAll[Area](s.baseSQL, areasTable)
}

func (s *AreasSQL) Create(createForm *CreateArea) (*Area, error) {
	return insert[Area](s.baseSQL, areasTable, *createForm)
}

func (s *AreasSQL) Edit(id int, editForm *EditArea) (*Area, error) {
	return update[Area](s.baseSQL, areasTable, id, *editForm)
}

func (s *AreasSQL) Delete(id int) (*Area, error) {
	return deleteById[Area](s.baseSQL, areasTable, id)
}

func (s *AreasSQL) Total() (int, error) {
	return total(s.baseSQL, areasTable)
}
