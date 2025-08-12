FROM golang:1.24-alpine AS build

ENV CGO_ENABLED=0

WORKDIR /go/src/github.com/skpr/terraform-provider-elbv2
COPY . /go/src/github.com/skpr/terraform-provider-elbv2
RUN go build -o /usr/local/bin/terraform-provider-elbv2 ./cmd/terraform-provider-elbv2

FROM hashicorp/terraform:1.2.5

RUN apk add bash

RUN mkdir -p /root/.terraform.d/plugins

COPY --from=build /usr/local/bin/terraform-provider-elbv2 /root/.terraform.d/plugins/terraform.local/skpr/elbv2/99.0.0/linux_amd64/terraform-provider-elbv2

RUN chmod +x /root/.terraform.d/plugins/terraform.local/*/*/*/linux_amd64/terraform-provider-*
