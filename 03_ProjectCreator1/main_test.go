package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadFile(t *testing.T) {
	content, errRe := readFile(_filePath)
	require.NoError(t, errRe)

	fmt.Println(strings.Join(content, "\n"))
}

func TestFile(t *testing.T) {
	content, errRe := readFile(_filePath)
	require.NoError(t, errRe)

	entries := parse(convertToEntries(content))

	fmt.Println(strings.Join(entries, "\n"))
}

func TestTypeofFile(t *testing.T) {
	got := []string{"! .", "# package main", "file.go", "folder"}
	want := []string{"path", "package", "file", "folder"}

	for i, _ := range got {
		assert.Equal(t, typeofFile(got[i]), want[i], "verify the type of file")
		fmt.Println(got[i], "is a", want[i])
	}
}
