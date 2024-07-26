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

	if err = checkType(args); err != nil {
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

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "Name\t Type\t Last Modified Time\t")
	// Print data rows
	for _, mr := range modelRecords {
		fmt.Fprintf(w, "%s\t %s\t %s\t\n", mr.Name, mr.Type, mr.UpdatedAt)
	}
	w.Flush()

	return nil
}
