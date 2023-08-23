CDK_DEFAULT_ACCOUNT := `aws sts get-caller-identity --query Account --output text`
CDK_DEFAULT_REGION := `aws configure get region`

synth:
  npx cdk synth

test:
  go test -v ./...

deploy:
  npx cdk deploy

diff:
  npx cdk diff

destroy:
  npx cdk destroy

bootstrap:
  npx cdk bootstrap
