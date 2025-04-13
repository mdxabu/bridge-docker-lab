#!/bin/sh

echo "Starting IPv6 client..."

# Optional: limit to 5-10 domains for testing
head -n 10 /top-domains.txt | while read -r domain; do
  echo "[IPv6 Client] Pinging $domain over IPv6..."
  ping6 -c 2 "$domain"
  echo ""
done

# Keep container alive for bridge interaction
tail -f /dev/null
