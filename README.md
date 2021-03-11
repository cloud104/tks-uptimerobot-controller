# TKS UptimeRobot Controller 

## Comands to generated

- `env GOPATH=$HOME/Workspace kubebuilder init --domain tks.sh`
- `env GOPATH=$HOME/Workspace kubebuilder create api --group monitors --version v1 --kind UptimeRobot --controller=true --resource=false`

## TODO

- [] Figure out the the queue stuff
- [] Make tests run on docker
- [x] Rename folder provider
- [] If u remove a host from de CRD it is not deleted, and will prevent is deletion
- [] Fix finalizers
- [] Add contact_id

## Adding alertContact by Monitor

In order to use Alert Contacts on Monitors, you need to have alert contacts added to your account. Once you add them via Dashboard, you will need their ID's. Fetching ID's is not something you can do via UpTime Robot's Dashboard. You will have to use their REST API to fetch alert contacts. To do that, run the following curl command on your terminal with your api key:

curl -d "api_key=your_api_key" -X POST https://api.uptimerobot.com/v2/getAlertContacts
You will get a response similar to what is shown below

```json
[
  {
    "stat": "ok",
    "offset": 0,
    "limit": 50,
    "total": 1,
    "alert_contacts": [
      {
        "id": "1234567",
        "friendly_name": "Operator Weekly",
        "type": 11,
        "status": 2,
        "value": "https://hooks.slack.com/services/RDR1TASY2/VBCG1Y2D/BZjcaM1crRYz2EzDu3Nabx4e"
      }
    ]
  }
]
```

Copy values of id field of your alert contacts which you want to use for TKS Uptime Robot Controller. Specify on spec like this:

```yaml
apiVersion: monitors.tks.sh/v1
kind: UptimeRobot
metadata:
  finalizers:
  - uptimerobot.k8s.io
  name: tks-monitors
  namespace: default
spec:
  hosts:
  - friendlyName: TKS-MONITOR
    url: https://tks.sk
  statusPage:
    friendlyName: TKS-MONITOR
    url: status-google-com.tks.sh
  alertContact:
  - id: "1234567"
    threshold: "1"
    recurrence: "1"
```