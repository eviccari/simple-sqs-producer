APP_NAME=$1
podman build -t "$APP_NAME" -f $(pwd)/Dockerfile && podman image rm -f $(podman images -f dangling=true -q)