package config

import (
	"strings"
	"testing"

	"github.com/mikepadge/photoprism/pkg/fs"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	ctx := CliTestContext()

	assert.True(t, ctx.IsSet("assets-path"))
	assert.False(t, ctx.Bool("debug"))

	c := NewConfig(ctx)

	assert.IsType(t, new(Config), c)

	assert.Equal(t, fs.Abs("../../assets"), c.AssetsPath())
	assert.False(t, c.Debug())
	assert.False(t, c.ReadOnly())
}

func TestConfig_Name(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	name := c.Name()
	assert.Equal(t, "config.test", name)
}

func TestConfig_Version(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	version := c.Version()
	assert.Equal(t, "0.0.0", version)
}

func TestConfig_TensorFlowVersion(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	version := c.TensorFlowVersion()
	assert.Equal(t, "1.14.0", version)
}

func TestConfig_TensorFlowDisabled(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	version := c.TensorFlowDisabled()
	assert.Equal(t, false, version)
}

func TestConfig_Copyright(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	copyright := c.Copyright()
	assert.Equal(t, "", copyright)
}

func TestConfig_ConfigFile(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	configFile := c.ConfigFile()
	assert.Equal(t, "", configFile)
}

func TestConfig_ConfigPath(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	configPath := c.ConfigPath()
	assert.Equal(t, "/go/src/github.com/mikepadge/photoprism/assets/config", configPath)
}

func TestConfig_PIDFilename(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	filename := c.PIDFilename()
	assert.Equal(t, "/go/src/github.com/mikepadge/photoprism/assets/photoprism.pid", filename)
}

func TestConfig_LogFilename(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	filename := c.LogFilename()
	assert.Equal(t, "/go/src/github.com/mikepadge/photoprism/assets/photoprism.log", filename)
}

func TestConfig_DetachServer(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	detachServer := c.DetachServer()
	assert.Equal(t, false, detachServer)
}

func TestConfig_SqlServerHost(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	host := c.SqlServerHost()
	assert.Equal(t, "127.0.0.1", host)
}

func TestConfig_SqlServerPort(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	port := c.SqlServerPort()
	assert.Equal(t, uint(4000), port)
}

func TestConfig_SqlServerPath(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	path := c.SqlServerPath()
	assert.Equal(t, "/go/src/github.com/mikepadge/photoprism/assets/resources/database", path)
}

func TestConfig_SqlServerPassword(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	password := c.SqlServerPassword()
	assert.Equal(t, "", password)
}

func TestConfig_HttpServerHost(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	host := c.HttpServerHost()
	assert.Equal(t, "0.0.0.0", host)
}

func TestConfig_HttpServerPort(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	port := c.HttpServerPort()
	assert.Equal(t, 2342, port)
}

func TestConfig_HttpServerMode(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	mode := c.HttpServerMode()
	assert.Equal(t, "release", mode)
}

func TestConfig_HttpServerPassword(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	password := c.HttpServerPassword()
	assert.Equal(t, "", password)
}

func TestConfig_OriginalsPath(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

<<<<<<< HEAD
<<<<<<< HEAD
	path := c.OriginalsPath()
	assert.Equal(t, "/go/src/github.com/mikepadge/photoprism/assets/testdata/originals", path)
=======
	result := c.OriginalsPath()
	assert.True(t, strings.HasPrefix(result, "/"))
	assert.True(t, strings.HasSuffix(result, "assets/testdata/originals"))
>>>>>>> 7cbdd31793e34cddb2c20a04d20d8ae5d25d7729
=======
	result := c.OriginalsPath()
	assert.True(t, strings.HasPrefix(result, "/"))
	assert.True(t, strings.HasSuffix(result, "assets/testdata/originals"))
>>>>>>> 5fba03844298ab501ce513a3f967b7578bc09707
}

func TestConfig_ImportPath(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	result := c.ImportPath()
	assert.True(t, strings.HasPrefix(result, "/"))
	assert.True(t, strings.HasSuffix(result, "assets/testdata/import"))
}

func TestConfig_ExportPath(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	result := c.ExportPath()
	assert.True(t, strings.HasPrefix(result, "/"))
	assert.True(t, strings.HasSuffix(result, "assets/testdata/export"))
}

func TestConfig_SipsBin(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	bin := c.SipsBin()
	assert.Equal(t, "", bin)
}

