#!/bin/bash

#Настрайваем откуда берем прото
PROJECT_ID="102" #id проекта в гите
PATH_TO_FILE="wallet_storage_ds.proto" # путь к файлу в гите (наример dir1/dir2/file.proto)
FILE_NAME="wallet_storage_ds.proto" # название файла (наример file.proto)
GIT_BRANCH="main" #ветка

GITLAB_TOKEN="nesecret" # токен для доступа к апи, этот - readonly мой(Айдархан Аркаев) можно им делится.

#Настрайваем куда сохраняем прото
OUTPUT_DIR="./pkg/proto/wallet_storage_ds" # куда сохранить прото локально

#Запускаем скрипт

urlencode() {
  # Usage: urlencode "string"
  local string="${1}"
  local length="${#string}"
  local url=""
  local c

  for (( i = 0; i < length; i++ )); do
c="${string:i:1}"
    case $c in
      [a-zA-Z0-9.~_-])
        url+="${c}"
        ;;
      *)
        printf -v c '%%%02x' "'$c"
        url+="${c}"
        ;;
    esac
done

  echo "${url}"
}

PATH_TO_FILE="$(urlencode "$PATH_TO_FILE")"
REQUEST_URL="https://gitlab.secret.kz/api/v4/projects/$PROJECT_ID/repository/files/$PATH_TO_FILE/raw?ref=$GIT_BRANCH"
mkdir -p "$OUTPUT_DIR"
# Download the proto file
curl --header "PRIVATE-TOKEN: $GITLAB_TOKEN" "$REQUEST_URL" --output "$OUTPUT_DIR/$FILE_NAME"

# Generate gRPC code from the proto file
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    "$OUTPUT_DIR/$FILE_NAME" --experimental_allow_proto3_optional

echo "gRPC proto file generated successfully in $OUTPUT_DIR"

