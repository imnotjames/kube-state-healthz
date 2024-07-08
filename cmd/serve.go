package cmd

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/urfave/negroni"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	list, err := getDeploymentsList()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var healthStatus = true

	for _, d := range list.Items {
		if d.Status.ReadyReplicas < d.Status.Replicas {
			healthStatus = false
		}
	}

	if healthStatus {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

type RecoveryMiddleware struct{}

func (rec *RecoveryMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		}
	}()

	next(rw, r)
}

type LoggerMiddleware struct {
	out io.Writer
}

func (l *LoggerMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()

	next(rw, r)

	res := rw.(negroni.ResponseWriter)

	fmt.Fprintf(
		l.out,
		"%s | %d | %16s | %s | %s %s\n",
		start.Format(time.RFC3339),
		res.Status(),
		time.Since(start),
		r.Host,
		r.Method,
		r.URL.Path,
	)
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
		out := cmd.OutOrStdout()

		router := http.NewServeMux()
		router.HandleFunc("/", statusHandler)
		router.HandleFunc("/healthz", healthHandler)
		router.HandleFunc("/readyz", readinessHandler)

		n := negroni.New()
		n.Use(&RecoveryMiddleware{})
		n.Use(&LoggerMiddleware{out: out})
		n.UseHandler(router)

		fmt.Fprintf(out, "%s | Starting server on %s:%d\n", time.Now().Format(time.RFC3339), Host, Port)
		err := http.ListenAndServe(fmt.Sprintf("%s:%d", Host, Port), n)
		if err != nil {
			return err
		}

		return nil
	},
}
