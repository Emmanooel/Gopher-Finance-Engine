package repository

import (
	"context"
	"gopher-finance-engine/internal/domain/entity"
	"gopher-finance-engine/internal/domain/infra/repository"
	"gopher-finance-engine/pkg/postgres"

	"go.uber.org/zap"
)

type UserRepository struct {
	logger *zap.Logger
}

func NewUserRepository(
	logger *zap.Logger,
) repository.UsersRepositoryI {
	return &UserRepository{
		logger: logger,
	}
}

func (u *UserRepository) CreateUser(ctx context.Context, body entity.Users) error {
	const query = `
		INSERT INTO users (id, name, email, password, role, created_at) VALUES ($1, $2, $3, $4, $5, $6)
	`
	defer postgres.Db.Close()

	_, err := postgres.Db.Exec(
		ctx,
		query,
		body.Id,
		body.Name,
		body.Email,
		body.Password,
		body.Role,
		body.CreatedAt,
	)

	if err != nil {
		u.logger.Error("error on create user", zap.Error(err))
		return err
	}

	u.logger.Info("user created successfully in database")
	return nil
}

func (u *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.Users, error) {
	const query = `
		SELECT id, name, email, password, role, created_at FROM users WHERE email = $1
	`

	resp := postgres.Db.QueryRow(ctx, query, email)

	var us entity.Users

	err := resp.Scan(
		&us.Id,
		&us.Name,
		&us.Email,
		&us.Password,
		&us.Role,
		&us.CreatedAt,
	)

	if err != nil {
		u.logger.Error("error on find user by email", zap.Error(err))
		return nil, err
	}

	return &us, nil
}

func (u *UserRepository) GetAllUsers(ctx context.Context) ([]*entity.Users, error) {
	const query = `
		SELECT * FROM users 
	`

	var users []*entity.Users

	rows, err := postgres.Db.Query(ctx, query)

	err = rows.Err()

	if err != nil {
		u.logger.Error("error on get all users", zap.Error(err))
		return nil, err
	}

	for rows.Next() {
		var user *entity.Users
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
		)
		if err != nil {
			u.logger.Error("error on scan user", zap.Error(err))
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
