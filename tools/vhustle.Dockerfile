FROM alpine:3.4

RUN apk add --update ca-certificates # Certificates for SSL

# Add the binary
COPY vhustle /bin/vhustle

# Run the axxoncloud command by default when the container starts.
CMD ["/bin/vhustle"]
