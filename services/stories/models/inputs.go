package models

type CreateStoryInput struct {
    UserID   string   `json:"userId"`
    Title    string   `json:"title"`
    Content  string   `json:"content"`
    Tags     []string `json:"tags"`
    Category string  `json:"category"`
    Mood     *string  `json:"mood,omitempty"`
}

type UpdateStoryInput struct {
    ID       string   `json:"id"`
    Title    *string   `json:"title"`
    Content  *string   `json:"content"`
    Tags     *[]string `json:"tags"`
    Category *string   `json:"category"` 
    Mood     *string   `json:"mood"`     
}

type DeleteStoryInput struct {
    ID *string `json:"id"`
    Title    **string   `json:"title"`
}