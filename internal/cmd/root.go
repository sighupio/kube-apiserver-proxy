package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/sighupio/kube-apiserver-proxy/internal/app"
	cobrax "github.com/sighupio/kube-apiserver-proxy/internal/x/cobra"
)

type RootCommand struct {
	*cobra.Command
}

func NewRootCommand(versions map[string]string) *RootCommand {
	const envPrefix = ""

	root := &RootCommand{
		Command: &cobra.Command{
			PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
				cobrax.BindFlags(cmd, cobrax.InitEnvs(envPrefix), log.Fatal, envPrefix)

				return nil
			},
			Use:           "kube-apiserver-proxy",
			SilenceUsage:  true,
			SilenceErrors: true,
		},
	}

	cobrax.BindFlags(root.Command, cobrax.InitEnvs(envPrefix), log.Fatal, envPrefix)

	root.AddCommand(NewVersionCommand(versions))
	root.AddCommand(NewServeCommand(app.NewContainer()))

	return root
}
