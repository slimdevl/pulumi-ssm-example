// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Deprecated: aws.getPrefixList has been deprecated in favor of aws.ec2.getPrefixList
func GetPrefixList(ctx *pulumi.Context, args *GetPrefixListArgs, opts ...pulumi.InvokeOption) (*GetPrefixListResult, error) {
	var rv GetPrefixListResult
	err := ctx.Invoke("aws:index/getPrefixList:getPrefixList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrefixList.
type GetPrefixListArgs struct {
	// Configuration block(s) for filtering. Detailed below.
	Filters []GetPrefixListFilter `pulumi:"filters"`
	// The name of the filter field. Valid values can be found in the [EC2 DescribePrefixLists API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribePrefixLists.html).
	Name *string `pulumi:"name"`
	// The ID of the prefix list to select.
	PrefixListId *string `pulumi:"prefixListId"`
}

// A collection of values returned by getPrefixList.
type GetPrefixListResult struct {
	// The list of CIDR blocks for the AWS service associated with the prefix list.
	CidrBlocks []string              `pulumi:"cidrBlocks"`
	Filters    []GetPrefixListFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the selected prefix list.
	Name         string  `pulumi:"name"`
	PrefixListId *string `pulumi:"prefixListId"`
}

func GetPrefixListOutput(ctx *pulumi.Context, args GetPrefixListOutputArgs, opts ...pulumi.InvokeOption) GetPrefixListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPrefixListResult, error) {
			args := v.(GetPrefixListArgs)
			r, err := GetPrefixList(ctx, &args, opts...)
			return *r, err
		}).(GetPrefixListResultOutput)
}

// A collection of arguments for invoking getPrefixList.
type GetPrefixListOutputArgs struct {
	// Configuration block(s) for filtering. Detailed below.
	Filters GetPrefixListFilterArrayInput `pulumi:"filters"`
	// The name of the filter field. Valid values can be found in the [EC2 DescribePrefixLists API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribePrefixLists.html).
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the prefix list to select.
	PrefixListId pulumi.StringPtrInput `pulumi:"prefixListId"`
}

func (GetPrefixListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPrefixListArgs)(nil)).Elem()
}

// A collection of values returned by getPrefixList.
type GetPrefixListResultOutput struct{ *pulumi.OutputState }

func (GetPrefixListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPrefixListResult)(nil)).Elem()
}

func (o GetPrefixListResultOutput) ToGetPrefixListResultOutput() GetPrefixListResultOutput {
	return o
}

func (o GetPrefixListResultOutput) ToGetPrefixListResultOutputWithContext(ctx context.Context) GetPrefixListResultOutput {
	return o
}

// The list of CIDR blocks for the AWS service associated with the prefix list.
func (o GetPrefixListResultOutput) CidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetPrefixListResult) []string { return v.CidrBlocks }).(pulumi.StringArrayOutput)
}

func (o GetPrefixListResultOutput) Filters() GetPrefixListFilterArrayOutput {
	return o.ApplyT(func(v GetPrefixListResult) []GetPrefixListFilter { return v.Filters }).(GetPrefixListFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPrefixListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPrefixListResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the selected prefix list.
func (o GetPrefixListResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetPrefixListResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetPrefixListResultOutput) PrefixListId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPrefixListResult) *string { return v.PrefixListId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPrefixListResultOutput{})
}
