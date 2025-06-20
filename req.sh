#!/bin/bash

# Define the server URL
SERVER_URL="http://localhost:8080"

# Loop to send 100 requests
for i in $(seq 1 1); do
    echo "Sending request $i"
    hey -n 5 -c 5 -m GET -d '{"Data":10000000}' -H "Content-Type: application/json" http://localhost:8080
    echo # New line for readability
done

echo "Completed sending 100 requests!"
