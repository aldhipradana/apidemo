package schemas

// User model
// @Model User
// @Description User model
type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"`
}
