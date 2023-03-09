FROM ubuntu:latest

EXPOSE 3000

WORKDIR /

COPY ./main main

CMD ["chmod +x ./main"]

CMD ["./main"]
