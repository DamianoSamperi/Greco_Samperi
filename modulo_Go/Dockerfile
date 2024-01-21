# Usa un'immagine di Go come base
FROM golang:1.17-alpine AS build

# Imposta la directory di lavoro nel container
WORKDIR /app

# Copia il codice sorgente nel container
COPY . .

# Build dell'applicazione Go
RUN go build -o main .

# Secondo stage per una dimensione finale più piccola
FROM alpine:latest

# Copia l'eseguibile dalla fase di build precedente
COPY --from=build /app/main /app/main

# Esponi la porta su cui l'applicazione ascolterà
EXPOSE 8080

# Comando per avviare l'applicazione
CMD ["/app/main"]
