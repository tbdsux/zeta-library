package lib

var CollectionTypes = []string{"series", "movies", "asian_drama", "anime", "books", "manga"}

type CollectionProps struct {
	Name        string `json:"name"`
	ID          string `json:"id,omitempty"`
	Key         string `json:"key,omitempty"`
	Type        string `json:"type"`
	Description string `json:"description"`
	CreatedAt   int    `json:"created_at"`
}
