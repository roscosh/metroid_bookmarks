package sql

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

func NewSkillsSQL(baseSQL *baseSQL) *SkillsSQL {
	return &SkillsSQL{baseSQL: baseSQL}
}

func (s *SkillsSQL) GetByID(id int) (*Skill, error) {
	return selectById[Skill](s.baseSQL, skillsTable, id)
}

func (s *SkillsSQL) GetAll() ([]Skill, error) {
	return selectAll[Skill](s.baseSQL, skillsTable)
}

func (s *SkillsSQL) Create(createForm *CreateSkill) (*Skill, error) {
	return insert[Skill](s.baseSQL, skillsTable, *createForm)
}

func (s *SkillsSQL) Edit(id int, editForm *EditSkill) (*Skill, error) {
	return update[Skill](s.baseSQL, skillsTable, id, *editForm)
}

func (s *SkillsSQL) Delete(id int) (*Skill, error) {
	return deleteById[Skill](s.baseSQL, skillsTable, id)
}

func (s *SkillsSQL) Total() (int, error) {
	return total(s.baseSQL, skillsTable)
}
