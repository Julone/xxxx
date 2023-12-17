package routes

import (
	"fmt"
	"gorm-mysql/database"
	"gorm-mysql/models"
	"gorm.io/gorm"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Hello
func Hello(c *fiber.Ctx) error {
	return c.SendString("fiber")
}

// AddBook
func AddBook(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DBConn.Create(&book)

	return c.Status(200).JSON(book)
}

func GetBook(c *fiber.Ctx) error {
	books := []models.Book{}

	database.DBConn.First(&books, c.Params("id"))

	return c.Status(200).JSON(books)
}

// AllBooks
func AllBooks(c *fiber.Ctx) error {
	var books []models.Book

	database.DBConn.Preload("PageInfo").
		Preload("PageInfo.WriterInfo").
		Preload("PageInfo.Comment", func(db *gorm.DB) *gorm.DB {
			return db.Order("id desc")
		}).
		Preload("PageInfo.Comment.SenderInfo").
		Find(&books)

	return c.Status(200).JSON(fiber.Map{"success": true, "data": books})
}

// Update
func Update(c *fiber.Ctx) error {
	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	id, _ := strconv.Atoi(c.Params("id"))

	database.DBConn.Model(&models.Book{}).Where("id = ?", id).Update("title", book.Title)

	return c.Status(400).JSON("updated")
}

// Delete
func Delete(c *fiber.Ctx) error {
	book := new(models.Book)

	id, _ := strconv.Atoi(c.Params("id"))

	database.DBConn.Where("id = ?", id).Delete(&book)

	return c.Status(200).JSON("deleted")
}

func buildTree(tree []models.Comments, all map[int][]models.Comments) {
	for _, comments := range tree {
		if _, ok := all[comments.ID]; ok {
			fmt.Printf("%vsdsdf", comments)
			comments.Children = append(comments.Children, all[comments.ID]...)
			buildTree(comments.Children, all)
		}
	}
	//return nil
}

func GetCommentsList(c *fiber.Ctx) error {
	var ds []models.Comments

	database.DBConn.Model(&models.Comments{}).
		Find(&ds)
	all := map[int][]models.Comments{}
	for _, d := range ds {
		all[d.ParentID] = append(all[d.ParentID], d)
	}
	fmt.Printf("%v", all[1])
	data := all[1]
	buildTree(data, all)
	fmt.Printf("%v", data)

	return c.JSON(data)
}
