package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceUniqueString_Format_Is_Correct(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `resource "arm2tf_unique_string" "test" {
							input = [ "test" ]
						}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("arm2tf_unique_string.test", "result", regexp.MustCompile(`^([a-zA-Z0-9_-]){13}`)),
				),
			},
		},
	})
}

func TestAccResourceUniqueString_Single_Value_Is_Correct(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config: `resource "arm2tf_unique_string" "test" {
					input = [ "test" ]
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("arm2tf_unique_string.test", "id", "rbgf3xv4ufgzg"),
					resource.TestCheckResourceAttr("arm2tf_unique_string.test", "result", "rbgf3xv4ufgzg"),
				),
			},
		},
	})
}

func TestAccResourceUniqueString_Multiple_Values_Are_Correct(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				Config: `resource "arm2tf_unique_string" "test" {
					input = [
						"test",
						"test2",
						"test3",
					]
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("arm2tf_unique_string.test", "id", "bqgd334z2uj64"),
					resource.TestCheckResourceAttr("arm2tf_unique_string.test", "result", "bqgd334z2uj64")),
			},
		},
	})
}
