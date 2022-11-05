package comment_test

import (
	"fmt"
	"testing"

	"cparse/pkg/comment"
	"github.com/stretchr/testify/assert"
)

func TestFromSingleLine(t *testing.T) {
	testCases := map[string]string{
		"":             "",
		"foo":          "",
		"//":           "",
		" // ":         "",
		"//foo":        "foo",
		"// foo ":      "foo",
		"// foo//bar ": "foo//bar",
	}

	for line, expect := range testCases {
		descr := fmt.Sprintf(`line "%s", expect "%s"`, line, expect)
		t.Run(descr, func(t *testing.T) {
			assert.Equal(t, expect, comment.FromSingleLine(line))
		})
	}
}
