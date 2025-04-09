from flask import Flask

app = Flask(__name__)

@app.route('/')
def hello():
    return "Hello from the IPv4 server!"

if __name__ == '__main__':
    # Running the Flask app on all IP addresses (0.0.0.0) so it can be accessed by other containers
    app.run(host='0.0.0.0', port=5000)
