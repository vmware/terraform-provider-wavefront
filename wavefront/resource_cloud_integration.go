package wavefront_plugin

import (
	"fmt"
	"github.com/WavefrontHQ/go-wavefront-management-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strings"
)

type DecodeCloudIntegration func(*schema.ResourceData, *wavefront.CloudIntegration) error
type EncodeCloudIntegration func(*schema.ResourceData, *wavefront.CloudIntegration) error

func decodeAwsIntegration(d *schema.ResourceData, integration *wavefront.CloudIntegration) error {
	baseCredentials := &wavefront.AWSBaseCredentials{
		RoleARN:    d.Get("role_arn").(string),
		ExternalID: d.Get("external_id").(string),
	}
	switch integration.Service {
	case "CLOUDWATCH":
		integration.CloudWatch = &wavefront.CloudWatchConfiguration{
			MetricFilterRegex:     d.Get("metric_filter_regex").(string),
			BaseCredentials:       baseCredentials,
			Namespaces:            decodeTypeListToString(d, "namespaces"),
			VolumeSelectionTags:   decodeTypeMapToStringMap(d, "volume_selection_tags"),
			InstanceSelectionTags: decodeTypeMapToStringMap(d, "instance_selection_tags"),
			PointTagFilterRegex:   d.Get("point_tag_filter_regex").(string),
		}
		break
	case "CLOUDTRAIL":
		integration.CloudTrail = &wavefront.CloudTrailConfiguration{
			Region:          d.Get("region").(string),
			Prefix:          d.Get("prefix").(string),
			BaseCredentials: baseCredentials,
			BucketName:      d.Get("bucket_name").(string),
			FilterRule:      d.Get("filter_rule").(string),
		}
		break
	case "EC2":
		var hostNameTags []string
		if encodedHostNameTags, ok := d.GetOk("hostname_tags"); ok {
			for _, v := range encodedHostNameTags.(*schema.Set).List() {
				hostNameTags = append(hostNameTags, v.(string))
			}
		}
		integration.EC2 = &wavefront.EC2Configuration{
			BaseCredentials: baseCredentials,
			HostNameTags:    hostNameTags,
		}
		break
	default:
		return fmt.Errorf("invalid service, expected one of CLOUDWATCH, CLOUDTRAIL, or EC2. got %s", integration.Service)
	}

	return nil
}

func decodeGcpIntegration(d *schema.ResourceData, integration *wavefront.CloudIntegration) error {
	jsonKey := d.Get("json_key").(string)
	projectId := d.Get("project_id").(string)
	switch integration.Service {
	case "GCP":
		var categories []string
		if encodedCategories, ok := d.GetOk("categories"); ok {
			for _, v := range encodedCategories.([]interface{}) {
				categories = append(categories, v.(string))
			}
		}
		integration.GCP = &wavefront.GCPConfiguration{
			MetricFilterRegex: d.Get("metric_filter_regex").(string),
			ProjectId:         projectId,
			GcpJSONKey:        jsonKey,
			CategoriesToFetch: categories,
		}
		break
	case "GCPBILLING":
		integration.GCPBilling = &wavefront.GCPBillingConfiguration{
			ProjectId:  projectId,
			GcpApiKey:  d.Get("api_key").(string),
			GcpJSONKey: jsonKey,
		}
		break
	default:
		return fmt.Errorf("invalid service, expected one of GCP or GCPBILLING. got %s", integration.Service)
	}
	return nil
}

func decodeNewRelicConfiguration(d *schema.ResourceData, integration *wavefront.CloudIntegration) error {
	if integration.Service != "NEWRELIC" {
		return fmt.Errorf("invalid service, expected NEWRELIC. got %s", integration.Service)
	}

	var newRelicMetricFilters []*wavefront.NewRelicMetricFilters
	if encodedMetricFilters, ok := d.GetOk("metric_filter"); ok {
		for _, v := range encodedMetricFilters.([]interface{}) {
			metricFilter := v.(map[string]interface{})
			newRelicMetricFilters = append(newRelicMetricFilters, &wavefront.NewRelicMetricFilters{
				AppName:           metricFilter["app_name"].(string),
				MetricFilterRegex: metricFilter["metric_filter_regex"].(string),
			})
		}
	}
	integration.NewRelic = &wavefront.NewRelicConfiguration{
		ApiKey:                d.Get("api_key").(string),
		AppFilterRegex:        d.Get("app_filter_regex").(string),
		HostFilterRegex:       d.Get("host_filter_regex").(string),
		NewRelicMetricFilters: newRelicMetricFilters,
	}

	return nil
}

