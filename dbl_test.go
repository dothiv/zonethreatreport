package zonethreatreport

import (
	"bufio"
	"bytes"
	assert "github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type MockCommandRunner struct {
	out []byte
}

func (r MockCommandRunner) Run(command string, args ...string) ([]byte, error) {
	return r.out, nil
}

func TestDblLookupNoMatch(t *testing.T) {
	assert := assert.New(t)

	var reportdata bytes.Buffer
	reportlog := bufio.NewWriter(&reportdata)

	var errordata bytes.Buffer
	errorlog := bufio.NewWriter(&errordata)

	alerts := Dbl(strings.NewReader("example.hiv"), "test.dbl.com", MockCommandRunner{[]byte("")}, errorlog, reportlog)
	assert.Equal(0, alerts)
	errorlog.Flush()
	assert.False(strings.Contains(errordata.String(), "example.hiv"))
}

func TestDblLookupMatch(t *testing.T) {
	assert := assert.New(t)

	var reportdata bytes.Buffer
	reportlog := bufio.NewWriter(&reportdata)

	var errordata bytes.Buffer
	errorlog := bufio.NewWriter(&errordata)

	alerts := Dbl(strings.NewReader("example.hiv"), "test.dbl.com", MockCommandRunner{[]byte("127.0.1.2")}, errorlog, reportlog)
	assert.Equal(1, alerts)
	errorlog.Flush()
	assert.True(strings.Contains(errordata.String(), "example.hiv"))
}
