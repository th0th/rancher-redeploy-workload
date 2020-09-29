`rancher-redeploy-workload` is a docker image that redeploys a kubernetes workload using Rancher's API.

## Running as a Github action

```yaml
  - name: Update rancher deployment
    uses: th0th/rancher-redeploy-workload@v0.1
    env:
      RANCHER_BEARER_TOKEN: ${{ secrets.RANCHER_BEARER_TOKEN }}
      RANCHER_NAMESPACE: 'namespace'
      RANCHER_PROJECT_ID: 'c-qyxkj:p-hn2z5'
      RANCHER_URL: 'hhttps://rancher.domain.tld'
      RANCHER_WORKLOAD: 'workload'

```shell script
$ docker run --rm -it \
    -e RANCHER_BEARER_TOKEN="token-xgskl:n45p7tmd47t9lfzh7xl8rw6rvtrfzzxrtdr6qvjg27r4sjcxvzss7d" \
    -e RANCHER_NAMESPACE="namespace" \
    -e RANCHER_PROJECT_ID="c-qyxkj:p-hn2z5" \
    -e RANCHER_URL="https://rancher.domain.tld" \
    -e RANCHER_WORKLOAD="workload" \
    rancher-redeploy-workload:latest
```
