// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	_ "github.com/joho/godotenv/autoload"

	"github.com/go-vela/vela-hugo/version"
)

func main() {
	// capture application version information
	v := version.New()

	// serialize the version information as pretty JSON
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		logrus.Fatal(err)
	}

	// output the version information to stdout
	fmt.Fprintf(os.Stdout, "%s\n", string(bytes))

	// create new CLI application
	app := cli.NewApp()

	// Plugin Information

	app.Name = "vela-hugo"
	app.HelpName = "vela-hugo"
	app.Usage = "Vela Hugo plugin for generating a static website"
	app.Copyright = "Copyright 2021 Target Brands, Inc. All rights reserved."
	app.Authors = []*cli.Author{
		{
			Name:  "Vela Admins",
			Email: "vela@target.com",
		},
	}

	// Plugin Metadata

	app.Action = run
	app.Compiled = time.Now()
	app.Version = v.Semantic()

	// Plugin Flags

	app.Flags = []cli.Flag{

		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_LOG_LEVEL", "HUGO_LOG_LEVEL"},
			FilePath: "/vela/parameters/hugo/log_level,/vela/secrets/hugo/log_level",
			Name:     "log.level",
			Usage:    "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
			Value:    "info",
		},
		&cli.BoolFlag{
			EnvVars:  []string{"PARAMETER_EXTENDED", "HUGO_EXTENDED"},
			FilePath: "/vela/parameters/hugo/extended,/vela/secrets/hugo/extended",
			Name:     "hugo.extended",
			Usage:    "sets whether to use the extended binary or not",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_VERSION", "HUGO_VERSION"},
			FilePath: "/vela/parameters/hugo/version,/vela/secrets/hugo/version",
			Name:     "hugo.version",
			Usage:    "set hugo version for plugin",
		},

		// Build Flags

		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_BASE_URL", "HUGO_BASE_URL"},
			FilePath: "/vela/parameters/hugo/base_url,/vela/secrets/hugo/base_url",
			Name:     "build.base_url",
			Usage:    "hostname (and path) to the root, e.g. http://spf13.com/",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_DRAFT", "HUGO_DRAFT"},
			FilePath: "/vela/parameters/hugo/draft,/vela/secrets/hugo/draft",
			Name:     "build.draft",
			Usage:    "include content marked as draft",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_EXPIRED", "HUGO_EXPIRED"},
			FilePath: "/vela/parameters/hugo/expired,/vela/secrets/hugo/expired",
			Name:     "build.expired",
			Usage:    "include expired content",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_FUTURE", "HUGO_FUTURE"},
			FilePath: "/vela/parameters/hugo/future,/vela/secrets/hugo/future",
			Name:     "build.future",
			Usage:    "include content with publishdate in the future",
		},

		// Config Flags

		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_CACHE_DIRECTORY", "HUGO_CACHE_DIRECTORY"},
			FilePath: "/vela/parameters/hugo/cache_directory,/vela/secrets/hugo/cache_directory",
			Name:     "config.cache_directory",
			Usage:    "filesystem path to cache directory",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_CONTENT_DIRECTORY", "HUGO_CONTENT_DIRECTORY"},
			FilePath: "/vela/parameters/hugo/content_directory,/vela/secrets/hugo/content_directory",
			Name:     "config.content_directory",
			Usage:    "filesystem path to content directory",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_CONFIG_DIRECTORY", "HUGO_CONFIG_DIRECTORY"},
			FilePath: "/vela/parameters/hugo/config_directory,/vela/secrets/hugo/config_directory",
			Name:     "config.directory",
			Usage:    "filesystem path to config directory",
			Value:    "config",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_ENVIRONMENT", "HUGO_ENVIRONMENT"},
			FilePath: "/vela/parameters/hugo/environment,/vela/secrets/hugo/environment",
			Name:     "config.environment",
			Usage:    "targeted build environment in config directory",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_CONFIG_FILE", "HUGO_CONFIG_FILE"},
			FilePath: "/vela/parameters/hugo/config_file,/vela/secrets/hugo/config_file",
			Name:     "config.file",
			Usage:    "name of config file in config directory",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_LAYOUT_DIRECTORY", "HUGO_LAYOUT_DIRECTORY"},
			FilePath: "/vela/parameters/hugo/layout_directory,/vela/secrets/hugo/layout_directory",
			Name:     "config.layout_directory",
			Usage:    "filesystem path to layout directory",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_OUTPUT_DIRECTORY", "HUGO_OUTPUT_DIRECTORY"},
			FilePath: "/vela/parameters/hugo/output_directory,/vela/secrets/hugo/output_directory",
			Name:     "config.output_directory",
			Usage:    "filesystem path to write files to",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_SOURCE_DIRECTORY", "HUGO_SOURCE_DIRECTORY"},
			FilePath: "/vela/parameters/hugo/source_directory,/vela/secrets/hugo/source_directory",
			Name:     "config.source_directory",
			Usage:    "filesystem path to read files from",
		},

		// Theme Flags

		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_THEME_NAME", "HUGO_THEME_NAME"},
			FilePath: "/vela/parameters/hugo/theme_name,/vela/secrets/hugo/theme_name",
			Name:     "theme.name",
			Usage:    "name of theme to use from theme directory",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_THEME_DIRECTORY", "HUGO_THEME_DIRECTORY"},
			FilePath: "/vela/parameters/hugo/theme_directory,/vela/secrets/hugo/theme_directory",
			Name:     "theme.directory",
			Usage:    "filesystem path to theme directory",
			Value:    "themes",
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

