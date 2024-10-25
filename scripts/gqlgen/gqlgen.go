package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
)

var packageName = flag.String("package", "", "package name for the resolver")

func init() {
	// paste this in `resolver.go` -> go:generate go run $PROJECT_ROOT/scripts/gqlgen/gqlgen.go -package=graphapi
	flag.Parse()
}

func main() {
	if packageName == nil || *packageName == "" {
		fmt.Fprintln(os.Stderr, "package name is required")
		os.Exit(2)
	}

	cfg := &config.Config{
		SchemaFilename: config.StringList{"schema/**/*.graphql"},
		Exec: config.ExecConfig{
			Filename: "gen/gen.go",
			Package:  "gen",
		},
		Federation: config.PackageConfig{
			// Filename: "gen/federation.go",
			// Package:  "graph",
		},
		Model: config.PackageConfig{
			Filename: "gen/model/models_gen.go",
			Package:  "model",
		},
		Resolver: config.ResolverConfig{
			Layout:           "follow-schema",
			DirName:          ".",
			Package:          *packageName,
			FilenameTemplate: "{name}.resolvers.go",
		},
		Models: config.TypeMap{
			"ID": config.TypeMapEntry{
				Model: config.StringList{
					"github.com/99designs/gqlgen/graphql.ID",
					"github.com/99designs/gqlgen/graphql.Int",
					"github.com/99designs/gqlgen/graphql.Int64",
					"github.com/99designs/gqlgen/graphql.Int32",
				},
			},
			"Int": config.TypeMapEntry{
				Model: config.StringList{
					"github.com/99designs/gqlgen/graphql.Int",
					"github.com/99designs/gqlgen/graphql.Int64",
					"github.com/99designs/gqlgen/graphql.Int32",
				},
			},
		},
		Directives: map[string]config.DirectiveConfig{},
	}

	if err := config.CompleteConfig(cfg); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := api.Generate(cfg); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
