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

func resourceDns() *schema.Resource {
	return &schema.Resource{
		CreateContext: createDomainContext,
		ReadContext:   getDomainContext,
		UpdateContext: updateDomainContext,
		DeleteContext: deleteDomainContext,

		Schema: map[string]*schema.Schema{
			"domain_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Domain type already defined to IPControl. If not specified, the “Default” domain type will be used.",
			},
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Domain name where resource records are to be added.",
			},
			"ttl": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Time to live",
			},
			"data": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The data portion of the resource record. The format is dependent on the type specified above.",
			},
		},
	}
}

func setIPCDomainResource(d *schema.ResourceData, rr *en.IPCDnsRR) {
	id := strconv.Itoa(rr.ID)
	d.SetId(id)
	d.Set("data", rr.Data)
	d.Set("domain_type", rr.DomainType)
	d.Set("owner", rr.Owner)
	d.Set("domain", rr.Domain)
	d.Set("ttl", rr.TTL)
	d.Set("comment", rr.Comment)
	d.Set("resource_record_type", rr.ResourceRecType)
	d.Set("resource_record_class", rr.ResourceRecClass)
	d.Set("device_rec_flag", rr.DeviceRecFlag)
}

func createDomainContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)

	owner := d.Get("owner").(string)
	domainType := d.Get("domain_type").(string)
	resourceRecType := d.Get("resource_record_type").(string)
	resourceRecClass := d.Get("resource_record_class").(string)
	domain := d.Get("domain").(string)
	data := d.Get("data").(string)
	comment := d.Get("comment").(string)
	ttl := d.Get("ttl").(string)
	// deviceRecFlag := d.Get("device_rec_flag").(bool)

	payload := en.NewDnsRRPost(en.IPCDnsRRPost{
		Domain:           domain,
		Owner:            owner,
		DomainType:       domainType,
		ResourceRecType:  resourceRecType,
		Data:             data,
		Comment:          comment,
		TTL:              ttl,
		ResourceRecClass: resourceRecClass,
		// DeviceRecFlag:    deviceRecFlag,
	})

	err := objMgr.CreateDnsRR(payload)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error when creating dns resource record",
			Detail:   fmt.Sprintf("Error when creating dns resource record: (%v)", err),
		})
		return diags
	}

	return getDomainContext(ctx, d, m)
}

func getDomainContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)

	owner := d.Get("owner").(string)
	resourceRecType := d.Get("resource_record_type").(string)
	domain := d.Get("domain").(string)
	data := d.Get("data").(string)

	query := map[string]string{
		"owner":      owner,
		"domainName": domain,
		"type":       resourceRecType,
		"rdata":      data,
	}

	rr, err := objMgr.GetDnsRR(query)

	if err != nil {
		diag := diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error when getting dns resource record",
			Detail:   fmt.Sprintf("Error when getting dns resource record: (%v)", err),
		}
		diags = append(diags, diag)
		return diags
	}

	setIPCDomainResource(d, rr)
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
			Summary:  "Can't format ID",
			Detail:   fmt.Sprintf("Error formatting ID: %v", err),
		})
		return diags
	}

	owner := d.Get("owner").(string)
	domainType := d.Get("domain_type").(string)
	resourceRecType := d.Get("resource_record_type").(string)
	resourceRecClass := d.Get("resource_record_class").(string)
	domain := d.Get("domain").(string)
	dataStr := d.Get("data").(string)
	comment := d.Get("comment").(string)
	ttl := d.Get("ttl").(string)
	// deviceRecFlag := d.Get("device_rec_flag").(bool)

	payload := en.NewDnsRRPost(en.IPCDnsRRPost{
		ID:               id,
		Domain:           domain,
		Owner:            owner,
		DomainType:       domainType,
		ResourceRecType:  resourceRecType,
		Data:             dataStr,
		Comment:          comment,
		TTL:              ttl,
		ResourceRecClass: resourceRecClass,
		// DeviceRecFlag:    deviceRecFlag,
	})

	err = objMgr.UpdateDnsRR(payload)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error when updating DNS resource record",
			Detail:   fmt.Sprintf("Error when updating DNS resource record: %v", err),
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

	owner := d.Get("owner").(string)
	resourceRecType := d.Get("resource_record_type").(string)
	domain := d.Get("domain").(string)
	data := d.Get("data").(string)

	err := objMgr.DeleteDnsRR(owner, domain, resourceRecType, data)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error when deleting dns resource record",
			Detail:   fmt.Sprintf("Error when deleting dns resource record: (%v)", err),
		})
		return diags
	}

	return diags
}
