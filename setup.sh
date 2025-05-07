#!/bin/bash
# Copyright 2025 Google Inc. All Rights Reserved
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS-IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Script that can be used by CI server for testing j2cl builds.
set -ex
sudo apt -y update && sudo apt -y install  libasound2
sudo apt -y install  libdbus-glib-1-2
sudo apt -y install  libatk-bridge2.0-0
sudo apt -y install  libatk1.0-0
sudo apt -y install  libc6
sudo apt -y install  libcairo2
sudo apt -y install  libcups2
sudo apt -y install  libdbus-1-3
sudo apt -y install  libexpat1
sudo apt -y install  libfontconfig1
sudo apt -y install  libgbm1
sudo apt -y install  libgcc1
sudo apt -y install  libglib2.0-0
sudo apt -y install  libgtk-3-0
sudo apt -y install  libnspr4
sudo apt -y install  libnss3
sudo apt -y install  libpango-1.0-0
sudo apt -y install  libpangocairo-1.0-0
sudo apt -y install  libstdc++6
sudo apt -y install  libx11-6
sudo apt -y install  libx11-xcb1
sudo apt -y install  libxcb1
sudo apt -y install  libxcomposite1
sudo apt -y install  libxcursor1
sudo apt -y install  libxdamage1
sudo apt -y install  libxext6
sudo apt -y install  libxfixes3
sudo apt -y install  libxi6
sudo apt -y install  libxrandr2
sudo apt -y install  libxrender1
sudo apt -y install  libxss1
sudo apt -y install  libxtst6
sudo apt install libgtk-3-0