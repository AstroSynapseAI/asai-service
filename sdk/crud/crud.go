package crud

type Repository[T any] interface {
	Create(model T) (T, error)
	Read(ID uint) (T, error)
	ReadAll() ([]T, error)
	Update(ID uint, model T) (T, error)
	Delete(ID uint) error
}
