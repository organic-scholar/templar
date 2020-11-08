package main

import (
	"errors"
	"os"
	"path"

	"github.com/organic-scholar/templar/command"
	"github.com/organic-scholar/templar/common"
	"github.com/organic-scholar/templar/rendering"
	"github.com/spf13/cobra"
)

var version = "1.0.0"

func main() {
	var rootCmd = &cobra.Command{
		Use:           "templar",
		Short:         "Templar is a project scaffolding tool",
		Version:       version,
		SilenceUsage:  true,
		SilenceErrors: true,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("repository source is required")
			}
			if !common.Pattern.Match([]byte(args[0])) {
				return errors.New("unable to parse source" + args[0])
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			out, err := common.CloneRepo(args)
			if err != nil {
				return nil
			}
			data, err := rendering.ParseTemplateFile(out)
			if err != nil {
				return err
			}
			err = command.PromptUserParameters(data)
			if err != nil {
				return err
			}
			for _, file := range data.Files {
				file = path.Join(out, file)
				err := rendering.RenderTemplateFile(file, data)
				if err != nil {
					return err
				}
			}
			fs := common.GetFs()
			err = fs.Remove(path.Join(out, "template.json"))
			if err != nil {
				return err
			}
			err = fs.RemoveAll(path.Join(out, ".git"))
			if err != nil {
				return err
			}
			return nil
		},
	}
	err := rootCmd.Execute()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
