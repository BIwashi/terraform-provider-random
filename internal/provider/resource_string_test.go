package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceString(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: protoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: `resource "random_string" "basic" {
							length = 12
						}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith("random_string.basic", "result", testCheckLen(12)),
				),
			},
			{
				ResourceName:      "random_string.basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceString_Keepers_Keep_EmptyMap(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "0"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_Keep_EmptyMapToNullValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_Keep_NullMap(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "0"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_Keep_NullMapToNullValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_Keep_NullValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_Keep_NullValues(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key1" = null
						"key2" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "2"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key1" = null
						"key2" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "2"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_Keep_Value(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_Keep_Values(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key1" = "123"
						"key2" = "456"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "2"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key1" = "123"
						"key2" = "456"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "2"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_Replace_EmptyMapToValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_Replace_NullMapToValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_Replace_NullValueToValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_Replace_ValueToEmptyMap(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "0"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_Replace_ValueToNullMap(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "0"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_Replace_ValueToNullValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_Replace_ValueToNewValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = "456"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_FrameworkMigration_NullMapToNullValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: providerVersion332(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_FrameworkMigration_NullMapToValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: providerVersion332(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key" = "123"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_FrameworkMigration_NullMapToMultipleNullValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: providerVersion332(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key1" = null
						"key2" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key1" = null
						"key2" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "2"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_FrameworkMigration_NullMapToMultipleValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: providerVersion332(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key1" = null
						"key2" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "0"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key1" = "123"
						"key2" = "456"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "2"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_FrameworkMigration_NullMapValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: providerVersion332(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key1" = "123"
						"key2" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key1" = "123"
						"key2" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesEqual(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "2"),
				),
			},
		},
	})
}

func TestAccResourceString_Keepers_FrameworkMigration_NullMapValueToValue(t *testing.T) {
	var id1, id2 string

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: providerVersion332(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key1" = "123"
						"key2" = null
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id1),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "1"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "test" {
					length = 12
					keepers = {
						"key1" = "123"
						"key2" = "456"
					}
				}`,
				Check: resource.ComposeTestCheckFunc(
					testExtractResourceAttr("random_string.test", "id", &id2),
					testCheckAttributeValuesDiffer(&id1, &id2),
					resource.TestCheckResourceAttr("random_string.test", "keepers.%", "2"),
				),
			},
		},
	})
}

func TestAccResourceString_Override(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: protoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: `resource "random_string" "override" {
							length = 4
							override_special = "!"
							lower = false
							upper = false
							numeric = false
						}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith("random_string.override", "result", testCheckLen(4)),
					resource.TestCheckResourceAttr("random_string.override", "result", "!!!!"),
				),
			},
		},
	})
}

func TestAccResourceString_Min(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: protoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: `resource "random_string" "min" {
							length = 12
							override_special = "!#@"
							min_lower = 2
							min_upper = 3
							min_special = 1
							min_numeric = 4
						}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith("random_string.min", "result", testCheckLen(12)),
					resource.TestMatchResourceAttr("random_string.min", "result", regexp.MustCompile(`([a-z].*){2,}`)),
					resource.TestMatchResourceAttr("random_string.min", "result", regexp.MustCompile(`([A-Z].*){3,}`)),
					resource.TestMatchResourceAttr("random_string.min", "result", regexp.MustCompile(`([0-9].*){4,}`)),
					resource.TestMatchResourceAttr("random_string.min", "result", regexp.MustCompile(`([!#@].*)`)),
				),
			},
		},
	})
}