func decodeAppDynamicsConfiguration(d *schema.ResourceData, integration *wavefront.CloudIntegration) error {
	if integration.Service != "APPDYNAMICS" {
		return fmt.Errorf("invalid service, expected APPDYNAMICS. got %s", integration.Service)
	}
	appFilterRegex := make([]string, 0)
	if encodedAppFilterRegex, ok := d.GetOk("app_filter_regex"); ok {
		for _, v := range encodedAppFilterRegex.([]interface{}) {
			appFilterRegex = append(appFilterRegex, v.(string))
		}
	}
	integration.AppDynamics = &wavefront.AppDynamicsConfiguration{
		UserName:                     d.Get("user_name").(string),
		ControllerName:               d.Get("controller_name").(string),
		EncryptedPassword:            d.Get("encrypted_password").(string),
		EnableRollup:                 d.Get("enable_rollup").(bool),
		EnableErrorMetrics:           d.Get("enable_error_metrics").(bool),
		EnableBusinessTrxMetrics:     d.Get("enable_business_trx_metrics").(bool),
		EnableBackendMetrics:         d.Get("enable_backend_metrics").(bool),
		EnableOverallPerfMetrics:     d.Get("enable_overall_perf_metrics").(bool),
		EnableIndividualNodeMetrics:  d.Get("enable_individual_node_metrics").(bool),
		EnableAppInfraMetrics:        d.Get("enable_app_infra_metrics").(bool),
		EnableServiceEndpointMetrics: d.Get("enable_service_endpoint_metrics").(bool),
		AppFilterRegex:               appFilterRegex,
	}

	return nil
}

func decodeAzureIntegration(d *schema.ResourceData, integration *wavefront.CloudIntegration) error {
	azureBaseCredentials := &wavefront.AzureBaseCredentials{
		ClientID:     d.Get("client_id").(string),
		ClientSecret: d.Get("client_secret").(string),
		Tenant:       d.Get("tenant").(string),
	}
	var categoryFilter []string
	if encodedCategoryFilters, ok := d.GetOk("category_filter"); ok {
		for _, v := range encodedCategoryFilters.([]interface{}) {
			categoryFilter = append(categoryFilter, v.(string))
		}
	}

	switch integration.Service {
	case "AZURE":
		var resourceGroupFilter []string
		if encodedResourceGroupFilter, ok := d.GetOk("resource_group_filter"); ok {
			for _, v := range encodedResourceGroupFilter.([]interface{}) {
				resourceGroupFilter = append(resourceGroupFilter, v.(string))
			}
		}
		integration.Azure = &wavefront.AzureConfiguration{
			MetricFilterRegex:   d.Get("metric_filter_regex").(string),
			BaseCredentials:     azureBaseCredentials,
			CategoryFilter:      categoryFilter,
			ResourceGroupFilter: resourceGroupFilter,
		}
		break
	case "AZUREACTIVITYLOG":
		integration.AzureActivityLog = &wavefront.AzureActivityLogConfiguration{
			BaseCredentials: azureBaseCredentials,
			CategoryFilter:  categoryFilter,
		}
		break
	default:
		return fmt.Errorf("invalid service, expected one of AZURE or AZUREACTIVITYLOG. got %s",
			integration.Service)
	}

	return nil
}

