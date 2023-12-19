package models

type Album struct {
	ID     string  `json:"id" xml:"id"`
	Title  string  `json:"title" xml:"title"`
	Artist string  `json:"artist" xml:"artist"`
	Price  float64 `json:"price" xml:"price"`
}
