package fact

import (
	"github.com/Ocelani/perdat/pkg/common"
	"github.com/Ocelani/perdat/pkg/common/database"
	"github.com/Ocelani/perdat/pkg/entity"
)

type Repository interface {
	Create(*[]entity.Fact) error
	Read(*[]entity.Fact) error
	Update(*entity.Fact) error
	Delete(*entity.Fact) error
}

type repo struct{ *common.Repository }

func NewRepository(db *database.DB) *repo { return &repo{common.NewRepository(db)} }

func (r *repo) Create(t *[]entity.Fact) error { return r.DB.Create(t).Error }
func (r *repo) Read(t *[]entity.Fact) error   { return r.DB.Find(t, t).Error }
func (r *repo) Update(t *entity.Fact) error   { return r.DB.Updates(t).Error }
func (r *repo) Delete(t *entity.Fact) error   { return r.DB.Delete(&entity.Fact{}, t.ID).Error }
