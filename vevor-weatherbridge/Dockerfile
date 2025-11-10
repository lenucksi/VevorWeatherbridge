ARG BUILD_FROM
FROM ${BUILD_FROM}

# Set shell
SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# Install Python dependencies
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy application files
COPY weatherstation.py .
COPY run.sh /

# Make run.sh executable
RUN chmod a+x /run.sh

# Expose port for weather station
EXPOSE 80

# Run the startup script
CMD ["/run.sh"]
