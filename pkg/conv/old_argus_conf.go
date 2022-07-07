package conv

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"strconv"
	"strings"
)

type intOrString int

func (is intOrString) MarshalText() ([]byte, error) {
	return []byte(strconv.Itoa(int(is))), nil
}

func (is *intOrString) UnmarshalYAML(
	unmarshal func(interface{}) error,
) error {
	var i int
	if err := unmarshal(&i); err == nil {
		*is = intOrString(i)
		return nil
	}
	var st string
	err := unmarshal(&st)
	if err == nil {
		i, err = strconv.Atoi(strings.Trim(strings.Trim(strings.TrimSpace(st), `"`), "'"))
		*is = intOrString(i)
		return nil
	}
	return err
}

type OldArgusConf struct {
	AccessID              string         `yaml:"accessID,omitempty" json:"accessID,omitempty"`
	AccessKey             string         `yaml:"accessKey,omitempty" json:"accessKey,omitempty"`
	Account               string         `yaml:"account,omitempty" json:"account,omitempty"`
	ClusterName           string         `yaml:"clusterName,omitempty" json:"clusterName,omitempty"`
	LogLevel              string         `yaml:"logLevel,omitempty" json:"logLevel,omitempty"`
	DeleteDevices         bool           `yaml:"deleteDevices,omitempty" json:"deleteDevices,omitempty"`
	DisableAlerting       bool           `yaml:"disableAlerting,omitempty" json:"disableAlerting,omitempty"`
	EnableRBAC            bool           `yaml:"enableRBAC,omitempty" json:"enableRBAC,omitempty"`
	ClusterGroupID        int            `yaml:"clusterGroupID,omitempty" json:"clusterGroupID,omitempty"`
	ResourceGroupID       int            `yaml:"resourceGroupID,omitempty" json:"resourceGroupID,omitempty"`
	EtcdDiscoveryToken    string         `yaml:"etcdDiscoveryToken,omitempty" json:"etcdDiscoveryToken,omitempty"`
	ImageRepository       string         `yaml:"imageRepository,omitempty" json:"imageRepository,omitempty"`
	ImageTag              string         `yaml:"imageTag,omitempty" json:"imageTag,omitempty"`
	ImagePullPolicy       string         `yaml:"imagePullPolicy,omitempty" json:"imagePullPolicy,omitempty"`
	ProxyURL              string         `yaml:"proxyURL,omitempty" json:"proxyURL,omitempty"`
	ProxyUser             string         `yaml:"proxyUser,omitempty" json:"proxyUser,omitempty"`
	ProxyPass             string         `yaml:"proxyPass,omitempty" json:"proxyPass,omitempty"`
	NodeSelector          map[string]any `yaml:"nodeSelector,omitempty" json:"nodeSelector,omitempty"`
	Affinity              map[string]any `yaml:"affinity,omitempty" json:"affinity,omitempty"`
	PriorityClassName     string         `yaml:"priorityClassName,omitempty" json:"priorityClassName,omitempty"`
	Tolerations           []interface{}  `yaml:"tolerations,omitempty" json:"tolerations,omitempty"`
	Labels                map[string]any `yaml:"labels,omitempty" json:"labels,omitempty"`
	Annotations           map[string]any `yaml:"annotations,omitempty" json:"annotations,omitempty"`
	IgnoreSsl             bool           `yaml:"ignore_ssl,omitempty" json:"ignore_ssl,omitempty"`
	RegisterGenericFilter bool           `yaml:"registerGenericFilter,omitempty" json:"registerGenericFilter,omitempty"`
	AppIntervals          struct {
		PeriodicSyncInterval   string `yaml:"periodic_sync_interval,omitempty" json:"periodic_sync_interval,omitempty"`
		PeriodicDeleteInterval string `yaml:"periodic_delete_interval,omitempty" json:"periodic_delete_interval,omitempty"`
		CacheSyncInterval      string `yaml:"cache_sync_interval,omitempty" json:"cache_sync_interval,omitempty"`
	} `yaml:"app_intervals,omitempty" json:"app_intervals,omitempty"`
	DeviceGroupProps struct {
		Cluster []struct {
			Name     string `yaml:"name,omitempty" json:"name,omitempty"`
			Value    string `yaml:"value,omitempty" json:"value,omitempty"`
			Override bool   `yaml:"override,omitempty" json:"override,omitempty"`
		} `yaml:"cluster,omitempty" json:"cluster,omitempty"`
		Pods []struct {
			Name     string `yaml:"name,omitempty" json:"name,omitempty"`
			Value    string `yaml:"value,omitempty" json:"value,omitempty"`
			Override bool   `yaml:"override,omitempty" json:"override,omitempty"`
		} `yaml:"pods,omitempty" json:"pods,omitempty"`
		Services []struct {
			Name     string `yaml:"name,omitempty" json:"name,omitempty"`
			Value    string `yaml:"value,omitempty" json:"value,omitempty"`
			Override bool   `yaml:"override,omitempty" json:"override,omitempty"`
		} `yaml:"services,omitempty" json:"services,omitempty"`
		Deployments []struct {
			Name     string `yaml:"name,omitempty" json:"name,omitempty"`
			Value    string `yaml:"value,omitempty" json:"value,omitempty"`
			Override bool   `yaml:"override,omitempty" json:"override,omitempty"`
		} `yaml:"deployments,omitempty" json:"deployments,omitempty"`
		Nodes []struct {
			Name     string `yaml:"name,omitempty" json:"name,omitempty"`
			Value    string `yaml:"value,omitempty" json:"value,omitempty"`
			Override bool   `yaml:"override,omitempty" json:"override,omitempty"`
		} `yaml:"nodes,omitempty" json:"nodes,omitempty"`
		Etcd []struct {
			Name     string `yaml:"name,omitempty" json:"name,omitempty"`
			Value    string `yaml:"value,omitempty" json:"value,omitempty"`
			Override bool   `yaml:"override,omitempty" json:"override,omitempty"`
		} `yaml:"etcd,omitempty" json:"etcd,omitempty"`
		Hpas []struct {
			Name     string `yaml:"name,omitempty" json:"name,omitempty"`
			Value    string `yaml:"value,omitempty" json:"value,omitempty"`
			Override bool   `yaml:"override,omitempty" json:"override,omitempty"`
		} `yaml:"hpas,omitempty" json:"hpas,omitempty"`
	} `yaml:"device_group_props,omitempty" json:"device_group_props,omitempty"`
	Filters     map[string]any `yaml:"filters,omitempty" json:"filters,omitempty"`
	Openmetrics struct {
		Port int `yaml:"port,omitempty" json:"port,omitempty"`
	} `yaml:"openmetrics,omitempty" json:"openmetrics,omitempty"`
	Collector struct {
		Replicas          intOrString    `yaml:"replicas,omitempty" json:"replicas,omitempty"`
		Size              string         `yaml:"size,omitempty" json:"size,omitempty"`
		ImageRepository   string         `yaml:"imageRepository,omitempty" json:"imageRepository,omitempty"`
		ImageTag          string         `yaml:"imageTag,omitempty" json:"imageTag,omitempty"`
		ImagePullPolicy   string         `yaml:"imagePullPolicy,omitempty" json:"imagePullPolicy,omitempty"`
		SecretName        string         `yaml:"secretName,omitempty" json:"secretName,omitempty"`
		GroupID           int            `yaml:"groupID,omitempty" json:"groupID,omitempty"`
		EscalationChainID intOrString    `yaml:"escalationChainID,omitempty" json:"escalationChainID,omitempty"`
		CollectorVersion  int            `yaml:"collectorVersion,omitempty" json:"collectorVersion,omitempty"`
		UseEA             bool           `yaml:"useEA,omitempty" json:"useEA,omitempty"`
		ProxyURL          string         `yaml:"proxyURL,omitempty" json:"proxyURL,omitempty"`
		ProxyUser         string         `yaml:"proxyUser,omitempty" json:"proxyUser,omitempty"`
		ProxyPass         string         `yaml:"proxyPass,omitempty" json:"proxyPass,omitempty"`
		Annotations       map[string]any `yaml:"annotations,omitempty" json:"annotations,omitempty"`
		Labels            map[string]any `yaml:"labels,omitempty" json:"labels,omitempty"`
		Statefulsetspec   struct {
			Template struct {
				Spec map[string]any `yaml:"spec,omitempty" json:"spec,omitempty"`
			} `yaml:"template,omitempty" json:"template,omitempty"`
		} `yaml:"statefulsetspec,omitempty" json:"statefulsetspec,omitempty"`
	} `yaml:"collector,omitempty" json:"collector,omitempty"`
	DisableResourceMonitoring []string `yaml:"disableResourceMonitoring,omitempty" json:"disableResourceMonitoring,omitempty"`
	DisableResourceAlerting   []string `yaml:"disableResourceAlerting,omitempty" json:"disableResourceAlerting,omitempty"`
	Replicas                  int      `yaml:"replicas,omitempty" json:"replicas,omitempty"`
	KubeStateMetrics          struct {
		Enabled     bool `yaml:"enabled,omitempty" json:"enabled,omitempty"`
		Replicas    int  `yaml:"replicas,omitempty" json:"replicas,omitempty"`
		SelfMonitor struct {
			Enabled       bool `yaml:"enabled,omitempty" json:"enabled,omitempty"`
			TelemetryPort int  `yaml:"telemetryPort,omitempty" json:"telemetryPort,omitempty"`
		} `yaml:"selfMonitor,omitempty" json:"selfMonitor,omitempty"`
		Collectors []string `yaml:"collectors,omitempty" json:"collectors,omitempty"`
	} `yaml:"kube-state-metrics,omitempty" json:"kube-state-metrics,omitempty"`
	Resources map[string]any `yaml:"resources,omitempty" json:"resources,omitempty"`
}

