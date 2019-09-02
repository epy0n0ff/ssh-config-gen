package cmd

import (
	"fmt"
	"os"

	"github.com/epy0n0ff/ssh-config-gen/sshconfig"
	"github.com/epy0n0ff/ssh-config-gen/util"
	"github.com/spf13/cobra"
)

func NewGenCommand() *cobra.Command {
	var tomlPath string
	var templatePath string
	var err error

	cmd := &cobra.Command{
		Use:   "gen",
		Short: "A brief description of your application",
		Run: func(cmd *cobra.Command, args []string) {
			g := &gen{
				tomlPath:     tomlPath,
				templatePath: templatePath,
			}

			err = util.ValidateGenCommandArguments(cmd.Flags())
			util.CheckErr(err)
			util.CheckErr(g.Run())
		},
	}

	cmd.PersistentFlags().StringVar(&tomlPath, "toml", "", "hosts.toml file path")
	cmd.PersistentFlags().StringVar(&templatePath, "tpl", "", "ssh configuration template file path")

	return cmd
}

type gen struct {
	tomlPath     string
	templatePath string
}

func (m *gen) Run() error {
	tml, err := sshconfig.LoadToml(m.tomlPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load toml file failed: %v", err)
	}
	config, err := sshconfig.GetSshConfig(m.templatePath, tml)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load template file failed: %v", err)
	}

	fmt.Fprintf(os.Stdout, "%s", config)
	return nil
}
