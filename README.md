# Simple Tenancy Aggregator

## Comands to generated

- `env GOPATH=$HOME/Workspace kubebuilder init --domain tks.sh`
- `env GOPATH=$HOME/Workspace kubebuilder create api --group monitors --version v1 --kind UptimeRobot --controller=true --resource=false`

## TODO

- Figure out the the queue stuff
- Make tests run on docker
