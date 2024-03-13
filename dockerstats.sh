#!/bin/bash

# File to store the stats
output_file="docker_stats.txt"

# Function to write header to the file if it doesn't exist
write_header() {
    if [ ! -f "$output_file" ]; then
        echo "Container ID      Name                CPU %   Mem Usage / Limit  Mem %   Net I/O    Block I/O  PIDs" > "$output_file"
        echo "===================================================================================================" >> "$output_file"
    fi
}

# Function to get docker stats and append to file
get_docker_stats() {
    # Fetch stats for containers containing "caching-performance" in their name
    docker stats --no-stream --format "table {{.Container}}\t{{.Name}}\t{{.CPUPerc}}\t{{.MemUsage}}\t{{.MemPerc}}\t{{.NetIO}}\t{{.BlockIO}}\t{{.PIDs}}" $(docker ps --filter "name=caching-performance" -q) | tail -n +2 >> "$output_file"
}

# Run indefinitely every 5 seconds
while true; do
    write_header
    get_docker_stats
    sleep 5
done
