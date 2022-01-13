package token

type Type string

const (
	Price   Type = "price"
	Comment Type = "comment"
	Date    Type = "date"
)

type Token struct {
	Position int
	Value    string
	Type     Type
}
