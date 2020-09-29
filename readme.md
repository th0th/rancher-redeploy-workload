`rancher-redeploy-workload` is a docker image that redeploys a kubernetes workload using Rancher's API.

## Running

```shell script
$ docker run --rm -it \
    -e RANCHER_BEARER_TOKEN="token-xgskl:n45p7tmd47t9lfzh7xl8rw6rvtrfzzxrtdr6qvjg27r4sjcxvzss7d" \
    -e RANCHER_NAMESPACE="namespace" \
    -e RANCHER_PROJECT_ID="c-qyxkj:p-hn2z5" \
    -e RANCHER_URL="https://rancher.domain.tld" \
    -e RANCHER_WORKLOAD="workload" \
    rancher-redeploy-workload:latest
```
