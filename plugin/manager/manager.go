package manager

import "sync"

type PluginManager struct {
	mu      sync.RWMutex
	Plugins map[string]Plugin
}

type Plugin struct {
	Name     string `json:"name"`
	Vsersion string `json:"version"`
	Hash     string `json:"hash"`
	Status   string `json:"status"`
	ExecPath string `json:"execPath"`
}

func (pm *PluginManager) addPlugin() {

}

func (pm *PluginManager) deletePlugin() {}

func (pm *PluginManager) runServer() {

	//listen 端口 ， 插件和 主进程通过grpc通信

	//plugin list
	//1. 系统基本信息的搜集
	//2. 桶和照片的定时同步
	//3. 定时发送邮件
}
