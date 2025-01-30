# Grab base image
FROM ubuntu AS base

# Copy binary
WORKDIR /app
COPY bin/tm_api /app/tm_api

# Setup and start server
EXPOSE 8317:8317
ENV CORS_ORIGIN="*"
ENV LOG_LEVEL="INFO"
CMD ["/app/tm_api", "--addr", "tm_api:8317", "--cors_origin", "$CORS_ORIGIN", "--log_level", "$LOG_LEVEL"]