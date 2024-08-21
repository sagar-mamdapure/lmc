package conv

type NewCscConf struct {
	AccessID  string `yaml:"accessID,omitempty" json:"accessID,omitempty"`
	AccessKey string `yaml:"accessKey,omitempty" json:"accessKey,omitempty"`
	Account   string `yaml:"account,omitempty" json:"account,omitempty"`
	Log       struct {
		Level string `yaml:"level,omitempty" json:"level,omitempty"`
	} `yaml:"log,omitempty" json:"log,omitempty"`
	NameOverride     string `yaml:"nameOverride,omitempty" json:"nameOverride,omitempty"`
	FullnameOverride string `yaml:"fullnameOverride,omitempty" json:"fullnameOverride,omitempty"`
	ServiceAccount   struct {
		Create bool `yaml:"create,omitempty" json:"create,omitempty"`
	} `yaml:"serviceAccount,omitempty" json:"serviceAccount,omitempty"`
	Rbac struct {
		Create bool `yaml:"create,omitempty" json:"create,omitempty"`
	} `yaml:"rbac,omitempty" json:"rbac,omitempty"`
	Image struct {
		Registry   string `yaml:"registry,omitempty" json:"registry,omitempty"`
		Repository string `yaml:"repository,omitempty" json:"repository,omitempty"`
		PullPolicy string `yaml:"pullPolicy,omitempty" json:"pullPolicy,omitempty"`
		Tag        string `yaml:"tag,omitempty" json:"tag,omitempty"`
	} `yaml:"image,omitempty" json:"image,omitempty"`
	Proxy struct {
		URL  string `yaml:"url,omitempty" json:"url,omitempty"`
		User string `yaml:"user,omitempty" json:"user,omitempty"`
		Pass string `yaml:"pass,omitempty" json:"pass,omitempty"`
	} `yaml:"proxy,omitempty" json:"proxy,omitempty"`
	NodeSelector      map[string]any    `yaml:"nodeSelector,omitempty" json:"nodeSelector,omitempty"`
	Affinity          map[string]any    `yaml:"affinity,omitempty" json:"affinity,omitempty"`
	PriorityClassName string            `yaml:"priorityClassName,omitempty" json:"priorityClassName,omitempty"`
	Tolerations       []map[string]any  `yaml:"tolerations,omitempty" json:"tolerations,omitempty"`
	Labels            map[string]string `yaml:"labels,omitempty" json:"labels,omitempty"`
	Annotations       map[string]string `yaml:"annotations,omitempty" json:"annotations,omitempty"`
	IgnoreSSL         bool              `yaml:"ignoreSSL,omitempty" json:"ignoreSSL,omitempty"`
	ImagePullSecrets  []string          `yaml:"imagePullSecrets,omitempty" json:"imagePullSecrets,omitempty"`
}
