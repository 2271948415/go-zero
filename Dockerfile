FROM ubuntu:21.10
RUN apt update
RUN apt install -y curl
WORKDIR /app
COPY ./ ./
CMD ["./go_zero"]

