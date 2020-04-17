package cmd

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"text/template"
	"github.com/spf13/cobra"
)

var versionTemplate = `Version:           {{.AppVersion}}
Go version:        {{.GoVersion}}
Git commit:        {{.GitCommit}}
Built:             {{.BuildTime}}
OS/Arch:           {{.GoOs}}/{{.GoArch}}`

var (
	AppVersion = "unknown-app-version"
	GitCommit  = "unknown-git-commit"
	BuildTime  = "unknown-build-time"
)

type Version struct {
	AppVersion string
	GoVersion  string
	GoOs       string
	GoArch     string
	GitCommit  string
	BuildTime  string
}

func NewVersionCommand() *cobra.Command {
	versionCommand := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of wait4x",
		Long:  "All software has versions. It's mine.",
		Run: runVersion,
	}

	return versionCommand
}

func runVersion(cmd *cobra.Command, args []string) {
	versionValues := Version{
		AppVersion: AppVersion,
		GoVersion: runtime.Version(),
		GoOs: runtime.GOOS,
		GoArch: runtime.GOARCH,
		GitCommit: GitCommit,
		BuildTime: BuildTime,
	}
	var tmplBytes bytes.Buffer

	t := template.Must(template.New("version").Parse(versionTemplate))
	err := t.Execute(&tmplBytes, versionValues)
	if err != nil {
		log.Println("executing template:", err)
	}

	fmt.Println(tmplBytes.String())
}
