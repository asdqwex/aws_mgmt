package main

import (
"fmt"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/ec2"
"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
/////////////////////////////////////////////////////////////////////////
// NEW EC2 Session
/////////////////////////////////////////////////////////////////////////
  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

/////////////////////////////////////////////////////////////////////////
// Describe EIPs
/////////////////////////////////////////////////////////////////////////

  params := &ec2.DescribeAddressesInput{
    DryRun: aws.Bool(false),
  }
  resp, err := svc.DescribeAddresses(params)

  if err != nil {
      // Print the error, cast err to awserr.Error to get the Code and
      // Message from an error.
    fmt.Println(err.Error())
    return
  }

  // Pretty-print the response data.
  fmt.Println(resp)
}
