// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialive

import (
	"context"

	awstypes "github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	fwflex "github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
	fwtypes "github.com/hashicorp/terraform-provider-aws/internal/framework/types"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @FrameworkDataSource("aws_medialive_input", name="Input")
// @Tags
// @Testing(tagsIdentifierAttribute="arn")
func newInputDataSource(_ context.Context) (datasource.DataSourceWithConfigure, error) {
	return &inputDataSource{}, nil
}

const (
	DSNameInput = "Input Data Source"
)

type inputDataSource struct {
	framework.DataSourceWithModel[inputDataSourceModel]
}

func (d *inputDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			names.AttrARN: framework.ARNAttributeComputedOnly(),
			"attached_channels": schema.ListAttribute{
				CustomType: fwtypes.ListOfStringType,
				Computed:   true,
			},
			"destinations": framework.DataSourceComputedListOfObjectAttribute[dsDestination](ctx),
			names.AttrID: schema.StringAttribute{
				Required: true,
			},
			"input_class": schema.StringAttribute{
				CustomType: fwtypes.StringEnumType[awstypes.InputClass](),
				Computed:   true,
			},
			"input_devices": framework.DataSourceComputedListOfObjectAttribute[dsInputDevice](ctx),
			"input_partner_ids": schema.ListAttribute{
				CustomType: fwtypes.ListOfStringType,
				Computed:   true,
			},
			"input_source_type": schema.StringAttribute{
				CustomType: fwtypes.StringEnumType[awstypes.InputSourceType](),
				Computed:   true,
			},
			"media_connect_flows": framework.DataSourceComputedListOfObjectAttribute[dsMediaConnectFlow](ctx),
			names.AttrName: schema.StringAttribute{
				Computed: true,
			},
			names.AttrRoleARN: schema.StringAttribute{
				Computed: true,
			},
			names.AttrSecurityGroups: schema.ListAttribute{
				CustomType: fwtypes.ListOfStringType,
				Computed:   true,
			},
			"sources": framework.DataSourceComputedListOfObjectAttribute[dsInputSource](ctx),
			names.AttrState: schema.StringAttribute{
				CustomType: fwtypes.StringEnumType[awstypes.InputState](),
				Computed:   true,
			},
			names.AttrTags: tftags.TagsAttributeComputedOnly(),
			names.AttrType: schema.StringAttribute{
				CustomType: fwtypes.StringEnumType[awstypes.InputType](),
				Computed:   true,
			},
		},
	}
}

func (d *inputDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	conn := d.Meta().MediaLiveClient(ctx)

	var data inputDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	out, err := FindInputByID(ctx, conn, data.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.MediaLive, create.ErrActionReading, DSNameInput, data.ID.String(), err),
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(fwflex.Flatten(ctx, out, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setTagsOut(ctx, out.Tags)

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

type inputDataSourceModel struct {
	framework.WithRegionModel
	ARN               types.String                                        `tfsdk:"arn"`
	AttachedChannels  fwtypes.ListValueOf[types.String]                   `tfsdk:"attached_channels"`
	Destinations      fwtypes.ListNestedObjectValueOf[dsDestination]      `tfsdk:"destinations"`
	ID                types.String                                        `tfsdk:"id"`
	InputClass        fwtypes.StringEnum[awstypes.InputClass]             `tfsdk:"input_class"`
	InputDevices      fwtypes.ListNestedObjectValueOf[dsInputDevice]      `tfsdk:"input_devices"`
	InputPartnerIDs   fwtypes.ListValueOf[types.String]                   `tfsdk:"input_partner_ids"`
	InputSourceType   fwtypes.StringEnum[awstypes.InputSourceType]        `tfsdk:"input_source_type"`
	MediaConnectFlows fwtypes.ListNestedObjectValueOf[dsMediaConnectFlow] `tfsdk:"media_connect_flows"`
	Name              types.String                                        `tfsdk:"name"`
	RoleARN           types.String                                        `tfsdk:"role_arn"`
	SecurityGroups    fwtypes.ListValueOf[types.String]                   `tfsdk:"security_groups"`
	Sources           fwtypes.ListNestedObjectValueOf[dsInputSource]      `tfsdk:"sources"`
	State             fwtypes.StringEnum[awstypes.InputState]             `tfsdk:"state"`
	Tags              tftags.Map                                          `tfsdk:"tags"`
	Type              fwtypes.StringEnum[awstypes.InputType]              `tfsdk:"type"`
}

type dsDestination struct {
	IP   types.String                           `tfsdk:"ip"`
	Port types.String                           `tfsdk:"port"`
	URL  types.String                           `tfsdk:"url"`
	VPC  fwtypes.ListNestedObjectValueOf[dsVPC] `tfsdk:"vpc"`
}

type dsVPC struct {
	AvailabilityZone   types.String `tfsdk:"availability_zone"`
	NetworkInterfaceID types.String `tfsdk:"network_interface_id"`
}

type dsInputDevice struct {
	ID types.String `tfsdk:"id"`
}

type dsMediaConnectFlow struct {
	FlowARN types.String `tfsdk:"flow_arn"`
}

type dsInputSource struct {
	PasswordParam types.String `tfsdk:"password_param"`
	URL           types.String `tfsdk:"url"`
	Username      types.String `tfsdk:"username"`
}
