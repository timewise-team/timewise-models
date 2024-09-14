package core_dtos

import "github.com/timewise-team/timewise-models/models"

type UpdateUserRequest struct {
	User          models.TwUser `json:"user"`
	UpdatedFields []string      `json:"updated_fields"`
}
