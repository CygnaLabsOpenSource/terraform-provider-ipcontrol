package ipcontrol

import (
	"context"
	"fmt"
	"log"
	"strconv"
	cc "terraform-provider-ipcontrol/ipcontrol/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAddresses() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAddressesRead,
		Schema: map[string]*schema.Schema{
			"aliases": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"container": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "The name of the container that contains the device.",
			},
			"rawcontainer": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Set to true to pass the container parameter through to the API without prefixing.",
			},
			"ip_address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The IP Address of the Device.",
			},
			"address_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The address type of this interface IP address",
			},
			"domain_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Domain name of the Device.",
			},
			"hostname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The host name of the Device.",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the Device.",
			},
			"device_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of the device.",
			},
			"domain_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the domain.",
			},
			"duid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The DHCP Unique Identifier for IPv6.",
			},
			"interfaces": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address_type": {
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Required:    true,
							Description: "The address type of this interface IP address.",
						},
						"container": {
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Optional:    true,
							Computed:    true,
							Description: "The name of the container that contains the device.",
						},
						"ip_address": {
							Type:        schema.TypeList,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Required:    true,
							Description: "The IP address of this interface.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The name of this interface.",
						},
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The address type of this interface IP address.",
						},
					},
				},
			},
			"resource_record_flag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAddressesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	ipAddress, ok := d.Get("ip_address").(string)
	container, ok := d.Get("container").(string)
	rawContainer := d.Get("rawcontainer").(bool)

	if !ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Can't read ip_address attribute",
			Detail:   fmt.Sprintf("Can't read ip_address attribute: (%v)", ipAddress),
		})
		return diags
	}

	query := map[string]string{
		"ipAddress": ipAddress,
	}

	if container != "" {
		query["container"] = container

		if rawContainer {
			query["rawcontainer"] = strconv.FormatBool(rawContainer)
		}
	}

	response, err := objMgr.GetAddress(query)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error when reading device address",
			Detail:   fmt.Sprintf("Error when reading device address (%v): %s", ipAddress, err.Error()),
		})
		return diags
	}

	setIPCAddressResource(d, *response)
	log.Printf("[DEBUG] %s: Completed reading device address", rsSubnetIdString(d))

	return diags
}
