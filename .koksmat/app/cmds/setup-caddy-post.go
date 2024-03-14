// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Caddy
---
*/
package cmds

import (
	"context"

	"github.com/365admin/koksmat-door/execution"
)

func SetupCaddyPost(ctx context.Context, args []string) (*string, error) {

	_, pwsherr := execution.ExecutePowerShell("john", "*", "koksmat-door", "10-setup", "caddy.ps1", "")
	if pwsherr != nil {
		return nil, pwsherr
	}
	return nil, nil

}