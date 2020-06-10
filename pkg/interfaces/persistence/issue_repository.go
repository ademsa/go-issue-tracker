package persistence

import (
	"github.com/jinzhu/gorm"
	"go-issue-tracker/pkg/domain"
)

// SQLiteIssueRepository is a repository
type SQLiteIssueRepository struct {
	db *gorm.DB
}

// NewSQLiteIssueRepository to create SQLiteIssueRepository
func NewSQLiteIssueRepository(db *gorm.DB) *SQLiteIssueRepository {
	return &SQLiteIssueRepository{
		db: db,
	}
}

// Add to add new issue
func (r *SQLiteIssueRepository) Add(issue *domain.Issue) (*domain.Issue, error) {
	if err := r.db.Create(issue).Error; err != nil {
		return nil, err
	}
	return issue, nil
}

// Update to update issue
func (r *SQLiteIssueRepository) Update(issue domain.Issue) (domain.Issue, error) {
	if err := r.db.Model(&issue).Association("Labels").Replace(issue.Labels).Error; err != nil {
		return issue, err
	}
	if err := r.db.Save(&issue).Error; err != nil {
		return issue, err
	}
	return issue, nil
}

// FindByID to find issue by ID
func (r *SQLiteIssueRepository) FindByID(id uint) (domain.Issue, error) {
	var item domain.Issue
	if err := r.db.Preload("Project").Preload("Labels").Where("ID = ?", id).First(&item).Error; err != nil {
		return item, err
	}
	return item, nil
}

// FindAll to find all issues
func (r *SQLiteIssueRepository) FindAll() ([]domain.Issue, error) {
	var items []domain.Issue
	if err := r.db.Preload("Project").Preload("Labels").Find(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}

// Remove to remove issue
func (r *SQLiteIssueRepository) Remove(id uint) (bool, error) {
	tx := r.db.Begin()
	if err := r.db.Exec("DELETE FROM \"issues_labels\" WHERE issue_id=?", id).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	if err := r.db.Where("ID = ?", id).Delete(domain.Issue{}).Error; err != nil {
		tx.Rollback()
		return false, err
	}
	tx.Commit()
	return true, nil
}
