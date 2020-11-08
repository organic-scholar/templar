package rendering

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/organic-scholar/templar/common"
	"github.com/spf13/afero"
)

func init() {
	os.Setenv("GO_ENV", "testing")
}

func TestItShouldRenderFile(t *testing.T) {
	loadFixtures(t)
	data, err := ParseTemplateFile("fixtures")
	if err != nil {
		t.Fatal(err)
	}
	err = RenderTemplateFile("fixtures/source.md", data)
	if err != nil {
		t.Fatal(err)
	}
	file, err := afero.ReadFile(common.GetFs(), "fixtures/source.md")
	if err != nil {
		t.Fatal(err)
	}
	var content = string(file)
	var contains = strings.Contains(content, data.Parameters["name"])
	if contains == false {
		t.Fail()
	}
}

func loadFixtures(t *testing.T) {
	dir := "fixtures/"
	fs := common.GetFs()
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Error(err)
	}
	for _, file := range files {
		if file.IsDir() == false {
			content, _ := ioutil.ReadFile(dir + file.Name())
			err = afero.WriteFile(fs, dir+file.Name(), content, 0644)
			if err != nil {
				t.Error(err)
			}
		}
	}

}
