// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an SSM Maintenance Window Task resource
//
// ## Example Usage
// ### Automation Tasks
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ssm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ssm.NewMaintenanceWindowTask(ctx, "example", &ssm.MaintenanceWindowTaskArgs{
// 			MaxConcurrency: pulumi.String("2"),
// 			MaxErrors:      pulumi.String("1"),
// 			Priority:       pulumi.Int(1),
// 			TaskArn:        pulumi.String("AWS-RestartEC2Instance"),
// 			TaskType:       pulumi.String("AUTOMATION"),
// 			WindowId:       pulumi.Any(aws_ssm_maintenance_window.Example.Id),
// 			Targets: ssm.MaintenanceWindowTaskTargetArray{
// 				&ssm.MaintenanceWindowTaskTargetArgs{
// 					Key: pulumi.String("InstanceIds"),
// 					Values: pulumi.StringArray{
// 						pulumi.Any(aws_instance.Example.Id),
// 					},
// 				},
// 			},
// 			TaskInvocationParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersArgs{
// 				AutomationParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersAutomationParametersArgs{
// 					DocumentVersion: pulumi.String(fmt.Sprintf("%v%v", "$", "LATEST")),
// 					Parameters: ssm.MaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameterArray{
// 						&ssm.MaintenanceWindowTaskTaskInvocationParametersAutomationParametersParameterArgs{
// 							Name: pulumi.String("InstanceId"),
// 							Values: pulumi.StringArray{
// 								pulumi.Any(aws_instance.Example.Id),
// 							},
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Run Command Tasks
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ssm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ssm.NewMaintenanceWindowTask(ctx, "example", &ssm.MaintenanceWindowTaskArgs{
// 			MaxConcurrency: pulumi.String("2"),
// 			MaxErrors:      pulumi.String("1"),
// 			Priority:       pulumi.Int(1),
// 			TaskArn:        pulumi.String("AWS-RunShellScript"),
// 			TaskType:       pulumi.String("RUN_COMMAND"),
// 			WindowId:       pulumi.Any(aws_ssm_maintenance_window.Example.Id),
// 			Targets: ssm.MaintenanceWindowTaskTargetArray{
// 				&ssm.MaintenanceWindowTaskTargetArgs{
// 					Key: pulumi.String("InstanceIds"),
// 					Values: pulumi.StringArray{
// 						pulumi.Any(aws_instance.Example.Id),
// 					},
// 				},
// 			},
// 			TaskInvocationParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersArgs{
// 				RunCommandParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersArgs{
// 					OutputS3Bucket:    pulumi.Any(aws_s3_bucket.Example.Bucket),
// 					OutputS3KeyPrefix: pulumi.String("output"),
// 					ServiceRoleArn:    pulumi.Any(aws_iam_role.Example.Arn),
// 					TimeoutSeconds:    pulumi.Int(600),
// 					NotificationConfig: &ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfigArgs{
// 						NotificationArn: pulumi.Any(aws_sns_topic.Example.Arn),
// 						NotificationEvents: pulumi.StringArray{
// 							pulumi.String("All"),
// 						},
// 						NotificationType: pulumi.String("Command"),
// 					},
// 					Parameters: ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameterArray{
// 						&ssm.MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameterArgs{
// 							Name: pulumi.String("commands"),
// 							Values: pulumi.StringArray{
// 								pulumi.String("date"),
// 							},
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Step Function Tasks
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ssm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ssm.NewMaintenanceWindowTask(ctx, "example", &ssm.MaintenanceWindowTaskArgs{
// 			MaxConcurrency: pulumi.String("2"),
// 			MaxErrors:      pulumi.String("1"),
// 			Priority:       pulumi.Int(1),
// 			TaskArn:        pulumi.Any(aws_sfn_activity.Example.Id),
// 			TaskType:       pulumi.String("STEP_FUNCTIONS"),
// 			WindowId:       pulumi.Any(aws_ssm_maintenance_window.Example.Id),
// 			Targets: ssm.MaintenanceWindowTaskTargetArray{
// 				&ssm.MaintenanceWindowTaskTargetArgs{
// 					Key: pulumi.String("InstanceIds"),
// 					Values: pulumi.StringArray{
// 						pulumi.Any(aws_instance.Example.Id),
// 					},
// 				},
// 			},
// 			TaskInvocationParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersArgs{
// 				StepFunctionsParameters: &ssm.MaintenanceWindowTaskTaskInvocationParametersStepFunctionsParametersArgs{
// 					Input: pulumi.String("{\"key1\":\"value1\"}"),
// 					Name:  pulumi.String("example"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// AWS Maintenance Window Task can be imported using the `window_id` and `window_task_id` separated by `/`.
//
// ```sh
//  $ pulumi import aws:ssm/maintenanceWindowTask:MaintenanceWindowTask task <window_id>/<window_task_id>
// ```
type MaintenanceWindowTask struct {
	pulumi.CustomResourceState

	// The description of the maintenance window task.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The maximum number of targets this task can be run for in parallel.
	MaxConcurrency pulumi.StringOutput `pulumi:"maxConcurrency"`
	// The maximum number of errors allowed before this task stops being scheduled.
	MaxErrors pulumi.StringOutput `pulumi:"maxErrors"`
	// The name of the maintenance window task.
	Name pulumi.StringOutput `pulumi:"name"`
	// The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The role that should be assumed when executing the task. If a role is not provided, Systems Manager uses your account's service-linked role. If no service-linked role for Systems Manager exists in your account, it is created for you.
	ServiceRoleArn pulumi.StringOutput `pulumi:"serviceRoleArn"`
	// The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
	Targets MaintenanceWindowTaskTargetArrayOutput `pulumi:"targets"`
	// The ARN of the task to execute.
	TaskArn pulumi.StringOutput `pulumi:"taskArn"`
	// Configuration block with parameters for task execution.
	TaskInvocationParameters MaintenanceWindowTaskTaskInvocationParametersPtrOutput `pulumi:"taskInvocationParameters"`
	// The type of task being registered. Valid values: `AUTOMATION`, `LAMBDA`, `RUN_COMMAND` or `STEP_FUNCTIONS`.
	TaskType pulumi.StringOutput `pulumi:"taskType"`
	// The Id of the maintenance window to register the task with.
	WindowId pulumi.StringOutput `pulumi:"windowId"`
}

// NewMaintenanceWindowTask registers a new resource with the given unique name, arguments, and options.
func NewMaintenanceWindowTask(ctx *pulumi.Context,
	name string, args *MaintenanceWindowTaskArgs, opts ...pulumi.ResourceOption) (*MaintenanceWindowTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MaxConcurrency == nil {
		return nil, errors.New("invalid value for required argument 'MaxConcurrency'")
	}
	if args.MaxErrors == nil {
		return nil, errors.New("invalid value for required argument 'MaxErrors'")
	}
	if args.TaskArn == nil {
		return nil, errors.New("invalid value for required argument 'TaskArn'")
	}
	if args.TaskType == nil {
		return nil, errors.New("invalid value for required argument 'TaskType'")
	}
	if args.WindowId == nil {
		return nil, errors.New("invalid value for required argument 'WindowId'")
	}
	var resource MaintenanceWindowTask
	err := ctx.RegisterResource("aws:ssm/maintenanceWindowTask:MaintenanceWindowTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMaintenanceWindowTask gets an existing MaintenanceWindowTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMaintenanceWindowTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MaintenanceWindowTaskState, opts ...pulumi.ResourceOption) (*MaintenanceWindowTask, error) {
	var resource MaintenanceWindowTask
	err := ctx.ReadResource("aws:ssm/maintenanceWindowTask:MaintenanceWindowTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MaintenanceWindowTask resources.
type maintenanceWindowTaskState struct {
	// The description of the maintenance window task.
	Description *string `pulumi:"description"`
	// The maximum number of targets this task can be run for in parallel.
	MaxConcurrency *string `pulumi:"maxConcurrency"`
	// The maximum number of errors allowed before this task stops being scheduled.
	MaxErrors *string `pulumi:"maxErrors"`
	// The name of the maintenance window task.
	Name *string `pulumi:"name"`
	// The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
	Priority *int `pulumi:"priority"`
	// The role that should be assumed when executing the task. If a role is not provided, Systems Manager uses your account's service-linked role. If no service-linked role for Systems Manager exists in your account, it is created for you.
	ServiceRoleArn *string `pulumi:"serviceRoleArn"`
	// The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
	Targets []MaintenanceWindowTaskTarget `pulumi:"targets"`
	// The ARN of the task to execute.
	TaskArn *string `pulumi:"taskArn"`
	// Configuration block with parameters for task execution.
	TaskInvocationParameters *MaintenanceWindowTaskTaskInvocationParameters `pulumi:"taskInvocationParameters"`
	// The type of task being registered. Valid values: `AUTOMATION`, `LAMBDA`, `RUN_COMMAND` or `STEP_FUNCTIONS`.
	TaskType *string `pulumi:"taskType"`
	// The Id of the maintenance window to register the task with.
	WindowId *string `pulumi:"windowId"`
}

type MaintenanceWindowTaskState struct {
	// The description of the maintenance window task.
	Description pulumi.StringPtrInput
	// The maximum number of targets this task can be run for in parallel.
	MaxConcurrency pulumi.StringPtrInput
	// The maximum number of errors allowed before this task stops being scheduled.
	MaxErrors pulumi.StringPtrInput
	// The name of the maintenance window task.
	Name pulumi.StringPtrInput
	// The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
	Priority pulumi.IntPtrInput
	// The role that should be assumed when executing the task. If a role is not provided, Systems Manager uses your account's service-linked role. If no service-linked role for Systems Manager exists in your account, it is created for you.
	ServiceRoleArn pulumi.StringPtrInput
	// The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
	Targets MaintenanceWindowTaskTargetArrayInput
	// The ARN of the task to execute.
	TaskArn pulumi.StringPtrInput
	// Configuration block with parameters for task execution.
	TaskInvocationParameters MaintenanceWindowTaskTaskInvocationParametersPtrInput
	// The type of task being registered. Valid values: `AUTOMATION`, `LAMBDA`, `RUN_COMMAND` or `STEP_FUNCTIONS`.
	TaskType pulumi.StringPtrInput
	// The Id of the maintenance window to register the task with.
	WindowId pulumi.StringPtrInput
}

func (MaintenanceWindowTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowTaskState)(nil)).Elem()
}

type maintenanceWindowTaskArgs struct {
	// The description of the maintenance window task.
	Description *string `pulumi:"description"`
	// The maximum number of targets this task can be run for in parallel.
	MaxConcurrency string `pulumi:"maxConcurrency"`
	// The maximum number of errors allowed before this task stops being scheduled.
	MaxErrors string `pulumi:"maxErrors"`
	// The name of the maintenance window task.
	Name *string `pulumi:"name"`
	// The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
	Priority *int `pulumi:"priority"`
	// The role that should be assumed when executing the task. If a role is not provided, Systems Manager uses your account's service-linked role. If no service-linked role for Systems Manager exists in your account, it is created for you.
	ServiceRoleArn *string `pulumi:"serviceRoleArn"`
	// The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
	Targets []MaintenanceWindowTaskTarget `pulumi:"targets"`
	// The ARN of the task to execute.
	TaskArn string `pulumi:"taskArn"`
	// Configuration block with parameters for task execution.
	TaskInvocationParameters *MaintenanceWindowTaskTaskInvocationParameters `pulumi:"taskInvocationParameters"`
	// The type of task being registered. Valid values: `AUTOMATION`, `LAMBDA`, `RUN_COMMAND` or `STEP_FUNCTIONS`.
	TaskType string `pulumi:"taskType"`
	// The Id of the maintenance window to register the task with.
	WindowId string `pulumi:"windowId"`
}

// The set of arguments for constructing a MaintenanceWindowTask resource.
type MaintenanceWindowTaskArgs struct {
	// The description of the maintenance window task.
	Description pulumi.StringPtrInput
	// The maximum number of targets this task can be run for in parallel.
	MaxConcurrency pulumi.StringInput
	// The maximum number of errors allowed before this task stops being scheduled.
	MaxErrors pulumi.StringInput
	// The name of the maintenance window task.
	Name pulumi.StringPtrInput
	// The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
	Priority pulumi.IntPtrInput
	// The role that should be assumed when executing the task. If a role is not provided, Systems Manager uses your account's service-linked role. If no service-linked role for Systems Manager exists in your account, it is created for you.
	ServiceRoleArn pulumi.StringPtrInput
	// The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
	Targets MaintenanceWindowTaskTargetArrayInput
	// The ARN of the task to execute.
	TaskArn pulumi.StringInput
	// Configuration block with parameters for task execution.
	TaskInvocationParameters MaintenanceWindowTaskTaskInvocationParametersPtrInput
	// The type of task being registered. Valid values: `AUTOMATION`, `LAMBDA`, `RUN_COMMAND` or `STEP_FUNCTIONS`.
	TaskType pulumi.StringInput
	// The Id of the maintenance window to register the task with.
	WindowId pulumi.StringInput
}

func (MaintenanceWindowTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowTaskArgs)(nil)).Elem()
}

type MaintenanceWindowTaskInput interface {
	pulumi.Input

	ToMaintenanceWindowTaskOutput() MaintenanceWindowTaskOutput
	ToMaintenanceWindowTaskOutputWithContext(ctx context.Context) MaintenanceWindowTaskOutput
}

func (*MaintenanceWindowTask) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindowTask)(nil))
}

func (i *MaintenanceWindowTask) ToMaintenanceWindowTaskOutput() MaintenanceWindowTaskOutput {
	return i.ToMaintenanceWindowTaskOutputWithContext(context.Background())
}

func (i *MaintenanceWindowTask) ToMaintenanceWindowTaskOutputWithContext(ctx context.Context) MaintenanceWindowTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowTaskOutput)
}