// TestAccResourceString_StateUpgradeV1toV2 covers the state upgrade from V1 to V2.
// This includes the deprecation of `number` and the addition of `numeric` attributes.
// v3.2.0 was used as this is the last version before `number` was deprecated and `numeric` attribute
// was added.
func TestAccResourceString_StateUpgradeV1toV2(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name                string
		configBeforeUpgrade string
		configDuringUpgrade string
		beforeStateUpgrade  []resource.TestCheckFunc
		afterStateUpgrade   []resource.TestCheckFunc
	}{
		{
			name: "number is absent before number and numeric are absent during",
			configBeforeUpgrade: `resource "random_string" "default" {
						length = 12
					}`,
			configDuringUpgrade: `resource "random_string" "default" {
						length = 12
					}`,
			beforeStateUpgrade: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("random_string.default", "number", "true"),
				resource.TestCheckNoResourceAttr("random_string.default", "numeric"),
			},
			afterStateUpgrade: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("random_string.default", "number", "true"),
				resource.TestCheckResourceAttrPair("random_string.default", "number", "random_string.default", "numeric"),
			},
		},
		{
			name: "number is absent before numeric is true during",
			configBeforeUpgrade: `resource "random_string" "default" {
						length = 12
					}`,
			configDuringUpgrade: `resource "random_string" "default" {
						length = 12
						numeric = true
					}`,
			beforeStateUpgrade: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("random_string.default", "number", "true"),
				resource.TestCheckNoResourceAttr("random_string.default", "numeric"),
			},
			afterStateUpgrade: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("random_string.default", "number", "true"),
				resource.TestCheckResourceAttrPair("random_string.default", "number", "random_string.default", "numeric"),
			},
		},
		{
			name: "number is absent before numeric is false during",
			configBeforeUpgrade: `resource "random_string" "default" {
						length = 12
					}`,
			configDuringUpgrade: `resource "random_string" "default" {
						length = 12
						numeric = false
					}`,
			beforeStateUpgrade: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("random_string.default", "number", "true"),
				resource.TestCheckNoResourceAttr("random_string.default", "numeric"),
			},
			afterStateUpgrade: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("random_string.default", "number", "false"),
				resource.TestCheckResourceAttrPair("random_string.default", "number", "random_string.default", "numeric"),
			},
		},
		{
			name: "number is true before numeric is true during",
			configBeforeUpgrade: `resource "random_string" "default" {
						length = 12
						number = true
					}`,
			configDuringUpgrade: `resource "random_string" "default" {
						length = 12
						numeric = true
					}`,
			beforeStateUpgrade: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("random_string.default", "number", "true"),
				resource.TestCheckNoResourceAttr("random_string.default", "numeric"),
			},
			afterStateUpgrade: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("random_string.default", "number", "true"),
				resource.TestCheckResourceAttrPair("random_string.default", "number", "random_string.default", "numeric"),
			},
		},
		{
			name: "number is true before number and numeric are absent during",
			configBeforeUpgrade: `resource "random_string" "default" {
						length = 12
						number = true
					}`,
			configDuringUpgrade: `resource "random_string" "default" {
						length = 12
					}`,
			beforeStateUpgrade: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("random_string.default", "number", "true"),
				resource.TestCheckNoResourceAttr("random_string.default", "numeric"),
			},
			afterStateUpgrade: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("random_string.default", "number", "true"),
				resource.TestCheckResourceAttrPair("random_string.default", "number", "random_string.default", "numeric"),
			},
		},
		{
			name: "number is true before numeric is false during",
			configBeforeUpgrade: `resource "random_string" "default" {
						length = 12
						number = true
					}`,
			configDuringUpgrade: `resource "random_string" "default" {
						length = 12
						numeric = false
					}`,
			beforeStateUpgrade: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("random_string.default", "number", "true"),
				resource.TestCheckNoResourceAttr("random_string.default", "numeric"),
			},
			afterStateUpgrade: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("random_string.default", "number", "false"),
				resource.TestCheckResourceAttrPair("random_string.default", "number", "random_string.default", "numeric"),
			},
		},
		{
			name: "number is false before numeric is false during",
			configBeforeUpgrade: `resource "random_string" "default" {
						length = 12
						number = false
					}`,
			configDuringUpgrade: `resource "random_string" "default" {
						length = 12
						numeric = false
					}`,
			beforeStateUpgrade: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("random_string.default", "number", "false"),
				resource.TestCheckNoResourceAttr("random_string.default", "numeric"),
			},
			afterStateUpgrade: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("random_string.default", "number", "false"),
				resource.TestCheckResourceAttrPair("random_string.default", "number", "random_string.default", "numeric"),
			},
		},
		{
			name: "number is false before number and numeric are absent during",
			configBeforeUpgrade: `resource "random_string" "default" {
						length = 12
						number = false
					}`,
			configDuringUpgrade: `resource "random_string" "default" {
						length = 12
					}`,
			beforeStateUpgrade: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("random_string.default", "number", "false"),
				resource.TestCheckNoResourceAttr("random_string.default", "numeric"),
			},
			afterStateUpgrade: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("random_string.default", "number", "true"),
				resource.TestCheckResourceAttrPair("random_string.default", "number", "random_string.default", "numeric"),
			},
		},
		{
			name: "number is false before numeric is true during",
			configBeforeUpgrade: `resource "random_string" "default" {
						length = 12
						number = false
					}`,
			configDuringUpgrade: `resource "random_string" "default" {
						length = 12
						numeric = true
					}`,
			beforeStateUpgrade: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("random_string.default", "number", "false"),
				resource.TestCheckNoResourceAttr("random_string.default", "numeric"),
			},
			afterStateUpgrade: []resource.TestCheckFunc{
				resource.TestCheckResourceAttr("random_string.default", "number", "true"),
				resource.TestCheckResourceAttrPair("random_string.default", "number", "random_string.default", "numeric"),
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			resource.UnitTest(t, resource.TestCase{
				Steps: []resource.TestStep{
					{
						ExternalProviders: map[string]resource.ExternalProvider{"random": {
							VersionConstraint: "3.2.0",
							Source:            "hashicorp/random",
						}},
						Config: c.configBeforeUpgrade,
						Check:  resource.ComposeTestCheckFunc(c.beforeStateUpgrade...),
					},
					{
						ProtoV5ProviderFactories: protoV5ProviderFactories(),
						Config:                   c.configDuringUpgrade,
						Check:                    resource.ComposeTestCheckFunc(c.afterStateUpgrade...),
					},
				},
			})
		})
	}
}

