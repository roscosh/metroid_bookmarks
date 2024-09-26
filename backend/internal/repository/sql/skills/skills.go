package skills

import (
	"metroid_bookmarks/internal/repository/sql/pgerr"
	"metroid_bookmarks/pkg/pgpool"
)

const skillsTable = "skills"

type SQL struct {
	sql pgpool.SQL[Skill]
}

func NewSQL(dbPool *pgpool.PgPool) *SQL {
	sql := pgpool.NewSQL[Skill](dbPool, skillsTable)
	return &SQL{sql: sql}
}

func (s *SQL) Create(createForm *CreateSkill) (*Skill, error) {
	entity, err := s.sql.Insert(createForm)
	if err != nil {
		err = pgerr.CreatePgError(err)
		return nil, err
	}

	return entity, nil
}

func (s *SQL) Delete(id int) (*Skill, error) {
	entity, err := s.sql.Delete(id)
	if err != nil {
		err = pgerr.DeletePgError(err, id)
		return nil, err
	}

	return entity, nil
}

func (s *SQL) Edit(id int, editForm *EditSkill) (*Skill, error) {
	entity, err := s.sql.Update(id, editForm)
	if err != nil {
		err = pgerr.EditPgError(err, id)
		return nil, err
	}

	return entity, nil
}

func (s *SQL) GetAll() ([]Skill, error) {
	return s.sql.SelectMany()
}

func (s *SQL) GetByID(id int) (*Skill, error) {
	entity, err := s.sql.SelectOne(id)
	if err != nil {
		err = pgerr.SelectPgError(err, id)
		return nil, err
	}

	return entity, nil
}

func (s *SQL) Total() (int, error) {
	return s.sql.Total()
}