func encodeAwsIntegration(d *schema.ResourceData, integration *wavefront.CloudIntegration) error {
	switch integration.Service {
	case "CLOUDWATCH":
		d.Set("metric_filter_regex", integration.CloudWatch.MetricFilterRegex)
		d.Set("instance_selection_tags", integration.CloudWatch.InstanceSelectionTags)
		d.Set("namespaces", integration.CloudWatch.Namespaces)
		d.Set("volume_selection_tags", integration.CloudWatch.VolumeSelectionTags)
		d.Set("instance_selection_tags", integration.CloudWatch.InstanceSelectionTags)
		d.Set("role_arn", integration.CloudWatch.BaseCredentials.RoleARN)
		d.Set("external_id", integration.CloudWatch.BaseCredentials.ExternalID)
		d.Set("point_tag_filter_regex", integration.CloudWatch.PointTagFilterRegex)
		break
	case "CLOUDTRAIL":
		d.Set("region", integration.CloudTrail.Region)
		d.Set("prefix", integration.CloudTrail.Prefix)
		d.Set("bucket_name", integration.CloudTrail.BucketName)
		d.Set("filter_rule", integration.CloudTrail.FilterRule)
		d.Set("role_arn", integration.CloudTrail.BaseCredentials.RoleARN)
		d.Set("external_id", integration.CloudTrail.BaseCredentials.ExternalID)
		break
	case "EC2":
		d.Set("hostname_tags", integration.EC2.HostNameTags)
		d.Set("role_arn", integration.EC2.BaseCredentials.RoleARN)
		d.Set("external_id", integration.EC2.BaseCredentials.ExternalID)
		break
	default:
		return fmt.Errorf("invalid service, expected one of CLOUDWATCH, CLOUDTRAIL, or EC2. got %s", integration.Service)
	}

	return nil
}

func encodeGcpIntegration(d *schema.ResourceData, integration *wavefront.CloudIntegration) error {
	switch integration.Service {
	case "GCP":
		d.Set("project_id", integration.GCP.ProjectId)
		d.Set("metric_filter_regex", integration.GCP.MetricFilterRegex)
		d.Set("categories", integration.GCP.CategoriesToFetch)
		break
	case "GCPBILLING":
		d.Set("project_id", integration.GCPBilling.ProjectId)
		break
	default:
		return fmt.Errorf("invalid service, expected one of GCP or GCPBILLING. got %s", integration.Service)
	}

	return nil
}

func encodeNewRelicConfiguration(d *schema.ResourceData, integration *wavefront.CloudIntegration) error {
	if integration.Service != "NEWRELIC" {
		return fmt.Errorf("invalid service, expected NEWRELIC. got %s", integration.Service)
	}
	d.Set("api_key", integration.NewRelic.ApiKey)
	d.Set("app_filter_regex", integration.NewRelic.AppFilterRegex)
	d.Set("host_filter_regex", integration.NewRelic.HostFilterRegex)
	if len(integration.NewRelic.NewRelicMetricFilters) > 0 {
		var metricFilters []map[string]string
		for _, v := range integration.NewRelic.NewRelicMetricFilters {
			metricFilters = append(metricFilters, map[string]string{
				"app_name":            v.AppName,
				"metric_filter_regex": v.MetricFilterRegex,
			})
		}
		d.Set("metric_filter", metricFilters)
	}
	return nil
}

func encodeAppDynamicsConfiguration(d *schema.ResourceData, integration *wavefront.CloudIntegration) error {
	if integration.Service != "APPDYNAMICS" {
		return fmt.Errorf("invalid service, expected APPDYNAMICS. got %s", integration.Service)
	}
	d.Set("user_name", integration.AppDynamics.UserName)
	d.Set("controller_name", integration.AppDynamics.ControllerName)
	d.Set("enable_rollup", integration.AppDynamics.EnableRollup)
	d.Set("enable_error_metrics", integration.AppDynamics.EnableErrorMetrics)
	d.Set("enable_business_trx_metrics", integration.AppDynamics.EnableBusinessTrxMetrics)
	d.Set("enable_backend_metrics", integration.AppDynamics.EnableBackendMetrics)
	d.Set("enable_overall_perf_metrics", integration.AppDynamics.EnableOverallPerfMetrics)
	d.Set("enable_individual_node_metrics", integration.AppDynamics.EnableIndividualNodeMetrics)
	d.Set("enable_app_infra_metrics", integration.AppDynamics.EnableAppInfraMetrics)
	d.Set("enable_service_endpoint_metrics", integration.AppDynamics.EnableServiceEndpointMetrics)
	d.Set("app_filter_regex", integration.AppDynamics.AppFilterRegex)

	return nil
}

