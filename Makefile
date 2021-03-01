.PHONY: setup third_party alpine build-alpine

TOP_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

CGO_CFLAGS=-I$(TOP_DIR)/third_party/mcl/include -I$(TOP_DIR)/third_party/bls/include $(CPPFLAGS)
CGO_LDFLAGS=-L$(TOP_DIR)/third_party/mcl/lib -L$(TOP_DIR)/third_party/bls/lib $(LDFLAGS)

BUILD_LDFLAGS :=
ifeq ($(OS),Windows_NT)
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		OS_RELEASE := $(shell . /etc/os-release && echo $$ID)
		ifeq ($(OS_RELEASE),alpine)
			BUILD_LDFLAGS = -ldflags '-s -w -extldflags=-static'
		endif
	else
	endif
endif

export CGO_CFLAGS CGO_LDFLAGS

third_party:
	- cd third_party/mcl && mkdir -p obj lib
	- cd third_party/bls && mkdir -p obj lib
	make LIB_DIR=lib OBJ_DIR=obj -C third_party/mcl lib/libmcl.a -j8
	make LIB_DIR=lib OBJ_DIR=obj -C third_party/bls lib/libbls384_256.a BLS_SWAP_G=1 -j8


setup:
	- (cd third_party/mcl && go mod init github.com/hyperion-hyn/mcl)
	- (cd third_party/bls && go mod init github.com/hyperion-hyn/bls)
	go get -u github.com/zouguangxian/modvendor@v0.3.2
	go mod tidy && go mod vendor
	modvendor -copy="**/*.c **/*.h **/*.proto" -v -include="github.com/ethereum/go-ethereum/crypto/secp256k1,github.com/karalabe/usb"
