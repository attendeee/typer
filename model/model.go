package model

type Book struct {
	Chapters []Chapter `json:"chapters"`
}

type Chapter struct {
	Title string   `json:"title"`
	Text  []string `json:"text"`
}
