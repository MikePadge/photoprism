package entity

import (
	"strings"
	"time"

	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
<<<<<<< HEAD
	"github.com/mikepadge/photoprism/internal/mutex"
	"github.com/mikepadge/photoprism/pkg/rnd"
=======
	"github.com/photoprism/photoprism/internal/mutex"
	"github.com/photoprism/photoprism/pkg/rnd"
	"github.com/photoprism/photoprism/pkg/txt"
<<<<<<< HEAD
>>>>>>> 7cbdd31793e34cddb2c20a04d20d8ae5d25d7729
=======
>>>>>>> 5fba03844298ab501ce513a3f967b7578bc09707
)

// Labels for photo, album and location categorization
type Label struct {
	ID               uint   `gorm:"primary_key"`
	LabelUUID        string `gorm:"type:varbinary(36);unique_index;"`
	LabelSlug        string `gorm:"type:varbinary(128);index;"`
	LabelName        string `gorm:"type:varchar(128);"`
	LabelPriority    int
	LabelFavorite    bool
	LabelDescription string   `gorm:"type:text;"`
	LabelNotes       string   `gorm:"type:text;"`
	LabelCategories  []*Label `gorm:"many2many:categories;association_jointable_foreignkey:category_id"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time `sql:"index"`
	New              bool       `gorm:"-"`
}

func (m *Label) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("LabelUUID", rnd.PPID('l')); err != nil {
		log.Errorf("label: %s", err)
		return err
	}

	return nil
}

func NewLabel(labelName string, labelPriority int) *Label {
	labelName = strings.TrimSpace(labelName)

	if labelName == "" {
		labelName = "Unknown"
	}

	labelSlug := slug.Make(labelName)

	result := &Label{
		LabelName:     labelName,
		LabelSlug:     labelSlug,
		LabelPriority: labelPriority,
	}

	return result
}

func (m *Label) FirstOrCreate(db *gorm.DB) *Label {
	mutex.Db.Lock()
	defer mutex.Db.Unlock()

	if err := db.FirstOrCreate(m, "label_slug = ?", m.LabelSlug).Error; err != nil {
		log.Errorf("label: %s", err)
	}

	return m
}

func (m *Label) AfterCreate(scope *gorm.Scope) error {
	return scope.SetColumn("New", true)
}

func (m *Label) Rename(name string) {
	name = txt.Clip(name, 128)

	if name == "" {
		name = txt.SlugToTitle(m.LabelSlug)
	}

	m.LabelName = name
}
