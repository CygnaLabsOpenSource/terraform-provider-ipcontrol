package ipcontrol

import (
	"fmt"
	cc "terraform-provider-ipcontrol/ipcontrol/utils"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDnsDomain(t *testing.T) {
	resourceName := "ipcontrol_dns_domain.my_dns"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDnsDestroy,
		Steps: []resource.TestStep{
			//  Step 1 create
			{

				Config: testAccConfigWithProviderIPC(
					fmt.Sprintf(`
					resource "ipcontrol_dns_domain" "my_dns" {
						domain_name  = "%s"
					}
					`, dns),
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "domain", dns),
				),
			},

			// Step 2 update data, add managed, delegated, refresh, default_ttl, retry, expire, negative_cache_ttl

			{

				Config: testAccConfigWithProviderIPC(
					fmt.Sprintf(`
					resource "ipcontrol_dns_domain" "my_dns" {
					  	domain_name = "%s"
						managed = true
						delegated = true
						refresh = "%s"
						default_ttl = "%s"
						retry = "%s"
						expire = "%s"
						negative_cache_ttl =  "%s"
					}
					`, dns, refresh, default_ttl, retry, expire, negative_cache_ttl),
				),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnsExists(resourceName),
					resource.TestCheckResourceAttr("data.ipcontrol_dns_domain.my_dns", "domain_name", dns),
					resource.TestCheckResourceAttr("data.ipcontrol_dns_domain.my_dns", "refresh", refresh),
					resource.TestCheckResourceAttr("data.ipcontrol_dns_domain.my_dns", "default_ttl", default_ttl),
					resource.TestCheckResourceAttr("data.ipcontrol_dns_domain.my_dns", "retry", retry),
					resource.TestCheckResourceAttr("data.ipcontrol_dns_domain.my_dns", "expire", expire),
					resource.TestCheckResourceAttr("data.ipcontrol_dns_domain.my_dns", "negative_cache_ttl", negative_cache_ttl),
				),
			},
		},
	})
}

// Helper function to check if domain exists
func testAccCheckDnsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Domain ID is set")
		}

		connector := testAccProvider.Meta().(*cc.Connector)
		objMgr := cc.NewObjectManager(connector)

		// Construct query based on the resource's attributes
		query := map[string]string{
			"name": rs.Primary.Attributes["domain_name"],
		}

		_, err := objMgr.GetDomain(query)
		if err != nil {
			return err
		}

		return nil
	}
}

// Helper function to check if domain is destroyed
func testAccCheckDnsDestroy(s *terraform.State) error {
	connector := testAccProvider.Meta().(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ipcontrol_dns_domain" {
			continue
		}

		// Construct query based on the resource's attributes
		query := map[string]string{
			"name": rs.Primary.Attributes["domain_name"],
		}

		_, err := objMgr.GetDomain(query)
		if err == nil {
			return fmt.Errorf("Domain still exists")
		}
	}

	return nil
}
