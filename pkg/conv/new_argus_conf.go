package conv

type NewArgusConf struct {
	AccessID            string `yaml:"accessID,omitempty" json:"accessID,omitempty"`
	AccessKey           string `yaml:"accessKey,omitempty" json:"accessKey,omitempty"`
	Account             string `yaml:"account,omitempty" json:"account,omitempty"`
	ClusterName         string `yaml:"clusterName,omitempty" json:"clusterName,omitempty"`
	ClusterTreeParentID int    `yaml:"clusterTreeParentID,omitempty" json:"clusterTreeParentID,omitempty"`
	Image               struct {
		Registry   string `yaml:"registry,omitempty" json:"registry,omitempty"`
		Repository string `yaml:"repository,omitempty" json:"repository,omitempty"`
		PullPolicy string `yaml:"pullPolicy,omitempty" json:"pullPolicy,omitempty"`
		Tag        string `yaml:"tag,omitempty" json:"tag,omitempty"`
	} `yaml:"image,omitempty" json:"image,omitempty"`
	NodeSelector struct {
	} `yaml:"nodeSelector,omitempty" json:"nodeSelector,omitempty"`
	Affinity struct {
	} `yaml:"affinity,omitempty" json:"affinity,omitempty"`
	PriorityClassName string        `yaml:"priorityClassName,omitempty" json:"priorityClassName,omitempty"`
	Tolerations       []interface{} `yaml:"tolerations,omitempty" json:"tolerations,omitempty"`
	Resources         struct {
	} `yaml:"resources,omitempty" json:"resources,omitempty"`
	Labels struct {
	} `yaml:"labels,omitempty" json:"labels,omitempty"`
	Annotations struct {
	} `yaml:"annotations,omitempty" json:"annotations,omitempty"`
	Replicas            int `yaml:"replicas,omitempty" json:"replicas,omitempty"`
	ResourceContainerID int `yaml:"resourceContainerID,omitempty" json:"resourceContainerID,omitempty"`
	Log                 struct {
		Level string `yaml:"level,omitempty" json:"level,omitempty"`
	} `yaml:"log,omitempty" json:"log,omitempty"`
	Collectorsetcontroller struct {
		Address string `yaml:"address,omitempty" json:"address,omitempty"`
		Port    int    `yaml:"port,omitempty" json:"port,omitempty"`
	} `yaml:"collectorsetcontroller,omitempty" json:"collectorsetcontroller,omitempty"`
	Proxy struct {
		URL  string `yaml:"url,omitempty" json:"url,omitempty"`
		User string `yaml:"user,omitempty" json:"user,omitempty"`
		Pass string `yaml:"pass,omitempty" json:"pass,omitempty"`
	} `yaml:"proxy,omitempty" json:"proxy,omitempty"`
	EtcdDiscoveryToken string `yaml:"etcdDiscoveryToken,omitempty" json:"etcdDiscoveryToken,omitempty"`
	IgnoreSSL          bool   `yaml:"ignoreSSL,omitempty" json:"ignoreSSL,omitempty"`
	Daemons            struct {
		LmResourceSweeper struct {
			Interval string `yaml:"interval,omitempty" json:"interval,omitempty"`
		} `yaml:"lmResourceSweeper,omitempty" json:"lmResourceSweeper,omitempty"`
		LmCacheSync struct {
			Interval string `yaml:"interval,omitempty" json:"interval,omitempty"`
		} `yaml:"lmCacheSync,omitempty" json:"lmCacheSync,omitempty"`
		Worker struct {
			PoolSize int `yaml:"poolSize,omitempty" json:"poolSize,omitempty"`
		} `yaml:"worker,omitempty" json:"worker,omitempty"`
		Watcher struct {
			BulkSyncInterval string `yaml:"bulkSyncInterval,omitempty" json:"bulkSyncInterval,omitempty"`
			Runner           struct {
				PoolSize                       int `yaml:"poolSize,omitempty" json:"poolSize,omitempty"`
				BackPressureQueueSizePerRunner int `yaml:"backPressureQueueSizePerRunner,omitempty" json:"backPressureQueueSizePerRunner,omitempty"`
			} `yaml:"runner,omitempty" json:"runner,omitempty"`
			SysIpsWaitTimeout string `yaml:"sysIpsWaitTimeout,omitempty" json:"sysIpsWaitTimeout,omitempty"`
		} `yaml:"watcher,omitempty" json:"watcher,omitempty"`
	} `yaml:"daemons,omitempty" json:"daemons,omitempty"`
	Monitoring struct {
		Disable []string `yaml:"disable,omitempty" json:"disable,omitempty"`
	} `yaml:"monitoring,omitempty" json:"monitoring,omitempty"`
	Lm struct {
		Lmlogs struct {
			K8Sevent struct {
				Enable bool `yaml:"enable,omitempty" json:"enable,omitempty"`
			} `yaml:"k8sevent,omitempty" json:"k8sevent,omitempty"`
			K8Spodlog struct {
				Enable bool `yaml:"enable,omitempty" json:"enable,omitempty"`
			} `yaml:"k8spodlog,omitempty" json:"k8spodlog,omitempty"`
		} `yaml:"lmlogs,omitempty" json:"lmlogs,omitempty"`
		Resource struct {
			GlobalDeleteAfterDuration string `yaml:"globalDeleteAfterDuration,omitempty" json:"globalDeleteAfterDuration,omitempty"`
			Alerting                  struct {
				Disable []string `yaml:"disable,omitempty" json:"disable,omitempty"`
			} `yaml:"alerting,omitempty" json:"alerting,omitempty"`
		} `yaml:"resource,omitempty" json:"resource,omitempty"`
		ResourceGroup struct {
			ExtraProps struct {
				Cluster []struct {
					Name     string `yaml:"name,omitempty" json:"name,omitempty"`
					Value    string `yaml:"value,omitempty" json:"value,omitempty"`
					Override bool   `yaml:"override,omitempty" json:"override,omitempty"`
				} `yaml:"cluster,omitempty" json:"cluster,omitempty"`
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
			} `yaml:"extraProps,omitempty" json:"extraProps,omitempty"`
		} `yaml:"resourceGroup,omitempty" json:"resourceGroup,omitempty"`
	} `yaml:"lm,omitempty" json:"lm,omitempty"`
	Filters     []string `yaml:"filters,omitempty" json:"filters,omitempty"`
	SelfMonitor struct {
		Enable bool `yaml:"enable,omitempty" json:"enable,omitempty"`
		Port   int  `yaml:"port,omitempty" json:"port,omitempty"`
	} `yaml:"selfMonitor,omitempty" json:"selfMonitor,omitempty"`
	Debug struct {
		Profiling struct {
			Enable bool `yaml:"enable,omitempty" json:"enable,omitempty"`
		} `yaml:"profiling,omitempty" json:"profiling,omitempty"`
	} `yaml:"debug,omitempty" json:"debug,omitempty"`
	Collector struct {
		Replicas int    `yaml:"replicas,omitempty" json:"replicas,omitempty"`
		Version  int    `yaml:"version,omitempty" json:"version,omitempty"`
		Size     string `yaml:"size,omitempty" json:"size,omitempty"`
		UseEA    bool   `yaml:"useEA,omitempty" json:"useEA,omitempty"`
		Lm       struct {
			GroupID           int `yaml:"groupID,omitempty" json:"groupID,omitempty"`
			EscalationChainID int `yaml:"escalationChainID,omitempty" json:"escalationChainID,omitempty"`
		} `yaml:"lm,omitempty" json:"lm,omitempty"`
		Image struct {
			Registry   string `yaml:"registry,omitempty" json:"registry,omitempty"`
			Repository string `yaml:"repository,omitempty" json:"repository,omitempty"`
			Tag        string `yaml:"tag,omitempty" json:"tag,omitempty"`
			PullPolicy string `yaml:"pullPolicy,omitempty" json:"pullPolicy,omitempty"`
		} `yaml:"image,omitempty" json:"image,omitempty"`
		Proxy struct {
			URL  string `yaml:"url,omitempty" json:"url,omitempty"`
			User string `yaml:"user,omitempty" json:"user,omitempty"`
			Pass string `yaml:"pass,omitempty" json:"pass,omitempty"`
		} `yaml:"proxy,omitempty" json:"proxy,omitempty"`
		Annotations struct {
		} `yaml:"annotations,omitempty" json:"annotations,omitempty"`
		Labels struct {
		} `yaml:"labels,omitempty" json:"labels,omitempty"`
		StatefulsetSpec struct {
			Template struct {
				Spec struct {
				} `yaml:"spec,omitempty" json:"spec,omitempty"`
			} `yaml:"template,omitempty" json:"template,omitempty"`
		} `yaml:"statefulsetSpec,omitempty" json:"statefulsetSpec,omitempty"`
	} `yaml:"collector,omitempty" json:"collector,omitempty"`
	KubeStateMetrics struct {
		Enabled     bool `yaml:"enabled,omitempty" json:"enabled,omitempty"`
		SelfMonitor struct {
			Enabled       bool `yaml:"enabled,omitempty" json:"enabled,omitempty"`
			TelemetryPort int  `yaml:"telemetryPort,omitempty" json:"telemetryPort,omitempty"`
		} `yaml:"selfMonitor,omitempty" json:"selfMonitor,omitempty"`
		Replicas   int      `yaml:"replicas,omitempty" json:"replicas,omitempty"`
		Collectors []string `yaml:"collectors,omitempty" json:"collectors,omitempty"`
	} `yaml:"kube-state-metrics,omitempty" json:"kube-state-metrics,omitempty"`
	ImagePullSecrets []interface{} `yaml:"imagePullSecrets,omitempty" json:"imagePullSecrets,omitempty"`
	Global           struct {
		AccessID  string `yaml:"accessID,omitempty" json:"accessID,omitempty"`
		AccessKey string `yaml:"accessKey,omitempty" json:"accessKey,omitempty"`
		Account   string `yaml:"account,omitempty" json:"account,omitempty"`
		Proxy     struct {
			URL  string `yaml:"url,omitempty" json:"url,omitempty"`
			User string `yaml:"user,omitempty" json:"user,omitempty"`
			Pass string `yaml:"pass,omitempty" json:"pass,omitempty"`
		} `yaml:"proxy,omitempty" json:"proxy,omitempty"`
		Image struct {
			Registry   string `yaml:"registry,omitempty" json:"registry,omitempty"`
			PullPolicy string `yaml:"pullPolicy,omitempty" json:"pullPolicy,omitempty"`
		} `yaml:"image,omitempty" json:"image,omitempty"`
		CollectorsetServiceNameSuffix string        `yaml:"collectorsetServiceNameSuffix,omitempty" json:"collectorsetServiceNameSuffix,omitempty"`
		ImagePullSecrets              []interface{} `yaml:"imagePullSecrets,omitempty" json:"imagePullSecrets,omitempty"`
	} `yaml:"global,omitempty" json:"global,omitempty"`
	NameOverride     string `yaml:"nameOverride,omitempty" json:"nameOverride,omitempty"`
	FullnameOverride string `yaml:"fullnameOverride,omitempty" json:"fullnameOverride,omitempty"`
	Rbac             *Rbac  `yaml:"rbac,omitempty" json:"rbac,omitempty"`
	ServiceAccount   struct {
		Create bool `yaml:"create,omitempty" json:"create,omitempty"`
	} `yaml:"serviceAccount,omitempty" json:"serviceAccount,omitempty"`
}
type Rbac struct {
	Create *bool `yaml:"create,omitempty" json:"create,omitempty"`
}
