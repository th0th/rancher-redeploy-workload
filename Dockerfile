FROM python:3.9

WORKDIR /rancher-redeploy-workload

COPY requirements.txt requirements.txt

RUN pip install -r requirements.txt

COPY redeploy_rancher_workload.py redeploy_rancher_workload.py

# docker-entrypoint.sh
COPY docker-entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/docker-entrypoint.sh

ENTRYPOINT ["docker-entrypoint.sh"]