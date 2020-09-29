from datetime import datetime
import logging
import os
import sys
from typing import List

import requests

logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s [%(levelname)s] %(message)s',
)

required_environment_variables: List[str] = [
    'RANCHER_BEARER_TOKEN',
    'RANCHER_NAMESPACE',
    'RANCHER_PROJECT_ID',
    'RANCHER_URL',
    'RANCHER_WORKLOAD',
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
rancher_namespace = os.environ['RANCHER_NAMESPACE']
rancher_project_id = os.environ['RANCHER_PROJECT_ID']
rancher_url = os.environ['RANCHER_URL']
rancher_workload = os.environ['RANCHER_WORKLOAD']

url = '{rancher_url}/v3/project/{rancher_project_id}/workloads/deployment:{rancher_namespace}:{rancher_workload}'.format(
    rancher_namespace=rancher_namespace,
    rancher_project_id=rancher_project_id,
    rancher_url=rancher_url,
    rancher_workload=rancher_workload,
)

headers = {
    'Authorization': 'Bearer {rancher_bearer_token}'.format(
        rancher_bearer_token=rancher_bearer_token,
    ),
}

workload = requests.get(
    headers={
        **headers
    },
    url=url,
).json()

workload['annotations']['cattle.io/timestamp'] = datetime.now().strftime('%Y-%m-%dT%H:%M:%SZ')

requests.put(
    headers={
        **headers,
    },
    json=workload,
    url=url,
)

logging.info("Workload {rancher_workload} is successfully redeployed.".format(
    rancher_workload=rancher_workload,
))
