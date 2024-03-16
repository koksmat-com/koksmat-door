// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Run Caddy
---
*/
package cmds

import (
	"context"

	"github.com/365admin/koksmat-door/execution"
)

func MagicCaddyPost(ctx context.Context, args []string) (*string, error) {

	_, pwsherr := execution.ExecutePowerShell("john", "*", "koksmat-door", "00-magic", "10-runcaddy.ps1", "")
	if pwsherr != nil {
		return nil, pwsherr
	}
	return nil, nil

}
