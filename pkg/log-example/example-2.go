package main

import (
	"flag"
	"fmt"
	"k8s.io/klog/v2"
)

// klog使用
func main2() {
	klog.InitFlags(flag.CommandLine)
	defer klog.Flush()
	flag.Parse()

	klog.InfoS("klog debug", "key11", "value1111", "key2222", "value2222")
	klog.Warningf("klog warning, warnning: %s, %s", "warning-key", "warning-value111")
	klog.ErrorS(fmt.Errorf("klog error"), "error message: ", "error-key", "error-value111")
	klog.V(1).InfoS("klog level-1 debug", "key1111", "level1")

	klog.V(2).InfoS("klog level-2 debug", "level-key", "level2", "message", "一般是http信息、生命周期的信息, pod创建和删除的流程打印")
	klog.V(3).InfoS("klog level-3 debug", "level-key", "level3", "message", "一般是系统日志")

	klog.V(4).InfoS("klog level-4 debug", "level-key", "level4", "message", "一般是debug日志")

	klog.V(7).InfoS("klog level-7 debug", "level-key", "level7")

	var _ int = 12

}
