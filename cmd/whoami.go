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

	fmt.Fprintf(os.Stdout, "your account name: %s\n", info.Data.Email)

	return nil
}
