package main

import (
  "fmt"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/ec2"
  "github.com/aws/aws-sdk-go/aws/session"
)

func main() {
/////////////////////////////////////////////////////////////////////////
// Get Account Attributes
/////////////////////////////////////////////////////////////////////////
  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

  params := &ec2.DescribeAccountAttributesInput{
    AttributeNames: []*string{
      aws.String("supported-platforms"),
      aws.String("default-vpc"),
      aws.String("max-instances"),
      aws.String("vpc-max-security-groups-per-interface"),
      aws.String("max-elastic-ips"),
      aws.String("vpc-max-elastic-ips"),
      },
      DryRun: aws.Bool(false),
      }
      resp, err := svc.DescribeAccountAttributes(params)

      if err != nil {
        // Print the error, cast err to awserr.Error to get the Code and
        // Message from an error.
        fmt.Println(err.Error())
        return
      }

  // Pretty-print the response data.
  fmt.Println(resp)
}
