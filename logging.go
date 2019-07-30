package logging

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"strconv"
	"strings"
	"time"
)

type HomeworkAccepted struct {
	Id    int
	Grade int
}

func (event HomeworkAccepted) GetId() int {
	return event.Id
}

func (event HomeworkAccepted) GetType() string {
	return "accepted"
}

func (event HomeworkAccepted) GetAdditionalFields() []string {
	return []string{strconv.Itoa(event.Grade)}
}

type HomeworkSubmitted struct {
	Id      int
	Code    string
	Comment string
}

func (event HomeworkSubmitted) GetId() int {
	return event.Id
}

func (event HomeworkSubmitted) GetType() string {
	return "submitted"
}

func (event HomeworkSubmitted) GetAdditionalFields() []string {
	return []string{event.Code, event.Comment}
}

type OtusEvent interface {
	GetId() int
	GetType() string
	GetAdditionalFields() []string
}

func LogOtusEvent(e OtusEvent, w io.Writer) error {
	_, err := fmt.Fprintf(
		w,
		"%v %v %v %s\n",
		time.Now().Format("02.01.2006"),
		e.GetType(),
		e.GetId(),
		strings.Join(e.GetAdditionalFields(), " "),
	)
	return errors.Wrap(err, "Couldnt write into writer")
}
