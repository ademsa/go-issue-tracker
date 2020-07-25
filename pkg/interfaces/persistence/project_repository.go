package persistence

import (
	"github.com/jinzhu/gorm"
	"go-issue-tracker/pkg/domain"
)

// SQLiteProjectRepository is a repository
type SQLiteProjectRepository struct {
	db *gorm.DB
}

// NewSQLiteProjectRepository to create SQLiteProjectRepository
func NewSQLiteProjectRepository(db *gorm.DB) *SQLiteProjectRepository {
	return &SQLiteProjectRepository{
		db: db,
	}
}

// Add to add new project
func (r *SQLiteProjectRepository) Add(project *domain.Project) (*domain.Project, error) {
	if err := r.db.Create(project).Error; err != nil {
		return nil, err
	}
	return project, nil
}

// Update to update project
func (r *SQLiteProjectRepository) Update(project domain.Project) (domain.Project, error) {
	if err := r.db.Save(&project).Error; err != nil {
		return project, err
	}
	return project, nil
}

// FindByID to find project by ID
func (r *SQLiteProjectRepository) FindByID(id uint) (domain.Project, error) {
	var item domain.Project
	if err := r.db.Where("ID = ?", id).First(&item).Error; err != nil {
		return item, err
	}
	return item, nil
}

// Find to find projects
func (r *SQLiteProjectRepository) Find(name string) ([]domain.Project, error) {
	var items []domain.Project
	if err := r.db.Where("name LIKE ?", "%"+name+"%").Find(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}

// FindAll to find all projects
func (r *SQLiteProjectRepository) FindAll() ([]domain.Project, error) {
	var items []domain.Project
	if err := r.db.Find(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}

// Remove to remove project
func (r *SQLiteProjectRepository) Remove(id uint) (bool, error) {
	var c int
	r.db.Table("issues").Where("project_id = ?", id).Count(&c)
	if c > 0 {
		return false, nil
	}
	if err := r.db.Where("ID = ?", id).Delete(domain.Project{}).Error; err != nil {
		return false, err
	}
	return true, nil
}
