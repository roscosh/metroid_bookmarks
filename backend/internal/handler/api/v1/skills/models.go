package skills

import (
	"metroid_bookmarks/internal/repository/sql/skills"
)

type createForm struct {
	*skills.CreateSkill
}

type createResponse struct {
	*skills.Skill
}
type editForm struct {
	*skills.EditSkill
}
type editResponse struct {
	*skills.Skill
}

type deleteResponse struct {
	*skills.Skill
}

type getAllResponse struct {
	Data  []skills.Skill `json:"data"`
	Total int            `json:"total"`
}
