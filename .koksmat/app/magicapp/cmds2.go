package magicapp

import (
	"github.com/365admin/koksmat-door/cmds"
	"github.com/spf13/cobra"
)

func RegisterCmds() {
	learnCmd := &cobra.Command{
		Use:   "learn",
		Short: "Learn",
		Long:  `Entry point to the kitchens`,
	}

	RootCmd.AddCommand(learnCmd)
	devCmd := &cobra.Command{
		Use:   "dev",
		Short: "Dev",
		Long:  `Entry point to the kitchens`,
	}
	DevCaddyPostCmd := &cobra.Command{
		Use:   "caddy",
		Short: "Caddy",
		Long:  `Caddy is a powerful, enterprise-ready, open source web server with automatic HTTPS written in Go.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.DevCaddyPost(ctx, args)
		},
	}
	devCmd.AddCommand(DevCaddyPostCmd)

	RootCmd.AddCommand(devCmd)
	customCmd := &cobra.Command{
		Use:   "custom",
		Short: "Custom Caddy",
		Long:  `Entry point to the kitchens`,
	}
	CustomStartPostCmd := &cobra.Command{
		Use:   "start",
		Short: "Start",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()

			cmds.CustomStartPost(ctx, args)
		},
	}
	customCmd.AddCommand(CustomStartPostCmd)

	RootCmd.AddCommand(customCmd)
}
