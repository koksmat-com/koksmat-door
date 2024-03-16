package magicapp

import (
	"github.com/365admin/koksmat-door/cmds"
	"github.com/365admin/koksmat-door/utils"
	"github.com/spf13/cobra"
)

func RegisterCmds() {
	RootCmd.PersistentFlags().StringVarP(&utils.Output, "output", "o", "", "Output format (json, yaml, xml, etc.)")

	magicCmd := &cobra.Command{
		Use:   "magic",
		Short: "Magic Buttons",
		Long:  `Entry point to the kitchens`,
	}
	MagicCaddyPostCmd := &cobra.Command{
		Use:   "caddy ",
		Short: "Run Caddy",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.MagicCaddyPost(ctx, args)
		},
	}
	magicCmd.AddCommand(MagicCaddyPostCmd)
	MagicShowPostCmd := &cobra.Command{
		Use:   "show ",
		Short: "Show Caddyfile",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.MagicShowPost(ctx, args)
		},
	}
	magicCmd.AddCommand(MagicShowPostCmd)

	RootCmd.AddCommand(magicCmd)
	setupCmd := &cobra.Command{
		Use:   "setup",
		Short: "Setup",
		Long:  `Entry point to the kitchens`,
	}

	RootCmd.AddCommand(setupCmd)
	codeCmd := &cobra.Command{
		Use:   "code",
		Short: "Code",
		Long:  `Entry point to the kitchens`,
	}

	RootCmd.AddCommand(codeCmd)
	provisionCmd := &cobra.Command{
		Use:   "provision",
		Short: "Provision",
		Long:  `Entry point to the kitchens`,
	}

	RootCmd.AddCommand(provisionCmd)
}
