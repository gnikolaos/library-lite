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
            email TEXT UNIQUE NOT NULL,
            password TEXT NOT NULL,
            role TEXT NOT NULL,
            name TEXT NOT NULL,
            surname TEXT NOT NULL,
            created_at DATETIME NOT NULL
        );
        CREATE TABLE IF NOT EXISTS roles (
            id INTEGER PRIMARY KEY,
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
        INSERT INTO roles (name) VALUES ('admin');
        INSERT INTO roles (name) VALUES ('librarian');
        INSERT INTO roles (name) VALUES ('patron');
        INSERT INTO roles (name) VALUES ('user');
    `)
	if err != nil {
		return err
	}

	// password = password
	_, err = s.db.Exec(`
        INSERT INTO users (email, password, role, name, surname, created_at) VALUES
            ('admin@lili.com', '$2a$08$xi46/MHHyZMz2.XHmlI1cemjyO.j48YKAgXEx9jW6ZqBdujg6x/GO', 1, 'Lili', 'Doe', datetime('now', 'localtime'));
    `)
	if err != nil {
		return err
	}

	return nil
}
