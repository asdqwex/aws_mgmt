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

func Az(svc *ec2.EC2) (azs *ec2.DescribeAvailabilityZonesOutput){

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
	azs = resp
  return
}

func Customer_gateway(svc *ec2.EC2) (customer_gateways *ec2.DescribeCustomerGatewaysOutput){

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
  customer_gateways = resp
  return
}

func Dhcp(svc *ec2.EC2) (dhcps *ec2.DescribeDhcpOptionsOutput){

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
  dhcps = resp
  return
}

func Eip(svc *ec2.EC2) (eips *ec2.DescribeAddressesOutput){

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
  eips = resp
  return
}

func Flowlog(svc *ec2.EC2) (flowlogs *ec2.DescribeFlowLogsOutput){

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
  flowlogs = resp
  return
}

func Igw(svc *ec2.EC2) (igws *ec2.DescribeInternetGatewaysOutput){

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
	igws = resp
  return
}

func Image(svc *ec2.EC2) (images *ec2.DescribeImagesOutput){

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
  images = resp
  return
}

func Instance(svc *ec2.EC2) (instances *ec2.DescribeInstancesOutput){

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
	instances = resp
  return
}

func Nat(svc *ec2.EC2) (nats *ec2.DescribeNatGatewaysOutput){

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
	nats = resp
  return
}

func Nic(svc *ec2.EC2) (nics *ec2.DescribeNetworkInterfacesOutput){

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
	nics = resp
  return
}

func Placement_group(svc *ec2.EC2) (placement_groups *ec2.DescribePlacementGroupsOutput){

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
  placement_groups = resp
  return
}

func Region() (regions *ec2.DescribeRegionsOutput){
  
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

func Route_table(svc *ec2.EC2) (route_tables *ec2.DescribeRouteTablesOutput){

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
	route_tables = resp
  return
}

func Security_group(svc *ec2.EC2) (security_groups *ec2.DescribeSecurityGroupsOutput){

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
	security_groups = resp
  return
}

func Snapshot(svc *ec2.EC2) (snapshots *ec2.DescribeSnapshotsOutput){

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
	snapshots = resp
  return
}

func Spot_price(svc *ec2.EC2) (spot_prices *ec2.DescribeSpotPriceHistoryOutput){

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
	spot_prices = resp
  return
}

func Volume(svc *ec2.EC2) (volumes *ec2.DescribeVolumesOutput){

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
	volumes = resp
  return
}

func Subnet(svc *ec2.EC2) (subnets *ec2.DescribeSubnetsOutput){

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
	subnets = resp
  return
}

func Tag(svc *ec2.EC2) (tags *ec2.DescribeTagsOutput){

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
	tags = resp
  return
}

func Vpn_gateway(svc *ec2.EC2) (vpn_gateways *ec2.DescribeVpnGatewaysOutput){

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
	vpn_gateways = resp
  return
}

func Vpn_connection(svc *ec2.EC2) (vpn_connections *ec2.DescribeVpnConnectionsOutput){

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
	vpn_connections = resp
  return
}

func Vpc_peer(svc *ec2.EC2) (vpc_peers *ec2.DescribeVpcPeeringConnectionsOutput){

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
	vpc_peers = resp
  return
}

func Vpc(svc *ec2.EC2) (vpcs *ec2.DescribeVpcsOutput){

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
	vpcs = resp
  return
}

