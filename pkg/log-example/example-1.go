package main

import (
	"fmt"
	"github.com/spf13/cobra"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/cli/globalflag"
	"k8s.io/component-base/logs"
	"k8s.io/component-base/term"
	"k8s.io/component-base/version/verflag"
	"k8s.io/klog/v2"
)

// k8s的日志管理
func main() {

	cmd := &cobra.Command{
		Use:   "project-golang",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For gin-example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
		Run: func(cmd *cobra.Command, args []string) {
			verflag.PrintAndExitIfRequested()
			cliflag.PrintFlags(cmd.Flags())

			fmt.Printf("\n\n\nrootrootssssss command in log-example is running\n\n\n")
			klog.InfoS("klog info", "key11", "value1111", "key2222", "value2222")

			klog.ErrorS(fmt.Errorf("klog error"), "error message: ", "error-key", "error-value111")
			klog.V(1).InfoS("klog level-1 debug", "key1111", "level1")

			klog.V(2).InfoS("klog level-2 debug", "level-key", "level2", "message", "一般是http信息、生命周期的信息, pod创建和删除的流程打印")
			klog.V(3).InfoS("klog level-3 debug", "level-key", "level3", "message", "一般是系统日志")

			klog.V(4).InfoS("klog level-4 debug", "level-key", "level4", "message", "一般是debug日志")

			klog.V(7).InfoS("klog level-7 debug", "level-key", "level7")

		},
	}

	//flag.Parse()

	klog.InfoS("klog debug", "key11", "value1111", "key2222", "value2222")

	klog.InfoS("klog debug", "key11", "value1111", "key2222", "value2222")
	klog.Warningf("klog warning, warnning: %s, %s", "warning-key", "warning-value111")
	klog.ErrorS(fmt.Errorf("klog error"), "error message: ", "error-key", "error-value111")
	klog.V(1).InfoS("klog level-1 debug", "key1111", "level1")

	klog.V(2).InfoS("klog level-2 debug", "level-key", "level2", "message", "一般是http信息、生命周期的信息, pod创建和删除的流程打印")
	klog.V(3).InfoS("klog level-3 debug", "level-key", "level3", "message", "一般是系统日志")

	klog.V(4).InfoS("klog level-4 debug", "level-key", "level4", "message", "一般是debug日志")

	klog.V(7).InfoS("klog level-7 debug", "level-key", "level7")

	fs := cmd.Flags()
	fss := cliflag.NamedFlagSets{}
	verflag.AddFlags(fss.FlagSet("global"))
	globalflag.AddGlobalFlags(fss.FlagSet("global"), cmd.Name())

	fmt.Println("fss: ", fss.FlagSets)

	for _, f := range fss.FlagSets {
		fs.AddFlagSet(f)
	}

	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cliflag.SetUsageAndHelpFunc(cmd, fss, cols)

	logs.InitLogs()
	defer logs.FlushLogs()

	cmd.Execute()

}
