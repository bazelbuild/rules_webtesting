resolved = [
     {
          "original_rule_class": "local_repository",
          "original_attributes": {
               "name": "bazel_tools",
               "path": "/var/tmp/_bazel_pcloudy/install/fd0e803cd377e5d6c999b9f309de8848/embedded_tools"
          },
          "native": "local_repository(name = \"bazel_tools\", path = __embedded_dir__ + \"/\" + \"embedded_tools\")"
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
          "definition_information": "Repository io_bazel_rules_go instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:29:13: in <toplevel>\nRepository rule http_archive defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in <toplevel>\n",
          "original_attributes": {
               "name": "io_bazel_rules_go",
               "urls": [
                    "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
                    "https://github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip"
               ],
               "sha256": "278b7ff5a826f3dc10f04feaf0b70d48b68748ccd512d7f98bf442077f043fe3"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
                    "attributes": {
                         "url": "",
                         "urls": [
                              "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
                              "https://github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip"
                         ],
                         "sha256": "278b7ff5a826f3dc10f04feaf0b70d48b68748ccd512d7f98bf442077f043fe3",
                         "integrity": "",
                         "netrc": "",
                         "auth_patterns": {},
                         "canonical_id": "",
                         "strip_prefix": "",
                         "add_prefix": "",
                         "type": "",
                         "patches": [],
                         "remote_file_urls": {},
                         "remote_file_integrity": {},
                         "remote_patches": {},
                         "remote_patch_strip": 0,
                         "patch_tool": "",
                         "patch_args": [
                              "-p0"
                         ],
                         "patch_cmds": [],
                         "patch_cmds_win": [],
                         "build_file_content": "",
                         "workspace_file_content": "",
                         "name": "io_bazel_rules_go"
                    },
                    "output_tree_hash": "e8bc41a36606fe0ee1474a5778109c480f9acd29f7e56c1a911b050646826ce5"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
          "definition_information": "Repository bazel_gazelle instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:44:13: in <toplevel>\nRepository rule http_archive defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in <toplevel>\n",
          "original_attributes": {
               "name": "bazel_gazelle",
               "urls": [
                    "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.32.0/bazel-gazelle-v0.32.0.tar.gz",
                    "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.32.0/bazel-gazelle-v0.32.0.tar.gz"
               ],
               "sha256": "29218f8e0cebe583643cbf93cae6f971be8a2484cdcfa1e45057658df8d54002"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
                    "attributes": {
                         "url": "",
                         "urls": [
                              "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.32.0/bazel-gazelle-v0.32.0.tar.gz",
                              "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.32.0/bazel-gazelle-v0.32.0.tar.gz"
                         ],
                         "sha256": "29218f8e0cebe583643cbf93cae6f971be8a2484cdcfa1e45057658df8d54002",
                         "integrity": "",
                         "netrc": "",
                         "auth_patterns": {},
                         "canonical_id": "",
                         "strip_prefix": "",
                         "add_prefix": "",
                         "type": "",
                         "patches": [],
                         "remote_file_urls": {},
                         "remote_file_integrity": {},
                         "remote_patches": {},
                         "remote_patch_strip": 0,
                         "patch_tool": "",
                         "patch_args": [
                              "-p0"
                         ],
                         "patch_cmds": [],
                         "patch_cmds_win": [],
                         "build_file_content": "",
                         "workspace_file_content": "",
                         "name": "bazel_gazelle"
                    },
                    "output_tree_hash": "5165655e670420f84fac7e8998ad01b9b311f0f472fe7ffe500894647982da82"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
          "definition_information": "Repository bazel_skylib instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:22:22: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/repositories.bzl:42:21: in web_test_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/repositories.bzl:67:17: in bazel_skylib\nRepository rule http_archive defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in <toplevel>\n",
          "original_attributes": {
               "name": "bazel_skylib",
               "generator_name": "bazel_skylib",
               "generator_function": "web_test_repositories",
               "generator_location": None,
               "urls": [
                    "https://github.com/bazelbuild/bazel-skylib/releases/download/1.0.3/bazel-skylib-1.0.3.tar.gz",
                    "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.0.3/bazel-skylib-1.0.3.tar.gz"
               ],
               "sha256": "1c531376ac7e5a180e0237938a2536de0c54d93f5c278634818e0efc952dd56c"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
                    "attributes": {
                         "url": "",
                         "urls": [
                              "https://github.com/bazelbuild/bazel-skylib/releases/download/1.0.3/bazel-skylib-1.0.3.tar.gz",
                              "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.0.3/bazel-skylib-1.0.3.tar.gz"
                         ],
                         "sha256": "1c531376ac7e5a180e0237938a2536de0c54d93f5c278634818e0efc952dd56c",
                         "integrity": "",
                         "netrc": "",
                         "auth_patterns": {},
                         "canonical_id": "",
                         "strip_prefix": "",
                         "add_prefix": "",
                         "type": "",
                         "patches": [],
                         "remote_file_urls": {},
                         "remote_file_integrity": {},
                         "remote_patches": {},
                         "remote_patch_strip": 0,
                         "patch_tool": "",
                         "patch_args": [
                              "-p0"
                         ],
                         "patch_cmds": [],
                         "patch_cmds_win": [],
                         "build_file_content": "",
                         "workspace_file_content": "",
                         "name": "bazel_skylib"
                    },
                    "output_tree_hash": "ec0173581163d32cb764072fa396fc158fb56ac7660d93f944ae7723401a67e2"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
          "definition_information": "Repository io_bazel_rules_scala instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:79:13: in <toplevel>\nRepository rule http_archive defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in <toplevel>\n",
          "original_attributes": {
               "name": "io_bazel_rules_scala",
               "url": "https://github.com/bazelbuild/rules_scala/releases/download/v6.5.0/rules_scala-v6.5.0.tar.gz",
               "sha256": "3b00fa0b243b04565abb17d3839a5f4fa6cc2cac571f6db9f83c1982ba1e19e5",
               "strip_prefix": "rules_scala-6.5.0"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
                    "attributes": {
                         "url": "https://github.com/bazelbuild/rules_scala/releases/download/v6.5.0/rules_scala-v6.5.0.tar.gz",
                         "urls": [],
                         "sha256": "3b00fa0b243b04565abb17d3839a5f4fa6cc2cac571f6db9f83c1982ba1e19e5",
                         "integrity": "",
                         "netrc": "",
                         "auth_patterns": {},
                         "canonical_id": "",
                         "strip_prefix": "rules_scala-6.5.0",
                         "add_prefix": "",
                         "type": "",
                         "patches": [],
                         "remote_file_urls": {},
                         "remote_file_integrity": {},
                         "remote_patches": {},
                         "remote_patch_strip": 0,
                         "patch_tool": "",
                         "patch_args": [
                              "-p0"
                         ],
                         "patch_cmds": [],
                         "patch_cmds_win": [],
                         "build_file_content": "",
                         "workspace_file_content": "",
                         "name": "io_bazel_rules_scala"
                    },
                    "output_tree_hash": "9351dbdfce7e7f74ed115a7b6b12cce4e11e2245dfca1026a52f2c34652c0d54"
               }
          ]
     },
     {
          "original_rule_class": "@@io_bazel_rules_scala//:scala_config.bzl%_config_repository",
          "definition_information": "Repository io_bazel_rules_scala_config instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:88:13: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scala_config.bzl:81:23: in scala_config\nRepository rule _config_repository defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scala_config.bzl:65:37: in <toplevel>\n",
          "original_attributes": {
               "name": "io_bazel_rules_scala_config",
               "generator_name": "io_bazel_rules_scala_config",
               "generator_function": "scala_config",
               "generator_location": None,
               "scala_version": "2.12.18",
               "enable_compiler_dependency_tracking": False
          },
          "repositories": [
               {
                    "rule_class": "@@io_bazel_rules_scala//:scala_config.bzl%_config_repository",
                    "attributes": {
                         "name": "io_bazel_rules_scala_config",
                         "generator_name": "io_bazel_rules_scala_config",
                         "generator_function": "scala_config",
                         "generator_location": None,
                         "scala_version": "2.12.18",
                         "enable_compiler_dependency_tracking": False
                    },
                    "output_tree_hash": "ec201e429492022076d570c4f8954d18c17c8c5071144b9f513faaede585c954"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:local.bzl%local_repository",
          "definition_information": "Repository rules_java_builtin instantiated at:\n  /DEFAULT.WORKSPACE:12:36: in <toplevel>\nRepository rule local_repository defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/local.bzl:64:35: in <toplevel>\n",
          "original_attributes": {
               "name": "rules_java_builtin",
               "path": "/var/tmp/_bazel_pcloudy/install/fd0e803cd377e5d6c999b9f309de8848/rules_java"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:local.bzl%local_repository",
                    "attributes": {
                         "name": "rules_java_builtin",
                         "path": "/var/tmp/_bazel_pcloudy/install/fd0e803cd377e5d6c999b9f309de8848/rules_java"
                    },
                    "output_tree_hash": "23156af102e8441d4b3e5358092fc1dce333786289d48b1df6503ecb8c735cf3"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
          "definition_information": "Repository rules_proto instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:92:18: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scala/private/macros/scala_repositories.bzl:113:21: in rules_scala_setup\nRepository rule http_archive defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in <toplevel>\n",
          "original_attributes": {
               "name": "rules_proto",
               "generator_name": "rules_proto",
               "generator_function": "rules_scala_setup",
               "generator_location": None,
               "urls": [
                    "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/refs/tags/5.3.0-21.7.tar.gz",
                    "https://github.com/bazelbuild/rules_proto/archive/refs/tags/5.3.0-21.7.tar.gz"
               ],
               "sha256": "dc3fb206a2cb3441b485eb1e423165b231235a1ea9b031b4433cf7bc1fa460dd",
               "strip_prefix": "rules_proto-5.3.0-21.7"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
                    "attributes": {
                         "url": "",
                         "urls": [
                              "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/refs/tags/5.3.0-21.7.tar.gz",
                              "https://github.com/bazelbuild/rules_proto/archive/refs/tags/5.3.0-21.7.tar.gz"
                         ],
                         "sha256": "dc3fb206a2cb3441b485eb1e423165b231235a1ea9b031b4433cf7bc1fa460dd",
                         "integrity": "",
                         "netrc": "",
                         "auth_patterns": {},
                         "canonical_id": "",
                         "strip_prefix": "rules_proto-5.3.0-21.7",
                         "add_prefix": "",
                         "type": "",
                         "patches": [],
                         "remote_file_urls": {},
                         "remote_file_integrity": {},
                         "remote_patches": {},
                         "remote_patch_strip": 0,
                         "patch_tool": "",
                         "patch_args": [
                              "-p0"
                         ],
                         "patch_cmds": [],
                         "patch_cmds_win": [],
                         "build_file_content": "",
                         "workspace_file_content": "",
                         "name": "rules_proto"
                    },
                    "output_tree_hash": "9534782371e091f25cc063e2930a188d6ae96f4e2bcb64c0bd53d2ef16caff8b"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
          "definition_information": "Repository rules_jvm_external instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:111:13: in <toplevel>\nRepository rule http_archive defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in <toplevel>\n",
          "original_attributes": {
               "name": "rules_jvm_external",
               "url": "https://github.com/bazelbuild/rules_jvm_external/releases/download/5.3/rules_jvm_external-5.3.tar.gz",
               "sha256": "d31e369b854322ca5098ea12c69d7175ded971435e55c18dd9dd5f29cc5249ac",
               "strip_prefix": "rules_jvm_external-5.3"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
                    "attributes": {
                         "url": "https://github.com/bazelbuild/rules_jvm_external/releases/download/5.3/rules_jvm_external-5.3.tar.gz",
                         "urls": [],
                         "sha256": "d31e369b854322ca5098ea12c69d7175ded971435e55c18dd9dd5f29cc5249ac",
                         "integrity": "",
                         "netrc": "",
                         "auth_patterns": {},
                         "canonical_id": "",
                         "strip_prefix": "rules_jvm_external-5.3",
                         "add_prefix": "",
                         "type": "",
                         "patches": [],
                         "remote_file_urls": {},
                         "remote_file_integrity": {},
                         "remote_patches": {},
                         "remote_patch_strip": 0,
                         "patch_tool": "",
                         "patch_args": [
                              "-p0"
                         ],
                         "patch_cmds": [],
                         "patch_cmds_win": [],
                         "build_file_content": "",
                         "workspace_file_content": "",
                         "name": "rules_jvm_external"
                    },
                    "output_tree_hash": "354d76a3c055314c9724db6981121cf44d39cef2129014fcb06c0d8f3c8869fe"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:local.bzl%local_repository",
          "definition_information": "Repository internal_platforms_do_not_use instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:153:6: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\nRepository rule local_repository defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/local.bzl:64:35: in <toplevel>\n",
          "original_attributes": {
               "name": "internal_platforms_do_not_use",
               "generator_name": "internal_platforms_do_not_use",
               "generator_function": "maybe",
               "generator_location": None,
               "path": "/var/tmp/_bazel_pcloudy/install/fd0e803cd377e5d6c999b9f309de8848/platforms"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:local.bzl%local_repository",
                    "attributes": {
                         "name": "internal_platforms_do_not_use",
                         "generator_name": "internal_platforms_do_not_use",
                         "generator_function": "maybe",
                         "generator_location": None,
                         "path": "/var/tmp/_bazel_pcloudy/install/fd0e803cd377e5d6c999b9f309de8848/platforms"
                    },
                    "output_tree_hash": "db797f5ddb49595460e727f2c71af1b3adfed4d65132bbe31bd9d3a06bd95dba"
               }
          ]
     },
     {
          "original_rule_class": "@@internal_platforms_do_not_use//host:extension.bzl%host_platform_repo",
          "definition_information": "Repository host_platform instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:165:6: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\nRepository rule host_platform_repo defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/internal_platforms_do_not_use/host/extension.bzl:51:37: in <toplevel>\n",
          "original_attributes": {
               "name": "host_platform",
               "generator_name": "host_platform",
               "generator_function": "maybe",
               "generator_location": None
          },
          "repositories": [
               {
                    "rule_class": "@@internal_platforms_do_not_use//host:extension.bzl%host_platform_repo",
                    "attributes": {
                         "name": "host_platform",
                         "generator_name": "host_platform",
                         "generator_function": "maybe",
                         "generator_location": None
                    },
                    "output_tree_hash": "dcbfb61be394f2b4fd27f49c2c538d0d87564c5c8baad14a0079063212442538"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:local.bzl%local_repository",
          "definition_information": "Repository platforms instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:147:6: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\nRepository rule local_repository defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/local.bzl:64:35: in <toplevel>\n",
          "original_attributes": {
               "name": "platforms",
               "generator_name": "platforms",
               "generator_function": "maybe",
               "generator_location": None,
               "path": "/var/tmp/_bazel_pcloudy/install/fd0e803cd377e5d6c999b9f309de8848/platforms"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:local.bzl%local_repository",
                    "attributes": {
                         "name": "platforms",
                         "generator_name": "platforms",
                         "generator_function": "maybe",
                         "generator_location": None,
                         "path": "/var/tmp/_bazel_pcloudy/install/fd0e803cd377e5d6c999b9f309de8848/platforms"
                    },
                    "output_tree_hash": "db797f5ddb49595460e727f2c71af1b3adfed4d65132bbe31bd9d3a06bd95dba"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remote_jdk8_linux_aarch64_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:370:22: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:349:34: in remote_jdk8_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remote_jdk8_linux_aarch64_toolchain_config_repo",
               "generator_name": "remote_jdk8_linux_aarch64_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_8\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"8\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_linux_aarch64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_linux_aarch64//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remote_jdk8_linux_aarch64_toolchain_config_repo",
                         "generator_name": "remote_jdk8_linux_aarch64_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_8\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"8\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_linux_aarch64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_linux_aarch64//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "c9c795851cffbf2a808bfc7cccea597c3b3fef46cfefa084f7e9de7e90b65447"
               }
          ]
     },
     {
          "original_rule_class": "@@io_bazel_rules_go//go/private:sdk.bzl%go_multiple_toolchains",
          "definition_information": "Repository go_sdk_toolchains instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:42:23: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_go/go/private/sdk.bzl:695:28: in go_register_toolchains\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_go/go/private/sdk.bzl:308:19: in go_download_sdk\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_go/go/private/sdk.bzl:296:27: in _go_toolchains\nRepository rule go_multiple_toolchains defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_go/go/private/sdk.bzl:283:41: in <toplevel>\n",
          "original_attributes": {
               "name": "go_sdk_toolchains",
               "generator_name": "go_sdk_toolchains",
               "generator_function": "go_register_toolchains",
               "generator_location": None,
               "prefixes": [
                    ""
               ],
               "sdk_repos": [
                    "go_sdk"
               ],
               "sdk_types": [
                    "remote"
               ],
               "sdk_versions": [
                    "1.20.5"
               ],
               "geese": [
                    ""
               ],
               "goarchs": [
                    ""
               ]
          },
          "repositories": [
               {
                    "rule_class": "@@io_bazel_rules_go//go/private:sdk.bzl%go_multiple_toolchains",
                    "attributes": {
                         "name": "go_sdk_toolchains",
                         "generator_name": "go_sdk_toolchains",
                         "generator_function": "go_register_toolchains",
                         "generator_location": None,
                         "prefixes": [
                              ""
                         ],
                         "sdk_repos": [
                              "go_sdk"
                         ],
                         "sdk_types": [
                              "remote"
                         ],
                         "sdk_versions": [
                              "1.20.5"
                         ],
                         "geese": [
                              ""
                         ],
                         "goarchs": [
                              ""
                         ]
                    },
                    "output_tree_hash": "7f22ee8fb31b495ed920a86a5e9980189a1462059c2fb132df84127b0714a64e"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remote_jdk8_linux_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:370:22: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:349:34: in remote_jdk8_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remote_jdk8_linux_toolchain_config_repo",
               "generator_name": "remote_jdk8_linux_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_8\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"8\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_linux//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_linux//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remote_jdk8_linux_toolchain_config_repo",
                         "generator_name": "remote_jdk8_linux_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_8\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"8\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_linux//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_linux//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "b6a178fc0ca08a4473490f1c5d0f9f633db0ca0f2834c69dd08ce8290cf9ca86"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remote_jdk8_macos_aarch64_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:370:22: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:349:34: in remote_jdk8_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remote_jdk8_macos_aarch64_toolchain_config_repo",
               "generator_name": "remote_jdk8_macos_aarch64_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_8\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"8\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_macos_aarch64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_macos_aarch64//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remote_jdk8_macos_aarch64_toolchain_config_repo",
                         "generator_name": "remote_jdk8_macos_aarch64_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_8\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"8\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_macos_aarch64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_macos_aarch64//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "4d721d8b0731cfb50f963f8b55c7bef9f572de0e2f251f07a12c722ef1acbb2f"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remote_jdk8_macos_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:370:22: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:349:34: in remote_jdk8_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remote_jdk8_macos_toolchain_config_repo",
               "generator_name": "remote_jdk8_macos_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_8\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"8\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_macos//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_macos//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remote_jdk8_macos_toolchain_config_repo",
                         "generator_name": "remote_jdk8_macos_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_8\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"8\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_macos//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_macos//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "e0d82dc2dbe8ec49d859811afe4973ec36226875a39ac7fc8419e91e7e9c89fb"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/sh:sh_configure.bzl%sh_config",
          "definition_information": "Repository local_config_sh instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:187:13: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/sh/sh_configure.bzl:83:14: in sh_configure\nRepository rule sh_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/sh/sh_configure.bzl:72:28: in <toplevel>\n",
          "original_attributes": {
               "name": "local_config_sh",
               "generator_name": "local_config_sh",
               "generator_function": "sh_configure",
               "generator_location": None
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/sh:sh_configure.bzl%sh_config",
                    "attributes": {
                         "name": "local_config_sh",
                         "generator_name": "local_config_sh",
                         "generator_function": "sh_configure",
                         "generator_location": None
                    },
                    "output_tree_hash": "e36855460b514225eac75f4abe2cb992c5455b7077a9028d213d269d11490744"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
          "definition_information": "Repository rules_java instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:92:18: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scala/private/macros/scala_repositories.bzl:104:21: in rules_scala_setup\nRepository rule http_archive defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in <toplevel>\n",
          "original_attributes": {
               "name": "rules_java",
               "generator_name": "rules_java",
               "generator_function": "rules_scala_setup",
               "generator_location": None,
               "urls": [
                    "https://github.com/bazelbuild/rules_java/releases/download/5.4.1/rules_java-5.4.1.tar.gz"
               ],
               "sha256": "a1f82b730b9c6395d3653032bd7e3a660f9d5ddb1099f427c1e1fe768f92e395"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
                    "attributes": {
                         "url": "",
                         "urls": [
                              "https://github.com/bazelbuild/rules_java/releases/download/5.4.1/rules_java-5.4.1.tar.gz"
                         ],
                         "sha256": "a1f82b730b9c6395d3653032bd7e3a660f9d5ddb1099f427c1e1fe768f92e395",
                         "integrity": "",
                         "netrc": "",
                         "auth_patterns": {},
                         "canonical_id": "",
                         "strip_prefix": "",
                         "add_prefix": "",
                         "type": "",
                         "patches": [],
                         "remote_file_urls": {},
                         "remote_file_integrity": {},
                         "remote_patches": {},
                         "remote_patch_strip": 0,
                         "patch_tool": "",
                         "patch_args": [
                              "-p0"
                         ],
                         "patch_cmds": [],
                         "patch_cmds_win": [],
                         "build_file_content": "",
                         "workspace_file_content": "",
                         "name": "rules_java"
                    },
                    "output_tree_hash": "1632b69490f133ad83431dd3fa5d7fc330770b38ce0401e6c8ed20e729ff4468"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remote_jdk8_windows_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:370:22: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:349:34: in remote_jdk8_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remote_jdk8_windows_toolchain_config_repo",
               "generator_name": "remote_jdk8_windows_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_8\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"8\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_windows//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_windows//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remote_jdk8_windows_toolchain_config_repo",
                         "generator_name": "remote_jdk8_windows_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_8\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"8\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_windows//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_windows//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "8d0b08c18f215c185d64efe72054a5ffef36325906c34ebf1d3c710d4ba5c685"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remote_jdk8_linux_s390x_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:370:22: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:349:34: in remote_jdk8_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remote_jdk8_linux_s390x_toolchain_config_repo",
               "generator_name": "remote_jdk8_linux_s390x_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_8\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"8\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:s390x\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_linux_s390x//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:s390x\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_linux_s390x//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remote_jdk8_linux_s390x_toolchain_config_repo",
                         "generator_name": "remote_jdk8_linux_s390x_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_8\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"8\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:s390x\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_linux_s390x//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:s390x\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remote_jdk8_linux_s390x//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "f1e3f0b4884e21863a7c19a3a12a8995ed4162e02bd07cbb61b42799fc2d7359"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk11_linux_aarch64_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:371:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:353:34: in remote_jdk11_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk11_linux_aarch64_toolchain_config_repo",
               "generator_name": "remotejdk11_linux_aarch64_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_11\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"11\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_linux_aarch64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_linux_aarch64//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk11_linux_aarch64_toolchain_config_repo",
                         "generator_name": "remotejdk11_linux_aarch64_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_11\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"11\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_linux_aarch64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_linux_aarch64//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "bef508c068dd47d605f62c53ab0628f1f7f5101fdcc8ada09b2067b36c47931f"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk11_linux_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:371:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:353:34: in remote_jdk11_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk11_linux_toolchain_config_repo",
               "generator_name": "remotejdk11_linux_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_11\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"11\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_linux//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_linux//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk11_linux_toolchain_config_repo",
                         "generator_name": "remotejdk11_linux_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_11\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"11\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_linux//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_linux//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "0a170bf4f31e6c4621aeb4d4ce4b75b808be2f3a63cb55dc8172c27707d299ab"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk11_macos_aarch64_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:371:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:353:34: in remote_jdk11_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk11_macos_aarch64_toolchain_config_repo",
               "generator_name": "remotejdk11_macos_aarch64_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_11\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"11\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_macos_aarch64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_macos_aarch64//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk11_macos_aarch64_toolchain_config_repo",
                         "generator_name": "remotejdk11_macos_aarch64_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_11\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"11\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_macos_aarch64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_macos_aarch64//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "ca1d067909669aa58188026a7da06d43bdec74a3ba5c122af8a4c3660acd8d8f"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk11_macos_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:371:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:353:34: in remote_jdk11_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk11_macos_toolchain_config_repo",
               "generator_name": "remotejdk11_macos_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_11\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"11\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_macos//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_macos//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk11_macos_toolchain_config_repo",
                         "generator_name": "remotejdk11_macos_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_11\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"11\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_macos//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_macos//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "45b3b36d22d3e614745e7a5e838351c32fe0eabb09a4a197bac0f4d416a950ce"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk11_win_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:371:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:353:34: in remote_jdk11_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk11_win_toolchain_config_repo",
               "generator_name": "remotejdk11_win_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_11\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"11\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_win//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_win//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk11_win_toolchain_config_repo",
                         "generator_name": "remotejdk11_win_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_11\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"11\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_win//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_win//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "d0587a4ecc9323d5cf65314b2d284b520ffb5ee1d3231cc6601efa13dadcc0f4"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk11_linux_ppc64le_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:371:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:353:34: in remote_jdk11_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk11_linux_ppc64le_toolchain_config_repo",
               "generator_name": "remotejdk11_linux_ppc64le_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_11\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"11\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:ppc\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_linux_ppc64le//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:ppc\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_linux_ppc64le//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk11_linux_ppc64le_toolchain_config_repo",
                         "generator_name": "remotejdk11_linux_ppc64le_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_11\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"11\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:ppc\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_linux_ppc64le//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:ppc\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_linux_ppc64le//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "3272b586976beea589d09ea8029fd5d714da40127c8850e3480991c2440c5825"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk11_linux_s390x_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:371:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:353:34: in remote_jdk11_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk11_linux_s390x_toolchain_config_repo",
               "generator_name": "remotejdk11_linux_s390x_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_11\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"11\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:s390x\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_linux_s390x//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:s390x\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_linux_s390x//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk11_linux_s390x_toolchain_config_repo",
                         "generator_name": "remotejdk11_linux_s390x_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_11\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"11\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:s390x\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_linux_s390x//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:s390x\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_linux_s390x//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "244e11245106a8495ac4744a90023b87008e3e553766ba11d47a9fe5b4bb408d"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk11_win_arm64_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:371:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:353:34: in remote_jdk11_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk11_win_arm64_toolchain_config_repo",
               "generator_name": "remotejdk11_win_arm64_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_11\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"11\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:arm64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_win_arm64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:arm64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_win_arm64//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk11_win_arm64_toolchain_config_repo",
                         "generator_name": "remotejdk11_win_arm64_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_11\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"11\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:arm64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_win_arm64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:arm64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk11_win_arm64//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "c237bd9668de9b6437c452c020ea5bc717ff80b1a5ffd581adfdc7d4a6c5fe03"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
          "definition_information": "Repository rules_cc instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:92:18: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scala/private/macros/scala_repositories.bzl:96:21: in rules_scala_setup\nRepository rule http_archive defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in <toplevel>\n",
          "original_attributes": {
               "name": "rules_cc",
               "generator_name": "rules_cc",
               "generator_function": "rules_scala_setup",
               "generator_location": None,
               "urls": [
                    "https://github.com/bazelbuild/rules_cc/releases/download/0.0.6/rules_cc-0.0.6.tar.gz"
               ],
               "sha256": "3d9e271e2876ba42e114c9b9bc51454e379cbf0ec9ef9d40e2ae4cec61a31b40",
               "strip_prefix": "rules_cc-0.0.6"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
                    "attributes": {
                         "url": "",
                         "urls": [
                              "https://github.com/bazelbuild/rules_cc/releases/download/0.0.6/rules_cc-0.0.6.tar.gz"
                         ],
                         "sha256": "3d9e271e2876ba42e114c9b9bc51454e379cbf0ec9ef9d40e2ae4cec61a31b40",
                         "integrity": "",
                         "netrc": "",
                         "auth_patterns": {},
                         "canonical_id": "",
                         "strip_prefix": "rules_cc-0.0.6",
                         "add_prefix": "",
                         "type": "",
                         "patches": [],
                         "remote_file_urls": {},
                         "remote_file_integrity": {},
                         "remote_patches": {},
                         "remote_patch_strip": 0,
                         "patch_tool": "",
                         "patch_args": [
                              "-p0"
                         ],
                         "patch_cmds": [],
                         "patch_cmds_win": [],
                         "build_file_content": "",
                         "workspace_file_content": "",
                         "name": "rules_cc"
                    },
                    "output_tree_hash": "269524f6b3349c144813cb06d881b33376d769381541dd615209f3b3f00d792f"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk17_linux_aarch64_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:372:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:357:34: in remote_jdk17_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk17_linux_aarch64_toolchain_config_repo",
               "generator_name": "remotejdk17_linux_aarch64_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_17\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"17\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_linux_aarch64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_linux_aarch64//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk17_linux_aarch64_toolchain_config_repo",
                         "generator_name": "remotejdk17_linux_aarch64_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_17\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"17\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_linux_aarch64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_linux_aarch64//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "b169b01ac1a169d7eb5e3525454c3e408e9127993ac0f578dc2c5ad183fd4e3e"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/cpp:cc_configure.bzl%cc_autoconf_toolchains",
          "definition_information": "Repository local_config_cc_toolchains instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:181:13: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/cpp/cc_configure.bzl:148:27: in cc_configure\nRepository rule cc_autoconf_toolchains defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/cpp/cc_configure.bzl:47:41: in <toplevel>\n",
          "original_attributes": {
               "name": "local_config_cc_toolchains",
               "generator_name": "local_config_cc_toolchains",
               "generator_function": "cc_configure",
               "generator_location": None
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/cpp:cc_configure.bzl%cc_autoconf_toolchains",
                    "attributes": {
                         "name": "local_config_cc_toolchains",
                         "generator_name": "local_config_cc_toolchains",
                         "generator_function": "cc_configure",
                         "generator_location": None
                    },
                    "output_tree_hash": "2c6c2998e70208a29847dd5420b3aff0b1e2f5ac956a0911addd090e92b83969"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk17_macos_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:372:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:357:34: in remote_jdk17_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk17_macos_toolchain_config_repo",
               "generator_name": "remotejdk17_macos_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_17\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"17\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_macos//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_macos//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk17_macos_toolchain_config_repo",
                         "generator_name": "remotejdk17_macos_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_17\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"17\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_macos//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_macos//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "41aa7b3317f8d9001746e908454760bf544ffaa058abe22f40711246608022ba"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk17_macos_aarch64_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:372:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:357:34: in remote_jdk17_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk17_macos_aarch64_toolchain_config_repo",
               "generator_name": "remotejdk17_macos_aarch64_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_17\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"17\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_macos_aarch64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_macos_aarch64//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk17_macos_aarch64_toolchain_config_repo",
                         "generator_name": "remotejdk17_macos_aarch64_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_17\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"17\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_macos_aarch64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_macos_aarch64//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "0eb17f6d969bc665a21e55d29eb51e88a067159ee62cf5094b17658a07d3accb"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk17_linux_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:372:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:357:34: in remote_jdk17_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk17_linux_toolchain_config_repo",
               "generator_name": "remotejdk17_linux_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_17\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"17\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_linux//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_linux//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk17_linux_toolchain_config_repo",
                         "generator_name": "remotejdk17_linux_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_17\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"17\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_linux//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_linux//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "f0f07fe0f645f2dc7b8c9953c7962627e1c7425cc52f543729dbff16cd20e461"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk17_win_arm64_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:372:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:357:34: in remote_jdk17_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk17_win_arm64_toolchain_config_repo",
               "generator_name": "remotejdk17_win_arm64_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_17\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"17\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:arm64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_win_arm64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:arm64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_win_arm64//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk17_win_arm64_toolchain_config_repo",
                         "generator_name": "remotejdk17_win_arm64_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_17\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"17\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:arm64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_win_arm64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:arm64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_win_arm64//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "86b129d9c464a9b08f97eca7d8bc5bdb3676b581f8aac044451dbdfaa49e69d3"
               }
          ]
     },
     {
          "original_rule_class": "local_config_platform",
          "original_attributes": {
               "name": "local_config_platform"
          },
          "native": "local_config_platform(name = 'local_config_platform')"
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk17_win_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:372:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:357:34: in remote_jdk17_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk17_win_toolchain_config_repo",
               "generator_name": "remotejdk17_win_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_17\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"17\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_win//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_win//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk17_win_toolchain_config_repo",
                         "generator_name": "remotejdk17_win_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_17\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"17\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_win//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_win//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "170c3c9a35e502555dc9f04b345e064880acbf7df935f673154011356f4aad34"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk21_win_arm64_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:373:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:361:34: in remote_jdk21_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk21_win_arm64_toolchain_config_repo",
               "generator_name": "remotejdk21_win_arm64_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_21\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"21\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:arm64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_win_arm64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:arm64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_win_arm64//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk21_win_arm64_toolchain_config_repo",
                         "generator_name": "remotejdk21_win_arm64_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_21\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"21\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:arm64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_win_arm64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:arm64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_win_arm64//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "9bbdbb329eeba27bc482582360abc6e3351d9a9a07ee11cba3a0026c90223e85"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk21_linux_aarch64_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:373:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:361:34: in remote_jdk21_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk21_linux_aarch64_toolchain_config_repo",
               "generator_name": "remotejdk21_linux_aarch64_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_21\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"21\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_linux_aarch64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_linux_aarch64//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk21_linux_aarch64_toolchain_config_repo",
                         "generator_name": "remotejdk21_linux_aarch64_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_21\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"21\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_linux_aarch64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_linux_aarch64//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "bb33021f243382d2fb849ec204c5c8be5083c37e081df71d34a84324687cf001"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk17_linux_s390x_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:372:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:357:34: in remote_jdk17_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk17_linux_s390x_toolchain_config_repo",
               "generator_name": "remotejdk17_linux_s390x_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_17\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"17\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:s390x\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_linux_s390x//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:s390x\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_linux_s390x//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk17_linux_s390x_toolchain_config_repo",
                         "generator_name": "remotejdk17_linux_s390x_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_17\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"17\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:s390x\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_linux_s390x//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:s390x\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_linux_s390x//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "6ba1870e09fccfdcd423f4169b966a73f8e9deaff859ec6fb3b626ed61ebd8b5"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk17_linux_ppc64le_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:372:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:357:34: in remote_jdk17_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk17_linux_ppc64le_toolchain_config_repo",
               "generator_name": "remotejdk17_linux_ppc64le_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_17\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"17\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:ppc\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_linux_ppc64le//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:ppc\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_linux_ppc64le//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk17_linux_ppc64le_toolchain_config_repo",
                         "generator_name": "remotejdk17_linux_ppc64le_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_17\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"17\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:ppc\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_linux_ppc64le//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:ppc\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk17_linux_ppc64le//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "fdc8ae00f2436bfc46b2f54c84f2bd84122787ede232a4d61ffc284bfe6f61ec"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk21_linux_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:373:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:361:34: in remote_jdk21_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk21_linux_toolchain_config_repo",
               "generator_name": "remotejdk21_linux_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_21\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"21\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_linux//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_linux//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk21_linux_toolchain_config_repo",
                         "generator_name": "remotejdk21_linux_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_21\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"21\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_linux//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_linux//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "ee548ad054c9b75286ff3cd19792e433a2d1236378d3a0d8076fca0bb1a88e05"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk21_macos_aarch64_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:373:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:361:34: in remote_jdk21_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk21_macos_aarch64_toolchain_config_repo",
               "generator_name": "remotejdk21_macos_aarch64_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_21\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"21\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_macos_aarch64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_macos_aarch64//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk21_macos_aarch64_toolchain_config_repo",
                         "generator_name": "remotejdk21_macos_aarch64_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_21\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"21\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_macos_aarch64//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:aarch64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_macos_aarch64//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "706d910cc6809ea7f77fa4f938a4f019dd90d9dad927fb804a14b04321300a36"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk21_win_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:373:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:361:34: in remote_jdk21_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk21_win_toolchain_config_repo",
               "generator_name": "remotejdk21_win_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_21\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"21\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_win//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_win//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk21_win_toolchain_config_repo",
                         "generator_name": "remotejdk21_win_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_21\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"21\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_win//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:windows\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_win//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "87012328b07a779503deec0ef47132a0de50efd69afe7df87619bcc07b1dc4ed"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk21_macos_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:373:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:361:34: in remote_jdk21_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk21_macos_toolchain_config_repo",
               "generator_name": "remotejdk21_macos_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_21\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"21\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_macos//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_macos//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk21_macos_toolchain_config_repo",
                         "generator_name": "remotejdk21_macos_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_21\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"21\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_macos//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:macos\", \"@platforms//cpu:x86_64\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_macos//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "434446eddb7f6a3dcc7a2a5330ed9ab26579c5142c19866b197475a695fbb32f"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk21_linux_s390x_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:373:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:361:34: in remote_jdk21_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk21_linux_s390x_toolchain_config_repo",
               "generator_name": "remotejdk21_linux_s390x_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_21\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"21\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:s390x\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_linux_s390x//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:s390x\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_linux_s390x//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk21_linux_s390x_toolchain_config_repo",
                         "generator_name": "remotejdk21_linux_s390x_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_21\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"21\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:s390x\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_linux_s390x//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:s390x\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_linux_s390x//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "30b78e0951c37c2d7ae1318f83045ff42ef261dbb93c5b4fd3ba963e12cf68d6"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
          "definition_information": "Repository remotejdk21_linux_ppc64le_toolchain_config_repo instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:373:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:361:34: in remote_jdk21_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:57:22: in remote_java_repository\nRepository rule _toolchain_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:27:36: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk21_linux_ppc64le_toolchain_config_repo",
               "generator_name": "remotejdk21_linux_ppc64le_toolchain_config_repo",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_21\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"21\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:ppc\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_linux_ppc64le//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:ppc\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_linux_ppc64le//:jdk\",\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:remote_java_repository.bzl%_toolchain_config",
                    "attributes": {
                         "name": "remotejdk21_linux_ppc64le_toolchain_config_repo",
                         "generator_name": "remotejdk21_linux_ppc64le_toolchain_config_repo",
                         "generator_function": "rules_java_dependencies",
                         "generator_location": None,
                         "build_file": "\nconfig_setting(\n    name = \"prefix_version_setting\",\n    values = {\"java_runtime_version\": \"remotejdk_21\"},\n    visibility = [\"//visibility:private\"],\n)\nconfig_setting(\n    name = \"version_setting\",\n    values = {\"java_runtime_version\": \"21\"},\n    visibility = [\"//visibility:private\"],\n)\nalias(\n    name = \"version_or_prefix_version_setting\",\n    actual = select({\n        \":version_setting\": \":version_setting\",\n        \"//conditions:default\": \":prefix_version_setting\",\n    }),\n    visibility = [\"//visibility:private\"],\n)\ntoolchain(\n    name = \"toolchain\",\n    target_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:ppc\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_linux_ppc64le//:jdk\",\n)\ntoolchain(\n    name = \"bootstrap_runtime_toolchain\",\n    # These constraints are not required for correctness, but prevent fetches of remote JDK for\n    # different architectures. As every Java compilation toolchain depends on a bootstrap runtime in\n    # the same configuration, this constraint will not result in toolchain resolution failures.\n    exec_compatible_with = [\"@platforms//os:linux\", \"@platforms//cpu:ppc\"],\n    target_settings = [\":version_or_prefix_version_setting\"],\n    toolchain_type = \"@bazel_tools//tools/jdk:bootstrap_runtime_toolchain_type\",\n    toolchain = \"@remotejdk21_linux_ppc64le//:jdk\",\n)\n"
                    },
                    "output_tree_hash": "7886e497d586c3f3c8225685281b0940e9aa699af208dc98de3db8839e197be3"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
          "definition_information": "Repository rules_python instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:97:25: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_proto/proto/repositories.bzl:29:14: in rules_proto_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\nRepository rule http_archive defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in <toplevel>\n",
          "original_attributes": {
               "name": "rules_python",
               "generator_name": "rules_python",
               "generator_function": "rules_proto_dependencies",
               "generator_location": None,
               "urls": [
                    "https://mirror.bazel.build/github.com/bazelbuild/rules_python/archive/4b84ad270387a7c439ebdccfd530e2339601ef27.tar.gz",
                    "https://github.com/bazelbuild/rules_python/archive/4b84ad270387a7c439ebdccfd530e2339601ef27.tar.gz"
               ],
               "sha256": "e5470e92a18aa51830db99a4d9c492cc613761d5bdb7131c04bd92b9834380f6",
               "strip_prefix": "rules_python-4b84ad270387a7c439ebdccfd530e2339601ef27"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
                    "attributes": {
                         "url": "",
                         "urls": [
                              "https://mirror.bazel.build/github.com/bazelbuild/rules_python/archive/4b84ad270387a7c439ebdccfd530e2339601ef27.tar.gz",
                              "https://github.com/bazelbuild/rules_python/archive/4b84ad270387a7c439ebdccfd530e2339601ef27.tar.gz"
                         ],
                         "sha256": "e5470e92a18aa51830db99a4d9c492cc613761d5bdb7131c04bd92b9834380f6",
                         "integrity": "",
                         "netrc": "",
                         "auth_patterns": {},
                         "canonical_id": "",
                         "strip_prefix": "rules_python-4b84ad270387a7c439ebdccfd530e2339601ef27",
                         "add_prefix": "",
                         "type": "",
                         "patches": [],
                         "remote_file_urls": {},
                         "remote_file_integrity": {},
                         "remote_patches": {},
                         "remote_patch_strip": 0,
                         "patch_tool": "",
                         "patch_args": [
                              "-p0"
                         ],
                         "patch_cmds": [],
                         "patch_cmds_win": [],
                         "build_file_content": "",
                         "workspace_file_content": "",
                         "name": "rules_python"
                    },
                    "output_tree_hash": "09b73abff05a660e6a6c039828e4511368c41a627c8b52f26833b1b57b7ce1a3"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
          "definition_information": "Repository org_seleniumhq_py instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:77:16: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/py_repositories.bzl:45:26: in py_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/py_repositories.bzl:61:17: in org_seleniumhq_py\nRepository rule http_archive defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in <toplevel>\n",
          "original_attributes": {
               "name": "org_seleniumhq_py",
               "generator_name": "org_seleniumhq_py",
               "generator_function": "py_repositories",
               "generator_location": None,
               "urls": [
                    "https://files.pythonhosted.org/packages/ed/9c/9030520bf6ff0b4c98988448a93c04fcbd5b13cd9520074d8ed53569ccfe/selenium-3.141.0.tar.gz"
               ],
               "sha256": "deaf32b60ad91a4611b98d8002757f29e6f2c2d5fcaf202e1c9ad06d6772300d",
               "strip_prefix": "selenium-3.141.0",
               "build_file": "//build_files:org_seleniumhq_py.BUILD"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
                    "attributes": {
                         "url": "",
                         "urls": [
                              "https://files.pythonhosted.org/packages/ed/9c/9030520bf6ff0b4c98988448a93c04fcbd5b13cd9520074d8ed53569ccfe/selenium-3.141.0.tar.gz"
                         ],
                         "sha256": "deaf32b60ad91a4611b98d8002757f29e6f2c2d5fcaf202e1c9ad06d6772300d",
                         "integrity": "",
                         "netrc": "",
                         "auth_patterns": {},
                         "canonical_id": "",
                         "strip_prefix": "selenium-3.141.0",
                         "add_prefix": "",
                         "type": "",
                         "patches": [],
                         "remote_file_urls": {},
                         "remote_file_integrity": {},
                         "remote_patches": {},
                         "remote_patch_strip": 0,
                         "patch_tool": "",
                         "patch_args": [
                              "-p0"
                         ],
                         "patch_cmds": [],
                         "patch_cmds_win": [],
                         "build_file": "//build_files:org_seleniumhq_py.BUILD",
                         "build_file_content": "",
                         "workspace_file_content": "",
                         "name": "org_seleniumhq_py"
                    },
                    "output_tree_hash": "04a91279bf332b458692be797b05403d894e2d90c2a6aa2d9443936015e56b9b"
               }
          ]
     },
     {
          "original_rule_class": "//web/internal:platform_archive.bzl%platform_archive",
          "definition_information": "Repository org_mozilla_geckodriver_macos_arm64 instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:59:21: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/versioned/browsers-0.3.3.bzl:31:32: in browser_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/versioned/browsers-0.3.3.bzl:266:21: in org_mozilla_geckodriver\nRepository rule platform_archive defined at:\n  /Users/pcloudy/workspace/rules_webtesting/web/internal/platform_archive.bzl:81:35: in <toplevel>\n",
          "original_attributes": {
               "name": "org_mozilla_geckodriver_macos_arm64",
               "generator_name": "org_mozilla_geckodriver_macos_arm64",
               "generator_function": "browser_repositories",
               "generator_location": None,
               "urls": [
                    "https://github.com/mozilla/geckodriver/releases/download/v0.29.1/geckodriver-v0.29.1-macos-aarch64.tar.gz",
                    "https://storage.googleapis.com/dev-infra-mirror/mozilla/geckodriver/0.29.1/geckodriver-v0.29.1-macos-aarch64.tar.gz"
               ],
               "sha256": "a1ec058b930fbfb684e30071ea47eec61bc18acb489914a9e0d095ede6088eea",
               "licenses": [
                    "reciprocal"
               ],
               "named_files": {
                    "GECKODRIVER": "geckodriver"
               }
          },
          "repositories": [
               {
                    "rule_class": "//web/internal:platform_archive.bzl%platform_archive",
                    "attributes": {
                         "name": "org_mozilla_geckodriver_macos_arm64",
                         "generator_name": "org_mozilla_geckodriver_macos_arm64",
                         "generator_function": "browser_repositories",
                         "generator_location": None,
                         "urls": [
                              "https://github.com/mozilla/geckodriver/releases/download/v0.29.1/geckodriver-v0.29.1-macos-aarch64.tar.gz",
                              "https://storage.googleapis.com/dev-infra-mirror/mozilla/geckodriver/0.29.1/geckodriver-v0.29.1-macos-aarch64.tar.gz"
                         ],
                         "sha256": "a1ec058b930fbfb684e30071ea47eec61bc18acb489914a9e0d095ede6088eea",
                         "licenses": [
                              "reciprocal"
                         ],
                         "named_files": {
                              "GECKODRIVER": "geckodriver"
                         }
                    },
                    "output_tree_hash": "52eb813734a17a779b5a2d4bee33187970d3df03ede745fee5fcb3d097c5f2b3"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_java_builtin//toolchains:local_java_repository.bzl%_local_java_repository_rule",
          "definition_information": "Repository local_jdk instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:85:6: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/local_java_repository.bzl:335:32: in local_java_repository\nRepository rule _local_java_repository_rule defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/local_java_repository.bzl:290:46: in <toplevel>\n",
          "original_attributes": {
               "name": "local_jdk",
               "generator_name": "local_jdk",
               "generator_function": "maybe",
               "generator_location": None,
               "build_file_content": "load(\"@rules_java//java:defs.bzl\", \"java_runtime\")\n\npackage(default_visibility = [\"//visibility:public\"])\n\nexports_files([\"WORKSPACE\", \"BUILD.bazel\"])\n\nfilegroup(\n    name = \"jre\",\n    srcs = glob(\n        [\n            \"jre/bin/**\",\n            \"jre/lib/**\",\n        ],\n        allow_empty = True,\n        # In some configurations, Java browser plugin is considered harmful and\n        # common antivirus software blocks access to npjp2.dll interfering with Bazel,\n        # so do not include it in JRE on Windows.\n        exclude = [\"jre/bin/plugin2/**\"],\n    ),\n)\n\nfilegroup(\n    name = \"jdk-bin\",\n    srcs = glob(\n        [\"bin/**\"],\n        # The JDK on Windows sometimes contains a directory called\n        # \"%systemroot%\", which is not a valid label.\n        exclude = [\"**/*%*/**\"],\n    ),\n)\n\n# This folder holds security policies.\nfilegroup(\n    name = \"jdk-conf\",\n    srcs = glob(\n        [\"conf/**\"],\n        allow_empty = True,\n    ),\n)\n\nfilegroup(\n    name = \"jdk-include\",\n    srcs = glob(\n        [\"include/**\"],\n        allow_empty = True,\n    ),\n)\n\nfilegroup(\n    name = \"jdk-lib\",\n    srcs = glob(\n        [\"lib/**\", \"release\"],\n        allow_empty = True,\n        exclude = [\n            \"lib/missioncontrol/**\",\n            \"lib/visualvm/**\",\n        ],\n    ),\n)\n\njava_runtime(\n    name = \"jdk\",\n    srcs = [\n        \":jdk-bin\",\n        \":jdk-conf\",\n        \":jdk-include\",\n        \":jdk-lib\",\n        \":jre\",\n    ],\n    # Provide the 'java` binary explicitly so that the correct path is used by\n    # Bazel even when the host platform differs from the execution platform.\n    # Exactly one of the two globs will be empty depending on the host platform.\n    # When --incompatible_disallow_empty_glob is enabled, each individual empty\n    # glob will fail without allow_empty = True, even if the overall result is\n    # non-empty.\n    java = glob([\"bin/java.exe\", \"bin/java\"], allow_empty = True)[0],\n    version = {RUNTIME_VERSION},\n)\n\nfilegroup(\n    name = \"jdk-jmods\",\n    srcs = glob(\n        [\"jmods/**\"],\n        allow_empty = True,\n    ),\n)\n\njava_runtime(\n    name = \"jdk-with-jmods\",\n    srcs = [\n        \":jdk-bin\",\n        \":jdk-conf\",\n        \":jdk-include\",\n        \":jdk-lib\",\n        \":jdk-jmods\",\n        \":jre\",\n    ],\n    java = glob([\"bin/java.exe\", \"bin/java\"], allow_empty = True)[0],\n    version = {RUNTIME_VERSION},\n)\n",
               "java_home": "",
               "version": ""
          },
          "repositories": [
               {
                    "rule_class": "@@rules_java_builtin//toolchains:local_java_repository.bzl%_local_java_repository_rule",
                    "attributes": {
                         "name": "local_jdk",
                         "generator_name": "local_jdk",
                         "generator_function": "maybe",
                         "generator_location": None,
                         "build_file_content": "load(\"@rules_java//java:defs.bzl\", \"java_runtime\")\n\npackage(default_visibility = [\"//visibility:public\"])\n\nexports_files([\"WORKSPACE\", \"BUILD.bazel\"])\n\nfilegroup(\n    name = \"jre\",\n    srcs = glob(\n        [\n            \"jre/bin/**\",\n            \"jre/lib/**\",\n        ],\n        allow_empty = True,\n        # In some configurations, Java browser plugin is considered harmful and\n        # common antivirus software blocks access to npjp2.dll interfering with Bazel,\n        # so do not include it in JRE on Windows.\n        exclude = [\"jre/bin/plugin2/**\"],\n    ),\n)\n\nfilegroup(\n    name = \"jdk-bin\",\n    srcs = glob(\n        [\"bin/**\"],\n        # The JDK on Windows sometimes contains a directory called\n        # \"%systemroot%\", which is not a valid label.\n        exclude = [\"**/*%*/**\"],\n    ),\n)\n\n# This folder holds security policies.\nfilegroup(\n    name = \"jdk-conf\",\n    srcs = glob(\n        [\"conf/**\"],\n        allow_empty = True,\n    ),\n)\n\nfilegroup(\n    name = \"jdk-include\",\n    srcs = glob(\n        [\"include/**\"],\n        allow_empty = True,\n    ),\n)\n\nfilegroup(\n    name = \"jdk-lib\",\n    srcs = glob(\n        [\"lib/**\", \"release\"],\n        allow_empty = True,\n        exclude = [\n            \"lib/missioncontrol/**\",\n            \"lib/visualvm/**\",\n        ],\n    ),\n)\n\njava_runtime(\n    name = \"jdk\",\n    srcs = [\n        \":jdk-bin\",\n        \":jdk-conf\",\n        \":jdk-include\",\n        \":jdk-lib\",\n        \":jre\",\n    ],\n    # Provide the 'java` binary explicitly so that the correct path is used by\n    # Bazel even when the host platform differs from the execution platform.\n    # Exactly one of the two globs will be empty depending on the host platform.\n    # When --incompatible_disallow_empty_glob is enabled, each individual empty\n    # glob will fail without allow_empty = True, even if the overall result is\n    # non-empty.\n    java = glob([\"bin/java.exe\", \"bin/java\"], allow_empty = True)[0],\n    version = {RUNTIME_VERSION},\n)\n\nfilegroup(\n    name = \"jdk-jmods\",\n    srcs = glob(\n        [\"jmods/**\"],\n        allow_empty = True,\n    ),\n)\n\njava_runtime(\n    name = \"jdk-with-jmods\",\n    srcs = [\n        \":jdk-bin\",\n        \":jdk-conf\",\n        \":jdk-include\",\n        \":jdk-lib\",\n        \":jdk-jmods\",\n        \":jre\",\n    ],\n    java = glob([\"bin/java.exe\", \"bin/java\"], allow_empty = True)[0],\n    version = {RUNTIME_VERSION},\n)\n",
                         "java_home": "",
                         "version": ""
                    },
                    "output_tree_hash": "a105a47d4f954665252e0c192181c2d4b0a12213ccec3651a34717605ba09729"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
          "definition_information": "Repository com_github_urllib3 instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:77:16: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/py_repositories.bzl:43:27: in py_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/py_repositories.bzl:50:17: in com_github_urllib3\nRepository rule http_archive defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in <toplevel>\n",
          "original_attributes": {
               "name": "com_github_urllib3",
               "generator_name": "com_github_urllib3",
               "generator_function": "py_repositories",
               "generator_location": None,
               "urls": [
                    "https://files.pythonhosted.org/packages/9a/8b/ea6d2beb2da6e331e9857d0a60b79ed4f72dcbc4e2c7f2d2521b0480fda2/urllib3-1.25.2.tar.gz"
               ],
               "sha256": "a53063d8b9210a7bdec15e7b272776b9d42b2fd6816401a0d43006ad2f9902db",
               "strip_prefix": "urllib3-1.25.2",
               "build_file": "//build_files:com_github_urllib3.BUILD"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
                    "attributes": {
                         "url": "",
                         "urls": [
                              "https://files.pythonhosted.org/packages/9a/8b/ea6d2beb2da6e331e9857d0a60b79ed4f72dcbc4e2c7f2d2521b0480fda2/urllib3-1.25.2.tar.gz"
                         ],
                         "sha256": "a53063d8b9210a7bdec15e7b272776b9d42b2fd6816401a0d43006ad2f9902db",
                         "integrity": "",
                         "netrc": "",
                         "auth_patterns": {},
                         "canonical_id": "",
                         "strip_prefix": "urllib3-1.25.2",
                         "add_prefix": "",
                         "type": "",
                         "patches": [],
                         "remote_file_urls": {},
                         "remote_file_integrity": {},
                         "remote_patches": {},
                         "remote_patch_strip": 0,
                         "patch_tool": "",
                         "patch_args": [
                              "-p0"
                         ],
                         "patch_cmds": [],
                         "patch_cmds_win": [],
                         "build_file": "//build_files:com_github_urllib3.BUILD",
                         "build_file_content": "",
                         "workspace_file_content": "",
                         "name": "com_github_urllib3"
                    },
                    "output_tree_hash": "e21bd22092242056e61bf4e7e62655be4e7f0b58c1e654210cae12468087fda5"
               }
          ]
     },
     {
          "original_rule_class": "//web/internal:platform_archive.bzl%platform_archive",
          "definition_information": "Repository org_chromium_chromedriver_macos_arm64 instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:59:21: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/versioned/browsers-0.3.3.bzl:27:34: in browser_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/versioned/browsers-0.3.3.bzl:107:21: in org_chromium_chromedriver\nRepository rule platform_archive defined at:\n  /Users/pcloudy/workspace/rules_webtesting/web/internal/platform_archive.bzl:81:35: in <toplevel>\n",
          "original_attributes": {
               "name": "org_chromium_chromedriver_macos_arm64",
               "generator_name": "org_chromium_chromedriver_macos_arm64",
               "generator_function": "browser_repositories",
               "generator_location": None,
               "urls": [
                    "https://storage.googleapis.com/chromium-browser-snapshots/Mac_Arm/902390/chromedriver_mac64.zip",
                    "https://storage.googleapis.com/dev-infra-mirror/chromium/902390/chromedriver_mac_arm64.zip"
               ],
               "sha256": "1f100aacf4bab4b3ac4218ecf654b17d66f2e07dd455f887bb3d9aa8d21862e1",
               "licenses": [
                    "reciprocal"
               ],
               "named_files": {
                    "CHROMEDRIVER": "chromedriver_mac64/chromedriver"
               }
          },
          "repositories": [
               {
                    "rule_class": "//web/internal:platform_archive.bzl%platform_archive",
                    "attributes": {
                         "name": "org_chromium_chromedriver_macos_arm64",
                         "generator_name": "org_chromium_chromedriver_macos_arm64",
                         "generator_function": "browser_repositories",
                         "generator_location": None,
                         "urls": [
                              "https://storage.googleapis.com/chromium-browser-snapshots/Mac_Arm/902390/chromedriver_mac64.zip",
                              "https://storage.googleapis.com/dev-infra-mirror/chromium/902390/chromedriver_mac_arm64.zip"
                         ],
                         "sha256": "1f100aacf4bab4b3ac4218ecf654b17d66f2e07dd455f887bb3d9aa8d21862e1",
                         "licenses": [
                              "reciprocal"
                         ],
                         "named_files": {
                              "CHROMEDRIVER": "chromedriver_mac64/chromedriver"
                         }
                    },
                    "output_tree_hash": "d56b3995a91a8dbe461305566518112f7525e9c1c1296aeaf148909bcabe8db7"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
          "definition_information": "Repository org_seleniumhq_selenium_selenium_remote_driver instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:73:18: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:70:55: in java_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:215:25: in org_seleniumhq_selenium_selenium_remote_driver\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/java.bzl:176:24: in java_import_external\nRepository rule jvm_import_external defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/jvm.bzl:236:38: in <toplevel>\n",
          "original_attributes": {
               "name": "org_seleniumhq_selenium_selenium_remote_driver",
               "generator_name": "org_seleniumhq_selenium_selenium_remote_driver",
               "generator_function": "java_repositories",
               "generator_location": None,
               "rule_name": "java_import",
               "licenses": [
                    "notice"
               ],
               "artifact_urls": [
                    "https://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-remote-driver/3.141.59/selenium-remote-driver-3.141.59.jar"
               ],
               "artifact_sha256": "9829fe57adf36743d785d0c2e7db504ba3ba0a3aacac652b8867cc854d2dfc45",
               "deps": [
                    "@com_google_guava_guava",
                    "@net_bytebuddy_byte_buddy",
                    "@com_squareup_okhttp3_okhttp",
                    "@com_squareup_okio_okio",
                    "@org_apache_commons_commons_exec",
                    "@org_seleniumhq_selenium_selenium_api"
               ],
               "testonly_": True
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
                    "attributes": {
                         "name": "org_seleniumhq_selenium_selenium_remote_driver",
                         "generator_name": "org_seleniumhq_selenium_selenium_remote_driver",
                         "generator_function": "java_repositories",
                         "generator_location": None,
                         "rule_name": "java_import",
                         "licenses": [
                              "notice"
                         ],
                         "artifact_urls": [
                              "https://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-remote-driver/3.141.59/selenium-remote-driver-3.141.59.jar"
                         ],
                         "artifact_sha256": "9829fe57adf36743d785d0c2e7db504ba3ba0a3aacac652b8867cc854d2dfc45",
                         "deps": [
                              "@com_google_guava_guava",
                              "@net_bytebuddy_byte_buddy",
                              "@com_squareup_okhttp3_okhttp",
                              "@com_squareup_okio_okio",
                              "@org_apache_commons_commons_exec",
                              "@org_seleniumhq_selenium_selenium_api"
                         ],
                         "testonly_": True
                    },
                    "output_tree_hash": "0dc2f861e68ce16535f4d3b0efc22aa32fbde5392af907121873a92e877306a8"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
          "definition_information": "Repository org_seleniumhq_selenium_selenium_api instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:73:18: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:68:45: in java_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:204:25: in org_seleniumhq_selenium_selenium_api\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/java.bzl:176:24: in java_import_external\nRepository rule jvm_import_external defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/jvm.bzl:236:38: in <toplevel>\n",
          "original_attributes": {
               "name": "org_seleniumhq_selenium_selenium_api",
               "generator_name": "org_seleniumhq_selenium_selenium_api",
               "generator_function": "java_repositories",
               "generator_location": None,
               "rule_name": "java_import",
               "licenses": [
                    "notice"
               ],
               "artifact_urls": [
                    "https://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-api/3.141.59/selenium-api-3.141.59.jar"
               ],
               "artifact_sha256": "8bfd5a736eccfc08866301ffc9b7f529e55976355c5799bed8392486df64dee5",
               "testonly_": True
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
                    "attributes": {
                         "name": "org_seleniumhq_selenium_selenium_api",
                         "generator_name": "org_seleniumhq_selenium_selenium_api",
                         "generator_function": "java_repositories",
                         "generator_location": None,
                         "rule_name": "java_import",
                         "licenses": [
                              "notice"
                         ],
                         "artifact_urls": [
                              "https://repo1.maven.org/maven2/org/seleniumhq/selenium/selenium-api/3.141.59/selenium-api-3.141.59.jar"
                         ],
                         "artifact_sha256": "8bfd5a736eccfc08866301ffc9b7f529e55976355c5799bed8392486df64dee5",
                         "testonly_": True
                    },
                    "output_tree_hash": "1b48829eb5af539eb5da5b2d9340655f2ab767bc056c21dada4ea934ed6de7f4"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
          "definition_information": "Repository com_google_guava_guava instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:73:18: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:50:31: in java_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:97:25: in com_google_guava_guava\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/java.bzl:176:24: in java_import_external\nRepository rule jvm_import_external defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/jvm.bzl:236:38: in <toplevel>\n",
          "original_attributes": {
               "name": "com_google_guava_guava",
               "generator_name": "com_google_guava_guava",
               "generator_function": "java_repositories",
               "generator_location": None,
               "rule_name": "java_import",
               "licenses": [
                    "notice"
               ],
               "artifact_urls": [
                    "https://repo1.maven.org/maven2/com/google/guava/guava/28.0-jre/guava-28.0-jre.jar"
               ],
               "artifact_sha256": "73e4d6ae5f0e8f9d292a4db83a2479b5468f83d972ac1ff36d6d0b43943b4f91",
               "exports": [
                    "@com_google_code_findbugs_jsr305",
                    "@com_google_errorprone_error_prone_annotations"
               ]
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
                    "attributes": {
                         "name": "com_google_guava_guava",
                         "generator_name": "com_google_guava_guava",
                         "generator_function": "java_repositories",
                         "generator_location": None,
                         "rule_name": "java_import",
                         "licenses": [
                              "notice"
                         ],
                         "artifact_urls": [
                              "https://repo1.maven.org/maven2/com/google/guava/guava/28.0-jre/guava-28.0-jre.jar"
                         ],
                         "artifact_sha256": "73e4d6ae5f0e8f9d292a4db83a2479b5468f83d972ac1ff36d6d0b43943b4f91",
                         "exports": [
                              "@com_google_code_findbugs_jsr305",
                              "@com_google_errorprone_error_prone_annotations"
                         ]
                    },
                    "output_tree_hash": "8a718ed1dc2b1ddcf0203224f566b47ecb407b12adb4ac600952ea912268b4bf"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
          "definition_information": "Repository org_apache_commons_commons_exec instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:73:18: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:60:40: in java_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:162:25: in org_apache_commons_commons_exec\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/java.bzl:176:24: in java_import_external\nRepository rule jvm_import_external defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/jvm.bzl:236:38: in <toplevel>\n",
          "original_attributes": {
               "name": "org_apache_commons_commons_exec",
               "generator_name": "org_apache_commons_commons_exec",
               "generator_function": "java_repositories",
               "generator_location": None,
               "rule_name": "java_import",
               "licenses": [
                    "notice"
               ],
               "artifact_urls": [
                    "https://repo1.maven.org/maven2/org/apache/commons/commons-exec/1.3/commons-exec-1.3.jar"
               ],
               "artifact_sha256": "cb49812dc1bfb0ea4f20f398bcae1a88c6406e213e67f7524fb10d4f8ad9347b"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
                    "attributes": {
                         "name": "org_apache_commons_commons_exec",
                         "generator_name": "org_apache_commons_commons_exec",
                         "generator_function": "java_repositories",
                         "generator_location": None,
                         "rule_name": "java_import",
                         "licenses": [
                              "notice"
                         ],
                         "artifact_urls": [
                              "https://repo1.maven.org/maven2/org/apache/commons/commons-exec/1.3/commons-exec-1.3.jar"
                         ],
                         "artifact_sha256": "cb49812dc1bfb0ea4f20f398bcae1a88c6406e213e67f7524fb10d4f8ad9347b"
                    },
                    "output_tree_hash": "d651d58b171f0e4664c28b684af348d81592fd32a40b2b3d4251791dca2f980d"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
          "definition_information": "Repository com_squareup_okio_okio instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:73:18: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:54:31: in java_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:125:25: in com_squareup_okio_okio\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/java.bzl:176:24: in java_import_external\nRepository rule jvm_import_external defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/jvm.bzl:236:38: in <toplevel>\n",
          "original_attributes": {
               "name": "com_squareup_okio_okio",
               "generator_name": "com_squareup_okio_okio",
               "generator_function": "java_repositories",
               "generator_location": None,
               "rule_name": "java_import",
               "licenses": [
                    "notice"
               ],
               "artifact_urls": [
                    "https://repo1.maven.org/maven2/com/squareup/okio/okio/2.3.0/okio-2.3.0.jar"
               ],
               "artifact_sha256": "1c52079b6159b096181a2fad4df7f15423ee6c66266d1dcb0264bf37c58178b0",
               "deps": [
                    "@com_google_code_findbugs_jsr305",
                    "@org_jetbrains_kotlin_kotlin_stdlib"
               ]
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
                    "attributes": {
                         "name": "com_squareup_okio_okio",
                         "generator_name": "com_squareup_okio_okio",
                         "generator_function": "java_repositories",
                         "generator_location": None,
                         "rule_name": "java_import",
                         "licenses": [
                              "notice"
                         ],
                         "artifact_urls": [
                              "https://repo1.maven.org/maven2/com/squareup/okio/okio/2.3.0/okio-2.3.0.jar"
                         ],
                         "artifact_sha256": "1c52079b6159b096181a2fad4df7f15423ee6c66266d1dcb0264bf37c58178b0",
                         "deps": [
                              "@com_google_code_findbugs_jsr305",
                              "@org_jetbrains_kotlin_kotlin_stdlib"
                         ]
                    },
                    "output_tree_hash": "cbd9ed87139fd1720363102a1786eaa8d60bdcd1e565c87f51405d23428ff5b5"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
          "definition_information": "Repository com_squareup_okhttp3_okhttp instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:73:18: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:52:36: in java_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:111:25: in com_squareup_okhttp3_okhttp\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/java.bzl:176:24: in java_import_external\nRepository rule jvm_import_external defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/jvm.bzl:236:38: in <toplevel>\n",
          "original_attributes": {
               "name": "com_squareup_okhttp3_okhttp",
               "generator_name": "com_squareup_okhttp3_okhttp",
               "generator_function": "java_repositories",
               "generator_location": None,
               "rule_name": "java_import",
               "licenses": [
                    "notice"
               ],
               "artifact_urls": [
                    "https://repo1.maven.org/maven2/com/squareup/okhttp3/okhttp/4.1.0/okhttp-4.1.0.jar"
               ],
               "artifact_sha256": "20f483a62087faa1dc8240150fa500b0a42c822735a12481ae32c5238d9922cc",
               "deps": [
                    "@com_squareup_okio_okio",
                    "@com_google_code_findbugs_jsr305"
               ]
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
                    "attributes": {
                         "name": "com_squareup_okhttp3_okhttp",
                         "generator_name": "com_squareup_okhttp3_okhttp",
                         "generator_function": "java_repositories",
                         "generator_location": None,
                         "rule_name": "java_import",
                         "licenses": [
                              "notice"
                         ],
                         "artifact_urls": [
                              "https://repo1.maven.org/maven2/com/squareup/okhttp3/okhttp/4.1.0/okhttp-4.1.0.jar"
                         ],
                         "artifact_sha256": "20f483a62087faa1dc8240150fa500b0a42c822735a12481ae32c5238d9922cc",
                         "deps": [
                              "@com_squareup_okio_okio",
                              "@com_google_code_findbugs_jsr305"
                         ]
                    },
                    "output_tree_hash": "6a752664aef1132093197f52d7807081fb40d19ff87eb9905ca9243129aea44f"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
          "definition_information": "Repository net_bytebuddy_byte_buddy instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:73:18: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:58:33: in java_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:151:25: in net_bytebuddy_byte_buddy\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/java.bzl:176:24: in java_import_external\nRepository rule jvm_import_external defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/jvm.bzl:236:38: in <toplevel>\n",
          "original_attributes": {
               "name": "net_bytebuddy_byte_buddy",
               "generator_name": "net_bytebuddy_byte_buddy",
               "generator_function": "java_repositories",
               "generator_location": None,
               "rule_name": "java_import",
               "licenses": [
                    "notice"
               ],
               "artifact_urls": [
                    "https://repo1.maven.org/maven2/net/bytebuddy/byte-buddy/1.9.16/byte-buddy-1.9.16.jar"
               ],
               "artifact_sha256": "6b71e4f70c96b67d420f592148aa4fd1966aba458b35d11f491ff13de97dc862",
               "deps": [
                    "@com_google_code_findbugs_jsr305"
               ]
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
                    "attributes": {
                         "name": "net_bytebuddy_byte_buddy",
                         "generator_name": "net_bytebuddy_byte_buddy",
                         "generator_function": "java_repositories",
                         "generator_location": None,
                         "rule_name": "java_import",
                         "licenses": [
                              "notice"
                         ],
                         "artifact_urls": [
                              "https://repo1.maven.org/maven2/net/bytebuddy/byte-buddy/1.9.16/byte-buddy-1.9.16.jar"
                         ],
                         "artifact_sha256": "6b71e4f70c96b67d420f592148aa4fd1966aba458b35d11f491ff13de97dc862",
                         "deps": [
                              "@com_google_code_findbugs_jsr305"
                         ]
                    },
                    "output_tree_hash": "6802a267c1abffa6f5f98782d10de233433a00ff6223554c1ec151d75e880f16"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
          "definition_information": "Repository com_google_errorprone_error_prone_annotations instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:73:18: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:48:54: in java_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:87:25: in com_google_errorprone_error_prone_annotations\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/java.bzl:176:24: in java_import_external\nRepository rule jvm_import_external defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/jvm.bzl:236:38: in <toplevel>\n",
          "original_attributes": {
               "name": "com_google_errorprone_error_prone_annotations",
               "generator_name": "com_google_errorprone_error_prone_annotations",
               "generator_function": "java_repositories",
               "generator_location": None,
               "rule_name": "java_import",
               "licenses": [
                    "notice"
               ],
               "artifact_urls": [
                    "https://repo1.maven.org/maven2/com/google/errorprone/error_prone_annotations/2.3.3/error_prone_annotations-2.3.3.jar"
               ],
               "artifact_sha256": "ec59f1b702d9afc09e8c3929f5c42777dec623a6ea2731ac694332c7d7680f5a"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
                    "attributes": {
                         "name": "com_google_errorprone_error_prone_annotations",
                         "generator_name": "com_google_errorprone_error_prone_annotations",
                         "generator_function": "java_repositories",
                         "generator_location": None,
                         "rule_name": "java_import",
                         "licenses": [
                              "notice"
                         ],
                         "artifact_urls": [
                              "https://repo1.maven.org/maven2/com/google/errorprone/error_prone_annotations/2.3.3/error_prone_annotations-2.3.3.jar"
                         ],
                         "artifact_sha256": "ec59f1b702d9afc09e8c3929f5c42777dec623a6ea2731ac694332c7d7680f5a"
                    },
                    "output_tree_hash": "07a4da7d6e7835a64b3dde37341306cdb907f18dc54bf5fa4288243baf6e8c95"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
          "definition_information": "Repository com_google_code_findbugs_jsr305 instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:73:18: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:43:40: in java_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:77:25: in com_google_code_findbugs_jsr305\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/java.bzl:176:24: in java_import_external\nRepository rule jvm_import_external defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/jvm.bzl:236:38: in <toplevel>\n",
          "original_attributes": {
               "name": "com_google_code_findbugs_jsr305",
               "generator_name": "com_google_code_findbugs_jsr305",
               "generator_function": "java_repositories",
               "generator_location": None,
               "rule_name": "java_import",
               "licenses": [
                    "notice"
               ],
               "artifact_urls": [
                    "https://repo1.maven.org/maven2/com/google/code/findbugs/jsr305/3.0.2/jsr305-3.0.2.jar"
               ],
               "artifact_sha256": "766ad2a0783f2687962c8ad74ceecc38a28b9f72a2d085ee438b7813e928d0c7"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
                    "attributes": {
                         "name": "com_google_code_findbugs_jsr305",
                         "generator_name": "com_google_code_findbugs_jsr305",
                         "generator_function": "java_repositories",
                         "generator_location": None,
                         "rule_name": "java_import",
                         "licenses": [
                              "notice"
                         ],
                         "artifact_urls": [
                              "https://repo1.maven.org/maven2/com/google/code/findbugs/jsr305/3.0.2/jsr305-3.0.2.jar"
                         ],
                         "artifact_sha256": "766ad2a0783f2687962c8ad74ceecc38a28b9f72a2d085ee438b7813e928d0c7"
                    },
                    "output_tree_hash": "15e7816a484354436230882b94dc6d7acfdb626ec4c95372decf822ad002e03a"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
          "definition_information": "Repository org_jetbrains_kotlin_kotlin_stdlib instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:73:18: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:64:43: in java_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/java_repositories.bzl:184:25: in org_jetbrains_kotlin_kotlin_stdlib\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/java.bzl:176:24: in java_import_external\nRepository rule jvm_import_external defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/jvm.bzl:236:38: in <toplevel>\n",
          "original_attributes": {
               "name": "org_jetbrains_kotlin_kotlin_stdlib",
               "generator_name": "org_jetbrains_kotlin_kotlin_stdlib",
               "generator_function": "java_repositories",
               "generator_location": None,
               "rule_name": "java_import",
               "licenses": [
                    "notice"
               ],
               "artifact_urls": [
                    "https://repo1.maven.org/maven2/org/jetbrains/kotlin/kotlin-stdlib/1.3.41/kotlin-stdlib-1.3.41.jar"
               ],
               "artifact_sha256": "6ea3d0921b26919b286f05cbdb906266666a36f9a7c096197114f7495708ffbc"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:jvm.bzl%jvm_import_external",
                    "attributes": {
                         "name": "org_jetbrains_kotlin_kotlin_stdlib",
                         "generator_name": "org_jetbrains_kotlin_kotlin_stdlib",
                         "generator_function": "java_repositories",
                         "generator_location": None,
                         "rule_name": "java_import",
                         "licenses": [
                              "notice"
                         ],
                         "artifact_urls": [
                              "https://repo1.maven.org/maven2/org/jetbrains/kotlin/kotlin-stdlib/1.3.41/kotlin-stdlib-1.3.41.jar"
                         ],
                         "artifact_sha256": "6ea3d0921b26919b286f05cbdb906266666a36f9a7c096197114f7495708ffbc"
                    },
                    "output_tree_hash": "f1765feec6391042b7dc91e206e109a1b7f1c449286bcc487ab92bb99a042245"
               }
          ]
     },
     {
          "original_rule_class": "@@io_bazel_rules_scala//scala:scala_maven_import_external.bzl%jvm_import_external",
          "definition_information": "Repository io_bazel_rules_scala_scalatest instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:106:23: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/testing/scalatest.bzl:4:18: in scalatest_repositories\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scalatest/scalatest.bzl:10:17: in scalatest_repositories\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/third_party/repositories/repositories.bzl:78:37: in repositories\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scala/scala_maven_import_external.bzl:253:30: in scala_maven_import_external\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scala/scala_maven_import_external.bzl:289:24: in jvm_maven_import_external\nRepository rule jvm_import_external defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scala/scala_maven_import_external.bzl:218:38: in <toplevel>\n",
          "original_attributes": {
               "name": "io_bazel_rules_scala_scalatest",
               "generator_name": "io_bazel_rules_scala_scalatest",
               "generator_function": "scalatest_repositories",
               "generator_location": None,
               "rule_name": "scala_import",
               "licenses": [
                    "notice"
               ],
               "jar_urls": [
                    "https://repo.maven.apache.org/maven2/org/scalatest/scalatest_2.12/3.2.9/scalatest_2.12-3.2.9.jar",
                    "https://maven-central.storage-download.googleapis.com/maven2/org/scalatest/scalatest_2.12/3.2.9/scalatest_2.12-3.2.9.jar",
                    "https://mirror.bazel.build/repo1.maven.org/maven2/org/scalatest/scalatest_2.12/3.2.9/scalatest_2.12-3.2.9.jar",
                    "https://jcenter.bintray.com/org/scalatest/scalatest_2.12/3.2.9/scalatest_2.12-3.2.9.jar"
               ],
               "artifact_sha256": "ed4a7e0a2373505ae5b9c4811fa2d2d167f5388556cdcb49bce11f27e18b90fa",
               "rule_load": "load(\"@io_bazel_rules_scala//scala:scala_import.bzl\", \"scala_import\")",
               "srcjar_urls": [
                    "https://repo.maven.apache.org/maven2/org/scalatest/scalatest_2.12/3.2.9/scalatest_2.12-3.2.9-sources.jar",
                    "https://maven-central.storage-download.googleapis.com/maven2/org/scalatest/scalatest_2.12/3.2.9/scalatest_2.12-3.2.9-sources.jar",
                    "https://mirror.bazel.build/repo1.maven.org/maven2/org/scalatest/scalatest_2.12/3.2.9/scalatest_2.12-3.2.9-sources.jar",
                    "https://jcenter.bintray.com/org/scalatest/scalatest_2.12/3.2.9/scalatest_2.12-3.2.9-sources.jar"
               ],
               "deps": [],
               "runtime_deps": [],
               "testonly_": False
          },
          "repositories": [
               {
                    "rule_class": "@@io_bazel_rules_scala//scala:scala_maven_import_external.bzl%jvm_import_external",
                    "attributes": {
                         "name": "io_bazel_rules_scala_scalatest",
                         "generator_name": "io_bazel_rules_scala_scalatest",
                         "generator_function": "scalatest_repositories",
                         "generator_location": None,
                         "rule_name": "scala_import",
                         "licenses": [
                              "notice"
                         ],
                         "jar_urls": [
                              "https://repo.maven.apache.org/maven2/org/scalatest/scalatest_2.12/3.2.9/scalatest_2.12-3.2.9.jar",
                              "https://maven-central.storage-download.googleapis.com/maven2/org/scalatest/scalatest_2.12/3.2.9/scalatest_2.12-3.2.9.jar",
                              "https://mirror.bazel.build/repo1.maven.org/maven2/org/scalatest/scalatest_2.12/3.2.9/scalatest_2.12-3.2.9.jar",
                              "https://jcenter.bintray.com/org/scalatest/scalatest_2.12/3.2.9/scalatest_2.12-3.2.9.jar"
                         ],
                         "artifact_sha256": "ed4a7e0a2373505ae5b9c4811fa2d2d167f5388556cdcb49bce11f27e18b90fa",
                         "rule_load": "load(\"@io_bazel_rules_scala//scala:scala_import.bzl\", \"scala_import\")",
                         "srcjar_urls": [
                              "https://repo.maven.apache.org/maven2/org/scalatest/scalatest_2.12/3.2.9/scalatest_2.12-3.2.9-sources.jar",
                              "https://maven-central.storage-download.googleapis.com/maven2/org/scalatest/scalatest_2.12/3.2.9/scalatest_2.12-3.2.9-sources.jar",
                              "https://mirror.bazel.build/repo1.maven.org/maven2/org/scalatest/scalatest_2.12/3.2.9/scalatest_2.12-3.2.9-sources.jar",
                              "https://jcenter.bintray.com/org/scalatest/scalatest_2.12/3.2.9/scalatest_2.12-3.2.9-sources.jar"
                         ],
                         "deps": [],
                         "runtime_deps": [],
                         "testonly_": False
                    },
                    "output_tree_hash": "3f0d4f66f62f557823fd819c4cd0fd884b784285cee4cd7bc78ed355ea0d8357"
               }
          ]
     },
     {
          "original_rule_class": "@@io_bazel_rules_scala//scala:scala_maven_import_external.bzl%jvm_import_external",
          "definition_information": "Repository io_bazel_rules_scala_scalatest_compatible instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:106:23: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/testing/scalatest.bzl:4:18: in scalatest_repositories\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scalatest/scalatest.bzl:10:17: in scalatest_repositories\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/third_party/repositories/repositories.bzl:78:37: in repositories\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scala/scala_maven_import_external.bzl:253:30: in scala_maven_import_external\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scala/scala_maven_import_external.bzl:289:24: in jvm_maven_import_external\nRepository rule jvm_import_external defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scala/scala_maven_import_external.bzl:218:38: in <toplevel>\n",
          "original_attributes": {
               "name": "io_bazel_rules_scala_scalatest_compatible",
               "generator_name": "io_bazel_rules_scala_scalatest_compatible",
               "generator_function": "scalatest_repositories",
               "generator_location": None,
               "rule_name": "scala_import",
               "licenses": [
                    "notice"
               ],
               "jar_urls": [
                    "https://repo.maven.apache.org/maven2/org/scalatest/scalatest-compatible/3.2.9/scalatest-compatible-3.2.9.jar",
                    "https://maven-central.storage-download.googleapis.com/maven2/org/scalatest/scalatest-compatible/3.2.9/scalatest-compatible-3.2.9.jar",
                    "https://mirror.bazel.build/repo1.maven.org/maven2/org/scalatest/scalatest-compatible/3.2.9/scalatest-compatible-3.2.9.jar",
                    "https://jcenter.bintray.com/org/scalatest/scalatest-compatible/3.2.9/scalatest-compatible-3.2.9.jar"
               ],
               "artifact_sha256": "7e5f1193af2fd88c432c4b80ce3641e4b1d062f421d8a0fcc43af9a19bb7c2eb",
               "rule_load": "load(\"@io_bazel_rules_scala//scala:scala_import.bzl\", \"scala_import\")",
               "srcjar_urls": [
                    "https://repo.maven.apache.org/maven2/org/scalatest/scalatest-compatible/3.2.9/scalatest-compatible-3.2.9-sources.jar",
                    "https://maven-central.storage-download.googleapis.com/maven2/org/scalatest/scalatest-compatible/3.2.9/scalatest-compatible-3.2.9-sources.jar",
                    "https://mirror.bazel.build/repo1.maven.org/maven2/org/scalatest/scalatest-compatible/3.2.9/scalatest-compatible-3.2.9-sources.jar",
                    "https://jcenter.bintray.com/org/scalatest/scalatest-compatible/3.2.9/scalatest-compatible-3.2.9-sources.jar"
               ],
               "deps": [],
               "runtime_deps": [],
               "testonly_": False
          },
          "repositories": [
               {
                    "rule_class": "@@io_bazel_rules_scala//scala:scala_maven_import_external.bzl%jvm_import_external",
                    "attributes": {
                         "name": "io_bazel_rules_scala_scalatest_compatible",
                         "generator_name": "io_bazel_rules_scala_scalatest_compatible",
                         "generator_function": "scalatest_repositories",
                         "generator_location": None,
                         "rule_name": "scala_import",
                         "licenses": [
                              "notice"
                         ],
                         "jar_urls": [
                              "https://repo.maven.apache.org/maven2/org/scalatest/scalatest-compatible/3.2.9/scalatest-compatible-3.2.9.jar",
                              "https://maven-central.storage-download.googleapis.com/maven2/org/scalatest/scalatest-compatible/3.2.9/scalatest-compatible-3.2.9.jar",
                              "https://mirror.bazel.build/repo1.maven.org/maven2/org/scalatest/scalatest-compatible/3.2.9/scalatest-compatible-3.2.9.jar",
                              "https://jcenter.bintray.com/org/scalatest/scalatest-compatible/3.2.9/scalatest-compatible-3.2.9.jar"
                         ],
                         "artifact_sha256": "7e5f1193af2fd88c432c4b80ce3641e4b1d062f421d8a0fcc43af9a19bb7c2eb",
                         "rule_load": "load(\"@io_bazel_rules_scala//scala:scala_import.bzl\", \"scala_import\")",
                         "srcjar_urls": [
                              "https://repo.maven.apache.org/maven2/org/scalatest/scalatest-compatible/3.2.9/scalatest-compatible-3.2.9-sources.jar",
                              "https://maven-central.storage-download.googleapis.com/maven2/org/scalatest/scalatest-compatible/3.2.9/scalatest-compatible-3.2.9-sources.jar",
                              "https://mirror.bazel.build/repo1.maven.org/maven2/org/scalatest/scalatest-compatible/3.2.9/scalatest-compatible-3.2.9-sources.jar",
                              "https://jcenter.bintray.com/org/scalatest/scalatest-compatible/3.2.9/scalatest-compatible-3.2.9-sources.jar"
                         ],
                         "deps": [],
                         "runtime_deps": [],
                         "testonly_": False
                    },
                    "output_tree_hash": "e6e63052b83212d711f0dac20d0ee5d94eaed3794a322604bf7f81bb189ba23c"
               }
          ]
     },
     {
          "original_rule_class": "@@io_bazel_rules_scala//scala:scala_maven_import_external.bzl%jvm_import_external",
          "definition_information": "Repository io_bazel_rules_scala_scalactic instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:106:23: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/testing/scalatest.bzl:4:18: in scalatest_repositories\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scalatest/scalatest.bzl:10:17: in scalatest_repositories\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/third_party/repositories/repositories.bzl:78:37: in repositories\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scala/scala_maven_import_external.bzl:253:30: in scala_maven_import_external\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scala/scala_maven_import_external.bzl:289:24: in jvm_maven_import_external\nRepository rule jvm_import_external defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scala/scala_maven_import_external.bzl:218:38: in <toplevel>\n",
          "original_attributes": {
               "name": "io_bazel_rules_scala_scalactic",
               "generator_name": "io_bazel_rules_scala_scalactic",
               "generator_function": "scalatest_repositories",
               "generator_location": None,
               "rule_name": "scala_import",
               "licenses": [
                    "notice"
               ],
               "jar_urls": [
                    "https://repo.maven.apache.org/maven2/org/scalactic/scalactic_2.12/3.2.9/scalactic_2.12-3.2.9.jar",
                    "https://maven-central.storage-download.googleapis.com/maven2/org/scalactic/scalactic_2.12/3.2.9/scalactic_2.12-3.2.9.jar",
                    "https://mirror.bazel.build/repo1.maven.org/maven2/org/scalactic/scalactic_2.12/3.2.9/scalactic_2.12-3.2.9.jar",
                    "https://jcenter.bintray.com/org/scalactic/scalactic_2.12/3.2.9/scalactic_2.12-3.2.9.jar"
               ],
               "artifact_sha256": "a5f01a0ecb7479b4f43e03147094279609d66fdaa04a9cb3238510d7c4dbc22a",
               "rule_load": "load(\"@io_bazel_rules_scala//scala:scala_import.bzl\", \"scala_import\")",
               "srcjar_urls": [
                    "https://repo.maven.apache.org/maven2/org/scalactic/scalactic_2.12/3.2.9/scalactic_2.12-3.2.9-sources.jar",
                    "https://maven-central.storage-download.googleapis.com/maven2/org/scalactic/scalactic_2.12/3.2.9/scalactic_2.12-3.2.9-sources.jar",
                    "https://mirror.bazel.build/repo1.maven.org/maven2/org/scalactic/scalactic_2.12/3.2.9/scalactic_2.12-3.2.9-sources.jar",
                    "https://jcenter.bintray.com/org/scalactic/scalactic_2.12/3.2.9/scalactic_2.12-3.2.9-sources.jar"
               ],
               "deps": [],
               "runtime_deps": [],
               "testonly_": False
          },
          "repositories": [
               {
                    "rule_class": "@@io_bazel_rules_scala//scala:scala_maven_import_external.bzl%jvm_import_external",
                    "attributes": {
                         "name": "io_bazel_rules_scala_scalactic",
                         "generator_name": "io_bazel_rules_scala_scalactic",
                         "generator_function": "scalatest_repositories",
                         "generator_location": None,
                         "rule_name": "scala_import",
                         "licenses": [
                              "notice"
                         ],
                         "jar_urls": [
                              "https://repo.maven.apache.org/maven2/org/scalactic/scalactic_2.12/3.2.9/scalactic_2.12-3.2.9.jar",
                              "https://maven-central.storage-download.googleapis.com/maven2/org/scalactic/scalactic_2.12/3.2.9/scalactic_2.12-3.2.9.jar",
                              "https://mirror.bazel.build/repo1.maven.org/maven2/org/scalactic/scalactic_2.12/3.2.9/scalactic_2.12-3.2.9.jar",
                              "https://jcenter.bintray.com/org/scalactic/scalactic_2.12/3.2.9/scalactic_2.12-3.2.9.jar"
                         ],
                         "artifact_sha256": "a5f01a0ecb7479b4f43e03147094279609d66fdaa04a9cb3238510d7c4dbc22a",
                         "rule_load": "load(\"@io_bazel_rules_scala//scala:scala_import.bzl\", \"scala_import\")",
                         "srcjar_urls": [
                              "https://repo.maven.apache.org/maven2/org/scalactic/scalactic_2.12/3.2.9/scalactic_2.12-3.2.9-sources.jar",
                              "https://maven-central.storage-download.googleapis.com/maven2/org/scalactic/scalactic_2.12/3.2.9/scalactic_2.12-3.2.9-sources.jar",
                              "https://mirror.bazel.build/repo1.maven.org/maven2/org/scalactic/scalactic_2.12/3.2.9/scalactic_2.12-3.2.9-sources.jar",
                              "https://jcenter.bintray.com/org/scalactic/scalactic_2.12/3.2.9/scalactic_2.12-3.2.9-sources.jar"
                         ],
                         "deps": [],
                         "runtime_deps": [],
                         "testonly_": False
                    },
                    "output_tree_hash": "3b5554adcb01e431a495b670f72c13d3e3c0da3e4eba04e0be47dada60b90614"
               }
          ]
     },
     {
          "original_rule_class": "@@io_bazel_rules_scala//scala:scala_maven_import_external.bzl%jvm_import_external",
          "definition_information": "Repository io_bazel_rules_scala_scalatest_core instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:106:23: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/testing/scalatest.bzl:4:18: in scalatest_repositories\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scalatest/scalatest.bzl:10:17: in scalatest_repositories\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/third_party/repositories/repositories.bzl:78:37: in repositories\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scala/scala_maven_import_external.bzl:253:30: in scala_maven_import_external\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scala/scala_maven_import_external.bzl:289:24: in jvm_maven_import_external\nRepository rule jvm_import_external defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_scala/scala/scala_maven_import_external.bzl:218:38: in <toplevel>\n",
          "original_attributes": {
               "name": "io_bazel_rules_scala_scalatest_core",
               "generator_name": "io_bazel_rules_scala_scalatest_core",
               "generator_function": "scalatest_repositories",
               "generator_location": None,
               "rule_name": "scala_import",
               "licenses": [
                    "notice"
               ],
               "jar_urls": [
                    "https://repo.maven.apache.org/maven2/org/scalatest/scalatest-core_2.12/3.2.9/scalatest-core_2.12-3.2.9.jar",
                    "https://maven-central.storage-download.googleapis.com/maven2/org/scalatest/scalatest-core_2.12/3.2.9/scalatest-core_2.12-3.2.9.jar",
                    "https://mirror.bazel.build/repo1.maven.org/maven2/org/scalatest/scalatest-core_2.12/3.2.9/scalatest-core_2.12-3.2.9.jar",
                    "https://jcenter.bintray.com/org/scalatest/scalatest-core_2.12/3.2.9/scalatest-core_2.12-3.2.9.jar"
               ],
               "artifact_sha256": "8d5bc6b847caaf221fa42cc214dcd1c70fd758aef384a2b6498463db0caf8e3c",
               "rule_load": "load(\"@io_bazel_rules_scala//scala:scala_import.bzl\", \"scala_import\")",
               "srcjar_urls": [
                    "https://repo.maven.apache.org/maven2/org/scalatest/scalatest-core_2.12/3.2.9/scalatest-core_2.12-3.2.9-sources.jar",
                    "https://maven-central.storage-download.googleapis.com/maven2/org/scalatest/scalatest-core_2.12/3.2.9/scalatest-core_2.12-3.2.9-sources.jar",
                    "https://mirror.bazel.build/repo1.maven.org/maven2/org/scalatest/scalatest-core_2.12/3.2.9/scalatest-core_2.12-3.2.9-sources.jar",
                    "https://jcenter.bintray.com/org/scalatest/scalatest-core_2.12/3.2.9/scalatest-core_2.12-3.2.9-sources.jar"
               ],
               "deps": [],
               "runtime_deps": [],
               "testonly_": False
          },
          "repositories": [
               {
                    "rule_class": "@@io_bazel_rules_scala//scala:scala_maven_import_external.bzl%jvm_import_external",
                    "attributes": {
                         "name": "io_bazel_rules_scala_scalatest_core",
                         "generator_name": "io_bazel_rules_scala_scalatest_core",
                         "generator_function": "scalatest_repositories",
                         "generator_location": None,
                         "rule_name": "scala_import",
                         "licenses": [
                              "notice"
                         ],
                         "jar_urls": [
                              "https://repo.maven.apache.org/maven2/org/scalatest/scalatest-core_2.12/3.2.9/scalatest-core_2.12-3.2.9.jar",
                              "https://maven-central.storage-download.googleapis.com/maven2/org/scalatest/scalatest-core_2.12/3.2.9/scalatest-core_2.12-3.2.9.jar",
                              "https://mirror.bazel.build/repo1.maven.org/maven2/org/scalatest/scalatest-core_2.12/3.2.9/scalatest-core_2.12-3.2.9.jar",
                              "https://jcenter.bintray.com/org/scalatest/scalatest-core_2.12/3.2.9/scalatest-core_2.12-3.2.9.jar"
                         ],
                         "artifact_sha256": "8d5bc6b847caaf221fa42cc214dcd1c70fd758aef384a2b6498463db0caf8e3c",
                         "rule_load": "load(\"@io_bazel_rules_scala//scala:scala_import.bzl\", \"scala_import\")",
                         "srcjar_urls": [
                              "https://repo.maven.apache.org/maven2/org/scalatest/scalatest-core_2.12/3.2.9/scalatest-core_2.12-3.2.9-sources.jar",
                              "https://maven-central.storage-download.googleapis.com/maven2/org/scalatest/scalatest-core_2.12/3.2.9/scalatest-core_2.12-3.2.9-sources.jar",
                              "https://mirror.bazel.build/repo1.maven.org/maven2/org/scalatest/scalatest-core_2.12/3.2.9/scalatest-core_2.12-3.2.9-sources.jar",
                              "https://jcenter.bintray.com/org/scalatest/scalatest-core_2.12/3.2.9/scalatest-core_2.12-3.2.9-sources.jar"
                         ],
                         "deps": [],
                         "runtime_deps": [],
                         "testonly_": False
                    },
                    "output_tree_hash": "2fc4866b6e71412739797302cd1b29412a169b171e0b7eda300ca8045bb7cb40"
               }
          ]
     },
     {
          "original_rule_class": "@@io_bazel_rules_go//go/private:nogo.bzl%go_register_nogo",
          "definition_information": "Repository io_bazel_rules_nogo instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:40:22: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_go/go/private/repositories.bzl:261:12: in go_rules_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_go/go/private/repositories.bzl:269:18: in _maybe\nRepository rule go_register_nogo defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_go/go/private/nogo.bzl:31:35: in <toplevel>\n",
          "original_attributes": {
               "name": "io_bazel_rules_nogo",
               "generator_name": "io_bazel_rules_nogo",
               "generator_function": "go_rules_dependencies",
               "generator_location": None,
               "nogo": "@io_bazel_rules_go//:default_nogo"
          },
          "repositories": [
               {
                    "rule_class": "@@io_bazel_rules_go//go/private:nogo.bzl%go_register_nogo",
                    "attributes": {
                         "name": "io_bazel_rules_nogo",
                         "generator_name": "io_bazel_rules_nogo",
                         "generator_function": "go_rules_dependencies",
                         "generator_location": None,
                         "nogo": "@io_bazel_rules_go//:default_nogo"
                    },
                    "output_tree_hash": "e4772e86da6e3bc0887a236dc36324e6b44be0e0315adf28750fcb363acae413"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
          "definition_information": "Repository remote_java_tools_darwin_arm64 instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:374:21: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:59:14: in java_tools_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\nRepository rule http_archive defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in <toplevel>\n",
          "original_attributes": {
               "name": "remote_java_tools_darwin_arm64",
               "generator_name": "remote_java_tools_darwin_arm64",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "urls": [
                    "https://mirror.bazel.build/bazel_java_tools/releases/java/v13.6.1/java_tools_darwin_arm64-v13.6.1.zip",
                    "https://github.com/bazelbuild/java_tools/releases/download/java_13.6.1/java_tools_darwin_arm64-v13.6.1.zip"
               ],
               "sha256": "eb54c4e5fa23d6e9e9fc14c106a682dbefc54659d8e389a2f3c0d61d51cae274"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
                    "attributes": {
                         "url": "",
                         "urls": [
                              "https://mirror.bazel.build/bazel_java_tools/releases/java/v13.6.1/java_tools_darwin_arm64-v13.6.1.zip",
                              "https://github.com/bazelbuild/java_tools/releases/download/java_13.6.1/java_tools_darwin_arm64-v13.6.1.zip"
                         ],
                         "sha256": "eb54c4e5fa23d6e9e9fc14c106a682dbefc54659d8e389a2f3c0d61d51cae274",
                         "integrity": "",
                         "netrc": "",
                         "auth_patterns": {},
                         "canonical_id": "",
                         "strip_prefix": "",
                         "add_prefix": "",
                         "type": "",
                         "patches": [],
                         "remote_file_urls": {},
                         "remote_file_integrity": {},
                         "remote_patches": {},
                         "remote_patch_strip": 0,
                         "patch_tool": "",
                         "patch_args": [
                              "-p0"
                         ],
                         "patch_cmds": [],
                         "patch_cmds_win": [],
                         "build_file_content": "",
                         "workspace_file_content": "",
                         "name": "remote_java_tools_darwin_arm64"
                    },
                    "output_tree_hash": "6049e72b170ca56e3084c6ecea53ef01eaedc3325ae472fc71a103d59779037c"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
          "definition_information": "Repository remote_java_tools instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:374:21: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:59:14: in java_tools_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\nRepository rule http_archive defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in <toplevel>\n",
          "original_attributes": {
               "name": "remote_java_tools",
               "generator_name": "remote_java_tools",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "urls": [
                    "https://mirror.bazel.build/bazel_java_tools/releases/java/v13.6.1/java_tools-v13.6.1.zip",
                    "https://github.com/bazelbuild/java_tools/releases/download/java_13.6.1/java_tools-v13.6.1.zip"
               ],
               "sha256": "74c978eab040ad4ec38ce0d0970ac813cc2c6f4f6f4f121c0414719487edc991"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
                    "attributes": {
                         "url": "",
                         "urls": [
                              "https://mirror.bazel.build/bazel_java_tools/releases/java/v13.6.1/java_tools-v13.6.1.zip",
                              "https://github.com/bazelbuild/java_tools/releases/download/java_13.6.1/java_tools-v13.6.1.zip"
                         ],
                         "sha256": "74c978eab040ad4ec38ce0d0970ac813cc2c6f4f6f4f121c0414719487edc991",
                         "integrity": "",
                         "netrc": "",
                         "auth_patterns": {},
                         "canonical_id": "",
                         "strip_prefix": "",
                         "add_prefix": "",
                         "type": "",
                         "patches": [],
                         "remote_file_urls": {},
                         "remote_file_integrity": {},
                         "remote_patches": {},
                         "remote_patch_strip": 0,
                         "patch_tool": "",
                         "patch_args": [
                              "-p0"
                         ],
                         "patch_cmds": [],
                         "patch_cmds_win": [],
                         "build_file_content": "",
                         "workspace_file_content": "",
                         "name": "remote_java_tools"
                    },
                    "output_tree_hash": "8672ba3444d84d0721ee3d4579a85b59b6e48a88c9ba8733dbbf8f1b641722f9"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/osx:xcode_configure.bzl%xcode_autoconf",
          "definition_information": "Repository local_config_xcode instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:184:16: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/osx/xcode_configure.bzl:312:19: in xcode_configure\nRepository rule xcode_autoconf defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/osx/xcode_configure.bzl:297:33: in <toplevel>\n",
          "original_attributes": {
               "name": "local_config_xcode",
               "generator_name": "local_config_xcode",
               "generator_function": "xcode_configure",
               "generator_location": None,
               "xcode_locator": "@bazel_tools//tools/osx:xcode_locator.m"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/osx:xcode_configure.bzl%xcode_autoconf",
                    "attributes": {
                         "name": "local_config_xcode",
                         "generator_name": "local_config_xcode",
                         "generator_function": "xcode_configure",
                         "generator_location": None,
                         "xcode_locator": "@bazel_tools//tools/osx:xcode_locator.m"
                    },
                    "output_tree_hash": "4cc7a67e53f3299c3df0ff7ff49add06398b74655432c8d0d4a5a8113821f27d"
               }
          ]
     },
     {
          "original_rule_class": "//web/internal:platform_archive.bzl%platform_archive",
          "definition_information": "Repository org_chromium_chromium_macos_arm64 instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:59:21: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/versioned/browsers-0.3.3.bzl:28:30: in browser_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/versioned/browsers-0.3.3.bzl:162:21: in org_chromium_chromium\nRepository rule platform_archive defined at:\n  /Users/pcloudy/workspace/rules_webtesting/web/internal/platform_archive.bzl:81:35: in <toplevel>\n",
          "original_attributes": {
               "name": "org_chromium_chromium_macos_arm64",
               "generator_name": "org_chromium_chromium_macos_arm64",
               "generator_function": "browser_repositories",
               "generator_location": None,
               "urls": [
                    "https://storage.googleapis.com/chromium-browser-snapshots/Mac_Arm/902390/chrome-mac.zip",
                    "https://storage.googleapis.com/dev-infra-mirror/chromium/902390/chrome-mac_arm64.zip"
               ],
               "sha256": "4845ce895d030aeb8bfd877a599f1f07d8c7a77d1e08513e80e60bb0093fca24",
               "licenses": [
                    "notice"
               ],
               "named_files": {
                    "CHROMIUM": "chrome-mac/Chromium.app/Contents/MacOS/Chromium"
               }
          },
          "repositories": [
               {
                    "rule_class": "//web/internal:platform_archive.bzl%platform_archive",
                    "attributes": {
                         "name": "org_chromium_chromium_macos_arm64",
                         "generator_name": "org_chromium_chromium_macos_arm64",
                         "generator_function": "browser_repositories",
                         "generator_location": None,
                         "urls": [
                              "https://storage.googleapis.com/chromium-browser-snapshots/Mac_Arm/902390/chrome-mac.zip",
                              "https://storage.googleapis.com/dev-infra-mirror/chromium/902390/chrome-mac_arm64.zip"
                         ],
                         "sha256": "4845ce895d030aeb8bfd877a599f1f07d8c7a77d1e08513e80e60bb0093fca24",
                         "licenses": [
                              "notice"
                         ],
                         "named_files": {
                              "CHROMIUM": "chrome-mac/Chromium.app/Contents/MacOS/Chromium"
                         }
                    },
                    "output_tree_hash": "7ad8040b1cc465bb31e54a06a4a80daed798860333791034cfb3c7d535ccd406"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
          "definition_information": "Repository remotejdk11_macos_aarch64 instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:371:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:353:34: in remote_jdk11_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:52:17: in remote_java_repository\nRepository rule http_archive defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk11_macos_aarch64",
               "generator_name": "remotejdk11_macos_aarch64",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "urls": [
                    "https://cdn.azul.com/zulu/bin/zulu11.72.19-ca-jdk11.0.23-macosx_aarch64.tar.gz",
                    "https://mirror.bazel.build/cdn.azul.com/zulu/bin/zulu11.72.19-ca-jdk11.0.23-macosx_aarch64.tar.gz"
               ],
               "sha256": "40fb1918385e03814b67b7608c908c7f945ccbeddbbf5ed062cdfb2602e21c83",
               "strip_prefix": "zulu11.72.19-ca-jdk11.0.23-macosx_aarch64",
               "build_file_content": "load(\"@rules_java//java:defs.bzl\", \"java_runtime\")\n\npackage(default_visibility = [\"//visibility:public\"])\n\nexports_files([\"WORKSPACE\", \"BUILD.bazel\"])\n\nfilegroup(\n    name = \"jre\",\n    srcs = glob(\n        [\n            \"jre/bin/**\",\n            \"jre/lib/**\",\n        ],\n        allow_empty = True,\n        # In some configurations, Java browser plugin is considered harmful and\n        # common antivirus software blocks access to npjp2.dll interfering with Bazel,\n        # so do not include it in JRE on Windows.\n        exclude = [\"jre/bin/plugin2/**\"],\n    ),\n)\n\nfilegroup(\n    name = \"jdk-bin\",\n    srcs = glob(\n        [\"bin/**\"],\n        # The JDK on Windows sometimes contains a directory called\n        # \"%systemroot%\", which is not a valid label.\n        exclude = [\"**/*%*/**\"],\n    ),\n)\n\n# This folder holds security policies.\nfilegroup(\n    name = \"jdk-conf\",\n    srcs = glob(\n        [\"conf/**\"],\n        allow_empty = True,\n    ),\n)\n\nfilegroup(\n    name = \"jdk-include\",\n    srcs = glob(\n        [\"include/**\"],\n        allow_empty = True,\n    ),\n)\n\nfilegroup(\n    name = \"jdk-lib\",\n    srcs = glob(\n        [\"lib/**\", \"release\"],\n        allow_empty = True,\n        exclude = [\n            \"lib/missioncontrol/**\",\n            \"lib/visualvm/**\",\n        ],\n    ),\n)\n\njava_runtime(\n    name = \"jdk\",\n    srcs = [\n        \":jdk-bin\",\n        \":jdk-conf\",\n        \":jdk-include\",\n        \":jdk-lib\",\n        \":jre\",\n    ],\n    # Provide the 'java` binary explicitly so that the correct path is used by\n    # Bazel even when the host platform differs from the execution platform.\n    # Exactly one of the two globs will be empty depending on the host platform.\n    # When --incompatible_disallow_empty_glob is enabled, each individual empty\n    # glob will fail without allow_empty = True, even if the overall result is\n    # non-empty.\n    java = glob([\"bin/java.exe\", \"bin/java\"], allow_empty = True)[0],\n    version = 11,\n)\n\nfilegroup(\n    name = \"jdk-jmods\",\n    srcs = glob(\n        [\"jmods/**\"],\n        allow_empty = True,\n    ),\n)\n\njava_runtime(\n    name = \"jdk-with-jmods\",\n    srcs = [\n        \":jdk-bin\",\n        \":jdk-conf\",\n        \":jdk-include\",\n        \":jdk-lib\",\n        \":jdk-jmods\",\n        \":jre\",\n    ],\n    java = glob([\"bin/java.exe\", \"bin/java\"], allow_empty = True)[0],\n    version = 11,\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
                    "attributes": {
                         "url": "",
                         "urls": [
                              "https://cdn.azul.com/zulu/bin/zulu11.72.19-ca-jdk11.0.23-macosx_aarch64.tar.gz",
                              "https://mirror.bazel.build/cdn.azul.com/zulu/bin/zulu11.72.19-ca-jdk11.0.23-macosx_aarch64.tar.gz"
                         ],
                         "sha256": "40fb1918385e03814b67b7608c908c7f945ccbeddbbf5ed062cdfb2602e21c83",
                         "integrity": "",
                         "netrc": "",
                         "auth_patterns": {},
                         "canonical_id": "",
                         "strip_prefix": "zulu11.72.19-ca-jdk11.0.23-macosx_aarch64",
                         "add_prefix": "",
                         "type": "",
                         "patches": [],
                         "remote_file_urls": {},
                         "remote_file_integrity": {},
                         "remote_patches": {},
                         "remote_patch_strip": 0,
                         "patch_tool": "",
                         "patch_args": [
                              "-p0"
                         ],
                         "patch_cmds": [],
                         "patch_cmds_win": [],
                         "build_file_content": "load(\"@rules_java//java:defs.bzl\", \"java_runtime\")\n\npackage(default_visibility = [\"//visibility:public\"])\n\nexports_files([\"WORKSPACE\", \"BUILD.bazel\"])\n\nfilegroup(\n    name = \"jre\",\n    srcs = glob(\n        [\n            \"jre/bin/**\",\n            \"jre/lib/**\",\n        ],\n        allow_empty = True,\n        # In some configurations, Java browser plugin is considered harmful and\n        # common antivirus software blocks access to npjp2.dll interfering with Bazel,\n        # so do not include it in JRE on Windows.\n        exclude = [\"jre/bin/plugin2/**\"],\n    ),\n)\n\nfilegroup(\n    name = \"jdk-bin\",\n    srcs = glob(\n        [\"bin/**\"],\n        # The JDK on Windows sometimes contains a directory called\n        # \"%systemroot%\", which is not a valid label.\n        exclude = [\"**/*%*/**\"],\n    ),\n)\n\n# This folder holds security policies.\nfilegroup(\n    name = \"jdk-conf\",\n    srcs = glob(\n        [\"conf/**\"],\n        allow_empty = True,\n    ),\n)\n\nfilegroup(\n    name = \"jdk-include\",\n    srcs = glob(\n        [\"include/**\"],\n        allow_empty = True,\n    ),\n)\n\nfilegroup(\n    name = \"jdk-lib\",\n    srcs = glob(\n        [\"lib/**\", \"release\"],\n        allow_empty = True,\n        exclude = [\n            \"lib/missioncontrol/**\",\n            \"lib/visualvm/**\",\n        ],\n    ),\n)\n\njava_runtime(\n    name = \"jdk\",\n    srcs = [\n        \":jdk-bin\",\n        \":jdk-conf\",\n        \":jdk-include\",\n        \":jdk-lib\",\n        \":jre\",\n    ],\n    # Provide the 'java` binary explicitly so that the correct path is used by\n    # Bazel even when the host platform differs from the execution platform.\n    # Exactly one of the two globs will be empty depending on the host platform.\n    # When --incompatible_disallow_empty_glob is enabled, each individual empty\n    # glob will fail without allow_empty = True, even if the overall result is\n    # non-empty.\n    java = glob([\"bin/java.exe\", \"bin/java\"], allow_empty = True)[0],\n    version = 11,\n)\n\nfilegroup(\n    name = \"jdk-jmods\",\n    srcs = glob(\n        [\"jmods/**\"],\n        allow_empty = True,\n    ),\n)\n\njava_runtime(\n    name = \"jdk-with-jmods\",\n    srcs = [\n        \":jdk-bin\",\n        \":jdk-conf\",\n        \":jdk-include\",\n        \":jdk-lib\",\n        \":jdk-jmods\",\n        \":jre\",\n    ],\n    java = glob([\"bin/java.exe\", \"bin/java\"], allow_empty = True)[0],\n    version = 11,\n)\n",
                         "workspace_file_content": "",
                         "name": "remotejdk11_macos_aarch64"
                    },
                    "output_tree_hash": "77f09d3a6f2a5588da29a0743f477cf6fca8f8340c32947963d14b69761106f8"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
          "definition_information": "Repository remotejdk21_macos_aarch64 instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:93:24: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:373:23: in rules_java_dependencies\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:361:34: in remote_jdk21_repos\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/java/repositories.bzl:333:14: in _remote_jdk_repos_for_version\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/utils.bzl:268:18: in maybe\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_java_builtin/toolchains/remote_java_repository.bzl:52:17: in remote_java_repository\nRepository rule http_archive defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/build_defs/repo/http.bzl:387:31: in <toplevel>\n",
          "original_attributes": {
               "name": "remotejdk21_macos_aarch64",
               "generator_name": "remotejdk21_macos_aarch64",
               "generator_function": "rules_java_dependencies",
               "generator_location": None,
               "urls": [
                    "https://cdn.azul.com/zulu/bin/zulu21.34.19-ca-jdk21.0.3-macosx_aarch64.tar.gz",
                    "https://mirror.bazel.build/cdn.azul.com/zulu/bin/zulu21.34.19-ca-jdk21.0.3-macosx_aarch64.tar.gz"
               ],
               "sha256": "4f42a561909d71868a700cf2efa1390e1b9e04863f3fa75ea30c4965e5a702f0",
               "strip_prefix": "zulu21.34.19-ca-jdk21.0.3-macosx_aarch64",
               "build_file_content": "load(\"@rules_java//java:defs.bzl\", \"java_runtime\")\n\npackage(default_visibility = [\"//visibility:public\"])\n\nexports_files([\"WORKSPACE\", \"BUILD.bazel\"])\n\nfilegroup(\n    name = \"jre\",\n    srcs = glob(\n        [\n            \"jre/bin/**\",\n            \"jre/lib/**\",\n        ],\n        allow_empty = True,\n        # In some configurations, Java browser plugin is considered harmful and\n        # common antivirus software blocks access to npjp2.dll interfering with Bazel,\n        # so do not include it in JRE on Windows.\n        exclude = [\"jre/bin/plugin2/**\"],\n    ),\n)\n\nfilegroup(\n    name = \"jdk-bin\",\n    srcs = glob(\n        [\"bin/**\"],\n        # The JDK on Windows sometimes contains a directory called\n        # \"%systemroot%\", which is not a valid label.\n        exclude = [\"**/*%*/**\"],\n    ),\n)\n\n# This folder holds security policies.\nfilegroup(\n    name = \"jdk-conf\",\n    srcs = glob(\n        [\"conf/**\"],\n        allow_empty = True,\n    ),\n)\n\nfilegroup(\n    name = \"jdk-include\",\n    srcs = glob(\n        [\"include/**\"],\n        allow_empty = True,\n    ),\n)\n\nfilegroup(\n    name = \"jdk-lib\",\n    srcs = glob(\n        [\"lib/**\", \"release\"],\n        allow_empty = True,\n        exclude = [\n            \"lib/missioncontrol/**\",\n            \"lib/visualvm/**\",\n        ],\n    ),\n)\n\njava_runtime(\n    name = \"jdk\",\n    srcs = [\n        \":jdk-bin\",\n        \":jdk-conf\",\n        \":jdk-include\",\n        \":jdk-lib\",\n        \":jre\",\n    ],\n    # Provide the 'java` binary explicitly so that the correct path is used by\n    # Bazel even when the host platform differs from the execution platform.\n    # Exactly one of the two globs will be empty depending on the host platform.\n    # When --incompatible_disallow_empty_glob is enabled, each individual empty\n    # glob will fail without allow_empty = True, even if the overall result is\n    # non-empty.\n    java = glob([\"bin/java.exe\", \"bin/java\"], allow_empty = True)[0],\n    version = 21,\n)\n\nfilegroup(\n    name = \"jdk-jmods\",\n    srcs = glob(\n        [\"jmods/**\"],\n        allow_empty = True,\n    ),\n)\n\njava_runtime(\n    name = \"jdk-with-jmods\",\n    srcs = [\n        \":jdk-bin\",\n        \":jdk-conf\",\n        \":jdk-include\",\n        \":jdk-lib\",\n        \":jdk-jmods\",\n        \":jre\",\n    ],\n    java = glob([\"bin/java.exe\", \"bin/java\"], allow_empty = True)[0],\n    version = 21,\n)\n"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/build_defs/repo:http.bzl%http_archive",
                    "attributes": {
                         "url": "",
                         "urls": [
                              "https://cdn.azul.com/zulu/bin/zulu21.34.19-ca-jdk21.0.3-macosx_aarch64.tar.gz",
                              "https://mirror.bazel.build/cdn.azul.com/zulu/bin/zulu21.34.19-ca-jdk21.0.3-macosx_aarch64.tar.gz"
                         ],
                         "sha256": "4f42a561909d71868a700cf2efa1390e1b9e04863f3fa75ea30c4965e5a702f0",
                         "integrity": "",
                         "netrc": "",
                         "auth_patterns": {},
                         "canonical_id": "",
                         "strip_prefix": "zulu21.34.19-ca-jdk21.0.3-macosx_aarch64",
                         "add_prefix": "",
                         "type": "",
                         "patches": [],
                         "remote_file_urls": {},
                         "remote_file_integrity": {},
                         "remote_patches": {},
                         "remote_patch_strip": 0,
                         "patch_tool": "",
                         "patch_args": [
                              "-p0"
                         ],
                         "patch_cmds": [],
                         "patch_cmds_win": [],
                         "build_file_content": "load(\"@rules_java//java:defs.bzl\", \"java_runtime\")\n\npackage(default_visibility = [\"//visibility:public\"])\n\nexports_files([\"WORKSPACE\", \"BUILD.bazel\"])\n\nfilegroup(\n    name = \"jre\",\n    srcs = glob(\n        [\n            \"jre/bin/**\",\n            \"jre/lib/**\",\n        ],\n        allow_empty = True,\n        # In some configurations, Java browser plugin is considered harmful and\n        # common antivirus software blocks access to npjp2.dll interfering with Bazel,\n        # so do not include it in JRE on Windows.\n        exclude = [\"jre/bin/plugin2/**\"],\n    ),\n)\n\nfilegroup(\n    name = \"jdk-bin\",\n    srcs = glob(\n        [\"bin/**\"],\n        # The JDK on Windows sometimes contains a directory called\n        # \"%systemroot%\", which is not a valid label.\n        exclude = [\"**/*%*/**\"],\n    ),\n)\n\n# This folder holds security policies.\nfilegroup(\n    name = \"jdk-conf\",\n    srcs = glob(\n        [\"conf/**\"],\n        allow_empty = True,\n    ),\n)\n\nfilegroup(\n    name = \"jdk-include\",\n    srcs = glob(\n        [\"include/**\"],\n        allow_empty = True,\n    ),\n)\n\nfilegroup(\n    name = \"jdk-lib\",\n    srcs = glob(\n        [\"lib/**\", \"release\"],\n        allow_empty = True,\n        exclude = [\n            \"lib/missioncontrol/**\",\n            \"lib/visualvm/**\",\n        ],\n    ),\n)\n\njava_runtime(\n    name = \"jdk\",\n    srcs = [\n        \":jdk-bin\",\n        \":jdk-conf\",\n        \":jdk-include\",\n        \":jdk-lib\",\n        \":jre\",\n    ],\n    # Provide the 'java` binary explicitly so that the correct path is used by\n    # Bazel even when the host platform differs from the execution platform.\n    # Exactly one of the two globs will be empty depending on the host platform.\n    # When --incompatible_disallow_empty_glob is enabled, each individual empty\n    # glob will fail without allow_empty = True, even if the overall result is\n    # non-empty.\n    java = glob([\"bin/java.exe\", \"bin/java\"], allow_empty = True)[0],\n    version = 21,\n)\n\nfilegroup(\n    name = \"jdk-jmods\",\n    srcs = glob(\n        [\"jmods/**\"],\n        allow_empty = True,\n    ),\n)\n\njava_runtime(\n    name = \"jdk-with-jmods\",\n    srcs = [\n        \":jdk-bin\",\n        \":jdk-conf\",\n        \":jdk-include\",\n        \":jdk-lib\",\n        \":jdk-jmods\",\n        \":jre\",\n    ],\n    java = glob([\"bin/java.exe\", \"bin/java\"], allow_empty = True)[0],\n    version = 21,\n)\n",
                         "workspace_file_content": "",
                         "name": "remotejdk21_macos_aarch64"
                    },
                    "output_tree_hash": "d66c9c1f43cabc4818a026850812f0d2c3f27fb1e2d1ea861bd1b1615ca44bfb"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_tools//tools/cpp:cc_configure.bzl%cc_autoconf",
          "definition_information": "Repository local_config_cc instantiated at:\n  /DEFAULT.WORKSPACE.SUFFIX:181:13: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/cpp/cc_configure.bzl:149:16: in cc_configure\nRepository rule cc_autoconf defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_tools/tools/cpp/cc_configure.bzl:109:30: in <toplevel>\n",
          "original_attributes": {
               "name": "local_config_cc",
               "generator_name": "local_config_cc",
               "generator_function": "cc_configure",
               "generator_location": None
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_tools//tools/cpp:cc_configure.bzl%cc_autoconf",
                    "attributes": {
                         "name": "local_config_cc",
                         "generator_name": "local_config_cc",
                         "generator_function": "cc_configure",
                         "generator_location": None
                    },
                    "output_tree_hash": "e7711dcc6d599fe5e5acc0a6678d46b79287164ed2c35f96e83c448012615b53"
               }
          ]
     },
     {
          "original_rule_class": "@@rules_jvm_external//:coursier.bzl%coursier_fetch",
          "definition_information": "Repository maven instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:120:14: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_jvm_external/private/rules/maven_install.bzl:104:19: in maven_install\nRepository rule coursier_fetch defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/rules_jvm_external/coursier.bzl:1211:33: in <toplevel>\n",
          "original_attributes": {
               "name": "maven",
               "generator_name": "maven",
               "generator_function": "maven_install",
               "generator_location": None,
               "repositories": [
                    "{ \"repo_url\": \"https://maven.google.com\" }",
                    "{ \"repo_url\": \"https://repo.maven.apache.org/maven2\" }"
               ],
               "artifacts": [
                    "{ \"group\": \"org.scalatest\", \"artifact\": \"scalatest-wordspec_2.12\", \"version\": \"3.2.9\" }"
               ],
               "fail_on_missing_checksum": True,
               "fetch_sources": False,
               "fetch_javadoc": False,
               "use_credentials_from_home_netrc_file": False,
               "excluded_artifacts": [],
               "generate_compat_repositories": False,
               "version_conflict_policy": "default",
               "override_targets": {},
               "strict_visibility": False,
               "strict_visibility_value": [
                    "//visibility:private"
               ],
               "resolve_timeout": 600,
               "jetify": False,
               "jetify_include_list": [
                    "*"
               ],
               "use_starlark_android_rules": False,
               "aar_import_bzl_label": "@build_bazel_rules_android//android:rules.bzl",
               "duplicate_version_warning": "warn"
          },
          "repositories": [
               {
                    "rule_class": "@@rules_jvm_external//:coursier.bzl%coursier_fetch",
                    "attributes": {
                         "name": "maven",
                         "generator_name": "maven",
                         "generator_function": "maven_install",
                         "generator_location": None,
                         "repositories": [
                              "{ \"repo_url\": \"https://maven.google.com\" }",
                              "{ \"repo_url\": \"https://repo.maven.apache.org/maven2\" }"
                         ],
                         "artifacts": [
                              "{ \"group\": \"org.scalatest\", \"artifact\": \"scalatest-wordspec_2.12\", \"version\": \"3.2.9\" }"
                         ],
                         "fail_on_missing_checksum": True,
                         "fetch_sources": False,
                         "fetch_javadoc": False,
                         "use_credentials_from_home_netrc_file": False,
                         "excluded_artifacts": [],
                         "generate_compat_repositories": False,
                         "version_conflict_policy": "default",
                         "override_targets": {},
                         "strict_visibility": False,
                         "strict_visibility_value": [
                              "//visibility:private"
                         ],
                         "resolve_timeout": 600,
                         "jetify": False,
                         "jetify_include_list": [
                              "*"
                         ],
                         "use_starlark_android_rules": False,
                         "aar_import_bzl_label": "@build_bazel_rules_android//android:rules.bzl",
                         "duplicate_version_warning": "warn"
                    },
                    "output_tree_hash": "c7b285fd7b8c0eb8dcd7c858d7d1b338655330d8bb9b2124c88146b7eadf06ea"
               }
          ]
     },
     {
          "original_rule_class": "@@io_bazel_rules_go//go/private:sdk.bzl%go_download_sdk_rule",
          "definition_information": "Repository go_sdk instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:42:23: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_go/go/private/sdk.bzl:695:28: in go_register_toolchains\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_go/go/private/sdk.bzl:307:25: in go_download_sdk\nRepository rule go_download_sdk_rule defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/io_bazel_rules_go/go/private/sdk.bzl:137:39: in <toplevel>\n",
          "original_attributes": {
               "name": "go_sdk",
               "generator_name": "go_sdk",
               "generator_function": "go_register_toolchains",
               "generator_location": None,
               "version": "1.20.5"
          },
          "repositories": [
               {
                    "rule_class": "@@io_bazel_rules_go//go/private:sdk.bzl%go_download_sdk_rule",
                    "attributes": {
                         "name": "go_sdk",
                         "generator_name": "go_sdk",
                         "generator_function": "go_register_toolchains",
                         "generator_location": None,
                         "version": "1.20.5"
                    },
                    "output_tree_hash": "374ca739dd71e5856077ce7de670922e9a079411982886d0d8153c3b2e3301de"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_gazelle//internal:go_repository_cache.bzl%go_repository_cache",
          "definition_information": "Repository bazel_gazelle_go_repository_cache instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:55:21: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_gazelle/deps.bzl:72:28: in gazelle_dependencies\nRepository rule go_repository_cache defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_gazelle/internal/go_repository_cache.bzl:71:38: in <toplevel>\n",
          "original_attributes": {
               "name": "bazel_gazelle_go_repository_cache",
               "generator_name": "bazel_gazelle_go_repository_cache",
               "generator_function": "gazelle_dependencies",
               "generator_location": None,
               "go_sdk_info": {
                    "go_sdk": "host"
               },
               "go_env": {}
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_gazelle//internal:go_repository_cache.bzl%go_repository_cache",
                    "attributes": {
                         "name": "bazel_gazelle_go_repository_cache",
                         "generator_name": "bazel_gazelle_go_repository_cache",
                         "generator_function": "gazelle_dependencies",
                         "generator_location": None,
                         "go_sdk_info": {
                              "go_sdk": "host"
                         },
                         "go_env": {}
                    },
                    "output_tree_hash": "5f1ed679d8f28074f1ce53fb983484a4422eddc1c6a46475ccff04a10fc53482"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_gazelle//internal:go_repository_tools.bzl%go_repository_tools",
          "definition_information": "Repository bazel_gazelle_go_repository_tools instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:55:21: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_gazelle/deps.bzl:78:24: in gazelle_dependencies\nRepository rule go_repository_tools defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_gazelle/internal/go_repository_tools.bzl:117:38: in <toplevel>\n",
          "original_attributes": {
               "name": "bazel_gazelle_go_repository_tools",
               "generator_name": "bazel_gazelle_go_repository_tools",
               "generator_function": "gazelle_dependencies",
               "generator_location": None,
               "go_cache": "@@bazel_gazelle_go_repository_cache//:go.env"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_gazelle//internal:go_repository_tools.bzl%go_repository_tools",
                    "attributes": {
                         "name": "bazel_gazelle_go_repository_tools",
                         "generator_name": "bazel_gazelle_go_repository_tools",
                         "generator_function": "gazelle_dependencies",
                         "generator_location": None,
                         "go_cache": "@@bazel_gazelle_go_repository_cache//:go.env"
                    },
                    "output_tree_hash": "2aae1b4d1fa2b493c281cd122b76aafcddc3deda9207fe830060b6523c458019"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_gazelle//internal:go_repository_config.bzl%go_repository_config",
          "definition_information": "Repository bazel_gazelle_go_repository_config instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:55:21: in <toplevel>\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_gazelle/deps.bzl:83:25: in gazelle_dependencies\nRepository rule go_repository_config defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_gazelle/internal/go_repository_config.bzl:69:39: in <toplevel>\n",
          "original_attributes": {
               "name": "bazel_gazelle_go_repository_config",
               "generator_name": "bazel_gazelle_go_repository_config",
               "generator_function": "gazelle_dependencies",
               "generator_location": None,
               "config": "//:WORKSPACE"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_gazelle//internal:go_repository_config.bzl%go_repository_config",
                    "attributes": {
                         "name": "bazel_gazelle_go_repository_config",
                         "generator_name": "bazel_gazelle_go_repository_config",
                         "generator_function": "gazelle_dependencies",
                         "generator_location": None,
                         "config": "//:WORKSPACE"
                    },
                    "output_tree_hash": "e0a452b97e40aed05812c8cfdddba4af470abe69c849777ecab30b2f9eeaabb5"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_gazelle//internal:go_repository.bzl%go_repository",
          "definition_information": "Repository com_github_gorilla_mux instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:69:25: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/go_repositories.bzl:67:31: in go_internal_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/go_repositories.bzl:83:18: in com_github_gorilla_mux\nRepository rule go_repository defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_gazelle/internal/go_repository.bzl:325:32: in <toplevel>\n",
          "original_attributes": {
               "name": "com_github_gorilla_mux",
               "generator_name": "com_github_gorilla_mux",
               "generator_function": "go_internal_repositories",
               "generator_location": None,
               "importpath": "github.com/gorilla/mux",
               "urls": [
                    "https://github.com/gorilla/mux/archive/v1.7.3.tar.gz"
               ],
               "strip_prefix": "mux-1.7.3",
               "sha256": "92adb9aea022f8b35686b75be50ba1206c4457c2f8a0e2a9d10d8721f35b3f11"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_gazelle//internal:go_repository.bzl%go_repository",
                    "attributes": {
                         "name": "com_github_gorilla_mux",
                         "generator_name": "com_github_gorilla_mux",
                         "generator_function": "go_internal_repositories",
                         "generator_location": None,
                         "importpath": "github.com/gorilla/mux",
                         "urls": [
                              "https://github.com/gorilla/mux/archive/v1.7.3.tar.gz"
                         ],
                         "strip_prefix": "mux-1.7.3",
                         "sha256": "92adb9aea022f8b35686b75be50ba1206c4457c2f8a0e2a9d10d8721f35b3f11"
                    },
                    "output_tree_hash": "26110bfe6a1875419dcbe55bb52344213cdc03889bffcf9ff1d54bb9cfafed4f"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_gazelle//internal:go_repository.bzl%go_repository",
          "definition_information": "Repository com_github_tebeka_selenium instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:67:16: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/go_repositories.bzl:45:35: in go_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/go_repositories.bzl:94:18: in com_github_tebeka_selenium\nRepository rule go_repository defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_gazelle/internal/go_repository.bzl:325:32: in <toplevel>\n",
          "original_attributes": {
               "name": "com_github_tebeka_selenium",
               "generator_name": "com_github_tebeka_selenium",
               "generator_function": "go_repositories",
               "generator_location": None,
               "importpath": "github.com/tebeka/selenium",
               "urls": [
                    "https://github.com/tebeka/selenium/archive/v0.9.9.tar.gz"
               ],
               "strip_prefix": "selenium-0.9.9",
               "sha256": "82846f237b742983a293619e712dcf167e3d7998df3239f3443303d9036ad412"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_gazelle//internal:go_repository.bzl%go_repository",
                    "attributes": {
                         "name": "com_github_tebeka_selenium",
                         "generator_name": "com_github_tebeka_selenium",
                         "generator_function": "go_repositories",
                         "generator_location": None,
                         "importpath": "github.com/tebeka/selenium",
                         "urls": [
                              "https://github.com/tebeka/selenium/archive/v0.9.9.tar.gz"
                         ],
                         "strip_prefix": "selenium-0.9.9",
                         "sha256": "82846f237b742983a293619e712dcf167e3d7998df3239f3443303d9036ad412"
                    },
                    "output_tree_hash": "81a195feef89e558b1b0193675f00bf102d54e22e3ee3e9fa1f6ba2ec4c1d185"
               }
          ]
     },
     {
          "original_rule_class": "@@bazel_gazelle//internal:go_repository.bzl%go_repository",
          "definition_information": "Repository com_github_blang_semver instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:67:16: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/go_repositories.bzl:43:32: in go_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/go_repositories.bzl:72:18: in com_github_blang_semver\nRepository rule go_repository defined at:\n  /private/var/tmp/_bazel_pcloudy/ffc9bb221cd9fd6ff19cab14eef0554e/external/bazel_gazelle/internal/go_repository.bzl:325:32: in <toplevel>\n",
          "original_attributes": {
               "name": "com_github_blang_semver",
               "generator_name": "com_github_blang_semver",
               "generator_function": "go_repositories",
               "generator_location": None,
               "importpath": "github.com/blang/semver",
               "urls": [
                    "https://github.com/blang/semver/archive/v3.6.1.tar.gz"
               ],
               "strip_prefix": "semver-3.6.1",
               "sha256": "dc85076e7c5a7a44e33fc24df320904b95e2fa12c94a3ac758a574dadd54ee53"
          },
          "repositories": [
               {
                    "rule_class": "@@bazel_gazelle//internal:go_repository.bzl%go_repository",
                    "attributes": {
                         "name": "com_github_blang_semver",
                         "generator_name": "com_github_blang_semver",
                         "generator_function": "go_repositories",
                         "generator_location": None,
                         "importpath": "github.com/blang/semver",
                         "urls": [
                              "https://github.com/blang/semver/archive/v3.6.1.tar.gz"
                         ],
                         "strip_prefix": "semver-3.6.1",
                         "sha256": "dc85076e7c5a7a44e33fc24df320904b95e2fa12c94a3ac758a574dadd54ee53"
                    },
                    "output_tree_hash": "3ae06ea2e6811c48c1006e716e39896331486645f8ccfa8d12afd73293df7c83"
               }
          ]
     },
     {
          "original_rule_class": "//web/internal:platform_archive.bzl%platform_archive",
          "definition_information": "Repository org_mozilla_firefox_macos_arm64 instantiated at:\n  /Users/pcloudy/workspace/rules_webtesting/WORKSPACE:59:21: in <toplevel>\n  /Users/pcloudy/workspace/rules_webtesting/web/versioned/browsers-0.3.3.bzl:30:28: in browser_repositories\n  /Users/pcloudy/workspace/rules_webtesting/web/versioned/browsers-0.3.3.bzl:220:21: in org_mozilla_firefox\nRepository rule platform_archive defined at:\n  /Users/pcloudy/workspace/rules_webtesting/web/internal/platform_archive.bzl:81:35: in <toplevel>\n",
          "original_attributes": {
               "name": "org_mozilla_firefox_macos_arm64",
               "generator_name": "org_mozilla_firefox_macos_arm64",
               "generator_function": "browser_repositories",
               "generator_location": None,
               "urls": [
                    "https://ftp.mozilla.org/pub/firefox/releases/90.0.1/mac/en-US/Firefox%2090.0.1.dmg",
                    "https://storage.googleapis.com/dev-infra-mirror/mozilla/firefox/Firefox%2090.0.1.dmg"
               ],
               "sha256": "76c1b9c42b52c7e5be4c112a98b7d3762a18841367f778a179679ac0de751f05",
               "licenses": [
                    "reciprocal"
               ],
               "named_files": {
                    "FIREFOX": "Firefox.app/Contents/MacOS/firefox"
               }
          },
          "repositories": [
               {
                    "rule_class": "//web/internal:platform_archive.bzl%platform_archive",
                    "attributes": {
                         "name": "org_mozilla_firefox_macos_arm64",
                         "generator_name": "org_mozilla_firefox_macos_arm64",
                         "generator_function": "browser_repositories",
                         "generator_location": None,
                         "urls": [
                              "https://ftp.mozilla.org/pub/firefox/releases/90.0.1/mac/en-US/Firefox%2090.0.1.dmg",
                              "https://storage.googleapis.com/dev-infra-mirror/mozilla/firefox/Firefox%2090.0.1.dmg"
                         ],
                         "sha256": "76c1b9c42b52c7e5be4c112a98b7d3762a18841367f778a179679ac0de751f05",
                         "licenses": [
                              "reciprocal"
                         ],
                         "named_files": {
                              "FIREFOX": "Firefox.app/Contents/MacOS/firefox"
                         }
                    },
                    "output_tree_hash": "4477959fc04f80e608870cdea87b33a3f0ec4a05df853df5c236ecd3e087979e"
               }
          ]
     }
]