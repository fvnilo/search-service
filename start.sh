#!/bin/bash
while ! nc -z $MS_ELASTICHOST $MS_ELASTICPORT
do
  echo "Waiting for Elastic Search on http://$MS_ELASTICHOST:$MS_ELASTICPORT to start..."
  sleep 10; 
done

echo "Service starting..."
./main