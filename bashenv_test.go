package bashenv

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestSource(t *testing.T) {
	fi, err := ioutil.TempFile("", "bashenv")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(fi.Name())

	fi.Write([]byte(`
        export STR="Hello World!"
        `))

	fi.Sync()
	fi.Close()

	err = Source(fi.Name())
	if err != nil {
		t.Fatal(err)
	}

	if os.Getenv("STR") != "Hello World!" {
		t.Error("env doesn't match")
	}
}
