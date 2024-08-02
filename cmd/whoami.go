package cmd

import (
	"fmt"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/siliconflow/siliconcloud-cli/lib"
	"github.com/siliconflow/siliconcloud-cli/meta"
	"github.com/urfave/cli/v2"
	"os"
)

func Whoami(c *cli.Context) error {
	args, err := globalArgs.Parse(c, meta.CmdWhoami)
	if err != nil {
		return cli.Exit(err, meta.LoadError)
	}
	setLogVerbose(args.Verbose)
	logs.Debugf("args: %#v\n", args)

	apiKey, err := lib.NewSfFolder().GetKey()
	if err != nil {
		return err
	}

	client := lib.NewClient(args.BaseDomain, apiKey)
	info, err := client.UserInfo()
	if err != nil {
		return err
	}

	if info.Data.Name != "" {
		fmt.Fprintf(os.Stdout, "Your account name: %s\n", info.Data.Name)
	} else {
		fmt.Fprintf(os.Stdout, "Your account email: %s\n", info.Data.Email)
	}

	return nil
}
