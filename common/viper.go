/*
Copyright 2024 BlackRock, Inc.

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

package common

import (
	"log/slog"
	"os"

	"github.com/spf13/viper"
)

func ViperWithLogging() *viper.Viper {
	v := viper.NewWithOptions(viper.WithLogger(slog.New(slog.NewJSONHandler(os.Stdout, nil))))
	return v
}
