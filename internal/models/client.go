package models

import (
	"github.com/sreok/kube-go/configs"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
)

func KubeClient() (*kubernetes.Clientset, error) {
	path := configs.KubeConfig
	var clientSet *kubernetes.Clientset
	_, err := os.Stat(path)
	// 文件存在，走外部集群
	if err == nil {
		log.Println("Out Cluster")
		clientSet = OutCluster()
	}
	// 文件不存在，走内部集群
	if os.IsNotExist(err) {
		log.Println("In Cluster")
		clientSet = InCluster()
	}
	return clientSet, err
}

func InCluster() *kubernetes.Clientset {
	// creates the in-cluster configs
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

func OutCluster() *kubernetes.Clientset {

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", configs.KubeConfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}
