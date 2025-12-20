package model

type Book struct {
	Chapters []Chapter `json:"chapters"`
}

type Chapter struct {
	Title string   `json:"title"`
	Text  []string `json:"text"`
}

type Model struct {
	Book      Book
	Chapter   int
	Text      string
	CursorPos int

	OffsetStep      int
	Offsets         []int
	UpperOffset     int
	UpperOffsetIdx  int
	BottomOffset    int
	BottomOffsetIdx int
}
