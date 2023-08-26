# Simple Private VPC

This is a simple private VPC written in the GoLang AWS CDK.
It has only private subnets with egress, but it does not provision
a NAT. Instead it provisions Gateway interfaces to lower cost
while prototyping resources that must have a VPC (such as Fargate)
but don't need to have egress to the internet at large.

> Please note, there VPC Endpoints are lot cost but are NOT free

## Dependencies
You will need to install the following for your machine

- [AWS Cli](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
  - You must configure your machine for development with AWS
- [GoLang](https://go.dev/doc/install)
- NodeJS
  - I personally find [the fnm project](https://github.com/Schniz/fnm) to be reliable and portable node version manager
  - Alternatively [nvm](https://github.com/nvm-sh/nvm) is popular
- [Just](https://github.com/casey/just) (optional)
  - A language agnostic command runner

## Running, Testing, Deploying

The primary commands are straight forward

`$ just bootstrap`
 * This is necessary the first time you use the CDK with your AWS account
 * It is only needed once, you don't need to run it ever again

`$ just deploy`
 * Will synthesize the CloudFormation template and deploy it to your account

`$ just destroy`
 * Will tear down and delete all the resources created when you deployed
 * Be sure to do this when you no longer need your VPC, the VPC Endpoints will incur costs

`$ just test`
 * Validates a VPC will be created
 * Validates all subnets are private with egress

`$ just synth`
 * Useful if you are curious about or want to debug the rendered CloudFormation template.



## Further Reading
- https://docs.aws.amazon.com/whitepapers/latest/aws-privatelink/what-are-vpc-endpoints.html
- https://aws.amazon.com/vpc/pricing/
- https://aws.amazon.com/privatelink/pricing/
