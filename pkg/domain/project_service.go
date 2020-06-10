package domain

// ProjectService interface
type ProjectService interface {
	Add(project *Project) (*Project, error)
	Update(project Project) (Project, error)
	FindByID(id uint) (Project, error)
	FindAll() ([]Project, error)
	Remove(id uint) (bool, error)
}

// projectService struct
type projectService struct {
	repository ProjectRepository
}

// GetDefaultProjectService alias to newProjectService
var GetDefaultProjectService = newProjectService

// ResetDefaultProjectService to reset GetDefaultIssueService value
func ResetDefaultProjectService() {
	GetDefaultProjectService = newProjectService
}

// newProjectService to create new ProjectService
func newProjectService(repository ProjectRepository) ProjectService {
	return &projectService{
		repository: repository,
	}
}

// Add to add new project
func (s *projectService) Add(project *Project) (*Project, error) {
	item, err := s.repository.Add(project)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Update to update project
func (s *projectService) Update(project Project) (Project, error) {
	item, err := s.repository.Update(project)
	if err != nil {
		return item, err
	}
	return item, nil
}

// FindByID to find project by ID
func (s *projectService) FindByID(id uint) (Project, error) {
	item, err := s.repository.FindByID(id)
	if err != nil {
		return item, err
	}
	return item, nil
}

// FindAll to find all projects
func (s *projectService) FindAll() ([]Project, error) {
	items, err := s.repository.FindAll()
	if err != nil {
		return items, err
	}
	return items, nil
}

// Remove to remove project
func (s *projectService) Remove(id uint) (bool, error) {
	status, err := s.repository.Remove(id)
	if err != nil {
		return status, err
	}
	return status, nil
}
