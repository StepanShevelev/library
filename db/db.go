package db

import (
	"fmt"
	cfg "github.com/StepanShevelev/library/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Client *client

type client struct {
	*gorm.DB
}

func New(config *cfg.Config) (*client, error) {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("cant connect to db: %w", err)
	}

	tmp := &client{db}
	Client = tmp
	return tmp, nil
}

func FillDbWithTestData(config *cfg.Config) (*client, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("cant connect to db: %w", err)
	}
	sql := "INSERT INTO authors(name, book_id) VALUES ('pushkin', 1)," +
		"('kin', 2)," +
		"('lermontov', 1)"

	db = db.Exec(sql)

	tmp := &client{db}
	Client = tmp
	return tmp, nil
}

func (db *client) GiveBookByAuthor(authorId int64) (string, error) {
	author := NewAuthor()
	var n string
	n = author.Name
	result := db.Where("author_id = ?", authorId).First(&author)
	if result.Error != nil {
		return "error occurred", result.Error
	}
	return n, nil
}
func (db *client) GiveAuthorByBook(bookId int64) (string, error) {
	book := NewBook()
	var n string
	n = book.Name
	result := db.Where("book_id = ?", bookId).First(&book)
	if result.Error != nil {
		return "error occurred", result.Error
	}
	return n, nil
}

func (db *client) SetDB() error {
	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Book{})
	return nil
}

type Author struct {
	gorm.Model
	Name   string
	BookId int
}

func NewAuthor() *Author {
	return &Author{}
}

type Book struct {
	gorm.Model
	Name     string
	AuthorId int
}

func NewBook() *Book {
	return &Book{}
}
