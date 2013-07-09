package blog

import (
	"github.com/ChuckHa/blog/db"
	"github.com/ChuckHa/blog/entry"
	"log"
)

const (
	blogTable = "blogs"
)

type Blog struct {
	Id      int
	Title   string
}

type Blogs []*Blog

// NewBlog will create a new blog.
func NewBlog(id int, title string) *Blog {
	return &Blog{
		Id:      id,
		Title:   title,
	}
}

// TODO: Write a ParseBlog function that returns a *Blog from a *sql.Row

// List gets all blogs in the database
func List() (blogs Blogs) {
	var err error
	fields := db.Fields("id", "title")
	rows := db.SelectAll(fields, blogTable)
	for rows.Next() {
		var id int
		var title string
		err = rows.Scan(&id, &title)
		if err != nil {
			log.Printf("Error in getting a row: %v", err)
		}
		blogs = append(blogs, NewBlog(id, title))
	}
	return blogs
}

// GetEntries gets all entries for a given blog
func (b *Blog) GetEntries() (entries entry.Entries) {
	var err error
	fields := db.Fields("author", "content", "title", "created", "edited")
	rows := db.GetEntriesFromBlog(fields, entry.EntriesTable, b.Id)
	for rows.Next() {
		var a, c, t, cr, ed string
		err = rows.Scan(&a, &c, &t, &cr, &ed)
		if err != nil {
			log.Printf("Error in getting a row: %v", err)
		}
		entries = append(entries, entry.ParseEntry(a, c, t, cr, ed))
	}
	return entries
}
