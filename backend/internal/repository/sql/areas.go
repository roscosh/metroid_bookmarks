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
	iBaseSQL[Area]
}

func NewAreasSQL(dbPool *DbPool) *AreasSQL {
	sql := newIBaseSQL[Area](dbPool, areasTable)
	return &AreasSQL{iBaseSQL: sql}
}

func (s *AreasSQL) Create(createForm *CreateArea) (*Area, error) {
	return s.insert(createForm)
}

func (s *AreasSQL) Delete(id int) (*Area, error) {
	return s.delete(id)
}

func (s *AreasSQL) Edit(id int, editForm *EditArea) (*Area, error) {
	return s.update(id, editForm)
}

func (s *AreasSQL) GetAll() ([]Area, error) {
	return s.selectMany()
}

func (s *AreasSQL) GetByID(id int) (*Area, error) {
	return s.selectOne(id)
}

func (s *AreasSQL) Total() (int, error) {
	return s.total()
}
