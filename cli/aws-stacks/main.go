package main

import (
  "fmt"
  "os"
  "github.com/urfave/cli"
  "github.com/scottbrown/aws-stack-tools/stackutil"
)

func main() {
  app := cli.NewApp()
  app.Name = "aws-stacks"
  app.Usage = "Displays active AWS Cloudformation stacks"
  app.Version = "1.0.0"
  app.HideVersion = true
  app.HideHelp = true
  app.Copyright = "2016"
  app.Authors = []cli.Author {
    cli.Author{
      Name: "Scott Brown",
    },
  }

  verbose := false
  showHelp := false
  showVersion := false
  app.Flags = []cli.Flag {
    cli.BoolFlag{
      Name: "verbose, v",
      Usage: "Adds more information to the output",
      Destination: &verbose,
    },
    cli.BoolFlag{
      Name: "help, h",
      Usage: "Displays this message",
      Destination: &showHelp,
    },
    cli.BoolFlag{
      Name: "version",
      Usage: "Displays the version",
      Destination: &showVersion,
    },
  }

  app.Action = func(c *cli.Context) {
    if (showHelp) {
      println("Help here")
      return
    }

    if (showVersion) {
      println(app.Version)
      return
    }

    region := ""
    if len(c.Args()) > 0 {
      region = c.Args()[0]
    }

    // die if the region doesn't exist
    stack_names, err := stackutil.ActiveStacks(region)

    if err != nil {
      fmt.Println(err)
      return
    }

    for _, i := range stack_names {
      fmt.Println(i)
    }
  }

  app.Run(os.Args)
}
