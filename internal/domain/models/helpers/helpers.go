package helpers

import (
	"github.com/xpanvictor/g3_template.git/internal/domain/models"
	"github.com/xpanvictor/g3_template.git/internal/graph/model"
)

func MapDomainUserToGraphQL(user *models.User) *model.User {
	return &model.User{
		ID:   user.ID.String(),
		Name: user.Name,
	}
}
