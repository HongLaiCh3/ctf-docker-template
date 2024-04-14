FROM golang:1.22

EXPOSE 80

ENV GOPROXY=https://goproxy.io

# add user
RUN useradd -m ch3
RUN groupadd -r ctfers
# change group
RUN usermod -aG ctfers ch3
RUN usermod -aG ctfers root


RUN mkdir -p /home/ch3/app && chmod 4777 /home/ch3/app

COPY ./src /home/ch3/app

COPY ./service/start.sh /

RUN chmod 777 /start.sh

WORKDIR /home/ch3/app

RUN chown -R ch3:ctfers /home/ch3


# set flag
RUN echo $GZCTF_FLAG > /home/ch3/app/fake_flag_path
RUN chmod 744 /home/ch3/app/fake_flag_path
RUN chown ch3:ctfers /home/ch3/app/fake_flag_path

# switch to user
USER ch3

# download the 
RUN go mod download && go mod verify

## hot deploy
# RUN go install github.com/pilu/fresh@latest

# run the server
CMD ["/start.sh"]