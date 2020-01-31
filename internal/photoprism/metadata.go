package photoprism

import (
	"github.com/mikepadge/photoprism/internal/meta"
)

// MetaData returns exif meta data of a media file.
func (m *MediaFile) MetaData() (result meta.Data, err error) {
	m.once.Do(func() { m.metaData, err = meta.Exif(m.Filename()) })
	return m.metaData, err
}
