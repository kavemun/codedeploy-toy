#!/bin/bash

docker run --rm -v "$PWD:$PWD" -w "$PWD" toy-robot --file $1
