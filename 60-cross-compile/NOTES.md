go get -u github.com/mitchellh/gox

gox \
-os="linux darwin windows" \
-arch="amd64 386" \ 
-output="dist/{{.os}}-{{.Arch}}/{{.Dir}}"
