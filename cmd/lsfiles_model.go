package cmd

import (
	"fmt"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/siliconflow/siliconcloud-cli/lib"
	"github.com/siliconflow/siliconcloud-cli/meta"
	"github.com/urfave/cli/v2"
	"os"
	"text/tabwriter"
)

func ListFilesModel(c *cli.Context) error {
	args, err := globalArgs.Parse(c, meta.CmdLs)
	if err != nil {
		return cli.Exit(err, meta.LoadError)
	}
	setLogVerbose(args.Verbose)
	logs.Debugf("args: %#v\n", args)

	if err = checkType(args); err != nil {
		return err
	}

	if err = checkName(args); err != nil {
		return err
	}

	apiKey, err := lib.NewSfFolder().GetKey()
	if err != nil {
		return err
	}

	client := lib.NewClient(args.BaseDomain, apiKey)

	modelFilesResp, err := client.ListModelFiles(args.Type, args.Name)
	if err != nil {
		return err
	}

	modelFiles := modelFilesResp.Data.Files

	if len(modelFiles) < 1 {
		fmt.Fprintln(os.Stdout, "No files found.")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "Paths:")
	// Print data rows
	for _, mf := range modelFiles {
		fmt.Fprintf(w, "%s\n", mf.Path)
	}
	w.Flush()

	return nil
}
