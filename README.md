# lmc

## Developing plugin

Replace hooks with following in plugin.yaml

```
install: "cd $HELM_PLUGIN_DIR; export PROFILE=LOCAL; export DIST_DIR=${PWD}/dist; export version=`git --git-dir=${PWD}//.git describe --tags --always --dirty=-SNAPSHOT-\\`git --git-dir=$(PWD)/.git rev-parse --short HEAD\\` | tr -d 'v'`; . ./scripts/install.sh"
update: "cd $HELM_PLUGIN_DIR; export PROFILE=LOCAL; export DIST_DIR=${PWD}/dist; export version=`git --git-dir=${PWD}//.git describe --tags --always --dirty=-SNAPSHOT-\\`git --git-dir=$(PWD)/.git rev-parse --short HEAD\\` | tr -d 'v'`; . ./scripts/install.sh"
```

Run `make deploy`