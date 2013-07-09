package entry

import (
	"github.com/ChuckHa/blog/db"
	"log"
	"time"
)

const (
	EntriesTable = "entries"
)

type Entry struct {
	Author, Title, Content string
	Created, Edited        time.Time
}

type Entries []*Entry

// NewEntry will create a new entry with default created and edited time
func NewEntry(author, content, title string) *Entry {
	return &Entry{
		Author:  author,
		Title:   title,
		Content: content,
		Created: time.Now(),
		Edited:  time.Now(),
	}
}

// ParseEntry is used to go from a *sql.Row to a *Entry
func ParseEntry(author, content, title, created, edited string) *Entry {
	createdT, err := time.Parse(db.MysqlForm, created)
	if err != nil {
		log.Printf("Error parsing time: %s", err)
	}
	editedT, _ := time.Parse(db.MysqlForm, edited)
	if err != nil {
		log.Printf("Error parsing time: %s", err)
	}
	return &Entry{
		Author:  author,
		Title:   title,
		Content: content,
		Created: createdT,
		Edited:  editedT,
	}
}

// Get all the entries in the database
func List() (entries Entries) {
	var err error
	fields := db.Fields("author", "content", "title", "created", "edited")
	rows := db.SelectAll(fields, EntriesTable)
	for rows.Next() {
		var a, c, t, cr, ed string
		err = rows.Scan(&a, &c, &t, &cr, &ed)
		if err != nil {
			log.Printf("Error in getting a row: %v", err)
		}
		entries = append(entries, ParseEntry(a, c, t, cr, ed))
	}
	return entries
}
