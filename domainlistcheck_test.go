package zonethreatreport

import (
	"bufio"
	"bytes"
	assert "github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDomainListCheckNoMatch(t *testing.T) {
	assert := assert.New(t)

	var reportdata bytes.Buffer
	reportlog := bufio.NewWriter(&reportdata)

	var errordata bytes.Buffer
	errorlog := bufio.NewWriter(&errordata)

	alerts := DomainListCheck(strings.NewReader("example.com\nexample.org"), "hiv", errorlog, reportlog)
	assert.Equal(0, alerts)
	errorlog.Flush()
	assert.False(strings.Contains(errordata.String(), "example.hiv"))
}

func TestDomainListCheckMatch(t *testing.T) {
	assert := assert.New(t)

	var reportdata bytes.Buffer
	reportlog := bufio.NewWriter(&reportdata)

	var errordata bytes.Buffer
	errorlog := bufio.NewWriter(&errordata)

	alerts := DomainListCheck(strings.NewReader("example.com\nexample.hiv\nexample.org"), "hiv", errorlog, reportlog)
	assert.Equal(1, alerts)
	errorlog.Flush()
	assert.True(strings.Contains(errordata.String(), "example.hiv"))
}
