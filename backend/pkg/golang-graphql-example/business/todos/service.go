package todos

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/daos"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/business/todos/models"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

type service struct {
	dao daos.Dao
}

func (s *service) MigrateDB(systemLogger log.Logger) error {
	systemLogger.Debug("Migrate database for Todos")
	return s.dao.MigrateDB()
}

func (s *service) GetAll() ([]*models.Todo, error) {
	return s.dao.GetAll()
}

func (s *service) Create(inp *InputCreateTodo) (*models.Todo, error) {
	tt := &models.Todo{
		Text: inp.Text,
	}

	return s.dao.CreateOrUpdate(tt)
}

func (s *service) Update(inp *InputUpdateTodo) (*models.Todo, error) {
	// Search by id first
	tt, err := s.dao.FindByID(inp.ID)
	if err != nil {
		return nil, err
	}
	// Update text in existing result
	tt.Text = inp.Text
	// Save
	return s.dao.CreateOrUpdate(tt)
}

func (s *service) Close(id string) (*models.Todo, error) {
	// Search by id first
	tt, err := s.dao.FindByID(id)
	if err != nil {
		return nil, err
	}
	// Update text in existing result
	tt.Done = true
	// Save
	return s.dao.CreateOrUpdate(tt)
}
