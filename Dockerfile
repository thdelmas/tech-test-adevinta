FROM alpine:latest  

WORKDIR /app

RUN apk add --no-cache libc6-compat

# Copy the Pre-built binary file from the previous stage
COPY ./bin/fizzbuzz-api /app/fizzbuzz-api


# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/app/fizzbuzz-api"]
