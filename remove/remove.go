package remove

import (
"fmt"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/cloudformation"
)

func Cloudformation(svc *cloudformation.CloudFormation, stackName string) (removed *cloudformation.DeleteStackOutput){

	params := &cloudformation.DeleteStackInput{
	    StackName: aws.String(stackName), // Required
	}
	resp, err := svc.DeleteStack(params)

	if err != nil {

	    fmt.Println(err.Error())
	    return
	}

	removed = resp
	return
}
