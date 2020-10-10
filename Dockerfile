FROM golang:1.14

WORKDIR /gorm_crud/src
# WORKDIR /go/src/app


ADD src /gorm_crud/src

COPY go.mod . 
COPY go.sum .

RUN go mod download


RUN go get github.com/stretchr/testify
RUN go get github.com/cosmtrek/air
RUN go get github.com/sqs/goreturns
RUN go get github.com/mdempsky/gocode
RUN go get github.com/uudashr/gopkgs/v2/cmd/gopkgs
RUN go get github.com/ramya-rao-a/go-outline
RUN go get github.com/stamblerre/gocode
RUN go get github.com/rogpeppe/godef
RUN go get golang.org/x/lint/golint
RUN go get golang.org/x/tools/gopls

COPY . .

EXPOSE 8080

ENTRYPOINT [ "air" ]

# CMD ["go", "run", "src/main.go"]
# CMD ["./main"]