func encodeAzureCloudIntegration(d *schema.ResourceData, integration *wavefront.CloudIntegration) error {
	switch integration.Service {
	case "AZURE":
		d.Set("metric_filter_regex", integration.Azure.MetricFilterRegex)
		d.Set("category_filter", integration.Azure.CategoryFilter)
		d.Set("resource_group_filter", integration.Azure.ResourceGroupFilter)
		d.Set("tenant", integration.Azure.BaseCredentials.Tenant)
		d.Set("client_id", integration.Azure.BaseCredentials.ClientID)
		break
	case "AZUREACTIVITYLOG":
		d.Set("category_filter", integration.AzureActivityLog.CategoryFilter)
		d.Set("tenant", integration.AzureActivityLog.BaseCredentials.Tenant)
		d.Set("client_id", integration.AzureActivityLog.BaseCredentials.ClientID)
		break
	default:
		return fmt.Errorf("invalid service, expected one of AZURE or AZUREACTIVITYLOG. got %s",
			integration.Service)
	}
	return nil
}

func decodeCloudIntegration(integration *wavefront.CloudIntegration, d *schema.ResourceData) error {
	b := map[string]DecodeCloudIntegration{
		"CLOUDWATCH":  decodeAwsIntegration,
		"CLOUDTRAIL":  decodeAwsIntegration,
		"EC2":         decodeAwsIntegration,
		"GCP":         decodeGcpIntegration,
		"GCPBILLING":  decodeGcpIntegration,
		"NEWRELIC":    decodeNewRelicConfiguration,
		"APPDYNAMICS": decodeAppDynamicsConfiguration,
		"TESLA": func(d *schema.ResourceData, integration *wavefront.CloudIntegration) error {
			integration.Tesla = &wavefront.TeslaConfiguration{
				Email:    d.Get("email").(string),
				Password: d.Get("password").(string),
			}
			return nil
		},
		"AZURE":            decodeAzureIntegration,
		"AZUREACTIVITYLOG": decodeAzureIntegration,
	}

	service := d.Get("service").(string)
	if decode, ok := b[service]; ok {
		return decode(d, integration)
	}

	return fmt.Errorf("invalid service \"%s\" specified", service)
}

func encodeCloudIntegration(integration *wavefront.CloudIntegration, d *schema.ResourceData) error {
	b := map[string]EncodeCloudIntegration{
		"CLOUDWATCH":  encodeAwsIntegration,
		"CLOUDTRAIL":  encodeAwsIntegration,
		"EC2":         encodeAwsIntegration,
		"GCP":         encodeGcpIntegration,
		"GCPBILLING":  encodeGcpIntegration,
		"NEWRELIC":    encodeNewRelicConfiguration,
		"APPDYNAMICS": encodeAppDynamicsConfiguration,
		"TESLA": func(d *schema.ResourceData, integration *wavefront.CloudIntegration) error {
			// The API will NEVER return a changed password this is problematic.
			d.Set("email", integration.Tesla.Email)
			return nil
		},
		"AZURE":            encodeAzureCloudIntegration,
		"AZUREACTIVITYLOG": encodeAzureCloudIntegration,
	}

	service := integration.Service
	if encode, ok := b[service]; ok {
		return encode(d, integration)
	}

	return fmt.Errorf("invalid service \"%s\" specified", service)
}

func resourceCloudIntegrationCreate(d *schema.ResourceData, meta interface{}) error {
	cloudIntegrations := meta.(*wavefrontClient).client.CloudIntegrations()

	pointTags := decodeTypeMapToStringMap(d, "additional_tags")
	integration := &wavefront.CloudIntegration{
		Name:                     d.Get("name").(string),
		ForceSave:                d.Get("force_save").(bool),
		Service:                  d.Get("service").(string),
		ServiceRefreshRateInMins: d.Get("service_refresh_rate_in_minutes").(int),
		AdditionalTags:           pointTags,
	}

	// configure the integration based on the service
	err := decodeCloudIntegration(integration, d)
	if err != nil {
		return fmt.Errorf("error binding state to wavefront.CloudIntegration. %s", err)
	}

	wfMutexKV.Lock("cloud_integration_create")
	err = cloudIntegrations.Create(integration)
	wfMutexKV.Unlock("cloud_integration_create")

	if err != nil {
		return fmt.Errorf("error creating Cloud Integration for service %s. got %s", d.Get("service"), err)
	}

	d.SetId(integration.Id)
	return resourceCloudIntegrationRead(d, meta)
}

