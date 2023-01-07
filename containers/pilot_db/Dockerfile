FROM postgres:15.1-alpine

ENV POSTGRES_USER agent_pilot
ENV POSTGRES_PASSWORD passwd
ENV POSTGRES_DB agent_pilot

COPY .profile /root/
COPY agent_pilot.sql /docker-entrypoint-initdb.d/
