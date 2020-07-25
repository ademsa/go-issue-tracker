package persistence

import (
	"github.com/jinzhu/gorm"
	"go-issue-tracker/pkg/domain"
)

// SQLiteLabelRepository is a repository
type SQLiteLabelRepository struct {
	db *gorm.DB
}

// NewSQLiteLabelRepository to create SQLiteLabelRepository
func NewSQLiteLabelRepository(db *gorm.DB) *SQLiteLabelRepository {
	return &SQLiteLabelRepository{
		db: db,
	}
}

// Add to add new label
func (r *SQLiteLabelRepository) Add(label *domain.Label) (*domain.Label, error) {
	if err := r.db.Create(label).Error; err != nil {
		return nil, err
	}
	return label, nil
}

// Update to update label
func (r *SQLiteLabelRepository) Update(label domain.Label) (domain.Label, error) {
	if err := r.db.Save(&label).Error; err != nil {
		return label, err
	}
	return label, nil
}

// FindByID to find label by ID
func (r *SQLiteLabelRepository) FindByID(id uint) (domain.Label, error) {
	var item domain.Label
	if err := r.db.Where("ID = ?", id).First(&item).Error; err != nil {
		return item, err
	}
	return item, nil
}

// FindByName to find label by name
func (r *SQLiteLabelRepository) FindByName(name string) (domain.Label, error) {
	var item domain.Label
	if err := r.db.Where("name = ?", name).First(&item).Error; err != nil {
		return item, err
	}
	return item, nil
}

// Find to find labels
func (r *SQLiteLabelRepository) Find(name string) ([]domain.Label, error) {
	var items []domain.Label
	if err := r.db.Where("name LIKE ?", "%"+name+"%").Find(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}

// FindAll to find all labels
func (r *SQLiteLabelRepository) FindAll() ([]domain.Label, error) {
	var items []domain.Label
	if err := r.db.Find(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}

// Remove to remove label
func (r *SQLiteLabelRepository) Remove(id uint) (bool, error) {
	var c int
	r.db.Table("issues_labels").Where("label_id = ?", id).Count(&c)
	if c > 0 {
		return false, nil
	}
	if err := r.db.Where("ID = ?", id).Delete(domain.Label{}).Error; err != nil {
		return false, err
	}
	return true, nil
}
