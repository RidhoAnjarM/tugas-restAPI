#!/bin/bash

MSG=""
STAGING='false'
PRESTAGING='false'
DEV='false'

while [[ $# -gt 0 ]]; do
    case $1 in
    -s | --staging)
        STAGING="true"
        shift
        shift
        ;;
    -ps | --pre-staging)
        PRESTAGING="true"
        shift
        shift
        ;;
    -d | --dev)
        DEV="true"
        shift
        shift
        ;;
    -m | --message)
        MSG="$2"
        shift
        shift
        ;;
    *)
        POSITIONAL_ARGS+=("$1")
        shift
        ;;
    esac
done

set -- "${POSITIONAL_ARGS[@]}"

if [ "$DEV" = 'true' ]; then
  git checkout dev &&
  sed -i '' 's/DEVICE = "local"/DEVICE = "remote"/g' constants/index.go &&
  git add . &&
  git commit -m "$MSG" &&
  git push origin dev
elif [ "$PRESTAGING" = 'true' ]; then
  git checkout dev &&
  git merge -s ours pre-staging -m "(+) Merge From dev" &&
  git checkout pre-staging &&
  git merge dev &&
  sed -i '' 's/ENV = "dev"/ENV = "pre-staging"/g' constants/index.go &&
  sed -i '' 's/DEVICE = "local"/DEVICE = "remote"/g' constants/index.go &&
  git add . &&
  git commit -m "(+) Merging Pre Staging" &&
  git push origin pre-staging &&
  git checkout dev
elif [ "$STAGING" = 'true' ]; then
  git checkout dev &&
  git merge -s ours staging -m "(+) Merge From Pre Staging" &&
  git checkout staging &&
  git merge dev &&
  sed -i '' 's/ENV = "dev"/ENV = "staging"/g' constants/index.go &&
  sed -i '' 's/DEVICE = "local"/DEVICE = "remote"/g' constants/index.go &&
  git add . &&
  git commit -m "(+) Merging Staging" &&
  git push origin staging &&
  git checkout dev
else
  git checkout staging &&
  git merge -s ours master -m "(+) Merge From Staging" &&
  git checkout master &&
  git merge staging &&
  sed -i '' 's/ENV = "staging"/ENV = "production"/g' constants/index.go &&
  sed -i '' 's/DEVICE = "local"/DEVICE = "remote"/g' constants/index.go &&
  git add . &&
  git commit -m "(+) Merging Staging" &&
  git push origin master &&
  git checkout dev
fi
