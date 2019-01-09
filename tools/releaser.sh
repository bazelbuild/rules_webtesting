#!/bin/bash

bazel build -c opt --stamp \
	--platforms=@io_bazel_rules_go//go/toolchain:darwin_amd64 \
	//go/metadata/main \
	//go/wsl/main \
	//go/wtl/main

bazel build -c opt --stamp \
	--platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 \
	//go/metadata/main \
	//go/wsl/main \
	//go/wtl/main

bazel build -c opt --stamp \
	--platforms=@io_bazel_rules_go//go/toolchain:windows_amd64 \
	//go/metadata/main \
	//go/wsl/main \
	//go/wtl/main	

 /google/data/ro/teams/copybara/copybara --folder-dir="$1" migrate tools/copy.bara.sky release .
