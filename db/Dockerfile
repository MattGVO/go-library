FROM postgres:15.1-bullseye

RUN mkdir -p /backup

RUN apt-get update && apt-get -y install apt-transport-https
RUN apt-get update && apt-get install curl -y

EXPOSE 5432
