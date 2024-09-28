package service

import (
	"metroid_bookmarks/internal/repository/sql/skills"
)

type SkillsService struct {
	sql skills.SQL
}

func newSkillsService(sql skills.SQL) *SkillsService {
	return &SkillsService{sql: sql}
}

func (s *SkillsService) Create(createForm *skills.CreateSkill) (*skills.Skill, error) {
	skill, err := s.sql.Create(createForm)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return skill, nil
}

func (s *SkillsService) Edit(skillID int, editForm *skills.EditSkill) (*skills.Skill, error) {
	skill, err := s.sql.Edit(skillID, editForm)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return skill, nil
}

func (s *SkillsService) Delete(skillID int) (*skills.Skill, error) {
	skill, err := s.sql.Delete(skillID)
	if err != nil {
		logger.Error(err.Error())
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

	if data == nil {
		data = []skills.Skill{}
	}

	return data, len(data), nil
}
