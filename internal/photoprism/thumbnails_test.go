package photoprism

import (
	"os"
	"testing"

	"github.com/disintegration/imaging"
	"github.com/mikepadge/photoprism/internal/classify"
	"github.com/mikepadge/photoprism/internal/entity"
	"github.com/mikepadge/photoprism/internal/nsfw"
	"github.com/mikepadge/photoprism/internal/thumb"

	"github.com/mikepadge/photoprism/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestThumbnails_Start(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	conf := config.TestConfig()

	if err := conf.CreateDirectories(); err != nil {
		t.Fatal(err)
	}

	conf.InitializeTestData(t)

	tf := classify.New(conf.ResourcesPath(), conf.TensorFlowDisabled())
	nd := nsfw.New(conf.NSFWModelPath())

	ind := NewIndex(conf, tf, nd)

	convert := NewConvert(conf)

	imp := NewImport(conf, ind, convert)

	imp.Start(conf.ImportPath())

	thumbnails := NewThumbnails(conf)

	err := thumbnails.Start(true)

	if err != nil {
		t.Fatal(err)
	}
}

func TestThumbnails_Filename(t *testing.T) {
	conf := config.TestConfig()

	thumbsPath := conf.CachePath() + "/_tmp"

	defer os.RemoveAll(thumbsPath)

	if err := conf.CreateDirectories(); err != nil {
		t.Error(err)
	}

	t.Run("", func(t *testing.T) {
		filename, err := thumb.Filename("99988", thumbsPath, 150, 150, thumb.ResampleFit, thumb.ResampleNearestNeighbor)
		assert.Nil(t, err)
		assert.Equal(t, "/go/src/github.com/mikepadge/photoprism/assets/testdata/cache/_tmp/9/9/9/99988_150x150_fit.jpg", filename)
	})
	t.Run("hash too short", func(t *testing.T) {
		_, err := thumb.Filename("999", thumbsPath, 150, 150, thumb.ResampleFit, thumb.ResampleNearestNeighbor)
		if err == nil {
			t.FailNow()
		}

		assert.Equal(t, "thumbs: file hash is empty or too short (\"999\")", err.Error())
	})
	t.Run("invalid width", func(t *testing.T) {
		_, err := thumb.Filename("99988", thumbsPath, -4, 150, thumb.ResampleFit, thumb.ResampleNearestNeighbor)
		if err == nil {
			t.FailNow()
		}
		assert.Equal(t, "thumbs: width exceeds limit (-4)", err.Error())
	})
	t.Run("invalid height", func(t *testing.T) {
		_, err := thumb.Filename("99988", thumbsPath, 200, -1, thumb.ResampleFit, thumb.ResampleNearestNeighbor)
		if err == nil {
			t.FailNow()
		}
		assert.Equal(t, "thumbs: height exceeds limit (-1)", err.Error())
	})
	t.Run("empty thumbpath", func(t *testing.T) {
		path := ""
		_, err := thumb.Filename("99988", path, 200, 150, thumb.ResampleFit, thumb.ResampleNearestNeighbor)
		if err == nil {
			t.FailNow()
		}
		assert.Equal(t, "thumbs: path is empty", err.Error())
	})
}

func TestThumbnails_ThumbnailFromFile(t *testing.T) {
	conf := config.TestConfig()

	thumbsPath := conf.CachePath() + "/_tmp"

	defer os.RemoveAll(thumbsPath)

	if err := conf.CreateDirectories(); err != nil {
		t.Error(err)
	}

	t.Run("valid parameter", func(t *testing.T) {
		fileModel := &entity.File{
			FileName: conf.ExamplesPath() + "/elephants.jpg",
			FileHash: "1234568889",
		}

		thumbnail, err := thumb.FromFile(fileModel.FileName, fileModel.FileHash, thumbsPath, 224, 224)
		assert.Nil(t, err)
		assert.FileExists(t, thumbnail)
	})

	t.Run("hash too short", func(t *testing.T) {
		fileModel := &entity.File{
			FileName: conf.ExamplesPath() + "/elephants.jpg",
			FileHash: "123",
		}

		_, err := thumb.FromFile(fileModel.FileName, fileModel.FileHash, thumbsPath, 224, 224)
		if err == nil {
			t.FailNow()
		}
		assert.Equal(t, "thumbs: file hash is empty or too short (\"123\")", err.Error())
	})
	t.Run("filename too short", func(t *testing.T) {
		fileModel := &entity.File{
			FileName: "xxx",
			FileHash: "12367890",
		}

		_, err := thumb.FromFile(fileModel.FileName, fileModel.FileHash, thumbsPath, 224, 224)
		if err == nil {
			t.FailNow()
		}
		assert.Equal(t, "thumbs: image filename is empty or too short (\"xxx\")", err.Error())
	})
}

func TestThumbnails_CreateThumbnail(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	conf := config.TestConfig()

	thumbsPath := conf.CachePath() + "/_tmp"

	defer os.RemoveAll(thumbsPath)

	if err := conf.CreateDirectories(); err != nil {
		t.Error(err)
	}

	t.Run("valid parameter", func(t *testing.T) {
		expectedFilename, err := thumb.Filename("12345", thumbsPath, 150, 150, thumb.ResampleFit, thumb.ResampleNearestNeighbor)

		if err != nil {
			t.Error(err)
		}

		img, err := imaging.Open(conf.ExamplesPath()+"/elephants.jpg", imaging.AutoOrientation(true))

		if err != nil {
			t.Errorf("can't open original: %s", err)
		}

		thumbnail, err := thumb.Create(img, expectedFilename, 150, 150, thumb.ResampleFit, thumb.ResampleNearestNeighbor)

		assert.Empty(t, err)

		assert.NotNil(t, thumbnail)

		bounds := thumbnail.Bounds()

		assert.Equal(t, 150, bounds.Dx())
		assert.Equal(t, 99, bounds.Dy())

		assert.FileExists(t, expectedFilename)
	})
	t.Run("invalid width", func(t *testing.T) {
		expectedFilename, err := thumb.Filename("12345", thumbsPath, 150, 150, thumb.ResampleFit, thumb.ResampleNearestNeighbor)

		if err != nil {
			t.Error(err)
		}

		img, err := imaging.Open(conf.ExamplesPath()+"/elephants.jpg", imaging.AutoOrientation(true))

		if err != nil {
			t.Errorf("can't open original: %s", err)
		}

		thumbnail, err := thumb.Create(img, expectedFilename, -1, 150, thumb.ResampleFit, thumb.ResampleNearestNeighbor)
		if err == nil {
			t.FailNow()
		}
		assert.Equal(t, "thumbs: width has an invalid value (-1)", err.Error())
		bounds := thumbnail.Bounds()
		assert.NotEqual(t, 150, bounds.Dx())
	})

	t.Run("invalid height", func(t *testing.T) {
		expectedFilename, err := thumb.Filename("12345", thumbsPath, 150, 150, thumb.ResampleFit, thumb.ResampleNearestNeighbor)

		if err != nil {
			t.Error(err)
		}

		img, err := imaging.Open(conf.ExamplesPath()+"/elephants.jpg", imaging.AutoOrientation(true))

		if err != nil {
			t.Errorf("can't open original: %s", err)
		}

		thumbnail, err := thumb.Create(img, expectedFilename, 150, -1, thumb.ResampleFit, thumb.ResampleNearestNeighbor)
		if err == nil {
			t.FailNow()
		}
		assert.Equal(t, "thumbs: height has an invalid value (-1)", err.Error())
		bounds := thumbnail.Bounds()
		assert.NotEqual(t, 150, bounds.Dx())
	})
}
