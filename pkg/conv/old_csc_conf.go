package conv

import (
	"gopkg.in/yaml.v2"
)

type OldCscConf struct {
	AccessID           string            `yaml:"accessID,omitempty" json:"accessID,omitempty"`
	AccessKey          string            `yaml:"accessKey,omitempty" json:"accessKey,omitempty"`
	Account            string            `yaml:"account,omitempty" json:"account,omitempty"`
	Debug              bool              `yaml:"debug,omitempty" json:"debug,omitempty"`
	EnableRBAC         bool              `yaml:"enableRBAC,omitempty" json:"enableRBAC,omitempty"`
	EtcdDiscoveryToken string            `yaml:"etcdDiscoveryToken,omitempty" json:"etcdDiscoveryToken,omitempty"`
	ImageRepository    string            `yaml:"imageRepository,omitempty" json:"imageRepository,omitempty"`
	ImageTag           string            `yaml:"imageTag,omitempty" json:"imageTag,omitempty"`
	ImagePullPolicy    string            `yaml:"imagePullPolicy,omitempty" json:"imagePullPolicy,omitempty"`
	ProxyURL           string            `yaml:"proxyURL,omitempty" json:"proxyURL,omitempty"`
	ProxyUser          string            `yaml:"proxyUser,omitempty" json:"proxyUser,omitempty"`
	ProxyPass          string            `yaml:"proxyPass,omitempty" json:"proxyPass,omitempty"`
	NodeSelector       map[string]any    `yaml:"nodeSelector,omitempty" json:"nodeSelector,omitempty"`
	Affinity           map[string]any    `yaml:"affinity,omitempty" json:"affinity,omitempty"`
	PriorityClassName  string            `yaml:"priorityClassName,omitempty" json:"priorityClassName,omitempty"`
	Tolerations        []map[string]any  `yaml:"tolerations,omitempty" json:"tolerations,omitempty"`
	Labels             map[string]string `yaml:"labels,omitempty" json:"labels,omitempty"`
	Annotations        map[string]string `yaml:"annotations,omitempty" json:"annotations,omitempty"`
	IgnoreSsl          bool              `yaml:"ignore_ssl,omitempty" json:"ignore_ssl,omitempty"`
}

func (oldConf *OldCscConf) ToNewCscConf() *NewCscConf {
	newCscConf := &NewCscConf{}
	newCscConf.Account = oldConf.Account
	newCscConf.AccessID = oldConf.AccessID
	newCscConf.AccessKey = oldConf.AccessKey

	if oldConf.Debug {
		newCscConf.Log.Level = "debug"
	}
	if oldConf.ImageRepository != "logicmonitor/collectorset-controller" {
		newCscConf.Image.Repository = oldConf.ImageRepository
	}
	newCscConf.Proxy.URL = oldConf.ProxyURL
	newCscConf.Proxy.User = oldConf.ProxyUser
	newCscConf.Proxy.Pass = oldConf.ProxyPass

	newCscConf.NodeSelector = oldConf.NodeSelector
	newCscConf.Affinity = oldConf.Affinity
	newCscConf.PriorityClassName = oldConf.PriorityClassName
	newCscConf.Tolerations = oldConf.Tolerations
	newCscConf.Labels = oldConf.Labels
	newCscConf.Annotations = oldConf.Annotations
	newCscConf.IgnoreSSL = oldConf.IgnoreSsl

	return newCscConf
}

func UnmarshalCscConf(values string) (*OldCscConf, error) {
	conf := &OldCscConf{}
	err := yaml.Unmarshal([]byte(values), conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
