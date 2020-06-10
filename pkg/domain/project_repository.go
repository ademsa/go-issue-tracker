package domain

// ProjectRepository repository
type ProjectRepository interface {
	Add(project *Project) (*Project, error)
	Update(project Project) (Project, error)
	FindByID(id uint) (Project, error)
	FindAll() ([]Project, error)
	Remove(id uint) (bool, error)
}
