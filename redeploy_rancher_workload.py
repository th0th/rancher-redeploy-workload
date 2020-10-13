from datetime import datetime
import logging
import os
import sys
from typing import List

import requests

logging.basicConfig(format='%(asctime)s [%(levelname)s] %(message)s', level=logging.INFO)

required_environment_variables: List[str] = [
    'RANCHER_BEARER_TOKEN',
    'RANCHER_CLUSTER_ID',
    'RANCHER_NAMESPACE',
    'RANCHER_PROJECT_ID',
    'RANCHER_URL',
    'RANCHER_WORKLOADS',
]

missing_environment_variables: List[str] = []

for required_environment_variable in required_environment_variables:
    if required_environment_variable not in os.environ:
        missing_environment_variables.append(required_environment_variable)

if len(missing_environment_variables) > 0:
    logging.error("These environment variables are required but not set: {missing_environment_variables}".format(
        missing_environment_variables=', '.join(missing_environment_variables),
    ))

    sys.exit(1)

rancher_bearer_token = os.environ['RANCHER_BEARER_TOKEN']
rancher_cluster_id = os.environ['RANCHER_CLUSTER_ID']
rancher_namespace = os.environ['RANCHER_NAMESPACE']
rancher_project_id = os.environ['RANCHER_PROJECT_ID']
rancher_url = os.environ['RANCHER_URL']
rancher_workloads = os.environ['RANCHER_WORKLOADS']


def generate_workload_url(r_workload: str) -> str:
    return (
        '{rancher_url}/v3/project/{rancher_cluster_id}:{rancher_project_id}'
        '/workloads/deployment:{rancher_namespace}:{rancher_workload}'
    ).format(
        rancher_cluster_id=rancher_cluster_id,
        rancher_namespace=rancher_namespace,
        rancher_project_id=rancher_project_id,
        rancher_url=rancher_url,
        rancher_workload=r_workload,
    )


headers = {
    'Authorization': 'Bearer {rancher_bearer_token}'.format(
        rancher_bearer_token=rancher_bearer_token,
    ),
}

for rancher_workload in rancher_workloads.split(','):
    url = generate_workload_url(rancher_workload)

    response_get = requests.get(
        headers={
            **headers
        },
        url=url,
    )

    response_get.raise_for_status()
    workload = response_get.json()

    workload['annotations']['cattle.io/timestamp'] = datetime.now().strftime('%Y-%m-%dT%H:%M:%SZ')

    response_put = requests.put(
        headers={
            **headers,
        },
        json=workload,
        url=url,
    )

    response_put.raise_for_status()

    logging.info("Workload {rancher_workload} is successfully redeployed.".format(
        rancher_workload=rancher_workload,
    ))