// run executes the plugin based off the configuration provided.
func run(c *cli.Context) error {
	// set the log level for the plugin
	switch c.String("log.level") {
	case "t", "trace", "Trace", "TRACE":
		logrus.SetLevel(logrus.TraceLevel)
	case "d", "debug", "Debug", "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "w", "warn", "Warn", "WARN":
		logrus.SetLevel(logrus.WarnLevel)
	case "e", "error", "Error", "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	case "f", "fatal", "Fatal", "FATAL":
		logrus.SetLevel(logrus.FatalLevel)
	case "p", "panic", "Panic", "PANIC":
		logrus.SetLevel(logrus.PanicLevel)
	case "i", "info", "Info", "INFO":
		fallthrough
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.WithFields(logrus.Fields{
		"code":     "https://github.com/go-vela/vela-hugo",
		"docs":     "https://go-vela.github.io/docs/plugins/registry/pipeline/hugo",
		"registry": "https://hub.docker.com/r/target/vela-hugo",
	}).Info("Vela Hugo Plugin")

	// capture extended binary configuration
	//
	// extended binary includes more features and functionality
	extended := c.Bool("hugo.extended")
	// capture custom hugo version requested
	version := c.String("hugo.version")

	// check if we should fetch extended binary or custom hugo version
	if len(version) > 0 || extended {
		// attempt to install the custom hugo version
		err := install(extended, version, os.Getenv("PLUGIN_HUGO_VERSION"))
		if err != nil {
			return err
		}
	}

	// create the plugin
	p := &Plugin{
		Build: &Build{
			BaseURL: c.String("build.base_url"),
			Draft:   c.Bool("build.draft"),
			Expired: c.Bool("build.expired"),
			Future:  c.Bool("build.future"),
		},
		Config: &Config{
			CacheDirectory:   c.String("config.cache_directory"),
			ContentDirectory: c.String("config.content_directory"),
			Directory:        c.String("config.directory"),
			Environment:      c.String("config.environment"),
			File:             c.String("config.file"),
			LayoutDirectory:  c.String("config.layout_directory"),
			OutputDirectory:  c.String("config.output_directory"),
			SourceDirectory:  c.String("config.source_directory"),
		},
		Theme: &Theme{
			Name:      c.String("theme.name"),
			Directory: c.String("theme.directory"),
		},
	}

	// validate the plugin
	err := p.Validate()
	if err != nil {
		return err
	}

	// execute the plugin
	return p.Exec()
}
