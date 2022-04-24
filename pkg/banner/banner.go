package banner

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

func Default(filepath string, data map[string]interface{}) {
	err := ShowFromFile(os.Stdout, filepath, data)
	if err != nil {
		log.Println(err)
	}
}

// ShowFromFile - load the banner from file and prints it to output
func ShowFromFile(out io.Writer, filepath string, data map[string]interface{}) (err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("error trying to open banner file: %v", err)
	}
	defer func() {
		if err == nil {
			err = file.Close()
		}
	}()

	return Show(out, file, data)
}

// Show - load the banner and prints it to output
func Show(out io.Writer, in io.Reader, data map[string]interface{}) error {
	if in == nil {
		return fmt.Errorf("the input is nil")
	}

	banner, err := ioutil.ReadAll(in)
	if err != nil {
		return fmt.Errorf("error trying to read the banner, err: %v", err)
	}

	t, err := template.New("banner").Parse(string(banner))
	if err != nil {
		return fmt.Errorf("error trying to parse the banner file, err: %v", err)
	}

	data["ansiColor"] = ansiColor{true}
	data["ansiBackground"] = ansiBackground{true}

	err = t.Execute(out, data)
	if err != nil {
		return fmt.Errorf("error trying to execute template: %v", err)
	}

	return nil
}
