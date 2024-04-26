#!/usr/bin/env bash

latest_version() {
  local owner="mschakulat"
  local repo="caddy"

  if ! command -v jq &> /dev/null; then
    echo "jq is not installed. Please install jq before running this script."
    exit 1
  fi
  
  local release_info=$(curl --silent "https://api.github.com/repos/$owner/$repo/releases/latest")
  local version_with_v=$(echo "$release_info" | jq -r .tag_name)
  local version=${version_with_v#v}

  echo "$version"
}

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

  if [ "$(has_caddy_home)" -eq "0" ]; then
    "$install_dir"/bin/caddy setup --skip-msg
  fi

  info "Setup" "complete"
}

create_tree() {
  info "Creating" "directory layout"
  local install_dir="$1"
  mkdir -p "$install_dir"/bin
}

has_caddy_home() {
  if [ -n "${CADDY_HOME-}" ] && [ -e "$CADDY_HOME" ]; then
    echo "1"
  else
    echo "0"
  fi
}

info() {
  local action="$1"
  local details="$2"
  command printf '\033[1;96m%12s\033[0m %s\n' "$action" "$details" 1>&2
}

version=$(latest_version)

install_dir="${CADDY_HOME:-"$HOME/.caddy"}"

install_release "$install_dir"