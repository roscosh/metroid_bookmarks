package sql

const roomsTable = "rooms"

type Room struct {
	Id     int    `json:"id"      db:"id"`
	NameEn string `json:"name_en" db:"name_en"`
	NameRu string `json:"name_ru" db:"name_ru"`
}

type CreateRoom struct {
	NameEn string `json:"name_en" db:"name_en" binding:"required"`
	NameRu string `json:"name_ru" db:"name_ru" binding:"required"`
}

type EditRoom struct {
	NameEn *string `json:"name_en" db:"name_en"`
	NameRu *string `json:"name_ru" db:"name_ru"`
}

type RoomsSQL struct {
	iBaseSQL[Room]
}

func NewRoomsSQL(dbPool *DbPool, table string) *RoomsSQL {
	sql := newIBaseSQL[Room](dbPool, table)
	return &RoomsSQL{iBaseSQL: sql}
}

func (s *RoomsSQL) Create(createForm *CreateRoom) (*Room, error) {
	return s.insert(*createForm)
}

func (s *RoomsSQL) Edit(id int, editForm *EditRoom) (*Room, error) {
	return s.update(id, *editForm)
}

func (s *RoomsSQL) Delete(id int) (*Room, error) {
	return s.delete(id)
}

func (s *RoomsSQL) GetAll() ([]Room, error) {
	return s.selectMany()
}

func (s *RoomsSQL) GetByID(id int) (*Room, error) {
	return s.selectOne(id)
}

func (s *RoomsSQL) Total() (int, error) {
	return s.total()
}
