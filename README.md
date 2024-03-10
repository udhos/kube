[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/udhos/kube/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/udhos/kube)](https://goreportcard.com/report/github.com/udhos/kube)
[![Go Reference](https://pkg.go.dev/badge/github.com/udhos/kube.svg)](https://pkg.go.dev/github.com/udhos/kube)

# kube

**kubeclient** create a kubernetes client.

# Usage

See: [cmd/kubeclient-example/main.go](cmd/kubeclient-example/main.go).

```golang
package main

import (
	"context"
	"log"

	"github.com/udhos/kube/kubeclient"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {

	options := kubeclient.Options{DebugLog: true}

	clientset, errClient := kubeclient.New(options)
	if errClient != nil {
		log.Fatalf("kubeclient error: %v", errClient)
	}

	namespace := ""
	labelSelector := "" // "key1=value1,key2=value2"

	log.Printf("namespace='%s' labelSelector='%s'", namespace, labelSelector)

	listOptions := metav1.ListOptions{LabelSelector: labelSelector}

	pods, errList := clientset.CoreV1().Pods(namespace).List(context.TODO(), listOptions)

	if errList != nil {
		log.Fatalf("list pods: %v", errList)
	}

	log.Printf("found pods: %d", len(pods.Items))

	for i, p := range pods.Items {
		log.Printf("%d/%d: namespace=%s name=%s",
			i, len(pods.Items), p.Namespace, p.Name)
	}
}
```