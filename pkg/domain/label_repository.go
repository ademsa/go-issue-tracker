package domain

// LabelRepository repository
type LabelRepository interface {
	Add(label *Label) (*Label, error)
	Update(label Label) (Label, error)
	FindByID(id uint) (Label, error)
	FindByName(name string) (Label, error)
	Find(name string) ([]Label, error)
	FindAll() ([]Label, error)
	Remove(id uint) (bool, error)
}
