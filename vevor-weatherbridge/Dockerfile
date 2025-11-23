ARG BUILD_FROM

# Stage 1: Install dependencies
# hadolint ignore=DL3006
FROM ${BUILD_FROM} AS builder

WORKDIR /build
COPY requirements.txt .

# Install dependencies to a separate directory for cleaner copying
RUN pip install --no-cache-dir --target=/install -r requirements.txt

# Stage 2: Runtime image
# hadolint ignore=DL3006
FROM ${BUILD_FROM}

# Set shell for pipefail
SHELL ["/bin/bash", "-o", "pipefail", "-c"]

# Copy installed packages from builder stage
# Note: Path depends on Python version in base image
COPY --from=builder /install /usr/local/lib/python3.12/site-packages/

# Set working directory
WORKDIR /app

# Copy application files
COPY weatherstation.py .
COPY run.sh /

# Make run.sh executable
RUN chmod a+x /run.sh

# Expose port for weather station HTTP endpoint
EXPOSE 80

# Health check - verify Flask is responding
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:80/ || exit 1

# Run the startup script
CMD ["/run.sh"]
