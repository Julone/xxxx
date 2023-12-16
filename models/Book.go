package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

// Book model
type Book struct {
	gorm.Model

	Title    string `json:"title"`
	Author   string `json:"author"`
	PageId   int    `json:"page_id" gorm:"type:bigint(0);default:0"`
	PageInfo *Page  `json:"page_info" gorm:"foreignKey:PageId"`
}

type Page struct {
	ID        uint `gorm:"type:bigint(20);autoIncrement;notNull;primaryKey;column:id"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Content    string `json:"content"`
	Writer     string `json:"writer" gorm:"column:writer;type:varchar(50)"`
	WriterID   uint   `json:"writer_id" gorm:"type:int;column:writer_id"`
	WriterInfo *User  `json:"writer_info" gorm:"foreignKey:WriterID;references:ID"`

	Comment []Comments `json:"comments" gorm:"foreignKey:PageID;"`
}

type Comments struct {
	ID        int `gorm:"index;type:bigint(20);autoIncrement;notNull;primaryKey;column:id"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Content    string `json:"content" gorm:"column:content"`
	PageID     int    `json:"page_id" gorm:"column:page_id"`
	ParentID   int    `json:"parent_id" gorm:"column:parent_id;default:0;type:int"`
	SenderID   uint   `json:"sender_id" gorm:"type:int;column:sender_id"`
	SenderInfo *User  `json:"sender_info" gorm:"foreignKey:SenderID;references:ID;"`
	IsDeleted  bool   `json:"IsDeleted" gorm:"column:isDeleted;default:true"`

	Children []Comments `json:"children" gorm:"-"'`
}

type User struct {
	ID        uint `gorm:"index;type:bigint(20);autoIncrement;notNull;primaryKey;column:id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserName  string `gorm:"index;type:varchar(50);primaryKey;distinct;column:username;notNull;" json:"username"`
	RealName  string `gorm:"type:varchar(50);column:real_name;default:''" json:"realname"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) error {
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	return nil
}

func (b *Book) AfterFind(db *gorm.DB) error {
	b.UpdatedAt = time.Now()
	return nil
}

func (b *Book) BeforeDelete(db *gorm.DB) error {
	log.Fatal("before delete")
	return nil
}
