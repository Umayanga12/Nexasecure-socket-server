# Use the official Redis image (Alpine variant)
FROM redis:7-alpine

# Expose the default Redis port
EXPOSE 6379

# Launch Redis with the password from the env var
# The shell form (sh -c) is needed so $REDIS_PASSWORD is expanded at container start
CMD ["sh", "-c", "redis-server --requirepass \"$REDIS_PASSWORD\""]
