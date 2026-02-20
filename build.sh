#!/bin/env bash
OS_TARGETS="windows linux freebsd openbsd darwin"
ARCH_TARGETS="arm64 amd64 ppc64 riscv64"

mkdir -p dist

for GOOS in $OS_TARGETS; do
    for GOARCH in $ARCH_TARGETS; do
        OUTPUT="Proxeye_${GOOS}-${GOARCH}"

        # Add .exe extension for Windows executables
        if [ "$GOOS" = "windows" ]; then
            OUTPUT="${OUTPUT}.exe"
        fi

        case "$GOOS/$GOARCH" in
            windows/ppc64|windows/riscv64|freebsd/ppc64|darwin/ppc64|darwin/riscv64)
                continue
                ;;
        esac
        env GOOS="$GOOS" GOARCH="$GOARCH" go build -o dist/"$OUTPUT"
    done
done
