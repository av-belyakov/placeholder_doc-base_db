
ARG BUILD_IMAGE_NAME=golang:1.24.4-alpine
ARG IMAGE_NAME=alpine

FROM ${BUILD_IMAGE_NAME} AS packages_image
ENV PATH=/usr/local/go/bin:$PATH
WORKDIR /go/src
COPY go.mod go.sum ./
RUN echo "packages_image" && \
    go mod download

FROM ${BUILD_IMAGE_NAME} AS build_image
LABEL temporary=""
ARG BRANCH
ARG VERSION
ARG USERNAME
ARG USERPASSWD
WORKDIR /go/
COPY --from=packages_image /go ./
RUN echo -e "build_image" && \
    rm -r ./src && \
    apk update && \
    apk add --no-cache git && \
    # брать исходный код с репозитория на gitlab.cloud.gcm 
    #git clone -b ${BRANCH} http://${USERNAME}:${USERPASSWD}@gitlab.cloud.gcm/a.belyakov/placeholder_doc-base_db.git ./src/${VERSION}/ && \
    git clone -b ${BRANCH} http://${USERNAME}:${USERPASSWD}@192.168.9.33/a.belyakov/placeholder_doc-base_db.git ./src/${VERSION}/ && \
    # брать исходный код с репозитория на github.com 
    #git clone -b ${BRANCH} https://github.com/av-belyakov/placeholder_doc-base_db.git  ./src/${VERSION}/ && \
    go build -C ./src/${VERSION}/cmd/ -o ../app

FROM ${IMAGE_NAME}
LABEL author="Artemij Belyakov"
# аргумент STATUS содержит режим запуска приложения prod или development
#если значение содержит запись development, то в таком режиме и будет
#работать приложение, во всех остальных случаях режим работы prod
ARG STATUS=""
ARG VERSION
ARG USERNAME=dockeruser
ARG US_DIR=/opt/placeholder_doc-base_db
ENV GO_PHDOCBASEDB_MAIN=${STATUS}
RUN addgroup --g 1500 groupcontainer && \
    adduser -u 1500 -G groupcontainer -D ${USERNAME} --home ${US_DIR}
USER ${USERNAME}
WORKDIR ${US_DIR}
RUN mkdir ./logs
COPY --from=build_image /go/src/${VERSION}/app ./
COPY --from=build_image /go/src/${VERSION}/README.md ./
COPY --from=build_image /go/src/${VERSION}/version ./ 
COPY config/* ./config/

ENTRYPOINT [ "./app" ]
