package main

import (
	chkp "api_go_sdk/APIFiles"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strings"
)


func resourceNetwork() *schema.Resource {
	return &schema.Resource{
		Create: createNetwork,
		Read:   readNetwork,
		Update: updateNetwork,
		Delete: deleteNetwork,
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: true,
				Description: "Object name. Should be unique in the domain.",
			},
			"subnet": {
				Type: schema.TypeString,
				Optional: true,
				Description: "IPv4 or IPv6 network address. If both addresses are required use subnet4 and subnet6 fields explicitly.",
			},
			"subnet4": {
				Type: schema.TypeString,
				Optional: true,
				Description: "IPv4 network address.",
			},
			"subnet6": {
				Type: schema.TypeString,
				Optional: true,
				Description: "IPv6 network address.",
			},
			"mask_length": {
				Type: schema.TypeInt,
				Optional: true,
				Description: "IPv4 or IPv6 network mask length. If both masks are required use mask-length4 and mask-length6 fields explicitly. Instead of IPv4 mask length it is possible to specify IPv4 mask itself in subnet-mask field.",
			},
			"mask_length4": {
				Type: schema.TypeInt,
				Optional: true,
				Description: "IPv4 network mask length.",
			},
			"mask_length6": {
				Type: schema.TypeInt,
				Optional: true,
				Description: "IPv6 network mask length.",
			},
			"subnet_mask": {
				Type: schema.TypeString,
				Optional: true,
				Description: "IPv4 network mask.",
			},
			"nat_settings" : {
				Type: schema.TypeSet,
				Optional: true,
				Description: "NAT settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_rule": {
							Type:     schema.TypeBool,
							Optional: true,
							Description: "Whether to add automatic address translation rules.",
						},
						"ip_address": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "IPv4 or IPv6 address. If both addresses are required use ipv4-address and ipv6-address fields explicitly. This parameter is not required in case \"method\" parameter is \"hide\" and \"hide-behind\" parameter is \"gateway\"",
						},
						"ipv4_address": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "IPv4 address.",
						},
						"ipv6_address": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "IPv6 address.",
						},
						"hide_behind": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".",
						},
						"install_on": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "Which gateway should apply the NAT translation.",
						},
						"method": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "NAT translation method.",
						},
					},
				},
			},
			"tags": {
				Type: schema.TypeList,
				Optional: true,
				Description: "Collection of tag name.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"broadcast": {
				Type: schema.TypeString,
				Optional: true,
				Description: "Allow broadcast address inclusion.",
			},
			"color": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"details_level": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.",
			},
			"groups": {
				Type: schema.TypeList,
				Optional: true,
				Description: "Collection of group name.",
				Elem: &schema.Schema {
					Type: schema.TypeString,
				},
			},
			"set_if_exists": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If another object with the same identifier already exists, it will be updated. The command behaviour will be the same as if originally a set command was called. Pay attention that original object's fields will be overwritten by the fields provided in the request payload!",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
		},
	}
}

func createNetwork(d *schema.ResourceData, m interface{}) error {
	client := m.(*chkp.ApiClient)
	network := make(map[string]interface{})
	if val, ok := d.GetOk("name"); ok {
		network["name"] = val.(string)
	}
	if val, ok := d.GetOk("subnet"); ok {
		network["subnet"] = val.(string)
	}
	if val, ok := d.GetOk("subnet4"); ok {
		network["subnet4"] = val.(string)
	}
	if val, ok := d.GetOk("subnet6"); ok {
		network["subnet6"] = val.(string)
	}
	if val, ok := d.GetOk("mask_length"); ok {
		network["mask-length"] = val.(int)
	}
	if val, ok := d.GetOk("mask_length4"); ok {
		network["mask-length4"] = val.(int)
	}
	if val, ok := d.GetOk("mask_length6"); ok {
		network["mask-length6"] = val.(int)
	}
	if val, ok := d.GetOk("subnet_mask"); ok {
		network["subnet-mask"] = val.(string)
	}
	if val, ok := d.GetOk("nat_settings"); ok {
		nat := val.(*schema.Set).List()
		if len(nat) > 0 {
			nat := nat[0].(map[string]interface{})
			network["nat-settings"] = expandNatSettings(nat)
		}
	}
	if val, ok := d.GetOk("tags"); ok {
		network["tags"] = val.([]interface{})
	}
	if val, ok := d.GetOk("groups"); ok {
		network["groups"] = val.([]interface{})
	}
	if val, ok := d.GetOk("broadcast"); ok {
		network["broadcast"] = val.(string)
	}
	if val, ok := d.GetOk("comments"); ok {
		network["comments"] = val.(string)
	}
	if val, ok := d.GetOk("set_if_exists"); ok {
		network["set-if-exists"] = val.(bool)
	}
	if val, ok := d.GetOk("color"); ok {
		network["color"] = val.(string)
	}
	if val, ok := d.GetOk("details_level"); ok {
		network["details-level"] = val.(bool)
	}
	if val, ok := d.GetOk("ignore_errors"); ok {
		network["ignore-errors"] = val.(bool)
	}
	if val, ok := d.GetOk("ignore_warnings"); ok {
		network["ignore-warnings"] = val.(bool)
	}
	log.Println("Create Network - Map = ", network)
	addNetworkRes, _ := client.ApiCall("add-network",network,client.GetSessionID(),true,false)
	if !addNetworkRes.Success {
		return fmt.Errorf(addNetworkRes.ErrorMsg)
	}
	d.SetId(addNetworkRes.GetData()["uid"].(string))
	return readNetwork(d, m)
}

