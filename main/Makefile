android:
	GOARCH=arm64 GOOS=android CGO_ENABLED=1 CC=/Users/spicely/Library/Android/sdk/ndk/21.4.7075529/toolchains/llvm/prebuilt/darwin-x86_64/bin/aarch64-linux-android21-clang go build -buildmode=c-shared -o ../build/arm64/openim_sdk_ffi.so openim_sdk_ffi.go

windows:
	GOARCH=arm64 go build -buildmode=c-shared -o ../build/windows/openim_sdk_ffi.dll openim_sdk_ffi.go