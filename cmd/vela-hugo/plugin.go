// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

var appFS = afero.NewOsFs()

type Plugin struct {
	// build arguments loaded for the plugin
	Build *Build
	// config arguments loaded for the plugin
	Config *Config
	// theme arguments loaded for the plugin
	Theme *Theme
}

func (p *Plugin) Command() *exec.Cmd {
	// variable to store flags for command
	var flags []string

	// check if a base url is provided
	if len(p.Build.BaseURL) > 0 {
		// add flag for the provided base url
		flags = append(flags, fmt.Sprintf("--baseURL=%s", p.Build.BaseURL))
	}

	// check if draft content should be included
	if p.Build.Draft {
		// add the flag for including draft content
		flags = append(flags, "--buildDrafts")
	}

	// check if expired content should be included
	if p.Build.Expired {
		// add the flag for including expired content
		flags = append(flags, "--buildExpired")
	}

	// check if future content should be included
	if p.Build.Future {
		// add the flag for including future content
		flags = append(flags, "--buildFuture")
	}

	// check if a cache directory is provided
	if len(p.Config.CacheDirectory) > 0 {
		// add flag for the provided cache directory
		flags = append(flags, fmt.Sprintf("--cacheDir=%s", p.Config.CacheDirectory))
	}

	// check if a config file is provided
	if len(p.Config.File) > 0 {
		// add flag for the provided config
		flags = append(flags, fmt.Sprintf("--config=%s", p.Config.File))
	}

	// check if a config directory is provided
	if len(p.Config.Directory) > 0 {
		// add flag for the provided config directory
		flags = append(flags, fmt.Sprintf("--configDir=%s", p.Config.Directory))
	}

	// check if a content directory is provided
	if len(p.Config.ContentDirectory) > 0 {
		// add flag for the provided content directory
		flags = append(flags, fmt.Sprintf("--contentDir=%s", p.Config.ContentDirectory))
	}

	// check if an environment is provided
	if len(p.Config.Environment) > 0 {
		// add flag for the provided Environment
		flags = append(flags, fmt.Sprintf("--environment=%s", p.Config.Environment))
	}

	// check if a layout directory is provided
	if len(p.Config.LayoutDirectory) > 0 {
		// add flag for the provided layout directory
		flags = append(flags, fmt.Sprintf("--layoutDir=%s", p.Config.LayoutDirectory))
	}

	// check if a output directory is provided
	if len(p.Config.OutputDirectory) > 0 {
		// add flag for the provided output directory
		flags = append(flags, fmt.Sprintf("--destination=%s", p.Config.OutputDirectory))
	}

	// check if a source directory is provided
	if len(p.Config.SourceDirectory) > 0 {
		// add flag for the provided source directory
		flags = append(flags, fmt.Sprintf("--source=%s", p.Config.SourceDirectory))
	}

	// check if a theme is provided
	if len(p.Theme.Name) > 0 {
		// add flag for the provided theme
		flags = append(flags, fmt.Sprintf("--theme=%s", p.Theme.Name))
	}

	// check if a theme directory is provided
	if len(p.Theme.Directory) > 0 {
		// add flag for the provided theme directory
		flags = append(flags, fmt.Sprintf("--themesDir=%s", p.Theme.Directory))
	}

	// run the hugo plugin with the provided flags
	return exec.Command(_hugo, flags...)
}

// Exec formats and runs the commands for the plugin.
func (p *Plugin) Exec() error {
	logrus.Debug("running plugin with provided configuration")

	// output hugo version for troubleshooting
	err := execCmd(versionCmd())
	if err != nil {
		return err
	}

	// run the hugo plugin with the provided flags
	err = execCmd(p.Command())
	if err != nil {
		return err
	}

	return nil
}

func (p *Plugin) Validate() error {
	logrus.Debug("validating plugin configuration")

	// validate config configuration
	err := p.Config.Validate()
	if err != nil {
		return err
	}

	// validate theme configuration
	err = p.Theme.Validate()
	if err != nil {
		return err
	}

	return nil
}
