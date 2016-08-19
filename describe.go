package main

import (
"fmt"
"github.com/asdqwex/aws_mgmt/util"
"github.com/asdqwex/aws_mgmt/describe"
"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/ec2"
"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	regions := util.Get_regions()
	for _, v := range regions {
		sess, err := session.NewSession(&aws.Config{Region: aws.String(v)})
		if err != nil {
		  fmt.Println("failed to create session,", err)
		  return
		}
		svc := ec2.New(sess)

		account_attrs := describe.Account(svc)
		fmt.Println(account_attrs)

		acls := describe.Acl(svc)
		fmt.Println(acls)

	}
}

