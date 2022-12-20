package provider

import (
	"context"

	unique_string "github.com/cloud-maker-ai/go-unique-string"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &UniqueStringResource{}

func NewUniqueStringResource() resource.Resource {
	return &UniqueStringResource{}
}

type UniqueStringResource struct{}

type UniqueStringModel struct {
	ID     types.String `tfsdk:"id"`
	Input  types.List   `tfsdk:"input"`
	Result types.String `tfsdk:"result"`
}

func (r *UniqueStringResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_unique_string"
}

func (r *UniqueStringResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The resource `arm2tf_unique_string` generates a deterministic hash string based on the values provided as parameters.\n" +
			"It is a Terraform Provider implementation of the [uniqueString()](https://learn.microsoft.com/en-us/azure/azure-resource-manager/templates/template-functions-string#uniquestring) Azure Resource Manager template function.\n" +
			"\n" +
			"This resource uses [cloud-maker-ai/go-unique-string](https://github.com/cloud-maker-ai/go-unique-string).",
		Attributes: map[string]schema.Attribute{
			"input": schema.ListAttribute{
				Description: "List of string values that are used to generate the unique string. When changed, resource recreation will " +
					"be triggered.",
				ElementType: types.StringType,
				Required:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
			},
			"result": schema.StringAttribute{
				Description: "The generated unique string.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				Description: "The generated unique string.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *UniqueStringResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan UniqueStringModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var input []string
	diags = plan.Input.ElementsAs(ctx, &input, false)
	if diags.HasError() {
		return
	}

	result := unique_string.GenerateUniqueString(input...)

	u := &UniqueStringModel{
		ID:     types.StringValue(result),
		Result: types.StringValue(result),
		Input:  plan.Input,
	}

	diags = resp.State.Set(ctx, u)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read does not need to perform any operations as the state in ReadResourceResponse is already populated.
func (r *UniqueStringResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update ensures the plan value is copied to the state to complete the update.
func (r *UniqueStringResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var model UniqueStringModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *UniqueStringResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
