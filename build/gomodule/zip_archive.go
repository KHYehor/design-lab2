package gomodule

import (
	"fmt"
	"os"
	"path"

	"github.com/google/blueprint"
	"github.com/roman-mazur/bood"
)

var (
	// Ninja rule to archive binary output file.
	makeArchive = pctx.StaticRule("makeArchive", blueprint.RuleParams{
		Command:     "$zipCommand",
		Description: "make archive from binary",
	}, "zipCommand")
)

type zipArchiveModule struct {
	blueprint.SimpleName

	properties struct {
		// Srcs []string //  no srcs needed
	}
}

func (gb *zipArchiveModule) GenerateBuildActions(ctx blueprint.ModuleContext) {
	// creating final zip name
	name := ctx.ModuleName()
	archiveName := name + ".zip"
	config := bood.ExtractConfig(ctx)
	outputPath := path.Join(config.BaseOutputDir, "archives", archiveName)
	// taking args from cli
	argsWithoutProg := os.Args[1:]
	var zipCommand = ""
	for _, element := range argsWithoutProg {
		zipCommand += "zip " + outputPath + " " + element + " "
	}
	ctx.Build(pctx, blueprint.BuildParams{
		Description: fmt.Sprintf("Build %s as zip archive", name),
		Rule:        makeArchive,
		Outputs:     []string{outputPath},
		Implicits:   argsWithoutProg,
		Args: map[string]string{
			"zipCommand": zipCommand,
		},
	})
}

func SimpleArchiveFactory() (blueprint.Module, []interface{}) {
	mType := &zipArchiveModule{}
	return mType, []interface{}{&mType.SimpleName.Properties, &mType.properties}
}
