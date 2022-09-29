package cli

import (
	"fmt"

	"github.com/suzuki-shunsuke/tfcmt/pkg/controller"
	"github.com/suzuki-shunsuke/tfcmt/pkg/terraform"
	"github.com/urfave/cli/v2"
)

func cmdPlan(ctx *cli.Context) error {
	logLevel := ctx.String("log-level")
	setLogLevel(logLevel)

	fmt.Println("hi")

	cfg, err := newConfig(ctx)
	if err != nil {
		return err
	}
	fmt.Println("hi")
	if logLevel == "" {
		logLevel = cfg.Log.Level
		setLogLevel(logLevel)
	}
	fmt.Println("hi")

	if err := parseOpts(ctx, &cfg); err != nil {
		return err
	}

	fmt.Println("hi")

	t := &controller.Controller{
		Config:             cfg,
		Parser:             terraform.NewPlanParser(),
		Template:           terraform.NewPlanTemplate(cfg.Terraform.Plan.Template),
		ParseErrorTemplate: terraform.NewPlanParseErrorTemplate(cfg.Terraform.Plan.WhenParseError.Template),
	}
	args := ctx.Args()

	fmt.Println("ready?")

	return t.Run(ctx.Context, controller.Command{
		Cmd:  args.First(),
		Args: args.Tail(),
	})
}
