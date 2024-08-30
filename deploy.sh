#!/bin/bash

DEV='false'
STAGING='false'
PRESTAGING='false'

while [[ $# -gt 0 ]]; do
    case $1 in
    -d | --dev)
        DEV='true'
        shift
        shift
        ;;
    -ps | --pre-staging)
        PRESTAGING='true'
        shift
        shift
        ;;
    -s | --staging)
        STAGING='true'
            shift
            shift
            ;;
    *)
        POSITIONAL_ARGS+=("$1")
        shift
        ;;
    esac
done

if [ $DEV = 'true' ]; then
  git checkout dev &&
  git merge -s ours deploy-dev &&
  git checkout deploy-dev &&
  git merge dev &&
  git add . &&
  git push origin deploy-dev &&
  git checkout dev
elif [ $PRESTAGING = 'true' ]; then
  git checkout pre-staging &&
  git merge -s ours deploy-pre-staging &&
  git checkout deploy-pre-staging &&
  git merge pre-staging &&
  git add . &&
  git push origin deploy-pre-staging &&
  git checkout dev
elif [ $STAGING = 'true' ]; then
  git checkout staging &&
  git merge -s ours deploy-staging &&
  git checkout deploy-staging &&
  git merge staging &&
  git add . &&
  git push origin deploy-staging &&
  git checkout dev
else
  git checkout master &&
  git merge -s ours deploy-production &&
  git checkout deploy-production &&
  git merge master &&
  git add . &&
  git push origin deploy-production &&
  git checkout dev
fi
