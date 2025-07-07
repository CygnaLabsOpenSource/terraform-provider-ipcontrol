package ipcontrol

import (
	"context"
	"fmt"
	"strconv"

	en "terraform-provider-ipcontrol/ipcontrol/entities"
	cc "terraform-provider-ipcontrol/ipcontrol/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDomain() *schema.Resource {
	return &schema.Resource{
		CreateContext: createDomainContext,
		ReadContext:   getDomainContext,
		UpdateContext: updateDomainContext,
		DeleteContext: deleteDomainContext,

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
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The user defined fields associated with this domain, as listed in the domain information template specified in parameter infoTemplate. Required, if for UDFs defined as required fields.",
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

func setIPCDomain(d *schema.ResourceData, rr *en.IPCDomain) {
	d.SetId(strconv.Itoa(rr.ID))
	d.Set("serial_number", rr.SerialNumber)
	d.Set("domain_type", rr.DomainType)
	d.Set("description", rr.Description)
	d.Set("refresh", rr.Refresh)
	d.Set("derivative", rr.Derivative)
	d.Set("serialformat", rr.SerialFormat)
	d.Set("reverse", rr.Reverse)
	d.Set("info_template", rr.InfoTemplate)

	// Convert []string to comma-separated string for Terraform schema.TypeString
	// d.Set("user_defined_fields", strings.Join(rr.UserDefinedFields, ","))

	// no need to set domain_name here as it is already set in the resource
	// d.Set("domain_name", rr.DomainName)

	d.Set("managed", rr.Managed)
	d.Set("negative_cache_ttl", rr.NegativeCacheTTL)
	d.Set("contact", rr.Contact)
	d.Set("expire", rr.Expire)
	d.Set("default_ttl", rr.DefaultTTL)
	d.Set("delegated", rr.Delegated)
	d.Set("local_rpz", rr.LocalRpz)
	d.Set("template_domain", rr.TemplateDomain)
	d.Set("retry", rr.Retry)
}

func createDomainContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)

	payload := &en.IPCDomainPost{
		SerialNumber:     d.Get("serial_number").(int),
		DomainType:       d.Get("domain_type").(string),
		Description:      d.Get("description").(string),
		Refresh:          d.Get("refresh").(string),
		Derivative:       d.Get("derivative").(string),
		SerialFormat:     d.Get("serialformat").(string),
		Reverse:          d.Get("reverse").(bool),
		InfoTemplate:     d.Get("info_template").(string),
		Managed:          d.Get("managed").(bool),
		NegativeCacheTTL: d.Get("negative_cache_ttl").(string),
		Contact:          d.Get("contact").(string),
		DomainName:       d.Get("domain_name").(string),
		Expire:           d.Get("expire").(string),
		DefaultTTL:       d.Get("default_ttl").(string),
		Delegated:        d.Get("delegated").(bool),
		LocalRpz:         d.Get("local_rpz").(bool),
		TemplateDomain:   d.Get("template_domain").(string),
		Retry:            d.Get("retry").(string),
	}

	// Parse user_defined_fields (comma-separated string → []string)
	// if udfStr, ok := d.GetOk("user_defined_fields"); ok && udfStr.(string) != "" {
	// 	payload.UserDefinedFields = strings.Split(udfStr.(string), ",")
	// }

	// Gửi request tạo domain
	err := objMgr.CreateDomain(payload)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error when creating DNS domain",
			Detail:   fmt.Sprintf("Error creating domain: %v", err),
		})
		return diags
	}

	return getDomainContext(ctx, d, m)
}

func getDomainContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
			Summary:  "Error when getting domain",
			Detail:   fmt.Sprintf("Error when getting domain: (%v)", err),
		}
		diags = append(diags, diag)
		return diags
	}

	setIPCDomain(d, domain)
	return diags
}

func updateDomainContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.Partial(true)
	var diags diag.Diagnostics

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)

	idStr := d.Id()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Invalid ID",
			Detail:   fmt.Sprintf("Failed to parse ID: %v", err),
		})
		return diags
	}

	payload := &en.IPCDomainPost{
		ID:               id,
		SerialNumber:     d.Get("serial_number").(int),
		DomainType:       d.Get("domain_type").(string),
		Description:      d.Get("description").(string),
		Refresh:          d.Get("refresh").(string),
		Derivative:       d.Get("derivative").(string),
		SerialFormat:     d.Get("serialformat").(string),
		Reverse:          d.Get("reverse").(bool),
		InfoTemplate:     d.Get("info_template").(string),
		Managed:          d.Get("managed").(bool),
		NegativeCacheTTL: d.Get("negative_cache_ttl").(string),
		Contact:          d.Get("contact").(string),
		DomainName:       d.Get("domain_name").(string),
		Expire:           d.Get("expire").(string),
		DefaultTTL:       d.Get("default_ttl").(string),
		Delegated:        d.Get("delegated").(bool),
		LocalRpz:         d.Get("local_rpz").(bool),
		TemplateDomain:   d.Get("template_domain").(string),
		Retry:            d.Get("retry").(string),
	}

	// if udfStr, ok := d.GetOk("user_defined_fields"); ok && udfStr.(string) != "" {
	// 	payload.UserDefinedFields = strings.Split(udfStr.(string), ",")
	// }

	err = objMgr.UpdateDomain(payload)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error updating DNS domain",
			Detail:   fmt.Sprintf("UpdateDomain failed: %v", err),
		})
		return diags
	}

	diags = getDomainContext(ctx, d, m)
	d.Partial(false)
	return diags
}

func deleteDomainContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)

	domainType := d.Get("domain_type").(string)
	domainName := d.Get("domain_name").(string)

	err := objMgr.DeleteDomain(domainName, domainType)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error when domain record",
			Detail:   fmt.Sprintf("Error when domain record: (%v)", err),
		})
		return diags
	}

	return diags
}
