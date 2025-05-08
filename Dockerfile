FROM golang:1.23.1

# COPY source destination
COPY . /bin/myapp

CMD ["/bin/myapp"]

#docker build . -t myapp:latest
#docker run -p 8080:8080 goserver
