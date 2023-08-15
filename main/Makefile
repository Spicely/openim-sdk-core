android:
	GOARCH=arm64 GOOS=android CGO_ENABLED=1 CC=/Users/muka/Library/Android/sdk/ndk/25.1.8937393/toolchains/llvm/prebuilt/darwin-x86_64/bin/aarch64-linux-android21-clang go build -buildmode=c-shared -ldflags="-s -w" -o ../build/libs/arm64/libopenim_sdk_ffi.so openim_sdk_ffi.go
	GOARCH=arm64 GOOS=android CGO_ENABLED=1 CC=/Users/muka/Library/Android/sdk/ndk/25.1.8937393/toolchains/llvm/prebuilt/darwin-x86_64/bin/aarch64-linux-android21-clang go build -buildmode=c-shared -ldflags="-s -w" -o ../build/libs/arm64-v8a/libopenim_sdk_ffi.so openim_sdk_ffi.go
	GOARCH=arm GOOS=android CGO_ENABLED=1 CC=/Users/muka/Library/Android/sdk/ndk/25.1.8937393/toolchains/llvm/prebuilt/darwin-x86_64/bin/armv7a-linux-androideabi21-clang go build -buildmode=c-shared -ldflags="-s -w" -o ../build/libs/armeabi-v7a/libopenim_sdk_ffi.so openim_sdk_ffi.go
	GOARCH=386 GOOS=android CGO_ENABLED=1 CC=/Users/muka/Library/Android/sdk/ndk/25.1.8937393/toolchains/llvm/prebuilt/darwin-x86_64/bin/i686-linux-android21-clang go build -buildmode=c-shared -ldflags="-s -w"  -o ../build/libs/x86/libopenim_sdk_ffi.so openim_sdk_ffi.go

	rm -rf ../build/libs/arm64/libopenim_sdk_ffi.h
	rm -rf ../build/libs/arm64-v8a/libopenim_sdk_ffi.h
	rm -rf ../build/libs/armeabi-v7a/libopenim_sdk_ffi.h
	rm -rf ../build/libs/x86/libopenim_sdk_ffi.h
	upx --android-shlib -f ../build/libs/arm64/libopenim_sdk_ffi.so
	upx --android-shlib --force-pie -f -o ../build/libs/arm64-v8a/libopenim_sdk_ffi.so
	upx --android-shlib --force-pie -f -o ../build/libs/armeabi-v7a/libopenim_sdk_ffi.so
	upx --android-shlib --force-pie -f -o ../build/libs/x86/libopenim_sdk_ffi.so
	rsync -av --delete ../build/libs /Users/muka/Desktop/www/flutter_openim_sdk_ffi/android

windows:
	SET CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64
	go build -buildmode=c-shared -o ../build/windows/openim_sdk_ffi.dll openim_sdk_ffi.go

macos:
	CGO_ENABLED=1 CC=clang CXX=clang++ GOOS=darwin GOARCH=arm64 go build -buildmode=c-shared -ldflags="-s -w" -o ../build/macos/libopenim_sdk_ffi_arm64.dylib openim_sdk_ffi.go
	CGO_ENABLED=1 CC=clang CXX=clang++ GOOS=darwin GOARCH=amd64 GOARM=7 go build -buildmode=c-shared -ldflags="-s -w" -o ../build/macos/libopenim_sdk_ffi_x86_64.dylib openim_sdk_ffi.go
	lipo -create ../build/macos/libopenim_sdk_ffi_arm64.dylib ../build/macos/libopenim_sdk_ffi_x86_64.dylib -output ../build/macos/libopenim_sdk_ffi.dylib
	rm -rf ../build/macos/libopenim_sdk_ffi_arm64.dylib  ../build/macos/libopenim_sdk_ffi_x86_64.dylib  ../build/macos/libopenim_sdk_ffi_x86_64.h
	install_name_tool -id @rpath/libopenim_sdk_ffi.dylib ../build/macos/libopenim_sdk_ffi.dylib

ios:
	export CFLAGS="-arch arm64 -miphoneos-version-min=9.0 -isysroot "$(xcrun -sdk iphoneos --show-sdk-path) 
	export CGO_LDFLAGS="-arch arm64 -miphoneos-version-min=9.0 -isysroot "$(xcrun -sdk iphoneos --show-sdk-path)  
	CGO_ENABLED=1 GOARCH=arm64 GOOS=ios CC="clang $CFLAGS $CGO_LDFLAGS" go build -tags ios -ldflags "-s -w" -trimpath -v -o ../build/ios/libopenim_sdk_ffi.a -buildmode c-archive openim_sdk_ffi.go
	xcrun -sdk iphoneos clang -arch arm64 -fpic -shared -Wl,-all_load ../build/ios/libopenim_sdk_ffi.a -framework CoreFoundation -framework Security -lresolv -o ../build/ios/libopenim_sdk_ffi.dylib
	install_name_tool -id @rpath/libopenim_sdk_ffi.dylib ../build/ios/libopenim_sdk_ffi.dylib


	export CFLAGS="-arch arm64 -miphoneos-version-min=9.0 -isysroot $(xcrun -sdk iphoneos --show-sdk-path)"
	export CGO_LDFLAGS="-arch arm64 -miphoneos-version-min=9.0 -isysroot $(xcrun -sdk iphoneos --show-sdk-path)"
	CGO_ENABLED=1 GOARCH=arm64 GOOS=darwin CC="clang $CFLAGS $CGO_LDFLAGS" go build -tags ios -buildmode=c-shared -ldflags="-w -s" -trimpath -o ../build/ios/libopenim_sdk_ffi_arm64.dylib openim_sdk_ffi.go
	export CFLAGS="-arch x86_64 -mios-simulator-version-min=9.0 -isysroot $(xcrun --sdk iphonesimulator --show-sdk-path)"
	export CGO_LDFLAGS="-arch x86_64 -mios-simulator-version-min=9.0 -isysroot $(xcrun --sdk iphonesimulator --show-sdk-path)"
	CGO_ENABLED=1 GOARCH=amd64 GOOS=darwin CC="clang $CFLAGS $CGO_LDFLAGS" go build -tags ios -buildmode=c-shared -ldflags="-w -s" -trimpath -o ../build/ios/libopenim_sdk_ffi_x86.dylib openim_sdk_ffi.go
	lipo -create ../build/ios/libopenim_sdk_ffi_arm64.dylib ../build/ios/libopenim_sdk_ffi_x86.dylib -output ../build/ios/libopenim_sdk_ffi.dylib
	rm -rf ../build/ios/libopenim_sdk_ffi_arm64.dylib ../build/ios/libopenim_sdk_ffi_x86.dylib  ../build/ios/libopenim_sdk_ffi_x86.h
	install_name_tool -id @rpath/libopenim_sdk_ffi.dylib ../build/ios/libopenim_sdk_ffi.dylib
	mv ../build/ios/libopenim_sdk_ffi.dylib ../../flutter_openim_sdk_ffi/ios/libopenim_sdk_ffi.dylib

linux:
	go build -tags linux -ldflags=-w -trimpath -v -o "../build/linux/libopenim_sdk_ffi.so" -buildmode c-shared openim_sdk_ffi.go
	
codesign -f -s "Apple Development: zheng tang (HS 53J 8AN55)" ../build/macos/libopenimSDK.dylib
xcrun -sdk iphoneos clang -arch armv7 -fpic -shared -Wl,-all_load libopenim_sdk_ffi.dylib -framework Corefoundation -o libopenim_sdk_ffi.dylib

upx -9 ../build/libs/arm64-v8a/libopenim_sdk_ffi.so