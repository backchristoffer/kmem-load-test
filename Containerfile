FROM golang:bullseye

WORKDIR /app
COPY . .

RUN go build .
EXPOSE 8080
CMD [ "./klt" ]
