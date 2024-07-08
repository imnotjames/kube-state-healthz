package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	v12 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
)

var KubeConfigPath string
var Selector string
var Namespace string

var rootCmd = &cobra.Command{
	Use:   "kube-state-healthz",
	Short: "kube-state-healthz determines the health of your cluster",
	Long:  `Check the state of a set of deployments in kubernetes`,
	Run:   func(cmd *cobra.Command, args []string) {},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// This is a nightmare but seems to work?
		// I'm not totally sure what's going on here
		// but it was suggested by spf.

		viper.SetEnvPrefix("ksh")
		viper.AutomaticEnv()

		// Needs to happen before the command runs but after
		// cobra runs parseFlags
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			if f.Changed {
				// Only override the flag of it is
				// set to the default value
				return
			}

			viperValue := viper.Get(f.Name)

			if viperValue != nil {
				strValue, err := cast.ToStringE(viperValue)
				if err == nil {
					err = f.Value.Set(strValue)
					if err != nil {
						log.Printf("err set pflag %s from viper err: %s", f.Name, err1)
					}
				} else {
					log.Printf("%s cast.ToStringE err %s", f.Name, err1)
				}
			}
		})

		viper.BindPFlags(cmd.Flags())
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&KubeConfigPath, "kubeconfig", "k", "", "Absolute path to the kubeconfig file")
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
	var config, err = getKubeConfig()

	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}

func getKubeConfig() (*rest.Config, error) {
	if KubeConfigPath != "" {
		// If the KubeconfigPath has been specified from the command line,
		// always use that path no matter what
		return clientcmd.BuildConfigFromFlags("", KubeConfigPath)
	}

	config, err := rest.InClusterConfig()
	if err == nil {
		// If we're in-cluster then return that config
		return config, nil
	}

	// If nothing else, try to use the config at the default path
	var defaultConfigPath = clientcmd.NewDefaultClientConfigLoadingRules().GetDefaultFilename()
	return clientcmd.BuildConfigFromFlags("", defaultConfigPath)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
