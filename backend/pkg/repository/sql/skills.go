package sql

import "github.com/jackc/pgx/v5"

const skillsTable = "skills"

type Skill struct {
	Id     int    `json:"id"      db:"id"`
	NameEn string `json:"name_en" db:"name_en"`
	NameRu string `json:"name_ru" db:"name_ru"`
}

type CreateSkill struct {
	NameEn string `json:"name_en" db:"name_en" binding:"required"`
	NameRu string `json:"name_ru" db:"name_ru" binding:"required"`
}

type EditSkill struct {
	NameEn *string `json:"name_en" db:"name_en"`
	NameRu *string `json:"name_ru" db:"name_ru"`
}

type SkillsSQL struct {
	*baseSQL
}

func NewSkillsSQL(pool *DbPool, table string) *SkillsSQL {
	sql := newBaseSQl(pool, table, Skill{})
	return &SkillsSQL{baseSQL: sql}
}

func (s *SkillsSQL) GetAll() ([]Skill, error) {
	rows, err := s.selectAll()
	if err != nil {
		return nil, err
	}
	return s.collectRows(rows)
}

func (s *SkillsSQL) GetByID(id int) (*Skill, error) {
	rows, err := s.selectById(id)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *SkillsSQL) Create(createForm *CreateSkill) (*Skill, error) {
	rows, err := s.insert(*createForm)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *SkillsSQL) Edit(id int, editForm *EditSkill) (*Skill, error) {
	rows, err := s.update(id, *editForm)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *SkillsSQL) Delete(id int) (*Skill, error) {
	rows, err := s.deleteById(id)
	if err != nil {
		return nil, err
	}
	return s.collectOneRow(rows)
}

func (s *SkillsSQL) Total() (int, error) {
	return s.total()
}

func (s *SkillsSQL) collectOneRow(rows pgx.Rows) (*Skill, error) {
	return collectOneRow[Skill](rows)
}

func (s *SkillsSQL) collectRows(rows pgx.Rows) ([]Skill, error) {
	return collectRows[Skill](rows)
}
