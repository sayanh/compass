package labeldef

import (
	"context"

	"github.com/kyma-incubator/compass/components/director/internal/model"
	"github.com/pkg/errors"
)

type service struct {
	repo       Repository
	uidService UIDService
}

func NewService(r Repository, uidService UIDService) *service {
	return &service{
		repo:       r,
		uidService: uidService,
	}
}

//go:generate mockery -name=Repository -output=automock -outpkg=automock -case=underscore
type Repository interface {
	Create(ctx context.Context, def model.LabelDefinition) error
}

//go:generate mockery -name=UIDService -output=automock -outpkg=automock -case=underscore
type UIDService interface {
	Generate() string
}

func (s *service) Create(ctx context.Context, def model.LabelDefinition) (model.LabelDefinition, error) {
	id := s.uidService.Generate()
	def.ID = id
	if err := def.Validate(); err != nil {
		return model.LabelDefinition{}, errors.Wrap(err, "while validation label definition")
	}

	if err := s.repo.Create(ctx, def); err != nil {
		return model.LabelDefinition{}, errors.Wrap(err, "while storing Label Definition")
	}
	// TODO get from DB?
	return def, nil

}
