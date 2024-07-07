package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	list, err := getDeploymentsList()

	if err != nil {
		w.WriteHeader(500)
		return
	}

	var healthStatus = true

	for _, d := range list.Items {
		if d.Status.ReadyReplicas < d.Status.Replicas {
			healthStatus = false
		}
	}

	var statusCode int
	
	if healthStatus {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusBadRequest
	}

	w.WriteHeader(statusCode)
}

var Host string
var Port int

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().StringVarP(&Host, "host", "H", "0.0.0.0", "Host address to serve on")
	serveCmd.PersistentFlags().IntVarP(&Port, "port", "p", 8000, "Port to serve on")
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve Health endpoint",
	RunE: func(cmd *cobra.Command, args []string) error {
		http.HandleFunc("/", statusHandler)
		http.HandleFunc("/healthz", healthHandler)
		http.HandleFunc("/readyz", readinessHandler)

		var Listen string = fmt.Sprintf("%s:%d", Host, Port)

		log.Printf("Starting Server on %s\n", Listen)
		if err := http.ListenAndServe(Listen, nil); err != nil {
			return err
		}

		return nil
	},
}
