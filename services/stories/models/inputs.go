package models

type CreateStoryInput struct {
    Title    string  `json:"title"`
    Content  string  `json:"content"`
    Category *string `json:"category,omitempty"`
    Mood     *string `json:"mood,omitempty"`
}

type UpdateStoryInput struct {
    ID       string  `json:"id"`
    Title    *string `json:"title,omitempty"`
    Content  *string `json:"content,omitempty"`
    Category *string `json:"category,omitempty"`
    Mood     *string `json:"mood,omitempty"`
}

type DeleteStoryInput struct {
    ID string `json:"id"`
}