package main

import (
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		variables := map[string]string{
			"/dev/jb/foundation/SAI_VPC_ID":                                  "1",
			"/dev/jb/foundation/SAI_SLIM_VPC_PRIVATE_SUBNET/1/ID":            "2",
			"/dev/jb/storage/SLIM_QUEUES/SAI_CONNECTORCACHE_QUEUE":           "3",
			"/dev/jb/storage/SLIM_DATABASE/dynamodb/SAI_LOGINIDENTITY_TABLE": "4",
			"/dev/jb/storage/SLIM_DATABASE/workflow/SAI_WORKFLOW_USERNAME":   "5",
		}
		for key, value := range variables {
			fmt.Sprintf("making %s :: %s", key, value)
			if _, err := ssm.NewParameter(ctx,
				strings.Trim(key, "/"),
				&ssm.ParameterArgs{
					Name:  pulumi.String(key),
					Type:  pulumi.String("String"),
					Value: pulumi.String(value),
				},
			); err != nil {
				return err
			}

		}
		return nil
	})
}
