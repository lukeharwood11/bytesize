#!/bin/bash

cd client
if ! command -v node &> /dev/null; then
    echo "Node.js is not installed. Please install Node.js before running this script."
    exit 1
fi

if [[ ! -d ./node_modules ]]; then
    echo "Project has not been initialized... installing dependencies..."
    sleep .25
    npm install;
fi

echo "Starting development server..."
npm start