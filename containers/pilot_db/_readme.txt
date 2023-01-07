[postgresql]
psql -h pilot-db-svc.agent-pilot.svc.cluster.local -p 5432 -U agent_pilot

# \dt
# \d+ pilot_deamons
# INSERT INTO pilot_deamons VALUES(DEFAULT, 'podname', '127.0.0.1', 'nodename', 0);
# SELECT * FROM pilot_deamons;
# \q
