FROM golang:1.18.3-stretch

RUN apt update && apt upgrade -y &&\
    apt install -y git \
    make openssh-client

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
# RUN go get -v github.com/cosmtrek/air
# RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
#     && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD air

# RUN go get -v github.com/cosmtrek/air
# ENTRYPOINT ["air"]