package ipcontrol

import (
	"context"
	"fmt"
	cc "terraform-provider-ipcontrol/ipcontrol/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDomain() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDomainRead,
		Schema: map[string]*schema.Schema{
			"domain_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Default",
				Description: "Domain type already defined to IPControl. If not specified, the “Default” domain type will be used.",
			},
			"domain_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the domain.",
			},
			"contact": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The contact email address in dotted format. If not specified a default contact name will be formed by prepending 'dnsadmin' to the domain name as in 'dnsadmin.dip.com.'",
			},
			"delegated": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates that this domain will be associated directly with a zone file. Accepted values are true or false. If not specified, defaults to true.",
			},
			"default_ttl": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "86400",
				Description: "Default time to live. If not specified, defaults to “86400”; ignored if Managed is “False”",
			},
			"derivative": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "Standard",
				Description: "Specify the role of this domain. This can be one of “STANDARD”, “TEMPLATE”, “CATALOG”, or “ALIAS”. If not specified, defaults to “STANDARD”.",
			},
			"expire": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "604800",
				Description: "Zone expire time. If not specified, defaults to “604800”; ignored if Managed is “False”.",
			},
			"info_template": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the information template to be associated with this container.",
			},
			"managed": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates that this domain is fully defined in IPControl. Accepted values are true or false. If not specified, defaults to true.",
			},
			"negative_cache_ttl": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "86400",
				Description: "Zone negative cache time to live. If not specified, defaults to “86400”; ignored if Managed is “False”.",
			},
			"refresh": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "10800",
				Description: "Zone refresh interval. If not specified, defaults to “10800”; ignored if Managed is “False”.",
			},
			"retry": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "3600",
				Description: "Zone retry interval. If not specified, defaults to “3600”; ignored if Managed is “False”.",
			},
			"reverse": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates that this domain is a reverse inaddr.arpa domain. Accepted values are true or false. If not specified, defaults to false.",
			},
			"serial_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				Description: "Zone serial number. If not specified or specified as “0”, defaults to “1”; ignored if Managed is “False”.",
			},
			"template_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of template domain. Required, if Derivative is “ALIAS”.",
			},
			"user_defined_fields": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of user-defined fields in key=value format, e.g. DHCPServerip=192.0.1.1",
			},
			"local_rpz": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "A value of true indicates that this domain will contain RPZ rules.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A description of the domain.",
			},
			"serialformat": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The format of the serial number of the domain if it is being overriden at the domain level",
			},
		},
	}
}

func dataSourceDomainRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)

	domainName := d.Get("domain_name").(string)
	query := map[string]string{
		"name": domainName,
	}
	domain, err := objMgr.GetDomain(query)

	if err != nil {
		diag := diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error when reading domain",
			Detail:   fmt.Sprintf("Error when reading domain: (%v)", err),
		}
		diags = append(diags, diag)
		return diags
	}

	setIPCDomain(d, domain)
	return diags
}
