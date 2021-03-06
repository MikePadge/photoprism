package classify

import (
	"io/ioutil"
	"testing"

	tensorflow "github.com/tensorflow/tensorflow/tensorflow/go"

	"github.com/mikepadge/photoprism/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestTensorFlow_LabelsFromFile(t *testing.T) {
	t.Run("chameleon_lime.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		tensorFlow := New(conf.ResourcesPath(), conf.TensorFlowDisabled())

		result, err := tensorFlow.File(conf.ExamplesPath() + "/chameleon_lime.jpg")

		assert.Nil(t, err)

		if err != nil {
			t.Log(err.Error())
			t.Fail()
		}

		assert.NotNil(t, result)
		assert.IsType(t, Labels{}, result)
		assert.Equal(t, 1, len(result))

		t.Log(result)

		assert.Equal(t, "chameleon", result[0].Name)

		assert.Equal(t, 7, result[0].Uncertainty)
	})
	t.Run("not existing file", func(t *testing.T) {
		conf := config.TestConfig()

		tensorFlow := New(conf.ResourcesPath(), conf.TensorFlowDisabled())

		result, err := tensorFlow.File(conf.ExamplesPath() + "/notexisting.jpg")
		assert.Contains(t, err.Error(), "no such file or directory")
		assert.Empty(t, result)
	})
}

func TestTensorFlow_Labels(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	t.Run("chameleon_lime.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		tensorFlow := New(conf.ResourcesPath(), conf.TensorFlowDisabled())

		if imageBuffer, err := ioutil.ReadFile(conf.ExamplesPath() + "/chameleon_lime.jpg"); err != nil {
			t.Error(err)
		} else {
			result, err := tensorFlow.Labels(imageBuffer)

			t.Log(result)

			assert.NotNil(t, result)

			assert.Nil(t, err)
			assert.IsType(t, Labels{}, result)
			assert.Equal(t, 1, len(result))

			assert.Equal(t, "chameleon", result[0].Name)

			assert.Equal(t, 100-93, result[0].Uncertainty)
		}
	})
	t.Run("dog_orange.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		tensorFlow := New(conf.ResourcesPath(), conf.TensorFlowDisabled())

		if imageBuffer, err := ioutil.ReadFile(conf.ExamplesPath() + "/dog_orange.jpg"); err != nil {
			t.Error(err)
		} else {
			result, err := tensorFlow.Labels(imageBuffer)

			t.Log(result)

			assert.NotNil(t, result)

			assert.Nil(t, err)
			assert.IsType(t, Labels{}, result)
			assert.Equal(t, 1, len(result))

			assert.Equal(t, "dog", result[0].Name)

			assert.Equal(t, 34, result[0].Uncertainty)
		}
	})
	t.Run("Random.docx", func(t *testing.T) {
		conf := config.TestConfig()

		tensorFlow := New(conf.ResourcesPath(), conf.TensorFlowDisabled())

		if imageBuffer, err := ioutil.ReadFile(conf.ExamplesPath() + "/Random.docx"); err != nil {
			t.Error(err)
		} else {
			result, err := tensorFlow.Labels(imageBuffer)
			assert.Empty(t, result)
			assert.Contains(t, err.Error(), "invalid image")
		}
	})
	t.Run("6720px_white.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		tensorFlow := New(conf.ResourcesPath(), conf.TensorFlowDisabled())

		if imageBuffer, err := ioutil.ReadFile(conf.ExamplesPath() + "/6720px_white.jpg"); err != nil {
			t.Error(err)
		} else {
			result, err := tensorFlow.Labels(imageBuffer)
			assert.Empty(t, result)
			assert.Nil(t, err)
		}
	})
}

func TestTensorFlow_LoadModel(t *testing.T) {
	t.Run("model path exists", func(t *testing.T) {
		conf := config.TestConfig()

		tensorFlow := New(conf.ResourcesPath(), conf.TensorFlowDisabled())

		result := tensorFlow.loadModel()
		assert.Nil(t, result)
	})
	t.Run("model path does not exist", func(t *testing.T) {
		conf := config.NewTestErrorConfig()

		tensorFlow := New(conf.ResourcesPath(), conf.TensorFlowDisabled())

		result := tensorFlow.loadModel()
		assert.Contains(t, result.Error(), "Could not find SavedModel")
	})
}

func TestTensorFlow_BestLabels(t *testing.T) {
	t.Run("labels not loaded", func(t *testing.T) {
		conf := config.TestConfig()

		tensorFlow := New(conf.ResourcesPath(), conf.TensorFlowDisabled())

		p := make([]float32, 1000)

		p[666] = 0.5

		result := tensorFlow.bestLabels(p)
		assert.Empty(t, result)
	})
	t.Run("labels loaded", func(t *testing.T) {
		conf := config.TestConfig()
		path := conf.TensorFlowModelPath()
		tensorFlow := New(conf.ResourcesPath(), conf.TensorFlowDisabled())
		tensorFlow.loadLabels(path)

		p := make([]float32, 1000)

		p[8] = 0.7
		p[1] = 0.5

		result := tensorFlow.bestLabels(p)
		assert.Equal(t, "chicken", result[0].Name)
		assert.Equal(t, "bird", result[0].Categories[0])
		assert.Equal(t, "animal", result[1].Categories[1])
		assert.Equal(t, "image", result[0].Source)
		assert.Equal(t, "fish", result[1].Name)
		assert.Equal(t, "image", result[1].Source)
		t.Log(result)
	})
}

func TestTensorFlow_MakeTensor(t *testing.T) {
	t.Run("cat_brown.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		tensorFlow := New(conf.ResourcesPath(), conf.TensorFlowDisabled())

		imageBuffer, err := ioutil.ReadFile(conf.ExamplesPath() + "/cat_brown.jpg")
		assert.Nil(t, err)
		result, err := tensorFlow.makeTensor(imageBuffer, "jpeg")
		assert.Equal(t, tensorflow.DataType(0x1), result.DataType())
		assert.Equal(t, int64(1), result.Shape()[0])
		assert.Equal(t, int64(224), result.Shape()[2])
	})
	t.Run("Random.docx", func(t *testing.T) {
		conf := config.TestConfig()

		tensorFlow := New(conf.ResourcesPath(), conf.TensorFlowDisabled())

		imageBuffer, err := ioutil.ReadFile(conf.ExamplesPath() + "/Random.docx")
		assert.Nil(t, err)
		result, err := tensorFlow.makeTensor(imageBuffer, "jpeg")
		assert.Empty(t, result)
		assert.Equal(t, "image: unknown format", err.Error())
	})
}

func Test_ConvertTF(t *testing.T) {
	result := convertTF(uint32(98765432))
	assert.Equal(t, float32(3024.898), result)
}
