# Base image
FROM python:3.9-slim

# Set working directory
WORKDIR /app

# Copy app files
COPY app.py /app/

# Install dependencies
RUN pip install flask redis

# Expose port
EXPOSE 5000

# Run the application
CMD ["python", "app.py"]
