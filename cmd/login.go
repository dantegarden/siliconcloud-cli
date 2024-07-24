package cmd

import (
	"fmt"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/siliconflow/siliconcloud-cli/lib"
	"github.com/siliconflow/siliconcloud-cli/meta"
	"github.com/urfave/cli/v2"
	"os"
)

func Login(c *cli.Context) error {
	args, err := globalArgs.Parse(c, meta.CmdLogin)
	if err != nil {
		return cli.Exit(err, meta.LoadError)
	}
	setLogVerbose(args.Verbose)
	logs.Debugf("args: %#v\n", args)

	if args.ApiKey == "" {
		return cli.Exit(fmt.Errorf("api key is required, you can specify \"--api_key\" to set"), meta.LoadError)
	}

	client := lib.NewClient(args.BaseDomain, args.ApiKey)
	_, err = client.UserInfo()
	if err != nil {
		return err
	}

	err = lib.NewSfFolder().SaveKey(args.ApiKey)
	if err != nil {
		return err
	}

	fmt.Fprint(os.Stdout, "Login success")
	return nil
}
