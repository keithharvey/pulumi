// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v2/codegen/internal/test/testdata/simple-enum-schema/go/plant"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type RubberTree struct {
	pulumi.CustomResourceState

	Container plant.ContainerPtrOutput `pulumi:"container"`
	Farm      pulumi.StringPtrOutput   `pulumi:"farm"`
	Type      pulumi.StringOutput      `pulumi:"type"`
}

// NewRubberTree registers a new resource with the given unique name, arguments, and options.
func NewRubberTree(ctx *pulumi.Context,
	name string, args *RubberTreeArgs, opts ...pulumi.ResourceOption) (*RubberTree, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if err := args.Type.Validate(); err != nil {
		return nil, fmt.Errorf("invalid value for enum 'Type': %w", err)
	}
	var resource RubberTree
	err := ctx.RegisterResource("plant-provider:tree/v1:RubberTree", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRubberTree gets an existing RubberTree resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRubberTree(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RubberTreeState, opts ...pulumi.ResourceOption) (*RubberTree, error) {
	var resource RubberTree
	err := ctx.ReadResource("plant-provider:tree/v1:RubberTree", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RubberTree resources.
type rubberTreeState struct {
	Container *plant.Container `pulumi:"container"`
	Farm      *string          `pulumi:"farm"`
	Type      *string          `pulumi:"type"`
}

type RubberTreeState struct {
	Container plant.ContainerPtrInput
	Farm      pulumi.StringPtrInput
	Type      RubberTreeVariety
}

func (RubberTreeState) ElementType() reflect.Type {
	return reflect.TypeOf((*rubberTreeState)(nil)).Elem()
}

type rubberTreeArgs struct {
	Container *plant.Container `pulumi:"container"`
	Farm      *string          `pulumi:"farm"`
	Type      string           `pulumi:"type"`
}

// The set of arguments for constructing a RubberTree resource.
type RubberTreeArgs struct {
	Container plant.ContainerPtrInput
	Farm      pulumi.StringPtrInput
	Type      RubberTreeVariety
}

func (RubberTreeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rubberTreeArgs)(nil)).Elem()
}
