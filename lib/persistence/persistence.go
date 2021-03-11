package persistence

type DatabaseHandler interface {
	AddEbook(Ebook) ([]byte, error)
	FindEbookByName(string) (Ebook, error)
}
