package register

import (
	cmdcore "github.com/lyft/flytectl/cmd/core"
	"github.com/spf13/cobra"
)

// RegisterCommand will return register command
func RegisterCommand() *cobra.Command {
	registerCmd := &cobra.Command{
		Use:   "register",
		Short: "Registers tasks/workflows/launchplans from list of generated serialized files.",
		Long: "Takes input files as serialized versions of the tasks/workflows/launchplans and registers them with flyteadmin.\n" +
			"Currently these input files are protobuf files generated as output from flytekit serialize.\n" +
			"Project & Domain are mandatory fields to be passed for registration and an optional version which defaults to v1\n" +
			"If the entities are already registered with flyte for the same version then registration would fail.\n",
	}
	registerResourcesFuncs := map[string]cmdcore.CommandEntry{
		"files": {CmdFunc: registerFromFilesFunc, Aliases: []string{"file"}, PFlagProvider: filesConfig, Short: "Registers file resources",
			Long:  "Registers all the serialized protobuf files including tasks, workflows and launchplans.\n" +
			"bin/flytectl register file  _pb_output_new/* -d development  -p flytesnacks -v v2\n"},
	}
	cmdcore.AddCommands(registerCmd, registerResourcesFuncs)
	return registerCmd
}