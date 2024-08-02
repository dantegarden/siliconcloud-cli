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

func ListModel(c *cli.Context) error {
	args, err := globalArgs.Parse(c, meta.CmdLs)
	if err != nil {
		return cli.Exit(err, meta.LoadError)
	}
	setLogVerbose(args.Verbose)
	logs.Debugf("args: %#v\n", args)

	if err = checkType(args, false); err != nil {
		return err
	}

	apiKey, err := lib.NewSfFolder().GetKey()
	if err != nil {
		return err
	}

	client := lib.NewClient(args.BaseDomain, apiKey)

	modelResp, err := client.ListModel(args.Type)
	if err != nil {
		return err
	}

	modelRecords := modelResp.Data.Models

	if len(modelRecords) < 1 {
		fmt.Fprintln(os.Stdout, "No models found.")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "Name\tType\tFile Count\tAvailable\tLast Modified Time\t")
	// Print data rows
	for _, mr := range modelRecords {
		fmt.Fprintf(w, "%s\t%s\t%d\t%s\t%s\t\n", mr.Name, mr.Type, mr.FileNum, func() string {
			if mr.Available {
				return "Yes"
			}
			return "No"
		}(), mr.UpdatedAt)
	}
	w.Flush()

	return nil
}
