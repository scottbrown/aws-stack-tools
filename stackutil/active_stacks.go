package stackutil

import (
  "fmt"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/cloudformation"
)

func ActiveStacks(region string) {
  //stack_names := []string{"a"}

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

  //return stack_names
}

