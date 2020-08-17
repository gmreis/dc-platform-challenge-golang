# bin/bash

if [[ "$(docker images -q url-aggregator-api:latest 2> /dev/null)" == "" ]]; then
  echo "Building url-aggregator-api..."
  docker build -t url-aggregator-api ./url-aggregator-api
  echo "Builded url-aggregator-api."
fi

if [[ "$(docker ps -f 'name=url-aggregator-api' --format '{{.ID}}')" == "" ]]; then
  echo "Starting url-aggregator-api..."
  docker run --rm -p 4567:4567 --name url-aggregator-api url-aggregator-api
  echo "Started url-aggregator-api."
else
  echo "url-aggregator-api is already running."
fi
