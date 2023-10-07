#! /bin/bash

TOTAL_PKGS=80
PKG_SECRET_FILE="sleep-values-secrets.yaml"

touch ${PKG_SECRET_FILE}
echo "" > ${PKG_SECRET_FILE}

for ((i=1;i<=$TOTAL_PKGS;i++))
do
    APP_NAME="sleep-app"$i
    APP_FULL_NAME="${APP_NAME}.corp.com"
    echo "processing app $APP_NAME"
#   create directory
    APP_PKG_DIR="sleep-app-repo/packages/${APP_FULL_NAME}"
    mkdir -p $APP_PKG_DIR
    # copy stuff
    # cp template/metadata.yml $APP_PKG_DIR/metadata.yml
    ytt -f template/metadata.yml -v app.name="$APP_NAME" > $APP_PKG_DIR/metadata.yml
    ytt -f template/package-template.yml -v app.name="$APP_NAME" > $APP_PKG_DIR/1.0.0.yml
    # ytt -f template/job-val-template.yml -v app.name="$APP_NAME" > $APP_PKG_DIR/values.yml

    echo "add secret values"
    cat <<EOF >> ${PKG_SECRET_FILE}
---
apiVersion: v1
kind: Secret
metadata:
  name: ${APP_NAME}-values
  namespace: sleeptest
stringData:
  values2.yml: |
    #@data/values
    ---
    app: 
      name: ${APP_NAME}
EOF

done

echo "generating the pkginstall"
ytt -f template/pkginstall.yml -f template/schema.yml --data-value-yaml count=$TOTAL_PKGS > pkg_install_multiple.yml
