#!/bin/sh

mkdir -p /hab/user/builder-datastore/config
cat <<EOT > /hab/user/builder-datastore/config/user.toml
dev_mode = "true"
EOT
