BINARY_PATH=./dist
BINARY_NAME=web_app
.PHONY:tidy build dist-path build-windows build-linux build-darwin

build:tidy clean build-darwin build-linux build-windows
# 这里是为了保证第三方依赖的兼容性，第一行找到非vendors文件夹下的所有go文件进行格式化处理
tidy:
	find . -type f -name "*.go" -not -path "./vendor/*" | xargs -n1 go fmt
	go mod tidy

dist-path:
	mkdir -p ${BINARY_PATH}
# 这里是给链接器加了两个参数，分别是禁用了生成DWARF文件和符号表生成，然后CGO_ENABLED=0可以保证静态链接，从而使得go程序的可移植性提高
build-darwin:dist-path
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ${BINARY_PATH}/${BINARY_NAME}
	cd ${BINARY_PATH} && zip ${BINARY_NAME}-darwin.zip ${BINARY_NAME} && rm ${BINARY_NAME}

build-linux:dist-path
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ${BINARY_PATH}/${BINARY_NAME}
	cd ${BINARY_PATH} && zip ${BINARY_NAME}-linux.zip ${BINARY_NAME} && rm ${BINARY_NAME}

build-windows:dist-path
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ${BINARY_PATH}/${BINARY_NAME}.exe

clean:
	rm -rf ${BINARY_PATH}