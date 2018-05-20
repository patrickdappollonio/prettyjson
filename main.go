package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// grab all arguments except the app name
	args := os.Args[1:]

	var buf bytes.Buffer

	switch len(args) {
	case 0:
		// if we have zero arguments, we read from stdin
		// into buf
		if _, err := io.Copy(&buf, os.Stdin); err != nil {
			errexit("unable to read from stdin: %s", err.Error())
			return
		}

	case 1:
		// with one arguments, a file was passed
		// make the path absolute and then read the file into
		// buf
		fullpath, err := filepath.Abs(args[0])
		if err != nil {
			errexit("unable to get full path for %q: %s", args[0], err.Error())
			return
		}

		// open the file and check if there was an error
		f, err := os.Open(fullpath)
		if err != nil {
			// ask whether or not the error was that file doesn't exist
			if os.IsNotExist(err) {
				errexit("file %q does not exist", args[0])
				return
			}

			errexit("unable to read %q: %s", fullpath, err.Error())
			return
		}

		if _, err := io.Copy(&buf, f); err != nil {
			errexit("unable to read %q: %s", fullpath, err.Error())
			return
		}

	default:
		errexit("input must be sent via stdin or by passing a file name")
		return
	}

	fmt.Print(buf.String())
}

func errexit(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(-1)
}
