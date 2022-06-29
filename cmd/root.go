package cmd

/*
Copyright © 2022 logicmonitor logicmonitor@gmail.com

*/

import (
	"errors"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"
	"os"
)

var (
	HelmConfigObj = HelmConfig{}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lmc",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var Quiet bool

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.lmc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	err := envconfig.Process("HELM", &HelmConfigObj)
	if err != nil {
		return
	}

	rootCmd.PersistentFlags().BoolVarP(&Quiet, "quiet", "q", false, "Stop informative message on stdout")
	//fmt.Println(HelmConfigObj)
}

type HelmConfig struct {
	DEBUG             bool   `envconfig:"DEBUG"`
	PLUGINS           string `envconfig:"PLUGINS"`
	PLUGIN_NAME       string `envconfig:"PLUGIN_NAME"`
	PLUGIN_DIR        string `envconfig:"PLUGIN_DIR"`
	BIN               string `envconfig:"BIN"`
	REGISTRY_CONFIG   string `envconfig:"REGISTRY_CONFIG"`
	REPOSITORY_CACHE  string `envconfig:"REPOSITORY_CACHE"`
	REPOSITORY_CONFIG string `envconfig:"REPOSITORY_CONFIG"`
	NAMESPACE         string `envconfig:"NAMESPACE"`
	KUBECONTEXT       string `envconfig:"KUBECONTEXT"`
}

type OutputFormat int

const (
	Yaml OutputFormat = iota
	Json
)

func (of *OutputFormat) String() string {
	switch *of {
	case Yaml:
		return "yaml"
	case Json:
		return "json"
	}
	return ""
}

func (of *OutputFormat) Set(v string) error {
	switch v {
	case "json":
		*of = Json
	case "yaml":
		*of = Yaml
	default:
		return errors.New(`must be one of "yaml", or "json". ( Default: yaml ) `)
	}
	return nil
}

// Type is only used in help text
func (of *OutputFormat) Type() string {
	return "OutputFormat"
}
