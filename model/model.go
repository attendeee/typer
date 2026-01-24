package model

type Book struct {
	Chapters []Chapter `json:"chapters"`
}

type Chapter struct {
	Title string   `json:"title"`
	Text  []string `json:"text"`
}

type State struct {
	Chapter      uint32 `json:"chapter"`
	CursorPos    uint32 `json:"cursor_pos"`
	ErrorCounter uint   `json:"error_counter"`
}
