#!/bin/bash

if [ -z "$1" ]; then
  echo "Please provide an action (create/delete)."
  exit 1
fi

if [ -z "$2" ]; then
  echo "Please provide a name for the project."
  exit 1
fi

action=$1
project_name=$2

if [ "$action" == "create" ]; then
  if [ -d "$project_name" ]; then
    echo "Domain already exists."
    exit 1
  fi

  mkdir -p "$project_name/delivery/http/handler"
  mkdir -p "$project_name/repository"
  mkdir -p "$project_name/usecase"
  if [ ! -d "model" ]; then
    mkdir model
  fi
  snake_case=$(echo "$project_name" | tr '-' '_')
  touch model/"$snake_case".go
  echo "Domain created successfully."

elif [ "$action" == "delete" ]; then
  if [ ! -d "$project_name" ]; then
    echo "Domain does not exist."
    exit 1
  fi

  rm -rf "$project_name"
  snake_case=$(echo "$project_name" | tr '-' '_')
  if [ -f "model/$snake_case.go" ]; then
    rm "model/$snake_case.go"
  fi
  echo "Domain deleted successfully."

else
  echo "Invalid action. Please use 'create' or 'delete'."
  exit 1
fi
