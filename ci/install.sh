#!/usr/bin/env bash

version="0.2.1"

os_architecture() {
  if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        echo "linux"
  elif [[ "$OSTYPE" == "darwin"* ]]; then
        ARCH=$(uname -m)
        if [[ "$ARCH" == "x86_64" ]]; then
            echo  "macos"
        elif [[ "$ARCH" == "arm64" ]]; then
            echo "macos-aarch64"
        fi
  else
        echo "Unknown system $OSTYPE"
        exit 1
  fi
}

download_release() {
  local os=$(os_architecture)
  local download_dir="$(mktemp -d)"
  local filename="caddy-$version-$os.tar.gz"

  local download_file="$download_dir/$filename"
  local asset_url="https://github.com/mschakulat/caddy/releases/download/v$version/$filename"
  curl --silent --show-error --location --fail --output "$download_file" "$asset_url" --write-out "$download_file"
}

install_release() {
  local install_dir="$1"
  create_tree "$install_dir"

  info "Downloading" "release"
  local download_file=$(download_release)
  local extract_to=$(mktemp -d)

  info "Extracting" "binaries"
  tar -xzf "$download_file" -C "$extract_to"

  cp "$extract_to"/release/* "$install_dir"/bin/

  "$install_dir"/bin/caddy setup --skip-msg

  info "Setup" "complete"
}

create_tree() {
  info "Creating" "directory layout"
  local install_dir="$1"
  mkdir -p "$install_dir"/bin
}

has_caddy_home() {
  if [ -n "${CADDY_HOME-}" ] && [ -e "$CADDY_HOME" ]; then
    return 1
  fi
  return 0
}

info() {
  local action="$1"
  local details="$2"
  command printf '\033[1;96m%12s\033[0m %s\n' "$action" "$details" 1>&2
}

install_dir="${CADDY_HOME:-"$HOME/.caddy"}"

install_release "$install_dir"