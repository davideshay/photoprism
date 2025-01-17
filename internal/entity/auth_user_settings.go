package entity

import (
	"fmt"
	"time"

	"github.com/photoprism/photoprism/internal/customize"
	"github.com/photoprism/photoprism/pkg/rnd"
)

// UserSettings represents user preferences.
type UserSettings struct {
	UserUID     string    `gorm:"type:VARBINARY(64);primary_key;auto_increment:false;" json:"-" yaml:"UserUID"`
	UIHome      string    `gorm:"type:VARBINARY(32);column:ui_home;" json:"UIHome,omitempty" yaml:"UIHome,omitempty"`
	UITheme     string    `gorm:"type:VARBINARY(32);column:ui_theme;" json:"UITheme,omitempty" yaml:"UITheme,omitempty"`
	UILanguage  string    `gorm:"type:VARBINARY(32);column:ui_language;" json:"UILanguage,omitempty" yaml:"UILanguage,omitempty"`
	UITimeZone  string    `gorm:"type:VARBINARY(64);column:ui_time_zone;" json:"UITimeZone,omitempty" yaml:"UITimeZone,omitempty"`
	MapsStyle   string    `gorm:"type:VARBINARY(32);" json:"MapsStyle,omitempty" yaml:"MapsStyle,omitempty"`
	MapsAnimate int       `json:"MapsAnimate,omitempty" yaml:"MapsAnimate,omitempty"`
	IndexPath   string    `gorm:"type:VARBINARY(500);" json:"IndexPath,omitempty" yaml:"IndexPath,omitempty"`
	IndexRescan int       `json:"IndexRescan,omitempty" yaml:"IndexRescan,omitempty"`
	ImportPath  string    `gorm:"type:VARBINARY(500);" json:"ImportPath,omitempty" yaml:"ImportPath,omitempty"`
	ImportMove  int       `json:"ImportMove,omitempty" yaml:"ImportMove,omitempty"`
	UploadPath  string    `gorm:"type:VARBINARY(500);" json:"UploadPath,omitempty" yaml:"UploadPath,omitempty"`
	CreatedAt   time.Time `json:"CreatedAt" yaml:"-"`
	UpdatedAt   time.Time `json:"UpdatedAt" yaml:"-"`
}

// TableName returns the entity table name.
func (UserSettings) TableName() string {
	return "auth_users_settings_dev"
}

// NewUserSettings creates new user preferences.
func NewUserSettings(uid string) *UserSettings {
	return &UserSettings{UserUID: uid}
}

// CreateUserSettings creates new user settings or returns nil on error.
func CreateUserSettings(user *User) error {
	if user == nil {
		return fmt.Errorf("user is nil")
	}

	if user.UID() == "" {
		return fmt.Errorf("empty user uid")
	}

	user.UserSettings = &UserSettings{}

	if err := Db().Where("user_uid = ?", user.UID()).First(user.UserSettings).Error; err == nil {
		return nil
	}

	return user.UserSettings.Create()
}

// HasID tests if the entity has a valid uid.
func (m *UserSettings) HasID() bool {
	return rnd.IsUID(m.UserUID, UserUID)
}

// Create new entity in the database.
func (m *UserSettings) Create() error {
	return Db().Create(m).Error
}

// Save entity properties.
func (m *UserSettings) Save() error {
	return Db().Save(m).Error
}

// Updates multiple properties in the database.
func (m *UserSettings) Updates(values interface{}) error {
	return UnscopedDb().Model(m).Updates(values).Error
}

// Apply applies the settings provided to the user preferences and keeps current values if they are not specified.
func (m *UserSettings) Apply(s *customize.Settings) *UserSettings {
	// UI preferences.
	if s.UI.Theme != "" {
		m.UITheme = s.UI.Theme
	}

	if s.UI.Language != "" {
		m.UILanguage = s.UI.Language
	}

	if s.UI.TimeZone != "" {
		m.UITimeZone = s.UI.TimeZone
	}

	// Maps preferences.
	if s.Maps.Style != "" {
		m.MapsStyle = s.Maps.Style

		if s.Maps.Animate > 0 {
			m.MapsAnimate = s.Maps.Animate
		} else {
			m.MapsAnimate = -1
		}
	}

	// Index preferences.
	if s.Index.Path != "" {
		m.IndexPath = s.Index.Path

		if s.Index.Rescan {
			m.IndexRescan = 1
		} else {
			m.IndexRescan = -1
		}
	}

	// Import preferences.
	if s.Import.Path != "" {
		m.ImportPath = s.Import.Path

		if s.Import.Move {
			m.ImportMove = 1
		} else {
			m.ImportMove = -1
		}
	}

	return m
}

// ApplyTo applies the user preferences to the client settings and keeps the default settings if they are not specified.
func (m *UserSettings) ApplyTo(s *customize.Settings) *customize.Settings {
	if m.UITheme != "" {
		s.UI.Theme = m.UITheme
	}

	if m.UILanguage != "" {
		s.UI.Language = m.UILanguage
	}

	if m.UITimeZone != "" {
		s.UI.TimeZone = m.UITimeZone
	}

	if m.MapsStyle != "" {
		s.Maps.Style = m.MapsStyle
	}

	if m.MapsAnimate > 0 {
		s.Maps.Animate = m.MapsAnimate
	} else if m.MapsAnimate < 0 {
		s.Maps.Animate = 0
	}

	if m.IndexPath != "" {
		s.Index.Path = m.IndexPath
	}

	if m.IndexRescan > 0 {
		s.Index.Rescan = true
	} else if m.IndexRescan < 0 {
		s.Index.Rescan = false
	}

	if m.ImportPath != "" {
		s.Import.Path = m.ImportPath
	}

	if m.ImportMove > 0 {
		s.Import.Move = true
	} else if m.ImportMove < 0 {
		s.Import.Move = false
	}

	return s
}
