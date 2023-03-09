FROM ubuntu:latest

EXPOSE 8000

WORKDIR /

COPY ./main main

RUN chmod +x main 

CMD ["./main"]
