`rancher-redeploy-workload` is a docker image that redeploys kubernetes workloads using Rancher's API.

## Running 

![Hint](hint.png)

### Running as a Github action

```yaml
  - name: Update rancher deployment
    uses: th0th/rancher-redeploy-workload@v0.9
    env:
      RANCHER_BEARER_TOKEN: ${{ secrets.RANCHER_BEARER_TOKEN }}
      RANCHER_CLUSTER_ID: 'c-qxyky'
      RANCHER_NAMESPACE: 'namespace'
      RANCHER_PROJECT_ID: 'p-hm2z1'
      RANCHER_URL: 'https://rancher.aperturescience.tld'
      RANCHER_WORKLOADS: 'wheatley1,wheatley2'
```

### Running as a docker container

```shell script
$ docker run --rm -it \
    -e RANCHER_BEARER_TOKEN="token-xgskl:n45p7tmd47t9lfzh7xl8rw6rvtrfzzxrtdr6qvjg27r4sjcxvzss7d" \
    -e RANCHER_CLUSTER_ID="c-qxyky" \
    -e RANCHER_NAMESPACE="namespace" \
    -e RANCHER_PROJECT_ID="p-hm2z1" \
    -e RANCHER_URL="https://rancher.aperturescience.tld" \
    -e RANCHER_WORKLOADS="wheatley1,wheatley2" \
    th0th/rancher-redeploy-workload:0.9
```

## Shameless plug

I am an indie hacker and I am running an uptime monitoring and analytics platform called [WebGazer](https://www.webgazer.io). You might want to check it out if you are running an online business and want to notice the incidents before your customers.

## License

Copyright © 2020, Gökhan Sarı. Released under the [MIT License](LICENSE).
