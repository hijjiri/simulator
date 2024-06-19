FROM envoyproxy/envoy:v1.19.1

# Install gettext to use envsubst
RUN apt-get update && apt-get install -y gettext-base && apt-get clean

COPY envoy.yaml.template /etc/envoy/envoy.yaml.template
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
