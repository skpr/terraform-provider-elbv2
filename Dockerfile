FROM hashicorp/terraform:1.2.5

RUN apk add bash

RUN mkdir -p /root/.terraform.d/plugins

ADD ./bin/terraform-provider-elbv2_linux-amd64 /root/.terraform.d/plugins/terraform.local/skpr/elbv2/99.0.0/linux_amd64/terraform-provider-elbv2_linux-amd64_v99.0.0

RUN chmod +x /root/.terraform.d/plugins/terraform.local/*/*/*/linux_amd64/terraform-provider-*
