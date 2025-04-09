# Use an official Python runtime as a base image
FROM python:3.11-slim

# Install necessary packages
RUN pip install requests

# Set the working directory
WORKDIR /app

# Copy the client script to the container
COPY client.py .

# Run the client
CMD ["python", "client.py"]
