package habit

import (
	"fmt"

	"github.com/Ocelani/perdat/pkg/entity"
)

type Service interface {
	Create(*[]entity.Habit) error
	Read(*[]entity.Habit) error
	Update(UpdateHabitNames) (*[]entity.Habit, error)
	Delete(*[]entity.Habit) error
}

type serv struct{ repo Repository }

func NewService(repo Repository) Service { return &serv{repo} }

func (s *serv) Read(habits *[]entity.Habit) error { return s.repo.Read(habits) }

func (s *serv) Delete(habits *[]entity.Habit) error {
	dels := []entity.Habit{}
	for _, f := range *habits {
		fmt.Printf("%+v", f)
		if err := s.repo.Delete(&f); err != nil {
			return err
		}
		dels = append(dels, f)
	}
	habits = &dels

	return nil
}

func (s *serv) Create(habits *[]entity.Habit) error {
	return s.repo.Create(habits)
}

func (s *serv) Update(update UpdateHabitNames) (*[]entity.Habit, error) {
	var (
		habits  = []entity.Habit{}
		habitCH = make(chan entity.Habit)
		quitCH  = make(chan bool)
		errCH   = make(chan error)
	)
	defer func() { close(habitCH); close(errCH) }()

	go func() {
		for is, tobe := range update {
			t := []entity.Habit{*entity.NewHabit(is)}
			if err := s.Read(&t); err != nil {
				errCH <- err
			}
			t[0].Name = tobe
			habitCH <- t[0]
		}
		quitCH <- true
	}()

	for {
		select {
		case t := <-habitCH:
			if err := s.repo.Update(&t); err != nil {
				errCH <- err
			}
			habits = append(habits, t)

		case <-quitCH:
			return &habits, nil

		case err := <-errCH:
			return nil, fmt.Errorf("habit.Update: %s", err)
		}
	}
}
