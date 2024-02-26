
# use official Golang image

FROM golang:1.22.0-alpine3.19


# set working directory

WORKDIR /app


#Copy the source code
COPY . .

# Download and install dependencies
# RUN go install github.com/codegangsta/gin@latest
RUN go install github.com/cosmtrek/air@latest

ENV PATH="/go/bin:${PATH}"

RUN go get -d -v ./...



# Build the application
RUN go build -o  api .

# Expose port 8080 to the outside world
EXPOSE 8888

# Command to run the executable
# CMD ["./api"]
# CMD ["gin", "run", "--port", "8080", "api"]
CMD ["air", "-c", ".air.toml"]
