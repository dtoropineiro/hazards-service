#!/bin/bash

# Configurable variables
process_name="emergencies"
app_binary="./$process_name"
log_path="../logs"
log_file="$log_path/hazards-go.log"

# Ensure log directory exists
mkdir -p "$log_path"

echo "Building app..."
if ! go build -o "$app_binary" emergencies.go; then
    echo "Build failed, exiting."
    exit 1
fi
echo "Build successful."

echo "Checking if process is active..."
if pgrep -f "$app_binary" > /dev/null; then
    echo "Process $process_name is active, killing it."
    pkill -f "$app_binary"
    echo "Process $process_name killed."
else
    echo "Process $process_name is not active."
fi

echo "Running app..."
nohup "$app_binary" > "$log_file" 2>&1 &

echo "App is running now!"
