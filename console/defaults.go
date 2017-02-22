/*
Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package console

import (
	"github.com/fatih/color"
	"github.com/wallix/awless/graph"
)

var DefaultsColumnDefinitions = map[graph.ResourceType][]ColumnDefinition{
	//EC2
	graph.Instance: {
		StringColumnDefinition{Prop: "Id"},
		//StringColumnDefinition{Prop: "SubnetId"},
		StringColumnDefinition{Prop: "Name", TruncateSize: 50, TruncateRight: false, DisableTruncate: false},
		ColoredValueColumnDefinition{
			StringColumnDefinition: StringColumnDefinition{Prop: "State", DisableTruncate: false, TruncateSize: 7},
			ColoredValues:          map[string]color.Attribute{"running": color.FgGreen, "stopped": color.FgRed},
		},
		StringColumnDefinition{Prop: "Type"},
		StringColumnDefinition{Prop: "KeyName", Friendly: "Access Key"},
		//StringColumnDefinition{Prop: "PublicIp", Friendly: "Public IP"},
		StringColumnDefinition{Prop: "PrivateIp", Friendly: "Private IP"},
		TimeColumnDefinition{StringColumnDefinition: StringColumnDefinition{Prop: "LaunchTime"}},
	},
	graph.Vpc: {
		StringColumnDefinition{Prop: "Id"},
		StringColumnDefinition{Prop: "Name"},
		ColoredValueColumnDefinition{
			StringColumnDefinition: StringColumnDefinition{Prop: "IsDefault", Friendly: "Default"},
			ColoredValues:          map[string]color.Attribute{"true": color.FgGreen},
		},
		StringColumnDefinition{Prop: "State"},
		StringColumnDefinition{Prop: "CidrBlock"},
	},
	graph.Subnet: {
		StringColumnDefinition{Prop: "Id"},
		StringColumnDefinition{Prop: "Name"},
		StringColumnDefinition{Prop: "CidrBlock"},
		StringColumnDefinition{Prop: "AvailabilityZone", Friendly: "Zone"},
		StringColumnDefinition{Prop: "VpcId"},
		ColoredValueColumnDefinition{
			StringColumnDefinition: StringColumnDefinition{Prop: "MapPublicIpOnLaunch", Friendly: "Public VMs"},
			ColoredValues:          map[string]color.Attribute{"true": color.FgYellow}},
		ColoredValueColumnDefinition{
			StringColumnDefinition: StringColumnDefinition{Prop: "State"},
			ColoredValues:          map[string]color.Attribute{"available": color.FgGreen}},
		ColoredValueColumnDefinition{
			StringColumnDefinition: StringColumnDefinition{Prop: "DefaultForAz", Friendly: "ZoneDefault"},
			ColoredValues:          map[string]color.Attribute{"true": color.FgGreen},
		},
	},
	graph.SecurityGroup: {
		StringColumnDefinition{Prop: "Id"},
		StringColumnDefinition{Prop: "VpcId"},
		FirewallRulesColumnDefinition{StringColumnDefinition: StringColumnDefinition{Prop: "InboundRules", Friendly: "Inbound"}},
		FirewallRulesColumnDefinition{StringColumnDefinition: StringColumnDefinition{Prop: "OutboundRules", Friendly: "Outbound"}},
		StringColumnDefinition{Prop: "Name", DisableTruncate: true},
		StringColumnDefinition{Prop: "Description", DisableTruncate: true},
	},
	graph.InternetGateway: {
		StringColumnDefinition{Prop: "Id"},
		StringColumnDefinition{Prop: "Name", DisableTruncate: true},
		StringColumnDefinition{Prop: "Vpcs", DisableTruncate: true},
	},
	graph.RouteTable: {
		StringColumnDefinition{Prop: "Id"},
		StringColumnDefinition{Prop: "Name", DisableTruncate: true},
		StringColumnDefinition{Prop: "VpcId"},
		StringColumnDefinition{Prop: "Main"},
		RoutesColumnDefinition{StringColumnDefinition: StringColumnDefinition{Prop: "Routes"}},
	},
	graph.Keypair: {
		StringColumnDefinition{Prop: "Id"},
		StringColumnDefinition{Prop: "KeyFingerprint", DisableTruncate: true},
	},
	graph.Volume: {
		StringColumnDefinition{Prop: "Id"},
		StringColumnDefinition{Prop: "Name", DisableTruncate: true},
		StringColumnDefinition{Prop: "VolumeType"},
		StringColumnDefinition{Prop: "State"},
		StringColumnDefinition{Prop: "Size", Friendly: "Size (Gb)"},
		StringColumnDefinition{Prop: "Encrypted"},
		TimeColumnDefinition{StringColumnDefinition: StringColumnDefinition{Prop: "CreateTime"}},
		StringColumnDefinition{Prop: "AvailabilityZone"},
	},
	//IAM
	graph.User: {
		StringColumnDefinition{Prop: "Id"},
		StringColumnDefinition{Prop: "Name", DisableTruncate: true},
		TimeColumnDefinition{StringColumnDefinition: StringColumnDefinition{Prop: "PasswordLastUsedDate", Friendly: "PasswordLastUsed"}},
		TimeColumnDefinition{StringColumnDefinition: StringColumnDefinition{Prop: "CreateDate"}},
	},
	graph.Role: {
		StringColumnDefinition{Prop: "Id"},
		StringColumnDefinition{Prop: "Name", DisableTruncate: true},
		TimeColumnDefinition{StringColumnDefinition: StringColumnDefinition{Prop: "CreateDate"}},
	},
	graph.Policy: {
		StringColumnDefinition{Prop: "Id"},
		StringColumnDefinition{Prop: "Name", DisableTruncate: true},
		TimeColumnDefinition{StringColumnDefinition: StringColumnDefinition{Prop: "CreateDate"}},
		TimeColumnDefinition{StringColumnDefinition: StringColumnDefinition{Prop: "UpdateDate"}},
	},
	graph.Group: {
		StringColumnDefinition{Prop: "Id"},
		StringColumnDefinition{Prop: "Name", DisableTruncate: true},
		TimeColumnDefinition{StringColumnDefinition: StringColumnDefinition{Prop: "CreateDate"}},
	},
	// S3
	graph.Bucket: {
		StringColumnDefinition{Prop: "Name", DisableTruncate: true},
		GrantsColumnDefinition{StringColumnDefinition: StringColumnDefinition{Prop: "Grants"}},
		TimeColumnDefinition{StringColumnDefinition: StringColumnDefinition{Prop: "CreateDate"}},
	},
	graph.Object: {
		StringColumnDefinition{Prop: "Key", TruncateRight: true},
		StringColumnDefinition{Prop: "BucketName"},
		TimeColumnDefinition{StringColumnDefinition: StringColumnDefinition{Prop: "ModifiedDate"}},
		StringColumnDefinition{Prop: "OwnerId", TruncateRight: true},
		StringColumnDefinition{Prop: "Size"},
		StringColumnDefinition{Prop: "Class"},
	},
}
