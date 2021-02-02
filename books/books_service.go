package books

type Book interface {
	Create(title string, author string) (*string, error)
}
