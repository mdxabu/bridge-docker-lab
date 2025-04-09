# Use an official Python runtime as a base image
FROM python:3.11-slim

# Install necessary packages
RUN pip install flask

# Set the working directory
WORKDIR /app

# Copy the Python server script to the container
COPY server.py .

# Expose port 5000 to allow access to the server
EXPOSE 5000

# Run the server
CMD ["python", "server.py"]
