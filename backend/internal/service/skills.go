package service

import (
	"errors"
	"metroid_bookmarks/internal/repository/sql"
)

type SkillsService struct {
	sql *sql.SkillsSQL
}

func newSkillsService(sql *sql.SkillsSQL) *SkillsService {
	return &SkillsService{sql: sql}
}

func (s *SkillsService) Create(createForm *sql.CreateSkill) (*sql.Skill, error) {
	skill, err := s.sql.Create(createForm)
	if err != nil {
		logger.Error(err.Error())
		err = createPgError(err)
		return nil, err
	}
	return skill, nil
}

func (s *SkillsService) Edit(id int, editForm *sql.EditSkill) (*sql.Skill, error) {
	if (editForm == &sql.EditSkill{}) {
		return nil, errors.New("Необходимо заполнить хотя бы один параметр в форме!")
	}
	skill, err := s.sql.Edit(id, editForm)
	if err != nil {
		logger.Error(err.Error())
		err = editPgError(err, id)
		return nil, err
	}
	return skill, nil
}

func (s *SkillsService) Delete(id int) (*sql.Skill, error) {
	skill, err := s.sql.Delete(id)
	if err != nil {
		logger.Error(err.Error())
		err = deletePgError(err, id)
		return nil, err
	}
	return skill, nil
}

func (s *SkillsService) GetAll() ([]sql.Skill, int, error) {
	data, err := s.sql.GetAll()
	if err != nil {
		logger.Error(err.Error())
		return nil, 0, err
	}
	total, err := s.sql.Total()
	if err != nil {
		logger.Error(err.Error())
		return nil, 0, err
	}
	return data, total, nil
}
