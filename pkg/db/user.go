package db

import "github.com/joefazee/autodeploy/pkg/domain"

// GetAllUsers return all the users in the database
func (q *SQLStore) GetAllUsers() ([]domain.User, *domain.AppError) {

	rows, err := q.db.Query("select id, name, email, created_at, updated_at from users")

	if err != nil {
		return nil, domain.NewAppError("GetAllUsers failed", err)
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, domain.NewAppError("GetAllUsers failed during scanning of rows", err)
		}
		users = append(users, user)
	}

	return users, nil
}

// CreateUser create and return a user record in the users table. No validation
func (q *SQLStore) CreateUser(input domain.User) (domain.User, *domain.AppError) {

	var u domain.User

	row := q.db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email, created_at, updated_at",
		input.Name, input.Email)

	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return u, domain.NewAppError("CreateUser failed during scanning", err)
	}

	return u, nil

}
