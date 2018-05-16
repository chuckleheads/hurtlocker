#!/bin/sh

mkdir -p /hab/user/datastore/config
cat <<EOT > /hab/user/datastore/config/user.toml
dev_mode = "true"
EOT
