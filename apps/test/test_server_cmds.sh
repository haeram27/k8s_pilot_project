#!/bin/sh

#SERVER_ADDR=pilot-server.agent-pilot.svc.cluster.local:80
SERVER_ADDR=$( kubectl get svc pilot-server-ext -n agent-pilot -o json | jq -r ".status.loadBalancer.ingress[].ip" )
echo ${SERVER_ADDR}

curl -X POST \
-H 'Content-Type: application/json' \
-d '{"header":{"version": "1","type":"COMMAND_MODE"},"body":{"command_mode":{"mode":"GET_PODS_PROCLIST"}}}' \
http://${SERVER_ADDR}/cmds
