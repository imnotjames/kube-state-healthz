package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	v12 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"

	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

var KubeconfigPath string
var Selector string
var Namespace string

var rootCmd = &cobra.Command{
	Use:   "kube-state-healthz",
	Short: "kube-state-healthz determines the health of your cluster",
	Long:  `Check the state of a set of deployments in kubernetes`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	var defaultConfigPath = clientcmd.NewDefaultClientConfigLoadingRules().GetDefaultFilename()

	rootCmd.PersistentFlags().StringVarP(&KubeconfigPath, "kubeconfig", "k", defaultConfigPath, "Absolute path to the kubeconfig file")
	rootCmd.PersistentFlags().StringVarP(&Selector, "selector", "l", "", "Kubernetes Label Selector query to filter on")
	rootCmd.PersistentFlags().StringVarP(&Namespace, "namespace", "n", "", "Kubernetes Namespace to inspect")
}

func getDeploymentsList() (*v12.DeploymentList, error) {
	deploymentsClient, err := getDeploymentsClient()

	if err != nil {
		return nil, err
	}

	selectors, err := labels.Parse(Selector)

	if err != nil {
		return nil, err
	}

	var listOptions = metav1.ListOptions{
		LabelSelector: selectors.String(),
	}

	return deploymentsClient.List(context.TODO(), listOptions)
}

func getDeploymentsClient() (v1.DeploymentInterface, error) {
	clientset, err := getClientset()
	if err != nil {
		return nil, err
	}

	return clientset.AppsV1().Deployments(Namespace), nil
}

func getClientset() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err == nil {
		return kubernetes.NewForConfig(config)
	}

	config, err = clientcmd.BuildConfigFromFlags("", KubeconfigPath)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