func (i *MaintenanceWindowTask) ToMaintenanceWindowTaskPtrOutput() MaintenanceWindowTaskPtrOutput {
	return i.ToMaintenanceWindowTaskPtrOutputWithContext(context.Background())
}

func (i *MaintenanceWindowTask) ToMaintenanceWindowTaskPtrOutputWithContext(ctx context.Context) MaintenanceWindowTaskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowTaskPtrOutput)
}

type MaintenanceWindowTaskPtrInput interface {
	pulumi.Input

	ToMaintenanceWindowTaskPtrOutput() MaintenanceWindowTaskPtrOutput
	ToMaintenanceWindowTaskPtrOutputWithContext(ctx context.Context) MaintenanceWindowTaskPtrOutput
}

type maintenanceWindowTaskPtrType MaintenanceWindowTaskArgs

func (*maintenanceWindowTaskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindowTask)(nil))
}

func (i *maintenanceWindowTaskPtrType) ToMaintenanceWindowTaskPtrOutput() MaintenanceWindowTaskPtrOutput {
	return i.ToMaintenanceWindowTaskPtrOutputWithContext(context.Background())
}

func (i *maintenanceWindowTaskPtrType) ToMaintenanceWindowTaskPtrOutputWithContext(ctx context.Context) MaintenanceWindowTaskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowTaskPtrOutput)
}

