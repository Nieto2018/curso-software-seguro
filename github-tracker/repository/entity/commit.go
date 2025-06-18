package entity

import "time"

type Commit struct {
	ID            int       `db:"id"`
	RepoName      string    `db:"repo_name"`      // Nombre del repositorio
	CommitID      string    `db:"commit_id"`      // ID del commit
	CommitMessage string    `db:"commit_message"` // Mensaje del commit
	AuthorName    string    `db:"author_name"`    // Nombre del autor del commit
	AuthorEmail   string    `db:"author_email"`   // Email del autor del commit
	Payload       string    `db:"payload"`        // Payload del commit
	CreatedAt     time.Time `db:"created_at"`     // Fecha de creación del commit
	UpdatedAt     time.Time `db:"updated_at"`     // Fecha de actualización del commit
}
