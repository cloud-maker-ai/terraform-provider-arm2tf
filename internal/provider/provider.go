package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure ARM2TFProvider satisfies various provider interfaces.
var _ provider.Provider = &ARM2TFProvider{}

// ARM2TFProvider defines the provider implementation.
type ARM2TFProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

func (p *ARM2TFProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "arm2tf"
	resp.Version = p.version
}

func (p *ARM2TFProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
}

func (p *ARM2TFProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *ARM2TFProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewGuidResource,
		NewUniqueStringResource,
	}
}

func (p *ARM2TFProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ARM2TFProvider{
			version: version,
		}
	}
}
