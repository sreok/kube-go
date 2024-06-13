package models

import (
	"github.com/sreok/kube-go/configs"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
)

// ClientSet 客户端。需要指定Group、指定Version，然后根据Resource获取
func ClientSet() *kubernetes.Clientset {
	clientSet, err := kubernetes.NewForConfig(GetKubeConfig())
	if err != nil {
		panic(err.Error())
	}
	return clientSet
}

// DiscoveryClient 发现客户端。主要用于发现Kubernetes API Server所支持的资源组、资源版本、资源信息。
// 除此之外，还可以将这些信息存储到本地，用户本地缓存，以减轻对Kubernetes API Server访问的压力。
func DiscoveryClient() *discovery.DiscoveryClient {
	// 新建discoveryClient实例
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(GetKubeConfig())
	if err != nil {
		panic(err.Error())
	}
	return discoveryClient
}

// DynamicClient 动态客户端，可以获取所有crd，只支持JSON
func DynamicClient() *dynamic.DynamicClient {
	dynamicClient, err := dynamic.NewForConfig(GetKubeConfig())
	if err != nil {
		panic(err.Error())
	}
	return dynamicClient
}

// GetKubeConfig 获取kubeconfig
func GetKubeConfig() *rest.Config {
	var config *rest.Config
	_, err := os.Stat(configs.KubeConfig)
	if os.IsNotExist(err) {
		// 文件不存在，走内部集群
		log.Println("In Cluster")
		config, err = rest.InClusterConfig()
	} else {
		// 文件存在，走外部集群
		log.Println("Out Cluster")
		config, err = clientcmd.BuildConfigFromFlags("", configs.KubeConfig)
	}
	if err != nil {
		panic(err.Error())
	}
	return config
}
