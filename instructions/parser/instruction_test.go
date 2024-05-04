package parser

import (
	"fmt"
	"regexp"
	"testing"
)

func TestParseInstruction(t *testing.T) {
	r, _ := regexp.Compile(`(\S*) ([^\s]*) (.*)`)

	m := r.FindStringSubmatch("ASD @xy #aa aa")

	fmt.Println(m[1], "-", m[2], "-", m[3])
}
