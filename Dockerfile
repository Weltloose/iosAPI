FROM golang

COPY  iosAPI/ $GOPATH/src/github.com/Weltloose/iosAPI/

ENV GOPROXY=https://goproxy.cn

WORKDIR $GOPATH/src/github.com/Weltloose/iosAPI

RUN go install

WORKDIR $GOPATH/bin

EXPOSE 8081

CMD ["./testGo"]