func TestConfig_DarktableBin(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	bin := c.DarktableBin()
	assert.Equal(t, "/usr/bin/darktable-cli", bin)
}

func TestConfig_HeifConvertBin(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	bin := c.HeifConvertBin()
	assert.Equal(t, "/usr/bin/heif-convert", bin)
}

func TestConfig_ExifToolBin(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	bin := c.ExifToolBin()
	assert.Equal(t, "/usr/bin/exiftool", bin)
}

func TestConfig_DatabaseDriver(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	driver := c.DatabaseDriver()
	assert.Equal(t, "internal", driver)
}

func TestConfig_DatabaseDsn(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	dsn := c.DatabaseDriver()
	assert.Equal(t, "internal", dsn)
}

func TestConfig_CachePath(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	assert.True(t, strings.HasSuffix(c.CachePath(), "assets/testdata/cache"))
}

func TestConfig_ThumbnailsPath(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	assert.True(t, strings.HasPrefix(c.ThumbnailsPath(), "/"))
	assert.True(t, strings.HasSuffix(c.ThumbnailsPath(), "assets/testdata/cache/thumbnails"))
}

func TestConfig_AssetsPath(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	path := c.AssetsPath()
	assert.Equal(t, "/go/src/github.com/mikepadge/photoprism/assets", path)
}

func TestConfig_ResourcesPath(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	path := c.ResourcesPath()
	assert.Equal(t, "/go/src/github.com/mikepadge/photoprism/assets/resources", path)
}

func TestConfig_DetectNSFW(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	result := c.DetectNSFW()
	assert.Equal(t, true, result)
}

func TestConfig_AdminPassword(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	result := c.AdminPassword()
	assert.Equal(t, "photoprism", result)
}

func TestConfig_NSFWModelPath(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	result := c.NSFWModelPath()
	assert.Equal(t, "/go/src/github.com/mikepadge/photoprism/assets/resources/nsfw", result)
}

func TestConfig_ExamplesPath(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	path := c.ExamplesPath()
	assert.Equal(t, "/go/src/github.com/mikepadge/photoprism/assets/resources/examples", path)
}

func TestConfig_TensorFlowModelPath(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	path := c.TensorFlowModelPath()
	assert.Equal(t, "/go/src/github.com/mikepadge/photoprism/assets/resources/nasnet", path)
}

func TestConfig_HttpTemplatesPath(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	path := c.HttpTemplatesPath()
	assert.Equal(t, "/go/src/github.com/mikepadge/photoprism/assets/resources/templates", path)
}

func TestConfig_HttpFaviconsPath(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	path := c.HttpFaviconsPath()
	assert.Equal(t, "/go/src/github.com/mikepadge/photoprism/assets/resources/static/favicons", path)
}

func TestConfig_HttpStaticPath(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	path := c.HttpStaticPath()
	assert.Equal(t, "/go/src/github.com/mikepadge/photoprism/assets/resources/static", path)
}

func TestConfig_HttpStaticBuildPath(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	path := c.HttpStaticBuildPath()
	assert.Equal(t, "/go/src/github.com/mikepadge/photoprism/assets/resources/static/build", path)
}

func TestConfig_CloseDb(t *testing.T) {
	c := NewTestConfig()

	assert.NotNil(t, c.Db())

	err := c.CloseDb()
	assert.Nil(t, err)
}

func TestConfig_ClientConfig(t *testing.T) {
	c := TestConfig()

	cc := c.ClientConfig()
	assert.NotEmpty(t, cc)
	assert.Contains(t, cc, "name")
	assert.Contains(t, cc, "version")
	assert.Contains(t, cc, "copyright")
	assert.Contains(t, cc, "debug")
	assert.Contains(t, cc, "readonly")
	assert.Contains(t, cc, "cameras")
	assert.Contains(t, cc, "countries")
	assert.Contains(t, cc, "thumbnails")
	assert.Contains(t, cc, "jsHash")
	assert.Contains(t, cc, "cssHash")
}

func TestConfig_Shutdown(t *testing.T) {
	c := NewTestConfig()
	c.Shutdown()
}

func TestConfig_Workers(t *testing.T) {
	ctx := CliTestContext()
	c := NewConfig(ctx)

	assert.GreaterOrEqual(t, c.Workers(), 1)
}
