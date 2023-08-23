package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestSimplePrivateVpcStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := NewSimpleIsolatedVpcStack(app, "MyStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack, &assertions.TemplateParsingOptions{})
	// - has the expected vpc
	template.HasResourceProperties(jsii.String("AWS::EC2::VPC"), map[string]interface{}{
		"Tags": assertions.Match_ArrayWith(&[]interface{}{
			assertions.Match_ObjectLike(&map[string]interface{}{
				"Key":   "Name",
				"Value": "SimplePrivateVPC",
			}),
		}),
	})
	// - has only private subnets with egress
	template.AllResourcesProperties(jsii.String("AWS::EC2::Subnet"), map[string]interface{}{
		"Tags": assertions.Match_ArrayWith(&[]interface{}{
			assertions.Match_ObjectLike(&map[string]interface{}{
				"Key":   "aws-cdk:subnet-type",
				"Value": "Private",
			}),
		}),
	})
}
