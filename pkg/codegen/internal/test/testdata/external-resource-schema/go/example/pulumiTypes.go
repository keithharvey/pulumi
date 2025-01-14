// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Pet struct {
	Age  *int              `pulumi:"age"`
	Name *random.RandomPet `pulumi:"name"`
}

// PetInput is an input type that accepts PetArgs and PetOutput values.
// You can construct a concrete instance of `PetInput` via:
//
//          PetArgs{...}
type PetInput interface {
	pulumi.Input

	ToPetOutput() PetOutput
	ToPetOutputWithContext(context.Context) PetOutput
}

type PetArgs struct {
	Age  pulumi.IntPtrInput    `pulumi:"age"`
	Name random.RandomPetInput `pulumi:"name"`
}

func (PetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Pet)(nil)).Elem()
}

func (i PetArgs) ToPetOutput() PetOutput {
	return i.ToPetOutputWithContext(context.Background())
}

func (i PetArgs) ToPetOutputWithContext(ctx context.Context) PetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PetOutput)
}

func (i PetArgs) ToPetPtrOutput() PetPtrOutput {
	return i.ToPetPtrOutputWithContext(context.Background())
}

func (i PetArgs) ToPetPtrOutputWithContext(ctx context.Context) PetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PetOutput).ToPetPtrOutputWithContext(ctx)
}

// PetPtrInput is an input type that accepts PetArgs, PetPtr and PetPtrOutput values.
// You can construct a concrete instance of `PetPtrInput` via:
//
//          PetArgs{...}
//
//  or:
//
//          nil
type PetPtrInput interface {
	pulumi.Input

	ToPetPtrOutput() PetPtrOutput
	ToPetPtrOutputWithContext(context.Context) PetPtrOutput
}

type petPtrType PetArgs

func PetPtr(v *PetArgs) PetPtrInput {
	return (*petPtrType)(v)
}

func (*petPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Pet)(nil)).Elem()
}

func (i *petPtrType) ToPetPtrOutput() PetPtrOutput {
	return i.ToPetPtrOutputWithContext(context.Background())
}

func (i *petPtrType) ToPetPtrOutputWithContext(ctx context.Context) PetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PetPtrOutput)
}

type PetOutput struct{ *pulumi.OutputState }

func (PetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Pet)(nil)).Elem()
}

func (o PetOutput) ToPetOutput() PetOutput {
	return o
}

func (o PetOutput) ToPetOutputWithContext(ctx context.Context) PetOutput {
	return o
}

func (o PetOutput) ToPetPtrOutput() PetPtrOutput {
	return o.ToPetPtrOutputWithContext(context.Background())
}

func (o PetOutput) ToPetPtrOutputWithContext(ctx context.Context) PetPtrOutput {
	return o.ApplyT(func(v Pet) *Pet {
		return &v
	}).(PetPtrOutput)
}
func (o PetOutput) Age() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Pet) *int { return v.Age }).(pulumi.IntPtrOutput)
}

func (o PetOutput) Name() random.RandomPetOutput {
	return o.ApplyT(func(v Pet) *random.RandomPet { return v.Name }).(random.RandomPetOutput)
}

type PetPtrOutput struct{ *pulumi.OutputState }

func (PetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pet)(nil)).Elem()
}

func (o PetPtrOutput) ToPetPtrOutput() PetPtrOutput {
	return o
}

func (o PetPtrOutput) ToPetPtrOutputWithContext(ctx context.Context) PetPtrOutput {
	return o
}

func (o PetPtrOutput) Elem() PetOutput {
	return o.ApplyT(func(v *Pet) Pet { return *v }).(PetOutput)
}

func (o PetPtrOutput) Age() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Pet) *int {
		if v == nil {
			return nil
		}
		return v.Age
	}).(pulumi.IntPtrOutput)
}

func (o PetPtrOutput) Name() random.RandomPetOutput {
	return o.ApplyT(func(v *Pet) *random.RandomPet {
		if v == nil {
			return nil
		}
		return v.Name
	}).(random.RandomPetOutput)
}

func init() {
	pulumi.RegisterOutputType(PetOutput{})
	pulumi.RegisterOutputType(PetPtrOutput{})
}
