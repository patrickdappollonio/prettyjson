package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// grab all arguments except the app name
	args := os.Args[1:]
	fileName := ""

	var buf bytes.Buffer

	switch len(args) {
	case 0:
		// if we have zero arguments, we read from stdin
		// into buf
		if _, err := io.Copy(&buf, os.Stdin); err != nil {
			errexit("unable to read from stdin: %s", err.Error())
			return
		}

		fileName = "stdin"

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

		fileName = args[0]

	default:
		errexit("input must be sent via stdin or by passing a file name")
		return
	}

	// check if the buffer has JSON
	var contents interface{}
	if err := json.NewDecoder(&buf).Decode(&contents); err != nil {
		errexit("unable to parse JSON contents from %q: %s", fileName, err.Error())
		return
	}

	// reset the buffer just in case
	// (it should've drain from the decoder above)
	buf.Reset()

	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")

	// re-encoding it will also let the JSON content be pretty-printed
	if err := enc.Encode(contents); err != nil {
		errexit("unable to re-indent JSON content: %s", err.Error())
	}

	// Write the contents to stdout
	if _, err := os.Stdout.Write(buf.Bytes()); err != nil {
		errexit("unable to write buffer contents to stdout: %s", err.Error())
	}
}

func errexit(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(-1)
}
