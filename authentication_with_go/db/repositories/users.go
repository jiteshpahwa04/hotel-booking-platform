package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetByID(id string) (*models.User, error)
	GetByEmail(id string) (*models.User, error)
	Create(username string, email string, hashedPassword string) (*models.User, error)
	GetAll() ([]*models.User, error)
	DeleteById(id string) error
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: _db,
	}
}

func (u *UserRepositoryImpl) GetByID(id string) (*models.User, error) {
	fmt.Println("Getting user by ID in user repository")
	
	// Step 1: Prepare the query
	query := "SELECT id, username, email FROM users WHERE id = ?"

	// Step 2: Execute the query
	row := u.db.QueryRow(query, id)

	// Step 3: Process the result
	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given ID")
			return nil, nil
		}
		fmt.Println("Error occurred while getting user by ID", err)
		return nil, err
	}

	// Step 4: Use the scanned values (for demonstration purposes)
	fmt.Printf("User found: ID=%d, Username=%s, Email=%s\n", user.ID, user.Username, user.Email)
	return user, nil
}

func (u *UserRepositoryImpl) GetByEmail(email string) (*models.User, error) {
	fmt.Println("Getting user by email from repository")

	query := "SELECT id, username, email, password from users where email = ?"

	row := u.db.QueryRow(query, email);

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given email")
			return nil, nil
		}
		fmt.Println("Error occurred while getting user by email", err)
		return nil, err
	}

	return user, nil
} 

func (u *UserRepositoryImpl) Create(username string, email string, hashedPassword string) (*models.User, error) {
	fmt.Println("Creating user in user repository")
	
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	result, err := u.db.Exec(query, username, email, hashedPassword)
	if err != nil {
		fmt.Println("Error occurred while creating user", err)
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error occurred while getting rows affected", err)
		return nil, err
	}
	if rowsAffected == 0 {
		fmt.Println("No rows were inserted")
		return nil, fmt.Errorf("no rows were inserted")
	}

	// Get the last inserted ID
	userID, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error occurred while getting last insert ID", err)
		return nil, err
	}

	// Create the user object
	user := &models.User{
		ID:       userID,
		Username: username,
		Email:    email,
	}
	return user, nil
}

func (u *UserRepositoryImpl) GetAll() ([]*models.User, error) {
	query := "SELECT id, username, email, created_at, updated_at FROM users"
	rows, err := u.db.Query(query)
	if err != nil {
		fmt.Println("Error occurred while getting all users", err)
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			fmt.Println("Error occurred while scanning user row", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error occurred while iterating over user rows", err)
		return nil, err
	}

	return users, nil
}

func (u *UserRepositoryImpl) DeleteById(id string) error {
	query := "DELETE FROM users WHERE id = ?"
	result, err := u.db.Exec(query, id)
	if err != nil {
		fmt.Println("Error occurred while deleting user", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error occurred while getting rows affected", err)
		return err
	}
	if rowsAffected == 0 {
		fmt.Println("No rows were deleted")
		return fmt.Errorf("no rows were deleted")
	}

	return nil
}