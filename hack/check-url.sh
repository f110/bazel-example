#!/usr/bin/env bash

while read -r url; do
    if [[ $url =~ ^https ]]; then
        echo "$url is ok"
    else
        echo "$url is not ok"
        exit 1
    fi
done < "$1"