// Package kubeclient creates kubernetes client.
package kubeclient

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Options define config for client.
type Options struct {
	DebugLog bool
	Logf     func(format string, v ...any)
}

// New creates kubernetes client.
func New(options Options) (*kubernetes.Clientset, error) {

	const me = "kubeclient.New"

	if options.Logf == nil {
		options.Logf = log.Printf
	}

	debugf := func(format string, v ...any) {
		if options.DebugLog {
			options.Logf(fmt.Sprintf("DEBUG: %s: ", me)+format, v...)
		}
	}

	//
	// 1. Attempt env var KUBECONFIG
	//

	kubeconfig := os.Getenv("KUBECONFIG")
	debugf("KUBECONFIG='%s'", kubeconfig)
	if kubeconfig == "" {
		//
		// 2. Attempt ~/.kube/config
		//
		home, errHome := os.UserHomeDir()
		if errHome != nil {
			debugf("could not get home dir: %v", errHome)
		}
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	config, errKubeconfig := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if errKubeconfig != nil {
		debugf("kubeconfig: %v", errKubeconfig)
		//
		// 3. Attempt in-cluster config
		//
		c, errInCluster := rest.InClusterConfig()
		if errInCluster != nil {
			debugf("in-cluster-config: %v", errInCluster)
		}
		config = c
	}

	if config == nil {
		return nil, errors.New("could not get cluster config")
	}

	clientset, errConfig := kubernetes.NewForConfig(config)

	return clientset, errConfig
}
