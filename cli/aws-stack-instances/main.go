package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "aws-stack-instances"
	app.Usage = "Displays instances in an AWS Cloudformation stack"

	// --region
	// --show=[public dns|private dns|instance-id]

	app.Action = func(c *cli.Context) {
		println("Hello")
	}

	app.Run(os.Args)
}
