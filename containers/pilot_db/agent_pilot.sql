CREATE ROLE agentdb WITH PASSWORD 'agentdb' LOGIN CREATEDB REPLICATION;
DROP DATABASE IF EXISTS agentdb;
CREATE DATABASE agentdb OWNER 'agentdb';
\c agentdb agentdb
CREATE TABLE node (
	node_name 	VARCHAR(255) PRIMARY KEY,
	timestamp	VARCHAR(25) NOT NULL,
	pod_info 	TEXT NOT NULL
);