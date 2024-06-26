package skills

import (
	"metroid_bookmarks/pkg/pgpool"
)

const skillsTable = "skills"

type SQL struct {
	sql pgpool.SQL[Skill]
}

func NewSQL(dbPool *pgpool.DbPool) *SQL {
	sql := pgpool.NewSQL[Skill](dbPool, skillsTable)
	return &SQL{sql: sql}
}

func (s *SQL) Create(createForm *CreateSkill) (*Skill, error) {
	return s.sql.Insert(createForm)
}

func (s *SQL) Delete(id int) (*Skill, error) {
	return s.sql.Delete(id)
}

func (s *SQL) Edit(id int, editForm *EditSkill) (*Skill, error) {
	return s.sql.Update(id, editForm)
}

func (s *SQL) GetAll() ([]Skill, error) {
	return s.sql.SelectMany()
}

func (s *SQL) GetByID(id int) (*Skill, error) {
	return s.sql.SelectOne(id)
}

func (s *SQL) Total() (int, error) {
	return s.sql.Total()
}
