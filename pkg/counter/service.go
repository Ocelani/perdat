package counter

import (
	"fmt"

	"github.com/Ocelani/perdat/pkg/entity"
)

type Service interface {
	Create(*[]entity.Counter) error
	Read(*[]entity.Counter) error
	UpdateCount(UpdateCounterInt) (*[]entity.Counter, error)
	UpdateNames(UpdateCounterNames) (*[]entity.Counter, error)
	Delete(*[]entity.Counter) error
}

type serv struct{ repo Repository }

func NewService(repo Repository) Service { return &serv{repo} }

func (s *serv) Read(counters *[]entity.Counter) error { return s.repo.Read(counters) }

func (s *serv) Delete(counters *[]entity.Counter) error {
	dels := []entity.Counter{}
	for _, f := range *counters {
		fmt.Printf("%+v", f)
		if err := s.repo.Delete(&f); err != nil {
			return err
		}
		dels = append(dels, f)
	}
	counters = &dels

	return nil
}

func (s *serv) Create(counters *[]entity.Counter) error {
	return s.repo.Create(counters)
}

func (s *serv) UpdateCount(update UpdateCounterInt) (*[]entity.Counter, error) {
	cs := []entity.Counter{}

	for str, num := range update {
		c := []entity.Counter{*entity.NewCounter(str)}

		if err := s.Read(&c); err != nil {
			continue
		}
		c[0].CountNum = num

		if err := s.repo.Update(&c[0]); err != nil {
			continue
		}
		cs = append(cs, entity.Counter{})
	}

	return &cs, nil
}

func (s *serv) UpdateNames(update UpdateCounterNames) (*[]entity.Counter, error) {
	var (
		counters  = []entity.Counter{}
		counterCH = make(chan entity.Counter)
		quitCH    = make(chan bool)
		errCH     = make(chan error)
	)
	defer func() { close(counterCH); close(errCH) }()

	go func() {
		for is, tobe := range update {
			t := []entity.Counter{*entity.NewCounter(is)}
			if err := s.Read(&t); err != nil {
				errCH <- err
			}
			t[0].Name = tobe
			counterCH <- t[0]
		}
		quitCH <- true
	}()

	for {
		select {
		case t := <-counterCH:
			if err := s.repo.Update(&t); err != nil {
				errCH <- err
			}
			counters = append(counters, t)

		case <-quitCH:
			return &counters, nil

		case err := <-errCH:
			return nil, fmt.Errorf("counter.Update: %s", err)
		}
	}
}
