package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// CreateK8sClient creates and returns a Kubernetes clientset

func CreateK8sClient() (*kubernetes.Clientset, error) {
	InfoLogger.Println("Creating Kubernetes client")

	// Try to use the in-cluster config first
	config, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		// If that fails, try to use the kubeconfig file
		kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
		if envvar := os.Getenv("KUBECONFIG"); envvar != "" {
			kubeconfig = envvar
		}
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			ErrorLogger.Printf("error building kubeconfig: %v", err)
			return nil, fmt.Errorf("error building kubeconfig: %v", err)
		}
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		ErrorLogger.Printf("error creating kubernetes client: %v", err)
		return nil, fmt.Errorf("error creating kubernetes client: %v", err)
	}

	return clientset, nil
}
