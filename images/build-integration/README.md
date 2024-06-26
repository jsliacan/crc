# Overview

The container includes the integration binary for all 3 platforms plus the required resources to run it.  

The container connects through ssh to the target host and copy the right binary for the platform, run integration tests and pick the results and logs back.

## Envs

**PLATFORM**:*define target platform (windows, macos, linux).*
**TARGET_HOST**:*dns or ip for the target host.*  
**TARGET_HOST_USERNAME**:*username for target host.*  
**TARGET_HOST_KEY_PATH**:*private key for user. (Mandatory if not TARGET_HOST_PASSWORD).*  
**TARGET_HOST_PASSWORD**:*password for user. (Mandatory if not TARGET_HOST_KEY_PATH).*  
**PULL_SECRET_FILE_PATH** pull secret file path (local to container).*  
**BUNDLE_LOCATION**:*(Optional). When testing crc with custom bundle set the bundle location on target server.*  
**RESULTS_PATH**:*(Optional). Path inside container to pick results and logs from integration execution.*

## Samples

```bash
# Run integration on macos platform with ssh key and custom bundle
podman run --rm -it --name crc-integration \
    -e PLATFORM=macos \
    -e TARGET_HOST=$IP \
    -e TARGET_HOST_USERNAME=$USER \
    -e TARGET_HOST_KEY_PATH=/opt/crc/id_rsa \
    -e PULL_SECRET_FILE_PATH=/opt/crc/pull-secret \
    -e BUNDLE_LOCATION=/bundles/crc_hyperv_4.8.0-rc.3.crcbundle \
    -v $PWD/pull-secret:/opt/crc/pull-secret:Z \
    -v $PWD/id_rsa:/opt/crc/id_rsa:Z \
    -v $PWD/output:/output:Z \
    quay.io/crcont/crc-integration:v1.29.0

# Run integration on windows platform with ssh password and crc released version
podman run --rm -it --name crc-integration \
    -e PLATFORM=windows \
    -e TARGET_HOST=$IP \
    -e TARGET_HOST_USERNAME=$USER \
    -e TARGET_HOST_PASSWORD=$PASSWORD \
    -e PULL_SECRET_FILE_PATH=/opt/crc/pull-secret \
    -v $PWD/pull-secret:/opt/crc/pull-secret:Z \
    -v $PWD/output:/output:Z \
    quay.io/crcont/crc-integration:v1.29.0
```
