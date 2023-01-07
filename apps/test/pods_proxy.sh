#!/bin/sh
 
K8S_APISERVER=https://${KUBERNETES_SERVICE_HOST}:${KUBERNETES_SERVICE_PORT}
ACCESS_TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)
 
echo $K8S_APISERVER
echo $ACCESS_TOKEN
 
curl --insecure -X GET \
-H "Authorization: Bearer $ACCESS_TOKEN" \
${K8S_APISERVER}/api/v1/namespaces/agent-pilot/services/pilot-test-svc
