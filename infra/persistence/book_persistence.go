package persistence

import (
	"time"

	"github.com/yyh-gl/go-api-server-by-ddd/domain/model"
)

type BookPersistence struct {}

func (book BookPersistence) GetAll() ([]*model.Book, error) {
	book1 := model.Book{}
	book1.Id = 1
	book1.Title = "Test1"
	book1.Author = "Tester1"
	book1.IssuedAt = time.Now()
	book1.CreatedAt = time.Now()
	book1.UpdatedAt = time.Now()

	book2 := model.Book{}
	book2.Id = 2
	book2.Title = "Test2"
	book2.Author = "Tester2"
	book2.IssuedAt = time.Now()
	book2.CreatedAt = time.Now()
	book2.UpdatedAt = time.Now()

	return []*model.Book{ &book1, &book2 }, nil
}
