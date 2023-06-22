package storer

type StorerRepository interface {
	Save(s *Storer) error
}
