package server

import (
	"errors"

	"github.com/futurexx/positions-service/app/storage"
)

// Mocked Storage

type MockStorage struct {
	positionsRepo storage.IPositionsRepo
}

func (s *MockStorage) Open() error {
	return nil
}

func (s *MockStorage) Close() error {
	return nil
}

func (s *MockStorage) Positions() storage.IPositionsRepo {
	if s.positionsRepo == nil {
		s.positionsRepo = &MockPositionsRepo{}
	}

	return s.positionsRepo
}

type MockPositionsRepo struct{}

func (p *MockPositionsRepo) Summary(domain string) (uint, error) {
	if domain == "existed.domain" {
		return 5, nil
	}
	return 0, nil
}

func (p *MockPositionsRepo) Positions(domain string, order string, limit uint64, offset uint64) ([]storage.DomainPosition, error) {
	var res []storage.DomainPosition

	if domain == "existed.domain" {
		t := storage.DomainPosition{Keyword: "test", Position: 5}
		res = append(res, t)
	}

	return res, nil
}

// Broken Mocked Storage

type MockBrokenStorage struct {
	positionsRepo storage.IPositionsRepo
}

func (s *MockBrokenStorage) Open() error {
	return nil
}

func (s *MockBrokenStorage) Close() error {
	return nil
}

func (s *MockBrokenStorage) Positions() storage.IPositionsRepo {
	if s.positionsRepo == nil {
		s.positionsRepo = &MockBrokenPositionsRepo{}
	}

	return s.positionsRepo
}

type MockBrokenPositionsRepo struct{}

func (p *MockBrokenPositionsRepo) Summary(domain string) (uint, error) {
	return 0, errors.New("test")
}

func (p *MockBrokenPositionsRepo) Positions(domain string, order string, limit uint64, offset uint64) ([]storage.DomainPosition, error) {
	return []storage.DomainPosition{}, errors.New("test")
}

var MockServer = Server{storage: &MockStorage{}}
var MockBrokenServer = Server{storage: &MockBrokenStorage{}}
