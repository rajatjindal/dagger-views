// A generated module for ModuleXyz functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/module-xyz/internal/dagger"
)

type ModuleXyz struct {
	DarwinCLIVersion string
}

// Returns a container that echoes whatever string argument is provided
func (m *ModuleXyz) ContainerEcho(stringArg string, skipTParseXyz string) *dagger.Container {
	return dag.ModuleAbc().ContainerEcho(stringArg, skipTParseXyz)
}

// Returns a container that echoes whatever string argument is provided
func (m *ModuleXyz) GetDarwinCLIVersion() *dagger.Container {
	return dag.ModuleAbc().GetKubectlCLIVersion()
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *ModuleXyz) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}
