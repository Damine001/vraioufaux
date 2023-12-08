package app_auth

type InsertUser struct {
	Name       string  `gorm:"column:name;not null" json:"name"`
	FamilyName string  `gorm:"column:family_name;not null" json:"family_name"`
	Email      string  `gorm:"column:email;not null;unique" json:"email"`
	Password   string  `gorm:"column:password;not null" json:"password"`
	Points     float32 `gorm:"column:points;not null" json:"points"`
	Role       string  `gorm:"column:role;not null" json:"role"`
}

type UserLogIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogedIn struct {
	Token string `json:"token"`
}
