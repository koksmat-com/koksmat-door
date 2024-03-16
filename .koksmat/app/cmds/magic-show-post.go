// -------------------------------------------------------------------
// Generated by 365admin-publish
// -------------------------------------------------------------------
/*
---
title: Show Caddyfile
---
*/
package cmds

import (
	"context"

	"github.com/365admin/koksmat-door/execution"
)

func MagicShowPost(ctx context.Context, args []string) (*string, error) {

	_, pwsherr := execution.ExecutePowerShell("john", "*", "koksmat-door", "00-magic", "20-showconfig.ps1", "")
	if pwsherr != nil {
		return nil, pwsherr
	}
	return nil, nil

}
