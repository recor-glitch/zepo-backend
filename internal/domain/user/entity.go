package user

type User struct {
	Id    string `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Email string `gorm:"unique_index" json:"email"`
	Image string `json:"image"`
	Role  string `json:"role"`
}