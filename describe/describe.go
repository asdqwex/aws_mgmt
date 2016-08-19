package describe

import (
"fmt"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/ec2"
"github.com/aws/aws-sdk-go/aws/session"
)

func Account(svc *ec2.EC2) (attributes *ec2.DescribeAccountAttributesOutput){

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
  attributes = resp
  return
}

func Acl(svc *ec2.EC2) (acls *ec2.DescribeNetworkAclsOutput){
	params := &ec2.DescribeNetworkAclsInput{
		DryRun: aws.Bool(false),
	}
	resp, err := svc.DescribeNetworkAcls(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	acls = resp
  return
}

func AZ() {

	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := ec2.New(sess)

	params := &ec2.DescribeAvailabilityZonesInput{
		DryRun: aws.Bool(false),
	}
	resp, err := svc.DescribeAvailabilityZones(params)

	if err != nil {
	    // Print the error, cast err to awserr.Error to get the Code and
	    // Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Customer_gateway() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

  params := &ec2.DescribeCustomerGatewaysInput{
    DryRun: aws.Bool(false),
  }
  resp, err := svc.DescribeCustomerGateways(params)

  if err != nil {
    // Print the error, cast err to awserr.Error to get the Code and
    // Message from an error.
    fmt.Println(err.Error())
    return
  }

  // Pretty-print the response data.
  fmt.Println(resp)
}

func Dhcp() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

  params := &ec2.DescribeDhcpOptionsInput{
    DryRun: aws.Bool(false),
  }
  resp, err := svc.DescribeDhcpOptions(params)

  if err != nil {
    // Print the error, cast err to awserr.Error to get the Code and
    // Message from an error.
    fmt.Println(err.Error())
    return
  }

  // Pretty-print the response data.
  fmt.Println(resp)
}

func Eip() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

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

func Flowlog() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

  params := &ec2.DescribeFlowLogsInput{
      MaxResults: aws.Int64(10),
  }
  resp, err := svc.DescribeFlowLogs(params)

  if err != nil {
      // Print the error, cast err to awserr.Error to get the Code and
      // Message from an error.
      fmt.Println(err.Error())
      return
  }

  // Pretty-print the response data.
  fmt.Println(resp)
}

func Igw() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

	params := &ec2.DescribeInternetGatewaysInput{
	  DryRun: aws.Bool(false),
	}
	resp, err := svc.DescribeInternetGateways(params)

	if err != nil {
	  // Print the error, cast err to awserr.Error to get the Code and
	  // Message from an error.
	  fmt.Println(err.Error())
	  return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Image() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

  params := &ec2.DescribeImagesInput{
    DryRun: aws.Bool(false),
  }
  resp, err := svc.DescribeImages(params)

  if err != nil {
    // Print the error, cast err to awserr.Error to get the Code and
    // Message from an error.
    fmt.Println(err.Error())
    return
  }

  // Pretty-print the response data.
  fmt.Println(resp)
}

