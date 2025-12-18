package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	en "terraform-provider-ipcontrol/ipcontrol/entities"
)

/* CreateSubnet */
func (objMgr *ObjectManager) CreateAddress(addr *en.IPCAddressPost) (string, error) {

	resp, err := objMgr.connector.CreateObject(addr, "ipcadddevice")
	if err != nil {
		return "", err
	}
	log.Println("[DEBUG] Address Resp: " + fmt.Sprintf("%v", resp))

	raw, err := json.Marshal(resp)
	if err != nil {
		log.Printf("[ERROR] marshal failed: %v", err)
		return "", fmt.Errorf("cannot marshal response: %v", err)
	}

	var ipAddress string
	if err := json.Unmarshal(raw, &ipAddress); err == nil {
		trimmed := strings.TrimSpace(ipAddress)
		if strings.HasPrefix(trimmed, "[") {
			var objs []struct {
				IpAddress string `json:"ipAddress"`
			}
			if err := json.Unmarshal([]byte(trimmed), &objs); err == nil {
				if len(objs) > 0 && objs[0].IpAddress != "" {
					return objs[0].IpAddress, nil
				}
			}
		}
		return ipAddress, nil
	}

	return "", fmt.Errorf("cannot extract IP from response: %s", string(raw))
}

/* get address */
func (objMgr *ObjectManager) GetAddress(query map[string]string) (*en.IPCAddress, error) {
	address := &en.IPCAddress{}
	queryParams := en.NewQueryParams(query)
	err := objMgr.connector.GetObject(nil, "ipcgetdevice", &address, queryParams)
	return address, err
}

/* delete address by ip address ref */
func (objMgr *ObjectManager) DeleteAddressRef(address string) (string, error) {
	sf := map[string]string{
		"ipAddress": address,
	}
	query := en.NewQueryParams(sf)
	str, err := objMgr.connector.DeleteObject(en.NewAddress(en.IPCAddress{}), "ipcdeletedevice", query)
	log.Printf("Delete Device Address %s", address)
	return str, err
}

/* UpdateAddress */
func (objMgr *ObjectManager) UpdateAddress(data *en.IPCAddressPost) error {
	// var address []en.IPCAddress
	_, err := objMgr.connector.UpdateObject(data, "ipcmodifydevice")
	if err != nil {
		return err
	}

	return nil

	// err = json.Unmarshal([]byte(resp), &address)

	// if err != nil {
	// 	log.Printf("Update Device Address Cannot unmarshall '%s', err: '%s'\n", string(resp), err)
	// 	return nil, err
	// }

	// return &address[0], err
}
