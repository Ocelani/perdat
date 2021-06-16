package fact

import (
	"fmt"

	"github.com/Ocelani/perdat/pkg/entity"
)

type Service interface {
	Create(*[]entity.Fact) error
	Read(*[]entity.Fact) error
	Update(UpdateFactNames) (*[]entity.Fact, error)
	Delete(*[]entity.Fact) error
}

type serv struct{ repo Repository }

func NewService(repo Repository) Service { return &serv{repo} }

func (s *serv) Read(facts *[]entity.Fact) error { return s.repo.Read(facts) }

func (s *serv) Delete(facts *[]entity.Fact) error {
	dels := []entity.Fact{}
	for _, f := range *facts {
		fmt.Printf("%+v", f)
		if err := s.repo.Delete(&f); err != nil {
			return err
		}
		dels = append(dels, f)
	}
	facts = &dels

	return nil
}

func (s *serv) Create(facts *[]entity.Fact) error {
	return s.repo.Create(facts)
}

func (s *serv) Update(update UpdateFactNames) (*[]entity.Fact, error) {
	var (
		facts  = []entity.Fact{}
		factCH = make(chan entity.Fact)
		quitCH = make(chan bool)
		errCH  = make(chan error)
	)
	defer func() { close(factCH); close(errCH) }()

	go func() {
		for is, tobe := range update {
			t := []entity.Fact{*entity.NewFact(is)}
			if err := s.Read(&t); err != nil {
				errCH <- err
			}
			t[0].Name = tobe
			factCH <- t[0]
		}
		quitCH <- true
	}()

	for {
		select {
		case t := <-factCH:
			if err := s.repo.Update(&t); err != nil {
				errCH <- err
			}
			facts = append(facts, t)

		case <-quitCH:
			return &facts, nil

		case err := <-errCH:
			return nil, fmt.Errorf("fact.Update: %s", err)
		}
	}
}
