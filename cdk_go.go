package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	cfn_inc "github.com/aws/aws-cdk-go/awscdk/v2/cloudformationinclude"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type config struct {
	// The name of the cloudformation template to include
	template string
	// The AWS region to deploy to
	Region string
	// The AWS account to deploy to
	Account string
	// Whether to deploy in debug mode
	Debug bool
}

type CdkGoStackProps struct {
	awscdk.StackProps
}

func NewCdkGoStack(scope constructs.Construct,
	id string, cfg config, props *CdkGoStackProps) awscdk.Stack {

	var sprops awscdk.StackProps

	if props != nil {
		sprops = props.StackProps
	}

	stack := awscdk.NewStack(scope, &id, &sprops)

	cfnTemplate := cfn_inc.NewCfnInclude(
		stack, jsii.String("Template"), &cfn_inc.CfnIncludeProps{
			TemplateFile: jsii.String(cfg.template),
		})

	fmt.Println(cfnTemplate.Stack())

	return stack
}

func main() {
	var cfg config

	flag.StringVar(&cfg.template, "template", "templates/s3.yaml", "The name of the cloudformation template to include")
	flag.StringVar(&cfg.Region, "region", "us-east-1", "The AWS region to deploy to")
	flag.StringVar(&cfg.Account, "account", "123456789012", "The AWS account to deploy to")

	defer jsii.Close()
	flag.Parse()

	app := awscdk.NewApp(nil)
	NewCdkGoStack(app, "CdkGoStack", cfg, &CdkGoStackProps{
		awscdk.StackProps{
			Env: &awscdk.Environment{
				Account: jsii.String(cfg.Account),
				Region:  jsii.String(cfg.Region),
			},
		},
	})
	app.Synth(nil)
}