func TestAccResourceString_LengthErrors(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV5ProviderFactories: protoV5ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: `resource "random_string" "invalid_length" {
  							length = 2
  							min_lower = 3
						}`,
				ExpectError: regexp.MustCompile(`.*Attribute length value must be at least sum of min_upper \+ min_lower \+\nmin_numeric \+ min_special, got: 2`),
			},
			{
				Config: `resource "random_string" "invalid_length" {
							length = 0
						}`,
				ExpectError: regexp.MustCompile(`.*Attribute length value must be at least 1, got: 0`),
			},
		},
	})
}

func TestAccResourceString_UpgradeFromVersion3_3_2(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				ExternalProviders: providerVersion332(),
				Config: `resource "random_string" "min" {
							length = 12
							override_special = "!#@"
							min_lower = 2
							min_upper = 3
							min_special = 1
							min_numeric = 4
						}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith("random_string.min", "result", testCheckLen(12)),
					resource.TestMatchResourceAttr("random_string.min", "result", regexp.MustCompile(`([a-z].*){2,}`)),
					resource.TestMatchResourceAttr("random_string.min", "result", regexp.MustCompile(`([A-Z].*){3,}`)),
					resource.TestMatchResourceAttr("random_string.min", "result", regexp.MustCompile(`([0-9].*){4,}`)),
					resource.TestMatchResourceAttr("random_string.min", "result", regexp.MustCompile(`([!#@])`)),
					resource.TestCheckResourceAttr("random_string.min", "special", "true"),
					resource.TestCheckResourceAttr("random_string.min", "upper", "true"),
					resource.TestCheckResourceAttr("random_string.min", "lower", "true"),
					resource.TestCheckResourceAttr("random_string.min", "number", "true"),
					resource.TestCheckResourceAttr("random_string.min", "numeric", "true"),
					resource.TestCheckResourceAttr("random_string.min", "min_special", "1"),
					resource.TestCheckResourceAttr("random_string.min", "min_upper", "3"),
					resource.TestCheckResourceAttr("random_string.min", "min_lower", "2"),
					resource.TestCheckResourceAttr("random_string.min", "min_numeric", "4"),
				),
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "min" {
							length = 12
							override_special = "!#@"
							min_lower = 2
							min_upper = 3
							min_special = 1
							min_numeric = 4
						}`,
				PlanOnly: true,
			},
			{
				ProtoV5ProviderFactories: protoV5ProviderFactories(),
				Config: `resource "random_string" "min" {
							length = 12
							override_special = "!#@"
							min_lower = 2
							min_upper = 3
							min_special = 1
							min_numeric = 4
						}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrWith("random_string.min", "result", testCheckLen(12)),
					resource.TestMatchResourceAttr("random_string.min", "result", regexp.MustCompile(`([a-z].*){2,}`)),
					resource.TestMatchResourceAttr("random_string.min", "result", regexp.MustCompile(`([A-Z].*){3,}`)),
					resource.TestMatchResourceAttr("random_string.min", "result", regexp.MustCompile(`([0-9].*){4,}`)),
					resource.TestMatchResourceAttr("random_string.min", "result", regexp.MustCompile(`([!#@])`)),
					resource.TestCheckResourceAttr("random_string.min", "special", "true"),
					resource.TestCheckResourceAttr("random_string.min", "upper", "true"),
					resource.TestCheckResourceAttr("random_string.min", "lower", "true"),
					resource.TestCheckResourceAttr("random_string.min", "number", "true"),
					resource.TestCheckResourceAttr("random_string.min", "numeric", "true"),
					resource.TestCheckResourceAttr("random_string.min", "min_special", "1"),
					resource.TestCheckResourceAttr("random_string.min", "min_upper", "3"),
					resource.TestCheckResourceAttr("random_string.min", "min_lower", "2"),
					resource.TestCheckResourceAttr("random_string.min", "min_numeric", "4"),
				),
			},
		},
	})
}

func testCheckLen(expectedLen int) func(input string) error {
	return func(input string) error {
		if len(input) != expectedLen {
			return fmt.Errorf("expected length %d, actual length %d", expectedLen, len(input))
		}

		return nil
	}
}

//nolint:unparam
func testCheckMinLen(minLen int) func(input string) error {
	return func(input string) error {
		if len(input) < minLen {
			return fmt.Errorf("minimum length %d, actual length %d", minLen, len(input))
		}

		return nil
	}
}
