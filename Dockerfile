# build artifacts stage
FROM public.ecr.aws/docker/library/golang:1.21.5 as Build
WORKDIR /build

# user build
USER root

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN chmod +x docker/entrypoint.sh

# build artifacts
RUN CGO_ENABLED=0 GOOS=linux go build -o artifacts .

# deploy stage
FROM public.ecr.aws/docker/library/alpine:3.18.5 as Deploy
WORKDIR /app

#  copy artifacts from build stage to deploy stage
COPY --from=Build /build/artifacts .
COPY --from=Build /build/docker/entrypoint.sh .
COPY --from=Build /build/databases ./databases
COPY --from=Build /build/.env .

EXPOSE 8088

ENTRYPOINT [ "./entrypoint.sh" ]