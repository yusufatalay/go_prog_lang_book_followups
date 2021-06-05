package clio

import (
	"strings"
	"io/ioutil"
	"os"
	"os/exec"
)

// DefaultEditor sets the default text editor as vim
const DefaultEditor = "vim"

// PrefferedEditorResolver is a function that returns an edior that the user prefers to use
// such as configured '$EDITOR' variable
type PrefferedEditorResolver func() string

// GetPrefferedEditorFromEnviroment returns the users preffered text editor if its not set then it returns the DefaultEditor which is vim
func GetPreferredEditorFromEnviroment() string {
	editor := os.Getenv("$EDITOR")

	if editor == "" {
		return DefaultEditor
	}

	return editor
}

func resolveEditorArguments(executable, filename string) []string {
	args := []string{filename}

	if strings.Contains(executable, "Visual Studio Code.app") {
		args = append([]string{"--wait"}, args...)
	}

	return args
}

// OpenFileInEditor opens filename in a text editor
func OpenFileInEditor(filename string, resolveEditor PrefferedEditorResolver) error {
	// get full path of the preffered editor
	executable, err := exec.LookPath(resolveEditor())

	if err != nil {
		return err
	}

	cmd := exec.Command(executable, resolveEditorArguments(executable, filename)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// CaptureInputFromEditor opens a temporary file with the preffered editor saves the input reads from there and deletes the temporary file
func CaptureInputFromEditor(resolveEditor PrefferedEditorResolver) ([]byte, error) {
	file, err := ioutil.TempFile(os.TempDir(), "*")

	if err != nil {
		return []byte{}, err
	}

	filename := file.Name()

	// Defer removal of the temporary file in any case of error

	defer os.Remove(filename)

	if err = file.Close(); err != nil {
		return []byte{}, err
	}

	if err = OpenFileInEditor(filename, resolveEditor); err != nil {
		return []byte{}, err
	}

	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return []byte{}, err
	}

	return bytes, nil

}
