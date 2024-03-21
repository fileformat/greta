FROM golang:1.22-alpine AS builder

RUN apk add --no-cache upx

RUN mkdir /build
ADD . /build/
WORKDIR /build
ARG COMMIT
ARG LASTMOD
ARG VERSION
ARG BUILTBY
RUN echo "INFO: building for $COMMIT on $LASTMOD"

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go build \
        -a \
        -ldflags "-s -w -X main.commit=$COMMIT -X main.date=$LASTMOD -X main.version=$VERSION -X main.builtBy=$BUILTBY -extldflags '-static'" \
        -o greta \
        ./greta.go \
    && upx greta

FROM scratch
COPY --from=builder /build/greta /bin/greta
WORKDIR /bin
ENTRYPOINT ["/bin/greta"]

