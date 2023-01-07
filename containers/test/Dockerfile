FROM sanmaul/agent_pilot:base

RUN mkdir -p /src
RUN mkdir -p /src/bin
COPY apps.exe /src/bin/.

ENTRYPOINT ["/src/bin/apps.exe"]
