/* Copyright 2019 VMware, Inc.
   SPDX-License-Identifier: MPL-2.0 */

package vmc

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"os"
	"testing"
)

func TestAccDataSourceVmcCustomerSubnets_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceVmcCustomerSubnetsConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.vmc_customer_subnets.my_subnets", "ids.#", "4"),
					resource.TestCheckResourceAttr("data.vmc_customer_subnets.my_subnets", "ids.0", "subnet-8cabb6f5"),
					resource.TestCheckResourceAttr("data.vmc_customer_subnets.my_subnets", "ids.1", "subnet-1ecff155"),
					resource.TestCheckResourceAttr("data.vmc_customer_subnets.my_subnets", "ids.2", "subnet-98fc13c5"),
					resource.TestCheckResourceAttr("data.vmc_customer_subnets.my_subnets", "ids.3", "subnet-c895f2e3"),
				),
			},
		},
	})
}

func testAccDataSourceVmcCustomerSubnetsConfig() string {
	return fmt.Sprintf(`
	
data "vmc_connected_accounts" "my_accounts" {
    account_number = %q
}

data "vmc_customer_subnets" "my_subnets" {
	connected_account_id = data.vmc_connected_accounts.my_accounts.id
	region = "US_WEST_2"
}
`,
		os.Getenv("AWS_ACCOUNT_NUMBER"))
}
