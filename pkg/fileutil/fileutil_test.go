package fileutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vikpe/cparse/pkg/fileutil"
)

func TestGetFilePaths(t *testing.T) {
	t.Run("path does not exist", func(t *testing.T) {
		expect := []string{}
		paths, err := fileutil.GetFilePaths("__INVALID_DIR_PATH__")
		assert.Equal(t, expect, paths)
		assert.ErrorContains(t, err, "no such file or directory")
	})

	t.Run("path exists", func(t *testing.T) {
		expect := []string{
			"test_files/foo.c",
			"test_files/subdir/bar.c",
		}
		paths, err := fileutil.GetFilePaths(".")
		assert.Equal(t, expect, paths)
		assert.Nil(t, err)
	})
}
