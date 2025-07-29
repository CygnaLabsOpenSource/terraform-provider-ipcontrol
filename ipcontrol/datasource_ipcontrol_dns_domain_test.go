package ipcontrol

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var (
	dns                = "com."
	refresh            = "10899"
	retry              = "3699"
	expire             = "604799"
	default_ttl        = "86499"
	negative_cache_ttl = "86499"
)

func TestAccDataSourceDomain(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfigWithProviderIPC(
					fmt.Sprintf(`
					
					resource "ipcontrol_dns_domain" "my_domain" {
						domain_name = "%s"
						managed = true
						delegated = true
						refresh = "%s"
  						default_ttl = "%s"
						retry = "%s"
						expire = "%s"
 						negative_cache_ttl =  "%s"
					}
					
					data "ipcontrol_dns_domain" "my_dns" {
					  domain               = "%s"
					}
					`, dns, refresh, default_ttl, retry, expire, negative_cache_ttl, dns),
				),
				Check: resource.ComposeTestCheckFunc(
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
