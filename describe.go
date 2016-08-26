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

		//account_attrs := describe.Account(svc)
		//fmt.Println(account_attrs)

		//acls := describe.Acl(svc)
		//fmt.Println(acls)

		//azs := describe.Az(svc)
		//fmt.Println(azs)

		//customer_gateways := describe.Customer_gateway(svc)
		//fmt.Println(customer_gateways)

		//dhcps := describe.Dhcp(svc)
		//fmt.Println(dhcps)

		//eips := describe.Eip(svc)
		//fmt.Println(eips)

		//flowlogs := describe.Flowlog(svc)
		//fmt.Println(flowlogs)

		//igws := describe.Igw(svc)
		//fmt.Println(igws)

		//images := describe.Image(svc)
		//fmt.Println(images)

		instances := describe.Instance(svc)
		fmt.Println(instances)

		//nats := describe.Nat(svc)
		//fmt.Println(nats)

		//nics := describe.Nic(svc)
		//fmt.Println(nics)

		//placement_groups := describe.Placement_group(svc)
		//fmt.Println(placement_groups)

		//route_tables := describe.Route_table(svc)
		//fmt.Println(route_tables)

		//security_groups := describe.Security_group(svc)
		//fmt.Println(security_groups)

		//snapshots := describe.Snapshot(svc)
		//fmt.Println(snapshots)

		//spot_prices := describe.Spot_price(svc)
		//fmt.Println(spot_prices)

		//volumes := describe.Volume(svc)
		//fmt.Println(volumes)

		//subnets := describe.Subnet(svc)
		//fmt.Println(subnets)

		//tag := describe.Tag(svc)
		//fmt.Println(tag)

		//vpn_gateways := describe.Vpn_gateway(svc)
		//fmt.Println(vpn_gateways)

		//vpn_connections := describe.Vpn_connection(svc)
		//fmt.Println(vpn_connections)

		//vpc_peers := describe.vpc_peer(svc)
		//fmt.Println(vpc_peers)

		//vpcs := describe.Vpc(svc)
		//fmt.Println(vpcs)

	}
}

