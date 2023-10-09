package common

import "fmt"

type Question struct {
	Text string
	Tags []string
}

func (q Question) String() string {
	return fmt.Sprintf("Text: '%s'\nTags: %s", q.Text, q.Tags)
}

