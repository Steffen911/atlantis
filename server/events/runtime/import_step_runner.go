package runtime

import (
	"github.com/hashicorp/go-version"
	"github.com/runatlantis/atlantis/server/events/models"
)

type ImportStepRunner struct {
	TerraformExecutor   TerraformExec
	DefaultTFVersion    *version.Version
	CommitStatusUpdater StatusUpdater
	AsyncTFExec         AsyncTFExec
}

func (i *ImportStepRunner) Run(ctx models.ProjectCommandContext, extraArgs []string, path string, envs map[string]string) (string, error) {
	tfVersion := i.DefaultTFVersion
	if ctx.TerraformVersion != nil {
		tfVersion = ctx.TerraformVersion
	}

	var importCmd []string
	output, err := i.TerraformExecutor.RunCommandWithVersion(ctx.Log, path, importCmd, envs, tfVersion, ctx.Workspace)

	if err != nil {
		return output, err
	}
	return output, nil
}
