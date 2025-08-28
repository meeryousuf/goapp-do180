# Use a multi-stage build to create a small and secure final image.

# ---- Builder Stage ----
# This stage compiles the Go application.
FROM https://registry.access.redhat.com/ubi8/go-toolset:1.19 AS builder

# Set the working directory inside the container.
WORKDIR /src

# Copy the Go source code into the container.
COPY main.go .

# Build the Go application.
# CGO_ENABLED=0 is important for creating a static binary.
# -o /app builds the binary and places it in the /app file.
RUN CGO_ENABLED=0 go build -o /app .

# ---- Runner Stage ----
# This stage creates the final, minimal image.
FROM https://registry.access.redhat.com/ubi8/ubi-minimal:latest

# Copy the compiled binary from the 'builder' stage.
COPY --from=builder /app /app

# Expose the port the application will run on.
EXPOSE 8080

# Set the command to run when the container starts.
CMD ["/app"]
