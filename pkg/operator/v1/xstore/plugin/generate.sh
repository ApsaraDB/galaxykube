#!/usr/bin/env bash

# Copyright 2021 Alibaba Group Holding Limited.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

TARGET_FILE="engines/plugin.generated.go"

PACKAGE_PREFIX="github.com/alibaba/polardbx-operator/pkg/operator/v1/xstore/plugin"

mkdir -p engines && rm -f $TARGET_FILE

echo "// Auto-generated by generate.sh. DO NOT MODIFY!!

/*
Copyright 2021 Alibaba Group Holding Limited.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package engines

// Plugins to provide multiple store backends.
import (" >$TARGET_FILE

for pkg in *; do
  if [ ! -d "${pkg}" ] || [ ! -f "${pkg}/engine.go" ]; then
    continue
  fi
  printf "\t_ \"%s/%s\"\n" ${PACKAGE_PREFIX} "${pkg}" >>$TARGET_FILE
done

echo ')' >>$TARGET_FILE
