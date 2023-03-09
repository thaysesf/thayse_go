FROM ubuntu:latest

EXPOSE 3000

WORKDIR /app

COPY ./main main

RUN chmod +x main 

CMD ["./main"]
