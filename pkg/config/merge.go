package config

import (
	"github.com/logicmonitor/lmc/pkg/conv"
)

type Global struct {
	AccessID  string `yaml:"accessID,omitempty" json:"accessID,omitempty"`
	AccessKey string `yaml:"accessKey,omitempty" json:"accessKey,omitempty"`
	Account   string `yaml:"account,omitempty" json:"account,omitempty"`
	Image     struct {
		Registry   string `yaml:"registry,omitempty" json:"registry,omitempty"`
		PullPolicy string `yaml:"pullPolicy,omitempty" json:"pullPolicy,omitempty"`
	} `yaml:"image,omitempty" json:"image,omitempty"`
	Proxy struct {
		URL  string `yaml:"url,omitempty" json:"url,omitempty"`
		User string `yaml:"user,omitempty" json:"user,omitempty"`
		Pass string `yaml:"pass,omitempty" json:"pass,omitempty"`
	} `yaml:"proxy,omitempty" json:"proxy,omitempty"`
}

type LMCConf struct {
	Argus                  *conv.NewArgusConf `yaml:"argus,omitempty" json:"argus,omitempty"`
	CollectorsetController *conv.NewCscConf   `yaml:"collectorset-controller,omitempty" json:"collectorset-controller,omitempty"`
	Global                 *Global            `yaml:"global,omitempty" json:"global,omitempty"`
}

func (lmc *LMCConf) combine() {
	if lmc.Global == nil {
		lmc.Global = &Global{}
	}
	if lmc.CollectorsetController != nil && lmc.Argus != nil {
		if lmc.CollectorsetController.Account == lmc.Argus.Account {
			lmc.Global.Account = lmc.Argus.Account
			lmc.CollectorsetController.Account = ""
			lmc.Argus.Account = ""
		}
		if lmc.CollectorsetController.AccessID == lmc.Argus.AccessID {
			lmc.Global.AccessID = lmc.Argus.AccessID
			lmc.CollectorsetController.AccessID = ""
			lmc.Argus.AccessID = ""
		}
		if lmc.CollectorsetController.AccessKey == lmc.Argus.AccessKey {
			lmc.Global.AccessKey = lmc.Argus.AccessKey
			lmc.CollectorsetController.AccessKey = ""
			lmc.Argus.AccessKey = ""
		}
		if lmc.CollectorsetController.Image.Registry == lmc.Argus.Image.Registry {
			lmc.Global.Image.Registry = lmc.CollectorsetController.Image.Registry
			lmc.CollectorsetController.Image.Registry = ""
			lmc.Argus.Image.Registry = ""
		}
		if lmc.CollectorsetController.Image.PullPolicy == lmc.Argus.Image.PullPolicy {
			lmc.Global.Image.PullPolicy = lmc.CollectorsetController.Image.PullPolicy
			lmc.CollectorsetController.Image.PullPolicy = ""
			lmc.Argus.Image.PullPolicy = ""
		}
		if lmc.CollectorsetController.Proxy.URL == lmc.Argus.Proxy.URL {
			lmc.Global.Proxy.URL = lmc.CollectorsetController.Proxy.URL
			lmc.CollectorsetController.Proxy.URL = ""
			lmc.Argus.Proxy.URL = ""
		}
		if lmc.CollectorsetController.Proxy.User == lmc.Argus.Proxy.User {
			lmc.Global.Proxy.User = lmc.CollectorsetController.Proxy.User
			lmc.CollectorsetController.Proxy.User = ""
			lmc.Argus.Proxy.User = ""
		}
		if lmc.CollectorsetController.Proxy.Pass == lmc.Argus.Proxy.Pass {
			lmc.Global.Proxy.Pass = lmc.CollectorsetController.Proxy.Pass
			lmc.CollectorsetController.Proxy.Pass = ""
			lmc.Argus.Proxy.Pass = ""
		}

	}
}

func Merge(csc *conv.NewCscConf, argus *conv.NewArgusConf) (*LMCConf, error) {
	lmc := &LMCConf{}
	//reflect.Copy()
	lmc.CollectorsetController = csc
	lmc.Argus = argus
	lmc.combine()
	return lmc, nil
}
