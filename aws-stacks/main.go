package main

import (
  "fmt"
  "os"
  "github.com/codegangsta/cli"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/cloudformation"
)

func main() {
  app := cli.NewApp()
  app.Name = "aws-stacks"
  app.Usage = "Displays active AWS Cloudformation stacks"
  app.Version = "1.0.0"
  app.Copyright = "2016"
  app.Authors = []cli.Author {
    cli.Author{
      Name: "Scott Brown",
    },
  }

  // --verbose
  verbose := false
  app.Flags = []cli.Flag {
    cli.BoolFlag{
      Name: "verbose, V",
      Usage: "Adds more information to the output",
      Destination: &verbose,
    },
  }

  // --region

  app.Action = func(c *cli.Context) {
    println("Verbose is:", verbose)

    region := ""
    if len(c.Args()) > 0 {
      region = c.Args()[0]
    }

    println("Region is:", region)

    // die if the region doesn't exist
    // ask aws for the list of stack names
    svc := cloudformation.New(session.New(), &aws.Config{Region: aws.String(region)})

    params := &cloudformation.ListStacksInput{
      StackStatusFilter: []*string{
        aws.String("CREATE_COMPLETE"),
        aws.String("UPDATE_COMPLETE"),
        aws.String("UPDATE_ROLLBACK_COMPLETE"),
      },
    }

    resp, err := svc.ListStacks(params)

    if err != nil {
      fmt.Println(err.Error())
      return
    }

    fmt.Println(resp)
  }

  app.Run(os.Args)
}
