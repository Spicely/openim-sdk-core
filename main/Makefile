android:
	GOARCH=arm64 GOOS=android CGO_ENABLED=1 CC=/Users/spicely/Library/Android/sdk/ndk/21.4.7075529/toolchains/llvm/prebuilt/darwin-x86_64/bin/aarch64-linux-android21-clang go build -buildmode=c-shared -o ../build/arm64/libopenimSDK.so openim_sdk_ffi.go

windows:
	SET CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64
	go build -buildmode=c-shared -o ../build/windows/openim_sdk_ffi.dll openim_sdk_ffi.go

macos_x86_64:
	export CFLAGS="-mmacosx-version-min=10.11 -isysroot "$(xcrun -sdk macosx --show-sdk-path) 
	export CGO_LDFLAGS="-mmacosx-version-min=10.11 -isysroot "$(xcrun -sdk macosx --show-sdk-path)  
	CGO_ENABLED=1 GOARCH=amd64 GOOS=darwin CC="clang $CFLAGS $CGO_LDFLAGS" go build -tags="darwin" -buildmode c-archive -ldflags=-w -trimpath -v -o ../build/macos/libopenim_sdk_ffi.dylib openim_sdk_ffi.go
	

	GOOS=darwin GOARCH=amd64 go build -buildmode=c-shared -o ../build/macos/libopenimSDK.dylib openim_sdk_ffi.go


macos_arm64:
	CGO_ENABLED=1 CC=clang CXX=clang++ GOOS=darwin GOARCH=arm64 go build -buildmode=c-shared -o ../build/macos/libopenim_sdk_ffi.dylib openim_sdk_ffi.go

ios:
	export CFLAGS="-arch arm64 -miphoneos-version-min=9.0 -isysroot "$(xcrun -sdk iphoneos --show-sdk-path) 
	export CGO_LDFLAGS="-arch arm64 -miphoneos-version-min=9.0 -isysroot "$(xcrun -sdk iphoneos --show-sdk-path)  
	CGO_ENABLED=1 GOARCH=arm64 GOOS=darwin CC="clang $CFLAGS $CGO_LDFLAGS" go build -tags ios -buildmode c-archive -ldflags=-w  -trimpath -v -o ../build/ios/libopenim_sdk_ffi.dylib openim_sdk_ffi.go

iphonesimulator:
	export CFLAGS="-arch x86_64 -miphoneos-version-min=9.0 -isysroot "$(xcrun -sdk iphonesimulator --show-sdk-path) 
	export CGO_LDFLAGS="-arch x86_64 -miphoneos-version-min=9.0 -isysroot "$(xcrun -sdk iphonesimulator --show-sdk-path) 
	CGO_ENABLED=1 GOARCH=amd64 GOOS=darwin CC="clang $CFLAGS $CGO_LDFLAGS" go build -tags ios -tags ios -buildmode c-archive -ldflags=-w  -trimpath -v -o ../build/ios/libopenim_sdk_ffi.dylib openim_sdk_ffi.go
	
codesign -f -s "Apple Development: zheng tang (HS53J8AN55)" ../build/macos/libopenimSDK.dylib
xcrun -sdk iphoneos clang -arch armv7 -fpic -shared -Wl,-all_load libopenim_sdk_ffi.dylib -framework Corefoundation -o libopenim_sdk_ffi.dylib
