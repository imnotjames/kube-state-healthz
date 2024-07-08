# kube-state-healthz

Provides a server that returns `200` when a set of matching deployments have enough `READY` replicas.

## Usage

### `kube-state-healthz serve`

Starts a web server with the following routes:

* `/` - Returns a 200 status code if the targeted deployments are all healthy.
* `/healthz` - Returns a 200 status code if the application can serve requests.  This is recommended to be used for liveness probes.
* `/readyz` - Returns a 200 status code if the application is ready to serve requests.  This is recommended to be used for readiness probes.

## Environment Variables

Most application configuration can be achieved
via environment variables.  Command line arguments take precedence over environment variables.

| Variable Name | Description | Default |
| --- | --- | --- |
| `KSH_NAMESPACE` | Kubernetes namespace to filter to | *none* |
| `KSH_SELECTOR` | Kubernetes label selector to filter to | *none* |
| `KSH_HOST` | Host to bind to when serving with `serve` subcommand | `0.0.0.0` |
| `KSH_PORT` | Port to bind to when serving with `serve` subcommands | `8000` |

## Limited Privileges Environment

If there's no cluster-reader role for security reasons, you can create a service account which has the `RoleBinding` with a `ClusterRole` `view`.