// MaintenanceWindowTaskArrayInput is an input type that accepts MaintenanceWindowTaskArray and MaintenanceWindowTaskArrayOutput values.
// You can construct a concrete instance of `MaintenanceWindowTaskArrayInput` via:
//
//          MaintenanceWindowTaskArray{ MaintenanceWindowTaskArgs{...} }
type MaintenanceWindowTaskArrayInput interface {
	pulumi.Input

	ToMaintenanceWindowTaskArrayOutput() MaintenanceWindowTaskArrayOutput
	ToMaintenanceWindowTaskArrayOutputWithContext(context.Context) MaintenanceWindowTaskArrayOutput
}

type MaintenanceWindowTaskArray []MaintenanceWindowTaskInput

func (MaintenanceWindowTaskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MaintenanceWindowTask)(nil)).Elem()
}

func (i MaintenanceWindowTaskArray) ToMaintenanceWindowTaskArrayOutput() MaintenanceWindowTaskArrayOutput {
	return i.ToMaintenanceWindowTaskArrayOutputWithContext(context.Background())
}

func (i MaintenanceWindowTaskArray) ToMaintenanceWindowTaskArrayOutputWithContext(ctx context.Context) MaintenanceWindowTaskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowTaskArrayOutput)
}

// MaintenanceWindowTaskMapInput is an input type that accepts MaintenanceWindowTaskMap and MaintenanceWindowTaskMapOutput values.
// You can construct a concrete instance of `MaintenanceWindowTaskMapInput` via:
//
//          MaintenanceWindowTaskMap{ "key": MaintenanceWindowTaskArgs{...} }
type MaintenanceWindowTaskMapInput interface {
	pulumi.Input

	ToMaintenanceWindowTaskMapOutput() MaintenanceWindowTaskMapOutput
	ToMaintenanceWindowTaskMapOutputWithContext(context.Context) MaintenanceWindowTaskMapOutput
}

