package skills

import "metroid_bookmarks/pkg/repository/sql"

type createForm struct {
	*sql.CreateSkill
}

type createResponse struct {
	*sql.Skill
}
type editForm struct {
	*sql.EditSkill
}
type editResponse struct {
	*sql.Skill
}

type deleteResponse struct {
	*sql.Skill
}

type getAllResponse struct {
	Data  []sql.Skill `json:"data"`
	Total int         `json:"total"`
}
