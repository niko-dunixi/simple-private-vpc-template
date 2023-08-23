package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type SimpleIsolatedVpcStackProps struct {
	awscdk.StackProps
}

func NewSimpleIsolatedVpcStack(scope constructs.Construct, id string, props *SimpleIsolatedVpcStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Create a VPC with private subnets with egress. We do not create a NAT.
	// So egress from the subnets is allowed, but not to the public internet.
	privateVpc := awsec2.NewVpc(stack, jsii.String("SimplePrivateVPC"), &awsec2.VpcProps{
		VpcName:               jsii.String("SimplePrivateVPC"),
		CreateInternetGateway: jsii.Bool(false),
		NatGateways:           jsii.Number(0),
		SubnetConfiguration: &[]*awsec2.SubnetConfiguration{
			{
				Name:       jsii.String("SingletonSubnet"),
				SubnetType: awsec2.SubnetType_PRIVATE_WITH_EGRESS,
			},
		},
		// Create Gateway Endpoints. These are free and common use-cases
		GatewayEndpoints: &map[string]*awsec2.GatewayVpcEndpointOptions{
			"S3Endpoint": {
				Service: awsec2.GatewayVpcEndpointAwsService_S3(),
			},
			"DynamoDBEndpoint": {
				Service: awsec2.GatewayVpcEndpointAwsService_DYNAMODB(),
			},
		},
	})
	// Create Interface Endpoints. These are not free, but are needed for common use cases.
	privateVpc.AddInterfaceEndpoint(jsii.String("EcrInterfaceGateway"), &awsec2.InterfaceVpcEndpointOptions{
		Service: awsec2.InterfaceVpcEndpointAwsService_ECR(),
	})
	privateVpc.AddInterfaceEndpoint(jsii.String("EcrDockerInterfaceGateway"), &awsec2.InterfaceVpcEndpointOptions{
		Service: awsec2.InterfaceVpcEndpointAwsService_ECR_DOCKER(),
	})
	privateVpc.AddInterfaceEndpoint(jsii.String("EcsInterfaceGateway"), &awsec2.InterfaceVpcEndpointOptions{
		Service: awsec2.InterfaceVpcEndpointAwsService_ECS(),
	})
	privateVpc.AddInterfaceEndpoint(jsii.String("SQSInterfaceGateway"), &awsec2.InterfaceVpcEndpointOptions{
		Service: awsec2.InterfaceVpcEndpointAwsService_SQS(),
	})
	privateVpc.AddInterfaceEndpoint(jsii.String("SNSInterfaceGateway"), &awsec2.InterfaceVpcEndpointOptions{
		Service: awsec2.InterfaceVpcEndpointAwsService_SNS(),
	})
	privateVpc.AddInterfaceEndpoint(jsii.String("CloudWatchInterfaceGateway"), &awsec2.InterfaceVpcEndpointOptions{
		Service: awsec2.InterfaceVpcEndpointAwsService_CLOUDWATCH(),
	})
	privateVpc.AddInterfaceEndpoint(jsii.String("CloudWatchLogsInterfaceGateway"), &awsec2.InterfaceVpcEndpointOptions{
		Service: awsec2.InterfaceVpcEndpointAwsService_CLOUDWATCH_LOGS(),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewSimpleIsolatedVpcStack(app, "SimplePrivateVpcStack", &SimpleIsolatedVpcStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
