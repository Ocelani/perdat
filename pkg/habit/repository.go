package habit

import (
	"github.com/Ocelani/perdat/pkg/common"
	"github.com/Ocelani/perdat/pkg/common/database"
	"github.com/Ocelani/perdat/pkg/entity"
)

type Repository interface {
	Create(*[]entity.Habit) error
	Read(*[]entity.Habit) error
	Update(*entity.Habit) error
	Delete(*entity.Habit) error
}

type repo struct{ *common.Repository }

func NewRepository(db *database.DB) *repo { return &repo{common.NewRepository(db)} }

func (r *repo) Create(t *[]entity.Habit) error { return r.DB.Create(t).Error }
func (r *repo) Read(t *[]entity.Habit) error   { return r.DB.Find(t, t).Error }
func (r *repo) Update(t *entity.Habit) error   { return r.DB.Updates(t).Error }
func (r *repo) Delete(t *entity.Habit) error   { return r.DB.Delete(&entity.Habit{}, t.ID).Error }
