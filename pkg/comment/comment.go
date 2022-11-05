package comment

import (
	"strings"
)

const SingleLinePrefix = "//"

func FromSingleLine(line string) string {
	indexCommentStart := strings.Index(line, SingleLinePrefix)
	if -1 == indexCommentStart {
		return ""
	}
	return strings.TrimSpace(line[indexCommentStart+len(SingleLinePrefix):])
}