func readNetwork(d *schema.ResourceData, m interface{}) error {
	client := m.(*chkp.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}
	showNetworkRes, _ := client.ApiCall("show-network",payload,client.GetSessionID(),true,false)
	if !showNetworkRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showNetworkRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showNetworkRes.ErrorMsg)
	}
	networkJson := showNetworkRes.GetData()
	log.Println("Read Network - Show JSON = ", networkJson)

	if _, ok := d.GetOk("name"); ok {
		_ = d.Set("name", networkJson["name"].(string))
	}

	if _, ok := d.GetOk("subnet"); ok {
		if v := networkJson["subnet4"].(string); v != ""{
			_ = d.Set("subnet", v)
		} else if v := networkJson["subnet6"].(string); v != "" {
			_ = d.Set("subnet", v)
		}
	}
	if _, ok := d.GetOk("subnet4"); ok {
		_ = d.Set("subnet4", networkJson["subnet4"].(string))
	}
	if _, ok := d.GetOk("subnet6"); ok {
		_ = d.Set("subnet6", networkJson["subnet6"].(string))
	}

	if _, ok := d.GetOk("mask_length"); ok {
		if v := int(networkJson["mask-length4"].(float64)); v != 0 {
			_ = d.Set("mask_length", v)
		} else if v := int(networkJson["mask-length6"].(float64)); v != 0 {
			_ = d.Set("mask_length", v)
		}
	}
	if _, ok := d.GetOk("mask_length4"); ok {
		_ = d.Set("mask_length4", int(networkJson["mask-length4"].(float64)))
	}
	if _, ok := d.GetOk("mask_length6"); ok {
		_ = d.Set("mask_length6", int(networkJson["mask-length6"].(float64)))
	}
	if _, ok := d.GetOk("subnet_mask"); ok {
		_ = d.Set("subnet_mask", networkJson["subnet-mask"].(string))
	}

	if _, ok := d.GetOk("broadcast"); ok {
		_ = d.Set("broadcast", networkJson["broadcast"].(string))
	}
	if _, ok := d.GetOk("color"); ok {
		_ = d.Set("color", networkJson["color"].(string))
	}
	if _, ok := d.GetOk("comments"); ok {
		_ = d.Set("comments", networkJson["comments"].(string))
	}
	if v, ok := d.GetOk("nat_settings"); ok {
		v := v.(*schema.Set).List()
		var nat interface{}
		if len(v) > 0 {
			nat = v[0]
		}
		_ = d.Set("nat_settings", flattenNatSettings(networkJson["nat-settings"], nat))
	}
	if _, ok := d.GetOk("groups"); ok {
		groupsJson := networkJson["groups"].([]interface{})
		groupsIds := make([]interface{}, len(groupsJson))
		if len(groupsJson) > 0 {
			// Create slice of group names
			for _, group := range groupsJson {
				group := group.(map[string]interface{})
				groupsIds = append(groupsIds, group["name"].(string))
			}
		}
		_ = d.Set("groups", groupsIds)
	}
	if _, ok := d.GetOk("tags"); ok {
		tagsJson := networkJson["tags"].([]interface{})
		var tagsIds = make([]interface{}, len(tagsJson))
		if len(tagsJson) > 0 {
			// Create slice of tag names
			for _, tag := range tagsJson {
				tag := tag.(map[string]interface{})
				tagsIds = append(tagsIds, tag["name"].(string))
			}
		}
		_ = d.Set("tags", tagsIds)
	}
	return nil
}

