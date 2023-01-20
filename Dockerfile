#
# svelte-builder
#

FROM node:alpine as app-builder

WORKDIR /app
COPY . ./

RUN npm install --package-lock-only
RUN npm prune
RUN npx vite build

#
# server-builder (CGO is required)
#

FROM golang:alpine as server-builder

RUN apk add build-base

WORKDIR /app

COPY internal ./internal/
COPY go.* ./
COPY cmd/server/main.go cmd/server/

RUN go build -mod=readonly -v cmd/server/main.go

#
# deploy
#

FROM alpine as deployment

WORKDIR /app

COPY --from=app-builder /app/build build/
COPY --from=server-builder /app ./

EXPOSE 8080

CMD ./main -docker