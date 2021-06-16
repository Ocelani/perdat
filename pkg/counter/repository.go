package counter

import (
	"github.com/Ocelani/perdat/pkg/common"
	"github.com/Ocelani/perdat/pkg/common/database"
	"github.com/Ocelani/perdat/pkg/entity"
)

type Repository interface {
	Create(*[]entity.Counter) error
	Read(*[]entity.Counter) error
	Update(*entity.Counter) error
	Delete(*entity.Counter) error
}

type repo struct{ *common.Repository }

func NewRepository(db *database.DB) *repo { return &repo{common.NewRepository(db)} }

func (r *repo) Create(t *[]entity.Counter) error { return r.DB.Create(t).Error }
func (r *repo) Read(t *[]entity.Counter) error   { return r.DB.Find(t, t).Error }
func (r *repo) Update(t *entity.Counter) error   { return r.DB.Updates(t).Error }
func (r *repo) Delete(t *entity.Counter) error   { return r.DB.Delete(&entity.Counter{}, t.ID).Error }