func resourceCloudIntegrationUpdate(d *schema.ResourceData, meta interface{}) error {
	cloudIntegrations := meta.(*wavefrontClient).client.CloudIntegrations()
	integrations, err := cloudIntegrations.Find([]*wavefront.SearchCondition{
		{
			Key:            "id",
			Value:          d.Id(),
			MatchingMethod: "EXACT",
		},
	})
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		} else {
			return fmt.Errorf("unable to find CloudIntegration with ID %s. %s", d.Id(), err)
		}
	}

	integration := integrations[0]
	integration.Name = d.Get("name").(string)
	integration.ForceSave = d.Get("force_save").(bool)
	integration.Service = d.Get("service").(string)
	integration.ServiceRefreshRateInMins = d.Get("service_refresh_rate_in_minutes").(int)
	if additionalTags := decodeTypeMapToStringMap(d, "additional_tags"); additionalTags != nil {
		integration.AdditionalTags = additionalTags
	}

	// We always have to pull the lastErrorEventOff at a minimum
	integration.LastErrorEvent = nil

	// configure the integration based on the service
	err = decodeCloudIntegration(integration, d)
	if err != nil {
		return fmt.Errorf("error binding state to wavefront.CloudIntegration. %s", err)
	}

	wfMutexKV.Lock("cloud_integration_update")
	err = cloudIntegrations.Update(integration)
	wfMutexKV.Unlock("cloud_integration_update")

	if err != nil {
		return fmt.Errorf("unable to update CloudIntegration with id %s", d.Id())
	}

	return resourceCloudIntegrationRead(d, meta)
}

func resourceCloudIntegrationRead(d *schema.ResourceData, meta interface{}) error {
	cloudIntegrations := meta.(*wavefrontClient).client.CloudIntegrations()
	integrations, err := cloudIntegrations.Find([]*wavefront.SearchCondition{
		{
			Key:            "id",
			Value:          d.Id(),
			MatchingMethod: "EXACT",
		},
	})
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		} else {
			return fmt.Errorf("unable to find CloudIntegration with ID %s. %s", d.Id(), err)
		}
	}

	var integration *wavefront.CloudIntegration
	if len(integrations) <= 0 {
		d.SetId("")
		return nil
	} else {
		integration = integrations[0]
	}
	/*
		You will notice we don't ever set the forceSave and that is because it is ALWAYS returned as false
		We'll use it when we create/update because it is necessary, but we won't set the state on read from it
		otherwise we can end up always showing a diff here
	*/
	d.Set("name", integration.Name)
	d.Set("service", integration.Service)
	d.Set("additional_tags", integration.AdditionalTags)
	d.Set("service_refresh_rate_in_minutes", integration.ServiceRefreshRateInMins)

	return encodeCloudIntegration(integration, d)
}

func resourceCloudIntegrationDelete(d *schema.ResourceData, meta interface{}) error {
	cloudIntegrations := meta.(*wavefrontClient).client.CloudIntegrations()
	integrations, err := cloudIntegrations.Find([]*wavefront.SearchCondition{
		{
			Key:            "id",
			Value:          d.Id(),
			MatchingMethod: "EXACT",
		},
	})
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		} else {
			return fmt.Errorf("unable to find CloudIntegration with ID %s. %s", d.Id(), err)
		}
	}

	integration := integrations[0]
	err = cloudIntegrations.Delete(integration, true)
	if err != nil {
		return fmt.Errorf("error deleting Cloud Integration. %s", err)
	}
	d.SetId("")
	return nil
}

func serviceSchemaDefinition(service string) *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
		ValidateFunc: func(i interface{}, s string) (strings []string, errors []error) {
			if i.(string) != service {
				return nil, []error{fmt.Errorf("invalid service, got %s", i.(string))}
			}
			return nil, nil
		},
		DefaultFunc: func() (i interface{}, err error) {
			return service, nil
		},
	}
}
