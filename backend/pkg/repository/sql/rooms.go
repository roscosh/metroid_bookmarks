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
	*baseSQL
}

func NewRoomsSQL(baseSQL *baseSQL) *RoomsSQL {
	return &RoomsSQL{baseSQL: baseSQL}
}

func (s *RoomsSQL) GetByID(id int) (*Room, error) {
	return selectById[Room](s.baseSQL, roomsTable, id)
}

func (s *RoomsSQL) GetAll() ([]Room, error) {
	return selectAll[Room](s.baseSQL, roomsTable)
}

func (s *RoomsSQL) Create(createForm *CreateRoom) (*Room, error) {
	return insert[Room](s.baseSQL, roomsTable, *createForm)
}

func (s *RoomsSQL) Edit(id int, editForm *EditRoom) (*Room, error) {
	return update[Room](s.baseSQL, roomsTable, id, *editForm)
}

func (s *RoomsSQL) Delete(id int) (*Room, error) {
	return deleteById[Room](s.baseSQL, roomsTable, id)
}

func (s *RoomsSQL) Total() (int, error) {
	return total(s.baseSQL, roomsTable)
}
