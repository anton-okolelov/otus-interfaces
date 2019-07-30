package logging

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccepted(t *testing.T) {
	var b bytes.Buffer
	event := HomeworkAccepted{Id: 100, Grade: 5}
	_ = LogOtusEvent(event, &b)
	assert.Equal(t, time.Now().Format("02.01.2006")+" accepted 100 5\n", b.String())
}

func TestSubmitted(t *testing.T) {
	var b bytes.Buffer
	event := HomeworkSubmitted{Id: 105, Code: "github.com", Comment: "done"}
	_ = LogOtusEvent(event, &b)
	assert.Equal(t, time.Now().Format("02.01.2006")+" submitted 105 github.com done\n", b.String())
}
