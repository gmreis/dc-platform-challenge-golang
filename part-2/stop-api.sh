# bin/bash

if [[ "$(docker ps -f 'name=url-aggregator-api' --format '{{.ID}}')" != "" ]]; then
  echo "Stopping url-aggregator-api..."
  docker stop url-aggregator-api
  echo "Stoped url-aggregator-api."
else
  echo "No running url-aggregator-api."
fi
