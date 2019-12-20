package rke

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

const (
	rkeClusterCloudProviderCustomName = "custom"
)

var (
	rkeClusterCloudProviderList = []string{
		rkeClusterCloudProviderAwsName,
		rkeClusterCloudProviderAzureName,
		rkeClusterCloudProviderCustomName,
		rkeClusterCloudProviderOpenstackName,
		rkeClusterCloudProviderVsphereName,
	}
)

//Schemas

func rkeClusterCloudProviderFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"aws_cloud_config": {
			Type:       schema.TypeList,
			MaxItems:   1,
			Optional:   true,
			Deprecated: "Use aws_cloud_provider instead",
			Elem: &schema.Resource{
				Schema: rkeClusterCloudProviderAwsFields(),
			},
		},
		"aws_cloud_provider": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "AWS Cloud Provider config",
			Elem: &schema.Resource{
				Schema: rkeClusterCloudProviderAwsFields(),
			},
		},
		"azure_cloud_config": {
			Type:       schema.TypeList,
			MaxItems:   1,
			Optional:   true,
			Deprecated: "Use azure_cloud_provider instead",
			Elem: &schema.Resource{
				Schema: rkeClusterCloudProviderAzureFields(),
			},
		},
		"azure_cloud_provider": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "Azure Cloud Provider config",
			Elem: &schema.Resource{
				Schema: rkeClusterCloudProviderAzureFields(),
			},
		},
		"custom_cloud_config": {
			Type:       schema.TypeString,
			Optional:   true,
			Deprecated: "Use custom_cloud_provider instead",
		},
		"custom_cloud_provider": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Custom Cloud Provider config",
		},
		"name": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validation.StringInSlice(rkeClusterCloudProviderList, true),
		},
		"openstack_cloud_config": {
			Type:       schema.TypeList,
			MaxItems:   1,
			Optional:   true,
			Deprecated: "Use openstack_cloud_provider instead",
			Elem: &schema.Resource{
				Schema: rkeClusterCloudProviderOpenstackFields(),
			},
		},
		"openstack_cloud_provider": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "Openstack Cloud Provider config",
			Elem: &schema.Resource{
				Schema: rkeClusterCloudProviderOpenstackFields(),
			},
		},
		"vsphere_cloud_config": {
			Type:       schema.TypeList,
			MaxItems:   1,
			Optional:   true,
			Deprecated: "Use vsphere_cloud_provider instead",
			Elem: &schema.Resource{
				Schema: rkeClusterCloudProviderVsphereFields(),
			},
		},
		"vsphere_cloud_provider": {
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Description: "Vsphere Cloud Provider config",
			Elem: &schema.Resource{
				Schema: rkeClusterCloudProviderVsphereFields(),
			},
		},
	}
	return s
}
