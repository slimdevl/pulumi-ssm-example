package main

import (
	"strings"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		values, err := GetParametersByPathAsEnvironmentVariables(ctx, "/dev/jb")
		if err != nil {
			return err
		}
		for key, value := range values {
			ctx.Export(key, pulumi.String(value))
		}
		return nil
	})
}

func GetParametersByPathAsEnvironmentVariables(ctx *pulumi.Context, path string) (map[string]string, error) {
	parameters := map[string]string{}
	opt := true
	pRes, err := ssm.GetParametersByPath(ctx, &ssm.GetParametersByPathArgs{
		Path:           path,
		Recursive:      &opt,
		WithDecryption: &opt,
	})
	if err != nil {
		return nil, err
	}
	for i, pNm := range pRes.Names {
		name := transformName(pNm)
		parameters[name] = pRes.Values[i]
	}
	return parameters, nil
}

func transformName(name string) string {
	parts := strings.Split(name, "/")
	// discard all parts until we find the key term SAI_
	for i, part := range parts {
		if !strings.HasPrefix(part, "SAI_") {
			continue
		}
		return strings.ToUpper(strings.Join(parts[i:], "_"))
	}
	return ""
}
