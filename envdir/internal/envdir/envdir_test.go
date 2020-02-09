package envdir

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Helper() (string, func()) {
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}
	return dir, func() {
		defer os.RemoveAll(dir)
	}
}

func TestReadDir(t *testing.T) {
	dir, cleanup := Helper()
	defer cleanup()
	tempFile, err := os.Create(dir + "/" + "TEST")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	tempFile.WriteString("TOST")
	envs, _ := ReadDir(dir)
	require.Equal(t, "TOST", envs["TEST"])
}
func TestReadDirEmptyFile(t *testing.T) {
	dir, cleanup := Helper()
	defer cleanup()
	tempFile, err := os.Create(dir + "/" + "TEST")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	envs, _ := ReadDir(dir)
	require.Equal(t, "", envs["TEST"])
}

func TestReadDirWithEqual(t *testing.T) {
	dir, cleanup := Helper()
	defer cleanup()
	tempFile, err := os.Create(dir + "/" + "TEST=")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	tempFile.WriteString("TOST")
	envs, _ := ReadDir(dir)
	require.Equal(t, 0, len(envs))
}

func TestReadDirWithEqual2(t *testing.T) {
	dir, cleanup := Helper()
	defer cleanup()
	tempFile, err := os.Create(dir + "/" + "TEST")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	tempFile.WriteString("TOST=")
	envs, _ := ReadDir(dir)
	require.Equal(t, 0, len(envs))
}

func TestReadDir_WithSpav(t *testing.T) {
	dir, cleanup := Helper()
	defer cleanup()
	tempFile, err := os.Create(dir + "/" + "TEST")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	tempFile.WriteString(" TOST ")
	envs, _ := ReadDir(dir)
	require.Equal(t, "TOST", envs["TEST"])
}

func TestRunCMD(t *testing.T) {
	r, w, _ := os.Pipe()
	os.Stdout = w
	RunCmd([]string{"env"}, map[string]string{"TEST_ARG": "TEST_VAL"})
	w.Close()
	out, _ := ioutil.ReadAll(r)
	require.Contains(t, string(out), "TEST_ARG=TEST_VAL")
}

func TestRunCMDEmpty(t *testing.T) {
	r, w, _ := os.Pipe()
	os.Stdout = w
	RunCmd([]string{"env"}, map[string]string{"HOME": ""})
	w.Close()
	out, _ := ioutil.ReadAll(r)
	require.Equal(t, false, strings.Contains(string(out), "HOME"))
}
