#
# Контейнер сборки
#
FROM golang:latest as builder

ENV CGO_ENABLED=0

ENV SERVICE="broker"

COPY . /go/src/github.com/MultiBanker/broker
WORKDIR /go/src/github.com/MultiBanker/broker/
RUN \
    version=`git describe --abbrev=6 --always --tag`; \
    echo "version=$version" && \
    cd src/app && \
    go build -a -tags broker -installsuffix broker -ldflags "-X main.version=${version} -s -w" -o /go/bin/broker -mod vendor

#
# Контейнер для получения актуальных SSL/TLS сертификатов
#
FROM alpine as alpine
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
RUN addgroup -S broker && adduser -S broker -G broker

# копируем документацию
RUN mkdir -p /usr/share/broker
COPY --from=builder /go/src/github.com/MultiBanker/broker/src/swagger /usr/share/broker
RUN chown -R broker:broker /usr/share/broker

ENTRYPOINT [ "/bin/broker" ]

#
# Контейнер рантайма
#
FROM scratch
COPY --from=builder /go/bin/broker /bin/broker

# копируем сертификаты из alpine
COPY --from=alpine /etc/ssl/certs /etc/ssl/certs

# копируем документацию
COPY --from=alpine /usr/share/broker /usr/share/broker

# копируем пользователя и группу из alpine
COPY --from=alpine /etc/passwd /etc/passwd
COPY --from=alpine /etc/group /etc/group

USER broker

ENTRYPOINT ["/bin/broker"]

