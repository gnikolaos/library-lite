package store

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db *sql.DB
}

func NewStore(dbFile string) (*Store, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	return &Store{db: db}, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) DB() *sql.DB {
	return s.db
}

func (s *Store) CreateTables() error {
	_, err := s.db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY,
            role INTEGER NOT NULL,
            email TEXT UNIQUE NOT NULL,
            name TEXT NOT NULL,
            surname TEXT NOT NULL,
            password TEXT NOT NULL,
            created_at DATETIME NOT NULL
        );
        CREATE TABLE IF NOT EXISTS roles (
            id INTEGER NOT NULL,
            name TEXT NOT NULL
        );
        CREATE TABLE IF NOT EXISTS books (
            id INTEGER PRIMARY KEY,
            title TEXT NOT NULL,
            subtitle TEXT,
            author TEXT NOT NULL,
            cover TEXT,
            publisher TEXT,
            isbn TEXT,
            isbn_13 TEXT,
            description TEXT,
            category INTEGER,
            original_language TEXT,
            original_title TEXT,
            physical_desctiption TEXT,
            pages INTEGER,
            size TEXT,
            total_copies INTEGER NOT NULL DEFAULT 0,
            available_copies INTEGER NOT NULL DEFAULT 0,
            borrowed_copies INTEGER NOT NULL DEFAULT 0,
            published_at TEXT,
            added_on TEXT NOT NULL
        );
        CREATE TABLE IF NOT EXISTS borrowings (
            id INTEGER PRIMARY KEY,
            user_id INTEGER NOT NULL,
            book_id INTEGER NOT NULL,
            copy_number INTEGER NOT NULL,
            borrowed_at TEXT NOT NULL,
            returned_at TEXT,
            FOREIGN KEY (user_id) REFERENCES user (id),
            FOREIGN KEY (book_id) REFERENCES books (id)
        );
    `)

	return err
}

func (s *Store) Seed() error {
	_, err := s.db.Exec(`
        INSERT INTO roles (id, name) VALUES (99, 'admin');
        INSERT INTO roles (id, name) VALUES (89, 'librarian');
        INSERT INTO roles (id, name) VALUES (9, 'patron');
        INSERT INTO roles (id, name) VALUES (1, 'user');
    `)
	if err != nil {
		return err
	}

	// password = password
	_, err = s.db.Exec(`
        INSERT INTO users (role, email, name, surname, password, created_at) VALUES
            (99, 'admin@lili.com','Lili', 'Doe', '$2a$08$xi46/MHHyZMz2.XHmlI1cemjyO.j48YKAgXEx9jW6ZqBdujg6x/GO', datetime('now', 'localtime'));
    `)
	if err != nil {
		return err
	}

	return nil
}
