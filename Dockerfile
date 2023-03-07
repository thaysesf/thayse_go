FROM ubuntu:latest

EXPOSE 3000

WORKDIR /

COPY ./main main

CMD ["./main"]