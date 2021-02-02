package books

type Book interface {
	Init() error
	Create(title string, author string) (*string, error)
}
