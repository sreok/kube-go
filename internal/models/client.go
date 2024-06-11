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
	config, err := clientcmd.BuildConfigFromFlags("", configs.KubeConfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

func KubeClient() *kubernetes.Clientset {
	var clientSet *kubernetes.Clientset
	_, err := os.Stat(configs.KubeConfig)
	if os.IsNotExist(err) { // 文件不存在，走内部集群
		log.Println("In Cluster")
		clientSet = InCluster()
	} else { // 文件存在，走外部集群
		log.Println("Out Cluster")
		clientSet = OutCluster()
	}
	return clientSet
}
