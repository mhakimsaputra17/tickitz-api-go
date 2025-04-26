package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mhakimsaputra17/tickitz-api-go/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository handles database operations for users
type UserRepository struct {
	db *pgxpool.Pool
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(db *pgxpool.Pool) *UserRepository{
	return &UserRepository{db: db}
}

// CreateUser adds a new user to the database with hashed password
func (r *UserRepository) CreateUser (ctx context.Context, user model.UserRegister ) (model.User, error){
	// Hash the password 
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}

	var createdUser model.User
	query:= `
		 INSERT INTO users (username, email, password_hash, first_name, last_name, created_at, updated_at) 
        VALUES ($1, $2, $3, $4, $5, NOW(), NOW()) 
        RETURNING user_id, username, email, first_name, last_name ,created_at, updated_at
    `
	err = r.db.QueryRow(ctx, query,
		user.Username,
		user.Email,
		string(hashedPassword),
		user.FirstName,
		user.LastName,

	).Scan(
		&createdUser.ID,
		&createdUser.Username,
		&createdUser.Email,
		&createdUser.FirstName,
		&createdUser.LastName,
		&createdUser.CreatedAt,
		&createdUser.UpdatedAt,
	)
	if err != nil {
		return model.User{}, err
	}

	return createdUser, nil

}

// GetUserByUsername retrieves a user by username
func (r *UserRepository) GetUserByEmail(ctx context.Context, email string ) (model.User, error){
	var user model.User
	query :=  `
        SELECT user_id, username, email, password_hash, first_name, last_name, created_at, updated_at 
        FROM users WHERE email = $1
    `

	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows){
			return model.User{}, errors.New("user not found")
		}
		return model.User{}, err
	}

	return user, nil



}



// VerifyPassword checks if the provided password matches the hashed one
func (r *UserRepository) VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}