#!/bin/bash

sed -i '' 's/DEVICE = "remote"/DEVICE = "local"/g' constants/index.go
