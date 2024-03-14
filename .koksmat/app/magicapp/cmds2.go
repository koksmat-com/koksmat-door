package magicapp

import (
	"github.com/365admin/koksmat-door/cmds"
	"github.com/spf13/cobra"
)

func RegisterCmds() {
	magicCmd := &cobra.Command{
		Use:   "magic",
		Short: "Magic Buttons",
		Long:  `Entry point to the kitchens`,
	}
	MagicCaddyPostCmd := &cobra.Command{
		Use:   "caddy",
		Short: "Run Caddy",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.MagicCaddyPost(ctx, args)
		},
	}
	magicCmd.AddCommand(MagicCaddyPostCmd)

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
	learnCmd := &cobra.Command{
		Use:   "learn",
		Short: "Learn",
		Long:  `Entry point to the kitchens`,
	}

	RootCmd.AddCommand(learnCmd)
	provisionCmd := &cobra.Command{
		Use:   "provision",
		Short: "Provision",
		Long:  `Entry point to the kitchens`,
	}

	RootCmd.AddCommand(provisionCmd)
	decommissionCmd := &cobra.Command{
		Use:   "decommission",
		Short: "Decommision",
		Long:  `Entry point to the kitchens`,
	}

	RootCmd.AddCommand(decommissionCmd)
}
