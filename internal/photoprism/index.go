package photoprism

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/jinzhu/gorm"
<<<<<<< HEAD
	"github.com/mikepadge/photoprism/internal/classify"
	"github.com/mikepadge/photoprism/internal/config"
	"github.com/mikepadge/photoprism/internal/event"
	"github.com/mikepadge/photoprism/internal/mutex"
	"github.com/mikepadge/photoprism/internal/nsfw"
=======
	"github.com/photoprism/photoprism/internal/classify"
	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/internal/mutex"
	"github.com/photoprism/photoprism/internal/nsfw"
	"github.com/photoprism/photoprism/pkg/fs"
<<<<<<< HEAD
>>>>>>> 7cbdd31793e34cddb2c20a04d20d8ae5d25d7729
=======
>>>>>>> 5fba03844298ab501ce513a3f967b7578bc09707
)

// Index represents an indexer that indexes files in the originals directory.
type Index struct {
	conf         *config.Config
	tensorFlow   *classify.TensorFlow
	nsfwDetector *nsfw.Detector
	db           *gorm.DB
}

// NewIndex returns a new indexer and expects its dependencies as arguments.
func NewIndex(conf *config.Config, tensorFlow *classify.TensorFlow, nsfwDetector *nsfw.Detector) *Index {
	i := &Index{
		conf:         conf,
		tensorFlow:   tensorFlow,
		nsfwDetector: nsfwDetector,
		db:           conf.Db(),
	}

	return i
}

func (ind *Index) originalsPath() string {
	return ind.conf.OriginalsPath()
}

func (ind *Index) thumbnailsPath() string {
	return ind.conf.ThumbnailsPath()
}

// Cancel stops the current indexing operation.
func (ind *Index) Cancel() {
	mutex.Worker.Cancel()
}

// Start indexes media files in the originals directory.
func (ind *Index) Start(options IndexOptions) map[string]bool {
	done := make(map[string]bool)
	originalsPath := ind.originalsPath()

	if !fs.PathExists(originalsPath) {
		event.Error(fmt.Sprintf("index: %s does not exist", originalsPath))
		return done
	}

	if err := mutex.Worker.Start(); err != nil {
		event.Error(fmt.Sprintf("index: %s", err.Error()))
		return done
	}

	defer mutex.Worker.Stop()

	if err := ind.tensorFlow.Init(); err != nil {
		log.Errorf("index: %s", err.Error())

		return done
	}

	jobs := make(chan IndexJob)

	// Start a fixed number of goroutines to index files.
	var wg sync.WaitGroup
	var numWorkers = ind.conf.Workers()
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			indexWorker(jobs) // HLc
			wg.Done()
		}()
	}

	err := filepath.Walk(originalsPath, func(fileName string, fileInfo os.FileInfo, err error) error {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("index: %s [panic]", err)
			}
		}()

		if mutex.Worker.Canceled() {
			return errors.New("indexing canceled")
		}

		if err != nil || done[fileName] {
			return nil
		}

		if fileInfo.IsDir() || strings.HasPrefix(filepath.Base(fileName), ".") {
			return nil
		}

		mf, err := NewMediaFile(fileName)

		if err != nil || !mf.IsPhoto() {
			return nil
		}

		related, err := mf.RelatedFiles()

		if err != nil {
			log.Warnf("index: %s", err.Error())

			return nil
		}

		var files MediaFiles

		for _, f := range related.files {
			if done[f.FileName()] {
				continue
			}

			files = append(files, f)
			done[f.FileName()] = true
		}

		done[mf.FileName()] = true

		related.files = files

		jobs <- IndexJob{
			filename: mf.FileName(),
			related:  related,
			opt:      options,
			ind:      ind,
		}

		return nil
	})

	close(jobs)
	wg.Wait()

	if err != nil {
		log.Error(err.Error())
	}

	return done
}
