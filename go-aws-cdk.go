package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type GoAwsCdkStackProps struct {
	awscdk.StackProps
}

func NewGoAwsCdkStack(scope constructs.Construct, id string, props *GoAwsCdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// create DB userTable
	userTable := awsdynamodb.NewTable(stack, jsii.String("user_table"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("username"),
			Type: dynamodb.AttributeType_STRING,
		},
		Stream: dynamodb.StreamViewType_NEW_IMAGE,
	})

	// lambda resource
	lambdaFunc := awslambda.NewFunction(stack, jsii.String("myLambdaFunction"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Code:    awslambda.AssetCode_FromAsset(jsii.String("./lambda/function.zip"), nil),
		Handler: jsii.String("main"),
	})

	// allow lambda to r/w to dynamodb table
	userTable.GrantReadWriteData(lambdaFunc)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewGoAwsCdkStack(app, "GoAwsCdkStack", &GoAwsCdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return nil
}
