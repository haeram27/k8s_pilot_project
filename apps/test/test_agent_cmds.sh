#!/bin/sh

AGENT_ADDR=$(./srvlookup.exe | head -n 1)
echo ${AGENT_ADDR}
#AGENT_ADDR=127.0.0.1:8080

curl -X POST \
-H 'Content-Type: application/json' \
-d '{"header":{"version": "1","type":"COMMAND_MODE"},"body":{"command_mode":{"mode":"GET_PODS_PROCLIST"}}}' \
http://${AGENT_ADDR}/cmds
