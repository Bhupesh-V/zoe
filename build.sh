#!/usr/bin/env bash

package_name="zoe"
platforms=("windows/amd64" "linux/amd64" "linux/arm64" "linux/386" "darwin/amd64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name="$package_name-$GOOS-$GOARCH"
    echo -e "Building for $GOOS-$GOARCH"
    if [ "$GOOS" = "windows" ]; then
        output_name+='.exe'
    elif [ "$GOOS" = "darwin" ]; then
    	output_name="$package_name-macos-$GOARCH"
    fi

    if env GOOS="$GOOS" GOARCH="$GOARCH" go build -o "$output_name"; then
    	echo -e "Built zoe for $platform ✅"
    else
    	echo -e "An error has occurred! Aborting ..."
        exit 1
    fi
done