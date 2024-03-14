// -------------------------------------------------------------------
// Generated by 365admin-publish/api/20 makeschema.ps1
// -------------------------------------------------------------------
/*
---
title: Run Caddy
---
*/
package endpoints

import (
	"context"

	"github.com/365admin/koksmat-door/execution"
	"github.com/swaggest/usecase"
)

func MagicCaddyPost() usecase.Interactor {
	type Request struct {
	}
	u := usecase.NewInteractor(func(ctx context.Context, input Request, output *string) error {

		_, err := execution.ExecutePowerShell("john", "*", "koksmat-door", "00-magic", "10-first.ps1", "")
		if err != nil {
			return err
		}

		return err

	})
	u.SetTitle("Run Caddy")
	// u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Magic Buttons")
	return u
}
