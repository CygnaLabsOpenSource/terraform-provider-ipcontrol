package utils

import (
	"log"

	en "terraform-provider-ipcontrol/ipcontrol/entities"
)

/* import domain record */
func (objMgr *ObjectManager) CreateDomain(payload *en.IPCDomainPost) error {

	_, err := objMgr.connector.CreateObject(payload, "ipcimportdomain")

	if err != nil {
		log.Printf("[DEBUG] Error when import domain , err: '%s'\n", err)
		return err
	}

	return nil
}

/* get domain record */
func (objMgr *ObjectManager) GetDomain(query map[string]string) (*en.IPCDomain, error) {
	domain := &en.IPCDomain{}
	queryParams := en.NewQueryParams(query)
	err := objMgr.connector.GetObject(nil, "ipcgetdomain", &domain, queryParams)
	return domain, err
}

/* import domain record */
func (objMgr *ObjectManager) UpdateDomain(payload *en.IPCDomainPost) error {

	_, err := objMgr.connector.CreateObject(payload, "ipcimportdomain")

	if err != nil {
		log.Printf("[DEBUG] Error when import domain , err: '%s'\n", err)
		return err
	}

	return nil
}

/* delete domain by id ref */
func (objMgr *ObjectManager) DeleteDomain(
	domainName string,
	domainType string,
) error {
	sf := map[string]string{
		"domainName": domainName,
		"domainType": domainType,
	}
	query := en.NewQueryParams(sf)
	log.Printf("[DEBUG] Delete Domain %s %s", domainName, domainType)
	_, err := objMgr.connector.DeleteObject(nil, "ipcdeletedomain", query)

	return err
}