func updateNetwork(d *schema.ResourceData, m interface{}) error {
	client := m.(*chkp.ApiClient)
	network := make(map[string]interface{})
	// Name is required
	network["name"] = d.Get("name").(string)
	if d.HasChange("name") {
		oldName , newName := d.GetChange("name")
		network["name"] = oldName.(string)
		network["new-name"] = newName.(string)
	}
	if ok := d.HasChange("subnet"); ok {
		network["subnet"] = d.Get("subnet").(string)
	}
	if ok := d.HasChange("subnet4"); ok {
		network["subnet4"] = d.Get("subnet4").(string)
	}
	if ok := d.HasChange("subnet6"); ok {
		network["subnet6"] = d.Get("subnet6").(string)
	}
	if ok := d.HasChange("mask_length"); ok {
		network["mask-length"] = d.Get("mask_length").(int)
	}
	if ok := d.HasChange("mask_length4"); ok {
		network["mask-length4"] = d.Get("mask_length4").(int)
	}
	if ok := d.HasChange("mask_length6"); ok {
		network["mask-length6"] = d.Get("mask_length6").(int)
	}
	if ok := d.HasChange("subnet_mask"); ok {
		network["subnet-mask"] = d.Get("subnet_mask").(string)
	}
	if ok := d.HasChange("nat_settings"); ok {
		nat := d.Get("nat_settings").(*schema.Set).List()
		if len(nat) > 0 {
			nat := nat[0].(map[string]interface{})
			network["nat-settings"] = expandNatSettings(nat)
		}
	}
	if ok := d.HasChange("tags"); ok {
		network["tags"] = d.Get("tags").([]interface{})
	}
	if ok := d.HasChange("groups"); ok {
		network["groups"] = d.Get("groups").([]interface{})
	}
	if ok := d.HasChange("broadcast"); ok {
		network["broadcast"] = d.Get("broadcast").(string)
	}
	if ok := d.HasChange("comments"); ok {
		network["comments"] = d.Get("comments").(string)
	}
	if ok := d.HasChange("color"); ok {
		network["color"] = d.Get("color").(string)
	}
	if ok := d.HasChange("details_level"); ok {
		network["details-level"] = d.Get("details_level").(bool)
	}
	if ok := d.HasChange("ignore_errors"); ok {
		network["ignore-errors"] = d.Get("ignore_errors").(bool)
	}
	if ok := d.HasChange("ignore_warnings"); ok {
		network["ignore-warnings"] = d.Get("ignore_warnings").(bool)
	}
	log.Println("Update Network - Map = ", network)
	setNetworkRes, _ := client.ApiCall("set-network",network,client.GetSessionID(),true,false)
	if !setNetworkRes.Success {
		return fmt.Errorf(setNetworkRes.ErrorMsg)
	}
	return readNetwork(d, m)
}

func deleteNetwork(d *schema.ResourceData, m interface{}) error {
	client := m.(*chkp.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}
	deleteNetworkRes, _ := client.ApiCall("delete-network",payload,client.GetSessionID(),true,false)
	if !deleteNetworkRes.Success {
		return fmt.Errorf(deleteNetworkRes.ErrorMsg)
	}
	d.SetId("")
	return nil
}

// Call from Create or Update
func expandNatSettings(natSchema map[string]interface{}) interface{} {
	if natSchema == nil {
		return nil
	}
	natSettingsConf := make(map[string]interface{})
	natSettingsConf["auto-rule"] = natSchema["auto_rule"].(bool)
	if v := natSchema["ip_address"].(string); v != "" {
		natSettingsConf["ip-address"] = v
	}
	if v := natSchema["ipv4_address"].(string); v != "" {
		natSettingsConf["ipv4-address"] = v
	}
	if v := natSchema["ipv6_address"].(string); v != ""{
		natSettingsConf["ipv6-address"] = v
	}
	if v := natSchema["hide_behind"].(string); v != "" {
		natSettingsConf["hide-behind"] = v
	}
	if v := natSchema["install_on"].(string); v != "" {
		natSettingsConf["install-on"] = v
	}
	if v := natSchema["method"].(string); v != "" {
		natSettingsConf["method"] = v
	}
	return natSettingsConf
}

// Call from Read
func flattenNatSettings(showNat interface{}, currNat interface{}) interface{} {
	if showNat == nil {
		return nil
	}
	res := make(map[string]interface{})
	for k, v := range showNat.(map[string]interface{}) {
		newKey := strings.ReplaceAll(k,"-","_")
		res[newKey] = v
	}

	// Normalize IP
	if currNat != nil {
		nat := currNat.(map[string]interface{})
		if v, ok := nat["ip_address"]; ok && v.(string) != "" {
			if v, ok := res["ipv4_address"]; ok && v.(string) != "" {
				v := v.(string)
				res["ip_address"] = v
				res["ipv4_address"] = ""
			} else if v, ok := res["ipv6_address"]; ok && v.(string) != "" {
				v := v.(string)
				res["ip_address"] = v
				res["ipv6_address"] = ""
			}
		}
	}
	var natSettingsConf []interface{}
	natSettingsConf = append(natSettingsConf, res)
	return natSettingsConf
}