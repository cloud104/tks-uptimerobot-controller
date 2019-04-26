# TKS UptimeRobot Controller 

## Comands to generated

- `env GOPATH=$HOME/Workspace kubebuilder init --domain tks.sh`
- `env GOPATH=$HOME/Workspace kubebuilder create api --group monitors --version v1 --kind UptimeRobot --controller=true --resource=false`

## TODO

- Figure out the the queue stuff
- Make tests run on docker
- Rename folder provider
- If u remove a host from de CRD it is not deleted, and will prevent is deletion
- Fix finalizers
- Add contact_id
