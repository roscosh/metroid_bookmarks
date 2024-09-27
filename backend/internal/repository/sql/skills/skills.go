package skills

import (
	"metroid_bookmarks/internal/repository/sql/pgerr"
	"metroid_bookmarks/pkg/pgpool"
)

const skillsTable = "skills"

type SQL interface {
	Create(createForm *CreateSkill) (*Skill, error)
	Delete(id int) (*Skill, error)
	Edit(id int, editForm *EditSkill) (*Skill, error)
	GetAll() ([]Skill, error)
	GetByID(id int) (*Skill, error)
	Total() (int, error)
}

type skillsSQL struct {
	sql pgpool.SQL[Skill]
}

func NewSQL(dbPool *pgpool.PgPool) SQL {
	sql := pgpool.NewSQL[Skill](dbPool, skillsTable)
	return &skillsSQL{sql: sql}
}

func (s *skillsSQL) Create(createForm *CreateSkill) (*Skill, error) {
	entity, err := s.sql.Insert(createForm)
	if err != nil {
		err = pgerr.CreatePgError(err)
		return nil, err
	}

	return entity, nil
}

func (s *skillsSQL) Delete(id int) (*Skill, error) {
	entity, err := s.sql.Delete(id)
	if err != nil {
		err = pgerr.DeletePgError(err, id)
		return nil, err
	}

	return entity, nil
}

func (s *skillsSQL) Edit(id int, editForm *EditSkill) (*Skill, error) {
	entity, err := s.sql.Update(id, editForm)
	if err != nil {
		err = pgerr.EditPgError(err, id)
		return nil, err
	}

	return entity, nil
}

func (s *skillsSQL) GetAll() ([]Skill, error) {
	return s.sql.SelectMany()
}

func (s *skillsSQL) GetByID(id int) (*Skill, error) {
	entity, err := s.sql.SelectOne(id)
	if err != nil {
		err = pgerr.SelectPgError(err, id)
		return nil, err
	}

	return entity, nil
}

func (s *skillsSQL) Total() (int, error) {
	return s.sql.Total()
}
