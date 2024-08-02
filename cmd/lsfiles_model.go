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
	args, err := globalArgs.Parse(c, meta.CmdLsFiles)
	if err != nil {
		return cli.Exit(err, meta.LoadError)
	}
	setLogVerbose(args.Verbose)
	logs.Debugf("args: %#v\n", args)

	if err = checkType(args, true); err != nil {
		return err
	}

	if err = checkName(args, false); err != nil {
		return err
	}

	apiKey, err := lib.NewSfFolder().GetKey()
	if err != nil {
		return err
	}

	client := lib.NewClient(args.BaseDomain, apiKey)

	modelFilesResp, err := client.ListModelFiles(args.Type, args.Name, args.ExtName)
	if err != nil {
		return err
	}

	modelFiles := modelFilesResp.Data.Files

	if len(modelFiles) < 1 {
		fmt.Fprintln(os.Stdout, "No files found.")
		return nil
	}

	if args.FormatTree {
		root := lib.NewNode("")
		for _, mf := range modelFiles {
			root.AddPath(mf.LabelPath)
		}
		root.PrintTree("")
	} else {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
		fmt.Fprintln(w, "Path\t Available\t")
		// Print data rows
		for _, mr := range modelFiles {
			fmt.Fprintf(w, "%s\t %s\t \n", mr.LabelPath, func() string {
				if mr.Available {
					return "Yes"
				}
				return "No"
			}())
		}
		w.Flush()
	}

	return nil
}
