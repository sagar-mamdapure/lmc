package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/logicmonitor/lmc/pkg/config"
	"github.com/logicmonitor/lmc/pkg/conv"
	"github.com/logicmonitor/lmc/pkg/exec"
	"github.com/logicmonitor/lmc/pkg/util"
	"gopkg.in/yaml.v2"
	"io/fs"
	"io/ioutil"
	"strings"
)

var (
	ar         string
	acf        string
	cscr       string
	csccf      string
	UsingFiles = false
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to migrate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		util.SetDebug(HelmConfigObj.DEBUG)
		util.SetQuiet(Quiet)
		execObj := exec.NewProcessExecutor(HelmConfigObj.DEBUG)
		cscConfStr := ""
		argusConfStr := ""
		if UsingFiles {
			util.PrintlnRunning(cmd, fmt.Sprintf("Reading Collectorset Controller Configuration from file %s", csccf))
			cscConfYaml, err := ioutil.ReadFile(csccf)
			if err != nil {
				util.PrintlnFailed(cmd, fmt.Sprintf("Failed to read collectorset controller configuration: %s", err))
				return
			}
			cscConfStr = string(cscConfYaml)
			argusConfYaml, err := ioutil.ReadFile(acf)
			if err != nil {
				util.PrintlnFailed(cmd, fmt.Sprintf("Failed to read collectorset controller configuration: %s", err))
				return
			}
			argusConfStr = string(argusConfYaml)

		} else {
			util.PrintlnRunning(cmd, fmt.Sprintf("Reading Collectorset Controller Configuration from release %s...", cscr))
			var err error
			cscConfStr, err = execObj.RunProcessAndCaptureOutput(HelmConfigObj.BIN, "get", "values", cscr, "-n", HelmConfigObj.NAMESPACE)
			if err != nil {
				util.PrintlnFailed(cmd, fmt.Sprintf("Failed to read Collectorset Controller Configuration from release %s: %s", cscr, err))
				return
			}
			util.PrintlnSuccess(cmd, fmt.Sprintf("Read Collectorset Controller Configuration from release %s", cscr))
			util.PrintlnRunning(cmd, fmt.Sprintf("Reading Argus Configuration from release %s...", ar))
			argusConfStr, err = execObj.RunProcessAndCaptureOutput(HelmConfigObj.BIN, "get", "values", ar, "-n", HelmConfigObj.NAMESPACE)
			if err != nil {
				util.PrintlnFailed(cmd, fmt.Sprintf("Failed to read Argus Configuration from release %s: %s", ar, err))
				return
			}
			util.PrintlnSuccess(cmd, fmt.Sprintf("Read Argus Configuration from release %s", ar))
		}

		util.PrintlnRunning(cmd, fmt.Sprintf("Unmarshalling Collectorset Controller Configuration..."))
		oldCSCConf, err := conv.UnmarshalCscConf(cscConfStr)
		if err != nil {
			util.PrintlnFailed(cmd, fmt.Sprintf("Failed to unmarshal collectorset controller config: %s", err))
			util.PrintlnDebug(cmd, cscConfStr)
			return
		}
		util.PrintlnSuccess(cmd, fmt.Sprintf("Successfully unmarshaled Collectorset Controller Configuration"))
		util.PrintlnDebugConfig(cmd, "Old Collectorset Controller Config: ", oldCSCConf)

		util.PrintlnRunning(cmd, fmt.Sprintf("Converting Collectorset Controller Configuration to new format..."))
		newCscConf := oldCSCConf.ToNewCscConf()
		util.PrintlnSuccess(cmd, fmt.Sprintf("Converted Collectorset Controller Configuration to new format"))

		util.PrintlnRunning(cmd, fmt.Sprintf("Unmarshalling Argus Configuration..."))
		oldArgusConf, err := conv.UnmarshalArgusConf(argusConfStr)
		if err != nil {
			util.PrintlnFailed(cmd, fmt.Sprintf("Failed to unmarshal argus config: %s", err))
			util.PrintlnDebug(cmd, argusConfStr)
			return
		}
		util.PrintlnSuccess(cmd, fmt.Sprintf("Successfully unmarshaled Argus Configuration"))
		util.PrintlnDebugConfig(cmd, "Old argus Config:", oldArgusConf)

		util.PrintlnRunning(cmd, fmt.Sprintf("Converting Argus Configuration to new format..."))
		newArgusConf := oldArgusConf.ToNewArgusConf()
		util.PrintlnSuccess(cmd, fmt.Sprintf("Converted Argus Configuration to new format"))

		util.PrintlnDebugConfig(cmd, "New Collectorset Controller Config:", newCscConf)
		util.PrintlnDebugConfig(cmd, "New Argus Config:", newArgusConf)
		util.PrintlnRunning(cmd, fmt.Sprintf("Merging Collectorset Controller & Argus Configuration to LM Container configuration..."))
		lmcConfig, err := config.Merge(newCscConf, newArgusConf)
		if err != nil {
			util.PrintlnFailed(cmd, fmt.Sprintf("Failed to merge collectorset controller & argus configs into LM Container config: %s", err))
			return
		}
		util.PrintlnSuccess(cmd, fmt.Sprintf("Merged Collectorset Controller & Argus Configuration to LM Container Configuration"))

		util.PrintlnDebugConfig(cmd, "Generated LM Container Config:", lmcConfig)
		var marshal []byte
		if output == Yaml {
			marshal, err = yaml.Marshal(lmcConfig)
			if err != nil {
				util.PrintlnFailed(cmd, fmt.Sprintf("Marshal LM Container Configuration failed: %s", err))
			}
		} else if output == Json {
			marshal, err = json.MarshalIndent(lmcConfig, "", "  ")
			if err != nil {
				util.PrintlnFailed(cmd, fmt.Sprintf("Marshal LM Container Configuration failed: %s", err))
			}
		}

		filePath := FilePath
		if output == Json {
			filePath = strings.Replace(filePath, ".yaml", ".json", -1)
		}
		util.PrintlnRunning(cmd, fmt.Sprintf("Storing generated LM Container configuration to file %s", filePath))
		err = ioutil.WriteFile(filePath, marshal, fs.ModePerm)
		if err != nil {
			util.PrintlnFailed(cmd, fmt.Sprintf("Failed to write generated configuration to file: %s", err))
			return
		}
		util.PrintlnSuccess(cmd, fmt.Sprintf("Stored generated LM Container configuration to file %s", filePath))
	},
}

var output = Yaml

func init() {
	configCmd.AddCommand(migrateCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	migrateCmd.Flags().StringVar(&ar, "argus-release", "argus", "Argus Helm Release Name")
	migrateCmd.Flags().StringVar(&cscr, "csc-release", "collectorset-controller", "Collectorset Controller Helm Release Name")
	migrateCmd.Flags().StringVar(&acf, "argus-conf-file", "argus-configuration.yaml", "Argus Configuration Yaml File Path")
	migrateCmd.Flags().StringVar(&csccf, "csc-conf-file", "collectorset-controller-configuration.yaml", "Collectorset Controller Configuration Yaml File Path")
	//migrateCmd.Flags().StringVarP(&output, "output", "o", "", "Output Format")
	migrateCmd.Flags().VarP(&output, "output", "o", "Output Format")
	migrateCmd.Flags().BoolVar(&UsingFiles, "using-files", false, "Set this flag to migrate configuration using yaml files")
	_ = migrateCmd.RegisterFlagCompletionFunc("output", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"json", "yaml"}, cobra.ShellCompDirectiveDefault
	})
}
