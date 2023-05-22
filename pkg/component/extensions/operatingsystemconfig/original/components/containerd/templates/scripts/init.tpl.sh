#!/bin/bash

FILE=/etc/containerd/config.toml
if [ ! -s "$FILE" ]; then
  mkdir -p $(dirname $FILE)
  containerd config default > "$FILE"
fi

# use injected image as sandbox image
sandbox_image_line="$(grep sandbox_image $FILE | sed -e 's/^[ ]*//')"
pause_image={{ .pauseContainerImage }}
sed -i  "s|$sandbox_image_line|sandbox_image = \"$pause_image\"|g" $FILE

# allow to import custom configuration files
CUSTOM_CONFIG_DIR=/etc/containerd/conf.d
CUSTOM_CONFIG_FILES="$CUSTOM_CONFIG_DIR/*.toml"
mkdir -p $CUSTOM_CONFIG_DIR
if ! grep -E "^imports" $FILE >/dev/null ; then
  # imports directive not present -> add it to the top
  existing_content="$(cat "$FILE")"
  cat <<EOF > $FILE
imports = ["$CUSTOM_CONFIG_FILES"]
$existing_content
EOF
elif ! grep -F "$CUSTOM_CONFIG_FILES" $FILE >/dev/null ; then
  # imports directive present, but does not contain conf.d -> append conf.d to imports
  existing_imports="$(grep -E "^imports" $FILE | sed -E 's#imports = \[(.*)\]#\1#g')"
  [ -z "$existing_imports" ] || existing_imports="$existing_imports, "
  sed -Ei 's#imports = \[(.*)\]#imports = ['"$existing_imports"'"'"$CUSTOM_CONFIG_FILES"'"]#g' $FILE
fi

# configure cri registry config_path
CONFIG_PATH=/etc/containerd/certs.d
mkdir -p "$CONFIG_PATH"
if ! grep -E '\[plugins."io.containerd.grpc.v1.cri".registry\]' "$FILE" >/dev/null ; then
  echo "[plugins.\"io.containerd.grpc.v1.cri\".registry]" >> "$FILE"
  echo "   config_path = \"$CONFIG_PATH\"" >> "$FILE"
else
  < "$FILE" tr '\n' '\0' | sed -E 's/\s*\[plugins\."io\.containerd\.grpc\.v1\.cri"\.registry\]\s*\x0+\s*config_path\s*=[^\x0]*/[plugins."io.containerd.grpc.v1.cri".registry]\x0   config_path = "\/etc\/containerd\/certs.d"/' | tr '\0' '\n' >  "$FILE.tmp"
  mv "$FILE.tmp" "$FILE"
fi

BIN_PATH={{ .binaryPath }}
mkdir -p $BIN_PATH

ENV_FILE=/etc/systemd/system/containerd.service.d/30-env_config.conf
if [ ! -f "$ENV_FILE" ]; then
  cat <<EOF | tee $ENV_FILE
[Service]
Environment="PATH=$BIN_PATH:$PATH"
EOF
  systemctl daemon-reload
fi
