package service

import (
	"metroid_bookmarks/internal/repository/sql/skills"
)

type SkillsService struct {
	sql *skills.SQL
}

func newSkillsService(sql *skills.SQL) *SkillsService {
	return &SkillsService{sql: sql}
}

func (s *SkillsService) Create(createForm *skills.CreateSkill) (*skills.Skill, error) {
	skill, err := s.sql.Create(createForm)
	if err != nil {
		logger.Error(err.Error())
		err = createPgError(err)
		return nil, err
	}
	return skill, nil
}

func (s *SkillsService) Edit(id int, editForm *skills.EditSkill) (*skills.Skill, error) {
	if (editForm == &skills.EditSkill{}) {
		return nil, ErrEmptyStruct
	}
	skill, err := s.sql.Edit(id, editForm)
	if err != nil {
		logger.Error(err.Error())
		err = editPgError(err, id)
		return nil, err
	}
	return skill, nil
}

func (s *SkillsService) Delete(id int) (*skills.Skill, error) {
	skill, err := s.sql.Delete(id)
	if err != nil {
		logger.Error(err.Error())
		err = deletePgError(err, id)
		return nil, err
	}
	return skill, nil
}

func (s *SkillsService) GetAll() ([]skills.Skill, int, error) {
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
	if data == nil {
		data = []skills.Skill{}
	}
	return data, total, nil
}
