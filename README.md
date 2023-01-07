# k8s_pilot_project  
    
1. Source  
  
1.1. source tree  
├── apps  
│   ├── pilot_agent  
│   ├── pilot_cppm  
│   ├── pilot_server  
│   └── test  
├── containers  
│   ├── agent_pilot.yaml  
│   ├── base  
│   ├── dvlp  
│   ├── pilot_agent  
│   ├── pilot_db  
│   ├── pilot_server  
│   ├── pilot_test  
│   └── test  
  
  
2.1. Namespace 생성 및 API 접근 권한 설정  
$ cd /src/agent_pilot/containers  
$ kubectl apply -f agent_pilot.yaml  
  
  
2.2. Dvlp Container에서 빌드  
'pilot-dvlp' 은 DaemonSet 형식입니다.  
소스 폴더 마운트를 위해 각 노드에 '/src/agent_pilot' 폴더가 있어야 합니다.  
  spec:  
    containers:  
      volumeMounts:  
      - name: pilot-dvlp-src  
        mountPath: /src/agent_pilot  
        readOnly: false  
    volumes:  
    - name: pilot-dvlp-src  
      hostPath:  
        path: /src/agent_pilot  
        type: Directory  
$ cd /src/agent_pilot/containers/dvlp  
   
$ kubectl apply -f pilot_dvlp.yaml  
$ kubectl get pods --all-namespaces  
NAMESPACE     NAME                 READY   STATUS    RESTARTS   AGE  
agent-pilot   pilot-dvlp-bpwgs     1/1     Running   0          82m  
   
$ kubectl exec -it pilot-dvlp-bpwgs -n agent-pilot -- /bin/sh -l  
   
[root@pilot-dvlp-bpwgs:/src]  
# cd agent_pilot/apps/test/  
[root@pilot-dvlp-bpwgs:/src/agent_pilot/apps/test]  
# make  
go build -o apps.exe apps.go  
go build -o srvlookup.exe srvlookup.go  
  
  
2.3. Database Container  
$ cd /src/agent_pilot/containers/pilot_db  
   
$ kubectl apply -f pilot_pv.yaml  
$ kubectl get pv  
NAME          CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS       
pilot-pv0     512Mi      RWO            Retain           Available  
   
$ kubectl apply -f pilot_db.yaml  
$ kubectl get pods --all-namespaces  
NAMESPACE     NAME           READY   STATUS    RESTARTS    AGE  
agent-pilot   pilot-db-0     1/1     Running   0           43m  
$ kubectl get pods --all-namespaces  
NAMESPACE     NAME                READY   STATUS    RESTARTS    AGE  
agent-pilot   pilot-dvlp-bpwgs    1/1     Running   0           90m  
   
$ kubectl exec -it pilot-dvlp-bpwgs -n agent-pilot -- /bin/sh -l  
   
[root@pilot-dvlp-bpwgs:/src]  
# psql -h pilot-db-svc.agent-pilot.svc.cluster.local -p 5432 -U agent_pilot  
Password for user agent_pilot: {passwd}  
   
psql (14.5, server 14.6)  
Type "help" for help.  
   
agent_pilot=# \dt  
agent_pilot=# \d+ pilot_daemons  
agent_pilot=# INSERT INTO pilot_daemons VALUES(DEFAULT, 'podname', '127.0.0.1', 'nodename', 0);  
agent_pilot=# SELECT * FROM pilot_daemons;  
agent_pilot=# \q  
