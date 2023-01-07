FROM sanmaul/agent_pilot:base

RUN apk --no-cache add bind-tools postgresql-client

COPY ./kubectl /usr/local/bin/kubectl
RUN chmod 0755 /usr/local/bin/kubectl

COPY ./.profile /root/.profile
RUN mkdir -p /src
WORKDIR /src

CMD ["sleep", "infinity"]