func Instance() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

	params := &ec2.DescribeInstancesInput{
	  DryRun: aws.Bool(false),
	  MaxResults: aws.Int64(100),
	}
	resp, err := svc.DescribeInstances(params)

	if err != nil {
	  // Print the error, cast err to awserr.Error to get the Code and
	  // Message from an error.
	  fmt.Println(err.Error())
	  return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Nat() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

	params := &ec2.DescribeNatGatewaysInput{
		MaxResults: aws.Int64(100),
	}
	resp, err := svc.DescribeNatGateways(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Nic() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

	params := &ec2.DescribeNetworkInterfacesInput{
		DryRun: aws.Bool(false),
	}
	resp, err := svc.DescribeNetworkInterfaces(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Placement_group() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

  params := &ec2.DescribePlacementGroupsInput{
    DryRun: aws.Bool(false),
  }
  resp, err := svc.DescribePlacementGroups(params)

  if err != nil {
    // Print the error, cast err to awserr.Error to get the Code and
    // Message from an error.
    fmt.Println(err.Error())
    return
  }

  // Pretty-print the response data.
  fmt.Println(resp)
}

func Region() (regions *ec2.DescribeRegionsOutput) {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

	params := &ec2.DescribeRegionsInput{
	  DryRun: aws.Bool(false),
	}
	resp, err := svc.DescribeRegions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	regions = resp
  return
}

func Route_table() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

	params := &ec2.DescribeRouteTablesInput{
	  DryRun: aws.Bool(false),
	}
	resp, err := svc.DescribeRouteTables(params)

	if err != nil {
	  // Print the error, cast err to awserr.Error to get the Code and
	  // Message from an error.
	  fmt.Println(err.Error())
	  return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Security_group() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

	params := &ec2.DescribeSecurityGroupsInput{
	  DryRun: aws.Bool(false),
	}
	resp, err := svc.DescribeSecurityGroups(params)

	if err != nil {
	  // Print the error, cast err to awserr.Error to get the Code and
	  // Message from an error.
	  fmt.Println(err.Error())
	  return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Snapshot() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

	params := &ec2.DescribeSnapshotsInput{
	  DryRun: aws.Bool(false),
	}
	resp, err := svc.DescribeSnapshots(params)

	if err != nil {
	  // Print the error, cast err to awserr.Error to get the Code and
	  // Message from an error.
	  fmt.Println(err.Error())
	  return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Spot_price() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

	params := &ec2.DescribeSpotPriceHistoryInput{
	  DryRun: aws.Bool(false),
	  MaxResults: aws.Int64(100),
	}
	resp, err := svc.DescribeSpotPriceHistory(params)

	if err != nil {
	  // Print the error, cast err to awserr.Error to get the Code and
	  // Message from an error.
	  fmt.Println(err.Error())
	  return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Volume() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

	params := &ec2.DescribeVolumesInput{
	  DryRun: aws.Bool(false),
	  MaxResults: aws.Int64(100),
	}
	resp, err := svc.DescribeVolumes(params)

	if err != nil {
	  // Print the error, cast err to awserr.Error to get the Code and
	  // Message from an error.
	  fmt.Println(err.Error())
	  return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Subnet() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

	params := &ec2.DescribeSubnetsInput{
	  DryRun: aws.Bool(false),
	}
	resp, err := svc.DescribeSubnets(params)

	if err != nil {
	  // Print the error, cast err to awserr.Error to get the Code and
	  // Message from an error.
	  fmt.Println(err.Error())
	  return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Tag() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

	params := &ec2.DescribeTagsInput{
	  DryRun: aws.Bool(false),
	  MaxResults: aws.Int64(100),
	}
	resp, err := svc.DescribeTags(params)

	if err != nil {
	  // Print the error, cast err to awserr.Error to get the Code and
	  // Message from an error.
	  fmt.Println(err.Error())
	  return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Vpn_gateway() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

	params := &ec2.DescribeVpnGatewaysInput{
	  DryRun: aws.Bool(false),
	}
	resp, err := svc.DescribeVpnGateways(params)

	if err != nil {
	  // Print the error, cast err to awserr.Error to get the Code and
	  // Message from an error.
	  fmt.Println(err.Error())
	  return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Vpn_connection() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

	params := &ec2.DescribeVpnConnectionsInput{
	  DryRun: aws.Bool(false),
	}
	resp, err := svc.DescribeVpnConnections(params)

	if err != nil {
	  // Print the error, cast err to awserr.Error to get the Code and
	  // Message from an error.
	  fmt.Println(err.Error())
	  return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Vpc_peer() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

	params := &ec2.DescribeVpcPeeringConnectionsInput{
	  DryRun: aws.Bool(false),
	}
	resp, err := svc.DescribeVpcPeeringConnections(params)

	if err != nil {
	  // Print the error, cast err to awserr.Error to get the Code and
	  // Message from an error.
	  fmt.Println(err.Error())
	  return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Vpc() {

  sess, err := session.NewSession()
  if err != nil {
    fmt.Println("failed to create session,", err)
    return
  }

  svc := ec2.New(sess)

	params := &ec2.DescribeVpcsInput{
	  DryRun: aws.Bool(false),
	}
	resp, err := svc.DescribeVpcs(params)

	if err != nil {
	  // Print the error, cast err to awserr.Error to get the Code and
	  // Message from an error.
	  fmt.Println(err.Error())
	  return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

