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
	iBaseSQL[Skill]
}

func NewSkillsSQL(dbPool *DbPool, table string) *SkillsSQL {
	sql := newIBaseSQL[Skill](dbPool, table)
	return &SkillsSQL{iBaseSQL: sql}
}

func (s *SkillsSQL) Create(createForm *CreateSkill) (*Skill, error) {
	return s.insert(*createForm)
}

func (s *SkillsSQL) Delete(id int) (*Skill, error) {
	return s.delete(id)
}

func (s *SkillsSQL) Edit(id int, editForm *EditSkill) (*Skill, error) {
	return s.update(id, *editForm)
}

func (s *SkillsSQL) GetAll() ([]Skill, error) {
	return s.selectMany()
}

func (s *SkillsSQL) GetByID(id int) (*Skill, error) {
	return s.selectOne(id)
}

func (s *SkillsSQL) Total() (int, error) {
	return s.total()
}
