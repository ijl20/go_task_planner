#!/bin/bash

GOOS=windows GOARCH=amd64 go build -o build/b21_task_planner.exe go_task_planner