func (oldConf *OldArgusConf) ToNewArgusConf() *NewArgusConf {
	newConf := &NewArgusConf{}
	newConf.AccessKey = oldConf.AccessKey
	newConf.AccessID = oldConf.AccessID
	newConf.Account = oldConf.Account

	newConf.ClusterName = oldConf.ClusterName

	if oldConf.ClusterGroupID > 1 {
		newConf.ClusterTreeParentID = oldConf.ClusterGroupID
	}

	if oldConf.ImageRepository != "logicmonitor/argus" {
		newConf.Image.Repository = oldConf.ImageRepository
	}

	if !oldConf.EnableRBAC {
		f := false
		newConf.Rbac = &Rbac{Create: &f}
	}

	newConf.NodeSelector = oldConf.NodeSelector
	newConf.Affinity = oldConf.Affinity
	newConf.PriorityClassName = oldConf.PriorityClassName
	newConf.Tolerations = oldConf.Tolerations
	newConf.Resources = oldConf.Resources
	newConf.Labels = oldConf.Labels
	newConf.Annotations = oldConf.Annotations
	if oldConf.ResourceGroupID > 1 {
		newConf.ResourceContainerID = oldConf.ResourceGroupID
	}

	if strings.ToLower(oldConf.LogLevel) != "info" {
		newConf.Log.Level = oldConf.LogLevel
	}

	newConf.Proxy.URL = oldConf.ProxyURL
	newConf.Proxy.User = oldConf.ProxyUser
	newConf.Proxy.Pass = oldConf.ProxyPass

	newConf.EtcdDiscoveryToken = oldConf.EtcdDiscoveryToken
	newConf.IgnoreSSL = oldConf.IgnoreSsl
	newConf.Monitoring.Disable = oldConf.DisableResourceMonitoring
	for _, prop := range oldConf.DeviceGroupProps.Cluster {
		if prop.Name == "lmlogs.k8sevent.enable" {
			if prop.Value == "true" {
				newConf.Lm.Lmlogs.K8Sevent.Enable = true
			}
		} else if prop.Name == "kubernetes.resourcedeleteafterduration" {
			newConf.Lm.Resource.GlobalDeleteAfterDuration = prop.Value
		} else {
			newConf.Lm.ResourceGroup.ExtraProps.Cluster = append(newConf.Lm.ResourceGroup.ExtraProps.Cluster, prop)
		}
	}
	for _, prop := range oldConf.DeviceGroupProps.Pods {
		if prop.Name == "lmlogs.k8spodlog.enable" && prop.Value == "true" {
			newConf.Lm.Lmlogs.K8Spodlog.Enable = true
		}
	}

	newConf.Lm.Resource.Alerting.Disable = oldConf.DisableResourceAlerting
	newConf.Lm.ResourceGroup.ExtraProps.Etcd = oldConf.DeviceGroupProps.Etcd
	newConf.Lm.ResourceGroup.ExtraProps.Nodes = oldConf.DeviceGroupProps.Nodes

	for k, v := range oldConf.Filters {
		if value, ok := v.(string); ok {
			if len(value) > 0 {
				arr := strings.Split(value, " || ")
				for _, r := range arr {
					//newConf.Filters = append(newConf.Filters, "type == \""+k+"\" && "+r)
					newConf.Filters = append(newConf.Filters, fmt.Sprintf("type == \"%s\" && %s", k, r))
				}
			}
		} else if value, ok := v.([]any); ok {
			if len(value) > 0 {
				for _, r := range value {
					newConf.Filters = append(newConf.Filters, fmt.Sprintf("type == \"%s\" && %s", k, r.(string)))
				}
			}
		}
	}

	newConf.Collector.Version = oldConf.Collector.CollectorVersion
	newConf.Collector.Size = oldConf.Collector.Size
	newConf.Collector.UseEA = oldConf.Collector.UseEA
	if oldConf.Collector.GroupID == -1 {
		newConf.Collector.Lm.GroupID = 0
	} else if oldConf.Collector.GroupID != 1 {
		newConf.Collector.Lm.GroupID = oldConf.Collector.GroupID
	}
	if oldConf.Collector.EscalationChainID > 0 {
		newConf.Collector.Lm.EscalationChainID = int(oldConf.Collector.EscalationChainID)
	}
	if oldConf.Collector.ImageRepository != "logicmonitor/collector" {
		newConf.Collector.Image.Repository = oldConf.Collector.ImageRepository
	}
	newConf.Collector.Proxy.URL = oldConf.Collector.ProxyURL
	newConf.Collector.Proxy.User = oldConf.Collector.ProxyUser
	newConf.Collector.Proxy.Pass = oldConf.Collector.ProxyPass

	if oldConf.Collector.Replicas > 1 {
		newConf.Collector.Replicas = int(oldConf.Collector.Replicas)
	}
	newConf.Collector.Annotations = oldConf.Collector.Annotations
	newConf.Collector.Labels = oldConf.Collector.Labels
	newConf.Collector.StatefulsetSpec.Template.Spec = oldConf.Collector.Statefulsetspec.Template.Spec

	return newConf
}

func UnmarshalArgusConf(values string) (*OldArgusConf, error) {
	conf := &OldArgusConf{}
	err := yaml.Unmarshal([]byte(values), conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
