package cmd

import (
	"fmt"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/siliconflow/siliconcloud-cli/lib"
	"github.com/siliconflow/siliconcloud-cli/meta"
	"github.com/urfave/cli/v2"
	"os"
)

func RemoveFileModel(c *cli.Context) error {
	args, err := globalArgs.Parse(c, meta.CmdRmFile)
	if err != nil {
		return cli.Exit(err, meta.LoadError)
	}
	setLogVerbose(args.Verbose)
	logs.Debugf("args: %#v\n", args)

	if err = checkType(args); err != nil {
		return err
	}

	if args.FilePath == "" {
		return cli.Exit("The following arguments are required: path", meta.LoadError)
	}

	apiKey, err := lib.NewSfFolder().GetKey()
	if err != nil {
		return err
	}

	client := lib.NewClient(args.BaseDomain, apiKey)

	_, err = client.RemoveModelFile(args.Type, args.Name, args.FilePath)
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, "File removed successfully.")

	return nil
}
