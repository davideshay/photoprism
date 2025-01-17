package customize

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/photoprism/photoprism/internal/acl"
)

func TestSettings_ApplyACL(t *testing.T) {
	original := NewDefaultSettings().Features

	t.Run("RoleAdmin", func(t *testing.T) {
		s := NewDefaultSettings()

		expected := FeatureSettings{
			Account:   true,
			Advanced:  true,
			Albums:    true,
			Archive:   true,
			Delete:    false,
			Download:  true,
			Edit:      true,
			Estimates: true,
			Favorites: true,
			Files:     true,
			Folders:   true,
			Import:    true,
			Labels:    true,
			Library:   true,
			Logs:      true,
			Moments:   true,
			People:    true,
			Places:    true,
			Private:   true,
			Ratings:   true,
			Reactions: true,
			Review:    true,
			Search:    true,
			Settings:  true,
			Share:     true,
			Sync:      true,
			Upload:    true,
			Videos:    true,
		}

		assert.Equal(t, original, s.Features)
		r := s.ApplyACL(acl.Resources, acl.RoleAdmin)

		t.Logf("RoleAdmin: %#v", r)
		assert.Equal(t, expected, r.Features)
	})

	t.Run("RoleVisitor", func(t *testing.T) {
		s := NewDefaultSettings()

		expected := FeatureSettings{
			Account:   false,
			Advanced:  false,
			Albums:    true,
			Archive:   false,
			Delete:    false,
			Download:  true,
			Edit:      false,
			Estimates: true,
			Favorites: false,
			Files:     false,
			Folders:   false,
			Import:    false,
			Labels:    false,
			Library:   false,
			Logs:      false,
			Moments:   false,
			People:    false,
			Places:    false,
			Private:   false,
			Ratings:   false,
			Reactions: false,
			Review:    false,
			Search:    false,
			Settings:  false,
			Share:     false,
			Sync:      false,
			Upload:    false,
			Videos:    false,
		}

		assert.Equal(t, original, s.Features)
		r := s.ApplyACL(acl.Resources, acl.RoleVisitor)
		t.Logf("RoleVisitor: %#v", r)
		assert.Equal(t, expected, r.Features)
	})
}
