package schemas

type Notes struct {
	Model
	NotesDTO
}

type NotesDTO struct {
	Title   string `json:"title" gorm:"not null"`
	Content string `json:"content" gorm:"not null"`
}
