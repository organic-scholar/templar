package rendering

import (
	"encoding/json"

	"github.com/cbroglie/mustache"
	"github.com/organic-scholar/templar/common"
	"github.com/spf13/afero"
)

type TemplateData struct {
	Parameters map[string]string
	Files      []string
}

// ParseTemplateFile : it is parsing template.json file
func ParseTemplateFile(dir string) (*TemplateData, error) {
	fs := common.GetFs()
	content, err := afero.ReadFile(fs, dir+"/template.json")
	if err != nil {
		return nil, err
	}
	var data = TemplateData{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func RenderTemplateFile(filePath string, data *TemplateData) error {
	fs := common.GetFs()
	content, err := afero.ReadFile(fs, filePath)
	if err != nil {
		return err
	}
	source, err := mustache.Render(string(content), data.Parameters)
	if err != nil {
		return err
	}
	err = afero.WriteFile(fs, filePath, []byte(source), 0644)
	if err != nil {
		return err
	}
	return nil
}
