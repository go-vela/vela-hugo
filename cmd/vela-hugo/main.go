// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/mail"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v3"

	_ "github.com/joho/godotenv/autoload"

	"github.com/go-vela/vela-hugo/version"
)

//nolint:funlen // function is long due to CLI flag configuration, breaking it up provides little value
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
	app := &cli.Command{
		Name:      "vela-hugo",
		Usage:     "Vela Hugo plugin for generating a static website",
		Copyright: "Copyright 2021 Target Brands, Inc. All rights reserved.",
		Authors: []any{
			&mail.Address{
				Name:    "Vela Admins",
				Address: "vela@target.com",
			},
		},
		Version: v.Semantic(),
		Action:  run,

		// Plugin Flags
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "log.level",
				Usage: "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
				Value: "info",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_LOG_LEVEL"),
					cli.EnvVar("HUGO_LOG_LEVEL"),
					cli.File("/vela/parameters/hugo/log_level"),
					cli.File("/vela/secrets/hugo/log_level"),
				),
			},
			&cli.BoolFlag{
				Name:  "hugo.extended",
				Usage: "sets whether to use the extended binary or not",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_EXTENDED"),
					cli.EnvVar("HUGO_EXTENDED"),
					cli.File("/vela/parameters/hugo/extended"),
					cli.File("/vela/secrets/hugo/extended"),
				),
			},
			&cli.StringFlag{
				Name:  "hugo.version",
				Usage: "set hugo version for plugin",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_VERSION"),
					cli.EnvVar("HUGO_VERSION"),
					cli.File("/vela/parameters/hugo/version"),
					cli.File("/vela/secrets/hugo/version"),
				),
			},

			// Build Flags
			&cli.StringFlag{
				Name:  "build.base_url",
				Usage: "hostname (and path) to the root, e.g. http://spf13.com/",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_BASE_URL"),
					cli.EnvVar("HUGO_BASE_URL"),
					cli.File("/vela/parameters/hugo/base_url"),
					cli.File("/vela/secrets/hugo/base_url"),
				),
			},
			&cli.StringFlag{
				Name:  "build.draft",
				Usage: "include content marked as draft",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_DRAFT"),
					cli.EnvVar("HUGO_DRAFT"),
					cli.File("/vela/parameters/hugo/draft"),
					cli.File("/vela/secrets/hugo/draft"),
				),
			},
			&cli.StringFlag{
				Name:  "build.expired",
				Usage: "include expired content",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_EXPIRED"),
					cli.EnvVar("HUGO_EXPIRED"),
					cli.File("/vela/parameters/hugo/expired"),
					cli.File("/vela/secrets/hugo/expired"),
				),
			},
			&cli.StringFlag{
				Name:  "build.future",
				Usage: "include content with publishdate in the future",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_FUTURE"),
					cli.EnvVar("HUGO_FUTURE"),
					cli.File("/vela/parameters/hugo/future"),
					cli.File("/vela/secrets/hugo/future"),
				),
			},

			// Config Flags
			&cli.StringFlag{
				Name:  "config.cache_directory",
				Usage: "filesystem path to cache directory",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_CACHE_DIRECTORY"),
					cli.EnvVar("HUGO_CACHE_DIRECTORY"),
					cli.File("/vela/parameters/hugo/cache_directory"),
					cli.File("/vela/secrets/hugo/cache_directory"),
				),
			},
			&cli.StringFlag{
				Name:  "config.content_directory",
				Usage: "filesystem path to content directory",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_CONTENT_DIRECTORY"),
					cli.EnvVar("HUGO_CONTENT_DIRECTORY"),
					cli.File("/vela/parameters/hugo/content_directory"),
					cli.File("/vela/secrets/hugo/content_directory"),
				),
			},
			&cli.StringFlag{
				Name:  "config.directory",
				Usage: "filesystem path to config directory",
				Value: "config",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_CONFIG_DIRECTORY"),
					cli.EnvVar("HUGO_CONFIG_DIRECTORY"),
					cli.File("/vela/parameters/hugo/config_directory"),
					cli.File("/vela/secrets/hugo/config_directory"),
				),
			},
			&cli.StringFlag{
				Name:  "config.environment",
				Usage: "targeted build environment in config directory",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_ENVIRONMENT"),
					cli.EnvVar("HUGO_ENVIRONMENT"),
					cli.File("/vela/parameters/hugo/environment"),
					cli.File("/vela/secrets/hugo/environment"),
				),
			},
			&cli.StringFlag{
				Name:  "config.file",
				Usage: "name of config file in config directory",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_CONFIG_FILE"),
					cli.EnvVar("HUGO_CONFIG_FILE"),
					cli.File("/vela/parameters/hugo/config_file"),
					cli.File("/vela/secrets/hugo/config_file"),
				),
			},
			&cli.StringFlag{
				Name:  "config.layout_directory",
				Usage: "filesystem path to layout directory",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_LAYOUT_DIRECTORY"),
					cli.EnvVar("HUGO_LAYOUT_DIRECTORY"),
					cli.File("/vela/parameters/hugo/layout_directory"),
					cli.File("/vela/secrets/hugo/layout_directory"),
				),
			},
			&cli.StringFlag{
				Name:  "config.output_directory",
				Usage: "filesystem path to write files to",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_OUTPUT_DIRECTORY"),
					cli.EnvVar("HUGO_OUTPUT_DIRECTORY"),
					cli.File("/vela/parameters/hugo/output_directory"),
					cli.File("/vela/secrets/hugo/output_directory"),
				),
			},
			&cli.StringFlag{
				Name:  "config.source_directory",
				Usage: "filesystem path to read files from",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_SOURCE_DIRECTORY"),
					cli.EnvVar("HUGO_SOURCE_DIRECTORY"),
					cli.File("/vela/parameters/hugo/source_directory"),
					cli.File("/vela/secrets/hugo/source_directory"),
				),
			},

			// Theme Flags
			&cli.StringFlag{
				Name:  "theme.name",
				Usage: "name of theme to use from theme directory",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_THEME_NAME"),
					cli.EnvVar("HUGO_THEME_NAME"),
					cli.File("/vela/parameters/hugo/theme_name"),
					cli.File("/vela/secrets/hugo/theme_name"),
				),
			},
			&cli.StringFlag{
				Name:  "theme.directory",
				Usage: "filesystem path to theme directory",
				Value: "themes",
				Sources: cli.NewValueSourceChain(
					cli.EnvVar("PARAMETER_THEME_DIRECTORY"),
					cli.EnvVar("HUGO_THEME_DIRECTORY"),
					cli.File("/vela/parameters/hugo/theme_directory"),
					cli.File("/vela/secrets/hugo/theme_directory"),
				),
			},
		},
	}

	err = app.Run(context.Background(), os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

// run executes the plugin based off the configuration provided.
func run(ctx context.Context, c *cli.Command) error {
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
		err := install(ctx, extended, version, os.Getenv("PLUGIN_HUGO_VERSION"))
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
	return p.Exec(ctx)
}