type MaintenanceWindowTaskMap map[string]MaintenanceWindowTaskInput

func (MaintenanceWindowTaskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MaintenanceWindowTask)(nil)).Elem()
}

func (i MaintenanceWindowTaskMap) ToMaintenanceWindowTaskMapOutput() MaintenanceWindowTaskMapOutput {
	return i.ToMaintenanceWindowTaskMapOutputWithContext(context.Background())
}

func (i MaintenanceWindowTaskMap) ToMaintenanceWindowTaskMapOutputWithContext(ctx context.Context) MaintenanceWindowTaskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowTaskMapOutput)
}

type MaintenanceWindowTaskOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindowTask)(nil))
}

func (o MaintenanceWindowTaskOutput) ToMaintenanceWindowTaskOutput() MaintenanceWindowTaskOutput {
	return o
}

func (o MaintenanceWindowTaskOutput) ToMaintenanceWindowTaskOutputWithContext(ctx context.Context) MaintenanceWindowTaskOutput {
	return o
}

func (o MaintenanceWindowTaskOutput) ToMaintenanceWindowTaskPtrOutput() MaintenanceWindowTaskPtrOutput {
	return o.ToMaintenanceWindowTaskPtrOutputWithContext(context.Background())
}

func (o MaintenanceWindowTaskOutput) ToMaintenanceWindowTaskPtrOutputWithContext(ctx context.Context) MaintenanceWindowTaskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MaintenanceWindowTask) *MaintenanceWindowTask {
		return &v
	}).(MaintenanceWindowTaskPtrOutput)
}

type MaintenanceWindowTaskPtrOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowTaskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindowTask)(nil))
}

func (o MaintenanceWindowTaskPtrOutput) ToMaintenanceWindowTaskPtrOutput() MaintenanceWindowTaskPtrOutput {
	return o
}

func (o MaintenanceWindowTaskPtrOutput) ToMaintenanceWindowTaskPtrOutputWithContext(ctx context.Context) MaintenanceWindowTaskPtrOutput {
	return o
}

func (o MaintenanceWindowTaskPtrOutput) Elem() MaintenanceWindowTaskOutput {
	return o.ApplyT(func(v *MaintenanceWindowTask) MaintenanceWindowTask {
		if v != nil {
			return *v
		}
		var ret MaintenanceWindowTask
		return ret
	}).(MaintenanceWindowTaskOutput)
}

type MaintenanceWindowTaskArrayOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowTaskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MaintenanceWindowTask)(nil))
}

func (o MaintenanceWindowTaskArrayOutput) ToMaintenanceWindowTaskArrayOutput() MaintenanceWindowTaskArrayOutput {
	return o
}

func (o MaintenanceWindowTaskArrayOutput) ToMaintenanceWindowTaskArrayOutputWithContext(ctx context.Context) MaintenanceWindowTaskArrayOutput {
	return o
}

func (o MaintenanceWindowTaskArrayOutput) Index(i pulumi.IntInput) MaintenanceWindowTaskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MaintenanceWindowTask {
		return vs[0].([]MaintenanceWindowTask)[vs[1].(int)]
	}).(MaintenanceWindowTaskOutput)
}

type MaintenanceWindowTaskMapOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowTaskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MaintenanceWindowTask)(nil))
}

func (o MaintenanceWindowTaskMapOutput) ToMaintenanceWindowTaskMapOutput() MaintenanceWindowTaskMapOutput {
	return o
}

func (o MaintenanceWindowTaskMapOutput) ToMaintenanceWindowTaskMapOutputWithContext(ctx context.Context) MaintenanceWindowTaskMapOutput {
	return o
}

func (o MaintenanceWindowTaskMapOutput) MapIndex(k pulumi.StringInput) MaintenanceWindowTaskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MaintenanceWindowTask {
		return vs[0].(map[string]MaintenanceWindowTask)[vs[1].(string)]
	}).(MaintenanceWindowTaskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowTaskInput)(nil)).Elem(), &MaintenanceWindowTask{})
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowTaskPtrInput)(nil)).Elem(), &MaintenanceWindowTask{})
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowTaskArrayInput)(nil)).Elem(), MaintenanceWindowTaskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowTaskMapInput)(nil)).Elem(), MaintenanceWindowTaskMap{})
	pulumi.RegisterOutputType(MaintenanceWindowTaskOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowTaskPtrOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowTaskArrayOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowTaskMapOutput{})
}
