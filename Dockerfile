#Using latest golang base image  
From golang:latest

#Maintainer
LABEL maintainer="Sai Kiran Ambati <saikiranambati942@gmail.com>"

#setting workdir
RUN mkdir -p /usr/local/go/src/Router
WORKDIR /usr/local/go/src/Router

#Copying files
COPY . /usr/local/go/src/Router

#Build app
RUN go build -o main .

#Ports
EXPOSE 8080

#Run Executable
CMD ["./main"]
