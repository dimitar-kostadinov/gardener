#!/bin/bash
# Copyright 2023 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

hostname=garden.local.gardener.cloud
config_path=/etc/containerd/certs.d

FILENAME=/etc/containerd/config.toml
if grep -q plugins.\"io.containerd.grpc.v1.cri\".registry.mirrors.\"localhost:5001\" "$FILENAME"; then
  # cleanup old configuration
  sed -i -E '/\[plugins\."io\.containerd\.grpc\.v1\.cri"\.registry\.mirrors\."localhost:5001"\]/,+11d' $FILENAME
  echo "Cleanup old registry mirrors configuration for local-setup."
fi

# configured containerd with registry mirrors for local-setup
namespaces=('localhost:5001' 'gcr.io' 'eu.gcr.io' 'ghcr.io' 'registry.k8s.io' 'quay.io')
servers=('http://localhost:5001' 'https://gcr.io' 'https://eu.gcr.io' 'https://ghcr.io' 'https://registry.k8s.io' 'https://quay.io')
ports=('5001' '5003' '5004' '5005' '5006' '5007')

for i in "${!namespaces[@]}"; do
  namespace=${namespaces[$i]}
  if [[ ! -f "${config_path}/${namespace}/hosts.toml" ]];then
    server=${servers[$i]}
    mkdir -p "${config_path}/${namespace}"
    cat <<EOF > "${config_path}/${namespace}/hosts.toml"
server="${server}"
EOF
    echo "Create hosts.toml for ${namespace} registry in local-setup."
  fi
  port=${ports[$i]}
  if ! grep -qF "http://${hostname}:${port}" "${config_path}/${namespace}/hosts.toml"; then
    cat <<EOF >> "${config_path}/${namespace}/hosts.toml"
[host."http://${hostname}:${port}"]
  capabilities = ["pull", "resolve"]
EOF
    echo "Append ${hostname}:${port} in hosts.toml for ${namespace} registry in local-setup."
  fi
done
