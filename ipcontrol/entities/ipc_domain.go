package entities

type IPCDomain struct {
	ObjBase
	SerialNumber        int      `json:"serialNumber"`
	DomainType          string   `json:"domainType"`
	Description         string   `json:"description"`
	Refresh             string   `json:"refresh"`
	Derivative          string   `json:"derivative"`
	SerialFormat        string   `json:"serialformat"`
	Reverse             bool     `json:"reverse"`
	InfoTemplate        string   `json:"infoTemplate"`
	UserDefinedFields   []string `json:"userDefinedFields"`
	RRTypeInfoTemplates []string `json:"rrTypeInfoTemplates"`
	Managed             bool     `json:"managed"`
	NegativeCacheTTL    string   `json:"negativeCacheTTL"`
	Contact             string   `json:"contact"`
	DomainName          string   `json:"domainName"`
	Expire              string   `json:"expire"`
	DefaultTTL          string   `json:"defaultTTL"`
	Delegated           bool     `json:"delegated"`
	LocalRpz            bool     `json:"localRpz"`
	ID                  int      `json:"id"`
	TemplateDomain      string   `json:"templateDomain"`
	Retry               string   `json:"retry"`
}

/*
 * Domain object constructor
 */
func NewDomain(sb IPCDomain) *IPCDomain {
	res := sb
	res.objectType = "ipc_domain"
	return &res
}

type IPCDomainPost struct {
	ObjBase             `json:"-"`
	SerialNumber        int      `json:"serialNumber,omitempty"`
	DomainType          string   `json:"domainType,omitempty"`
	Description         string   `json:"description,omitempty"`
	Refresh             string   `json:"refresh,omitempty"`
	Derivative          string   `json:"derivative,omitempty"`
	SerialFormat        string   `json:"serialformat,omitempty"`
	Reverse             bool     `json:"reverse,omitempty"`
	InfoTemplate        string   `json:"infoTemplate,omitempty"`
	UserDefinedFields   []string `json:"userDefinedFields,omitempty"`
	RRTypeInfoTemplates []string `json:"rrTypeInfoTemplates,omitempty"`
	Managed             bool     `json:"managed,omitempty"`
	NegativeCacheTTL    string   `json:"negativeCacheTTL,omitempty"`
	Contact             string   `json:"contact,omitempty"`
	DomainName          string   `json:"domainName"`
	Expire              string   `json:"expire,omitempty"`
	DefaultTTL          string   `json:"defaultTTL,omitempty"`
	Delegated           bool     `json:"delegated,omitempty"`
	LocalRpz            bool     `json:"localRpz,omitempty"`
	ID                  int      `json:"id,omitempty"`
	TemplateDomain      string   `json:"templateDomain,omitempty"`
	Retry               string   `json:"retry,omitempty"`
}

/*
 * Resource record object constructor
 */
func NewDomainPost(sb IPCDomainPost) *IPCDomainPost {
	res := sb
	res.objectType = "ipc_domain"
	return &res
}
