package model

type CreatePost struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdatePost struct {
	Title       string `json:"title" `
	Description string `json:"description" `
}
