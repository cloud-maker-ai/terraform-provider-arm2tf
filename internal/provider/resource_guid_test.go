package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceGuid_Format_Is_Correct(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `resource "arm2tf_guid" "test" {
							input = [ "test" ]
						}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("arm2tf_guid.test", "result", regexp.MustCompile(`[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}`)),
				),
			},
		},
	})
}

func TestAccResourceGuid_Single_Value_Is_Correct(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config: `resource "arm2tf_guid" "test" {
					input = [ "test" ]
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("arm2tf_guid.test", "id", "d6a5baf1-8b9b-542a-8525-2982f1f98a0c"),
					resource.TestCheckResourceAttr("arm2tf_guid.test", "result", "d6a5baf1-8b9b-542a-8525-2982f1f98a0c"),
				),
			},
		},
	})
}

func TestAccResourceGuid_Multiple_Values_Are_Correct(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config: `resource "arm2tf_guid" "test" {
					input = [
						"test",
						"test2",
						"test3",
					]
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("arm2tf_guid.test", "id", "ba91b5ff-f126-51c5-a93f-d37f8292ee79"),
					resource.TestCheckResourceAttr("arm2tf_guid.test", "result", "ba91b5ff-f126-51c5-a93f-d37f8292ee79")),
			},
		},
	})
}
