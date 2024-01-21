# # Usa un'immagine di Go come base
# FROM golang:1.21.6-alpine AS build


# # Imposta la directory di lavoro nel container
# WORKDIR /app

# # Copia il codice sorgente nel container
# COPY . .

# RUN go install go.mod

# # Build dell'applicazione Go
# RUN go build -o main .

# # Secondo stage per una dimensione finale più piccola
# FROM alpine:latest

# # Copia l'eseguibile dalla fase di build precedente
# COPY --from=build /app/main /app/main

# # Esponi la porta su cui l'applicazione ascolterà
# EXPOSE 8080

# # Comando per avviare l'applicazione
# CMD ["/app/main"]

# #build stage
# FROM golang:alpine AS builder
# RUN apk add --no-cache git
# WORKDIR /go/src/app
# COPY . .
# RUN go get -d -v ./...
# RUN go build -o /go/bin/app -v ./...

# #final stage
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# COPY --from=builder /go/bin/app /app
# ENTRYPOINT /app
# LABEL Name=grecosamperi Version=0.0.1
# EXPOSE 8080

FROM golang:alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
## Add this go mod download command to pull in any dependencies
RUN go mod download
## Our project will now successfully build with the necessary go libraries included.
RUN go build -o main .
## Our start command which kicks off
## our newly created binary executable
CMD ["/app/main"]