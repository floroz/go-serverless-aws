# Get dependencies

```sh
$ go get
```

# Bootstrap for Lambda Deployment

```sh
$ cd lambda
$ make build
```

# AWS CDK

This project uses AWS CDK development with Go.

The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests
