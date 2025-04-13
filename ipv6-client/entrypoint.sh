#!/bin/sh

echo "[client] Waiting for bridge to start..."

# Wait until /bridge/ready file exists
while [ ! -f /bridge/ready ]; do
  sleep 1
done

echo "[client] Bridge is ready! Starting full domain ping tests..."

# Ping all domains one by one, don't stop if one fails
while read -r domain; do
  if [ -n "$domain" ]; then
    echo "[client] Pinging $domain..."
    ping6 -c 2 "$domain" || echo "[client] ⚠️ Failed to ping $domain"
    echo ""
  fi
done < /top-domains.txt

echo "[client] All domain pings attempted. Client will now idle."
tail -f /dev/null
