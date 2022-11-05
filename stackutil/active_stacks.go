package stackutil

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"sort"
)

type StackSummary struct {
	CreationTime        string
	StackId             string
	StackName           string
	StackStatus         string
	TemplateDescription string
}

type StackResults struct {
	StackSummaries []StackSummary
}

func ActiveStacks(region string) ([]string, error) {
	var stackNames []string

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
		return nil, err
	}

	for _, element := range resp.StackSummaries {
		stackNames = append(stackNames, *element.StackName)
	}

	sort.Strings(stackNames)

	return stackNames, nil
}
