#!/bin/bash

sed -i '' 's/DEVICE = "local"/DEVICE = "remote"/g' constants/index.go
