import requests

# Replace with the IPv4 address of the server container
server_url = "http://172.20.0.10:5000"  # Assuming the server container is named ipv4_server

try:
    response = requests.get(server_url)
    if response.status_code == 200:
        print(f"Received from server: {response.text}")
    else:
        print(f"Failed to connect to the server. Status code: {response.status_code}")
except requests.exceptions.RequestException as e:
    print(f"Error connecting to the server: {e}")
