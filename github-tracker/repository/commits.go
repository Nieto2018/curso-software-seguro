package repository

import (
	"context"
	"database/sql"
	"github-tracker/github-tracker/repository/entity"
)

type Commit interface {
	Insert(ctx context.Context, commit *entity.Commit) (err error)
	GetCommitsByAuthorEmail(ctx context.Context, email string) (commits []entity.Commit, err error)
	// GetCommitsByAuthorName(ctx context.Context, nameame string) (commits []*entity.Commit, err error)
	// GetCommitByID(ctx context.Context, commitID string) (commit *entity.Commit, err error)
}

type commit struct {
	Conn *sql.DB
}

func NewCommit(conn *sql.DB) Commit {
	return commit{
		Conn: conn
	}
}

func (m commit) Insert(ctx context.Context, commit *entity.Commit) (err error) {
	query := `INSERT INTO commits (repo_name, commit_id, commit_message, author_name, author_email, payload) 
					 VALUES ($1, $2, $3, $4, $5, $6, $7)`

	stmt, err = m.Conn.ExecContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	
	err = stmt.QueryRowContext(
		ctx, 
		commit.RepoName, 
		commit.CommitID, 
		commit.CommitMessage, 
		commit.AuthorName, 
		commit.AuthorEmail, 
		commit.Payload
		commit.CreatedAt
		commit.UpdatedAt,
		).Err()

}