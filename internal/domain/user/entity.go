package user

type User struct {
	Id        string `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Email     string `gorm:"unique" json:"email"`
	Image     string `json:"image"`
	Role      string `gorm:"default:user" json:"role"`
	CreatedAt string `gorm:"type:timestamptz;not null;default:NOW()" json:"created_at"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
