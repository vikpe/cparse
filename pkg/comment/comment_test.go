package comment_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vikpe/cparse/pkg/comment"
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
