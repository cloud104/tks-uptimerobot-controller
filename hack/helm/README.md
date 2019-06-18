# TKS UPTIMEROBOT CONTROLLER

## Install

```
helm upgrade tks-uptimerobot-controller . \
  -f ./values.yaml \
  --set secrets.apiKey="@TODO" \
  --namespace=tks-uptimerobot-controller \
  --debug \
  --install
```
## Delete

```
helm delete --purge tks-uptimerobot-controller
```
