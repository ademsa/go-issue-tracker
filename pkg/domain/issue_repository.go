package domain

// IssueRepository repository
type IssueRepository interface {
	Add(issue *Issue) (*Issue, error)
	Update(issue Issue) (Issue, error)
	FindByID(id uint) (Issue, error)
	FindAll() ([]Issue, error)
	Remove(id uint) (bool, error)
}
