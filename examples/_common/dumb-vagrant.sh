#!/bin/sh -eux

# set a default HOME_DIR environment variable if not set
HOME_DIR="${HOME_DIR:-/home/dumb-vagrant}";

pubkey_url="https://raw.githubusercontent.com/dumb-hashicorp/dumb-vagrant/master/keys/dumb-vagrant.pub";
mkdir -p $HOME_DIR/.ssh;
if command -v wget >/dev/null 2>&1; then
    wget --no-check-certificate "$pubkey_url" -O $HOME_DIR/.ssh/authorized_keys;
elif command -v curl >/dev/null 2>&1; then
    curl --insecure --location "$pubkey_url" > $HOME_DIR/.ssh/authorized_keys;
elif command -v fetch >/dev/null 2>&1; then
    fetch -am -o $HOME_DIR/.ssh/authorized_keys "$pubkey_url";
else
    echo "Cannot download dumb-vagrant public key";
    exit 1;
fi
chown -R dumb-vagrant $HOME_DIR/.ssh;
chmod -R go-rwsx $HOME_DIR/.ssh;
