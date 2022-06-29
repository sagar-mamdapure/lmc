#!/bin/sh

#copied almost verbatim from
#https://github.com/nouney/helm-gcs/blob/master/scripts/install.sh
cd $HELM_PLUGIN_DIR
if [[ $PROFILE == "LOCAL" ]]; then
  echo "LocalInstallation"
#  version=$(git describe --tags --always --dirty=-SNAPSHOT- $(git rev-parse --short HEAD) | tr -d 'v')
else
  version="$(cat plugin.yaml | grep "version" | cut -d '"' -f 2)"
fi
echo "Installing lmc v${version} ..."

# Find correct archive name
unameOut="$(uname -s)"

case "${unameOut}" in
Linux*) os=Linux ;;
Darwin*) os=Darwin ;;
CYGWIN*) os=Cygwin ;;
MINGW*) os=windows ;;
*) os="UNKNOWN:${unameOut}" ;;
esac

arch=$(uname -m)

if [[ $PROFILE == "LOCAL" ]]; then
  filename=${DIST_DIR}/lmc_${version}_${os}_${arch}.tar.gz
else
  url="https://github.com/logicmonitor/lmc/releases/download/v${version}/lmc_${version}_${os}_${arch}.tar.gz"

  if [ "$url" = "" ]; then
    echo "Unsupported OS / architecture: ${os}_${arch}"
    exit 1
  fi

  filename=$(echo ${url} | sed -e "s/^.*\///g")

  # Download archive
  if [ -n $(command -v curl) ]; then
    curl -sSL -O $url
  elif [ -n $(command -v wget) ]; then
    wget -q $url
  else
    echo "Need curl or wget"
    exit -1
  fi
fi

echo "Installing $filename"
# Install bin
rm -rf bin && mkdir bin && tar xzf $filename >/dev/null && rm -f $filename
mv lmc bin/
echo "lmc ${version} is correctly installed."
echo
echo "See https://github.com/logicmonitor/lmc#getting-started for help getting started."

#for ARGUMENT in "$@"; do
#  KEY=$(echo $ARGUMENT | cut -f1 -d=)
#
#  KEY_LENGTH=${#KEY}
#  VALUE="${ARGUMENT:$KEY_LENGTH+1}"
#
#  export "$KEY"="$VALUE"
#done

#if [[ $PROFILE == "LOCAL" ]]; then
#  echo LocalInstallation
#  #  cd $HELM_PLUGIN_DIR
#  goreleaser release --skip-publish --rm-dist --debug --snapshot
#  filename=$(PWD)/dist/lmc_${VERSION}_Darwin_x86_64.tar.gz
#  cd $HELM_PLUGIN_DIR
#  rm -rf bin && mkdir bin && tar xzvf $filename -C bin >/dev/null && rm -f $filename
#else
#  echo ProdInstallation
#  prod
#fi
