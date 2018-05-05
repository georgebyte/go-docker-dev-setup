FROM golang

ARG app_env
ENV APP_ENV $app_env

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY . /go/src/github.com/jurebajt/personal-finance/server
WORKDIR /go/src/github.com/jurebajt/personal-finance/server

RUN dep ensure
RUN go install

# In development use pilu/fresh for hot reloading, in production run the binnary
CMD if [ ${APP_ENV} = production ]; \
    then \
        $GOPATH/bin/server; \
    else \
        go get github.com/pilu/fresh && \
        fresh; \
    fi

EXPOSE 8080
