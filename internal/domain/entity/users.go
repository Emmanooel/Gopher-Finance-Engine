package entity

import "time"

type Users struct {
	Id        string     `json:"id"`
	Name      string     `json:"name" binding:"required"`
	Email     string     `json:"email" binding:"required"`
	Password  *string    `json:"password" binding:"required,min=8"`
	Rule      string     `json:"rule"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Rule  string `json:"rule"`
}

type Meta struct {
	Page       int  `json:"page"`
	PerPage    int  `json:"per_page"`
	TotalPages int  `json:"total_pages"`
	TotalItems int  `json:"total_items"`
	NextPage   *int `json:"next_page,omitempty"`
	PrevPage   *int `json:"prev_page,omitempty"`
}

type UsersResponse struct {
	Data []*Users `json:"data"`
	Meta Meta     `json:"meta"`
}

type UserLogin struct {
	Email    string `json:"email"binding:"required,email"`
	Password string `json:"password"binding:"required,min=8"`
}

func (u *Users) BuildUser() *User {
	return &User{
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email,
		Rule:  u.Rule,
	}
}
