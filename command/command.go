package command

import (
	"errors"

	"github.com/manifoldco/promptui"
	"github.com/organic-scholar/templar/rendering"
)

func PromptUserParameters(data *rendering.TemplateData) error {
	for key, value := range data.Parameters {
		prompt := promptui.Prompt{
			Label:   key,
			Default: value,
			Validate: func(s string) error {
				if len(s) != 0 {
					return nil
				}
				return errors.New("This is required")
			},
		}
		input, err := prompt.Run()
		if err != nil {
			return err
		}
		println(input)
		data.Parameters[key] = input
	}
	return nil
}
