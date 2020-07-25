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

// Find to find issues
func (r *SQLiteIssueRepository) Find(title string, projectID uint, labels []string) ([]domain.Issue, error) {
	var items []domain.Issue
	query := ""
	args := []interface{}{}
	if title != "" {
		query = "title LIKE ?"
		args = append(args, "%"+title+"%")
	}
	if projectID != uint(0) {
		if query != "" {
			query += " AND project_id = ?"
		} else {
			query += "project_id = ?"
		}
		args = append(args, projectID)
	}
	if len(labels) > 0 {
		if query != "" {
			query += " AND \"issues_labels\".\"label_id\" IN (?)"
		} else {
			query += "\"issues_labels\".\"label_id\" IN (?)"
		}
		args = append(args, labels)
	}
	if query == "" {
		if err := r.db.Preload("Project").Preload("Labels").Find(&items).Error; err != nil {
			return items, err
		}
	} else if len(labels) > 0 {
		if err := r.db.Preload("Project").Preload("Labels").Joins("INNER JOIN \"issues_labels\" ON \"issues_labels\".\"issue_id\" = \"issues\".\"id\"").Where(query, args...).Select("DISTINCT \"issues\".*").Find(&items).Error; err != nil {
			return items, err
		}
	} else {
		if err := r.db.Preload("Project").Preload("Labels").Where(query, args...).Find(&items).Error; err != nil {
			return items, err
		}
	}
	return items, nil
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
