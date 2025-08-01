package models

type User struct {
	ID       int64
	email    string `binding:"required"`
	password string `binding:"required"`
}

func (u User) Save() {
	query := `
		INSERT INTO users (email, password)
		VALUES (?, ?)
	`

}
