/*
 * Copyright 2018-2020 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package springboot

import (
	"github.com/buildpacks/libbuildpack/v2/application"
	"github.com/cloudfoundry/libcfbuildpack/v2/logger"
	"github.com/cloudfoundry/libcfbuildpack/v2/manifest"
)

// Metadata describes the application's metadata.
type Metadata struct {
	// Classes indicates the Spring-Boot-Classes of a Spring Boot application.
	Classes string `mapstructure:"classes" properties:"Spring-Boot-Classes,default=" toml:"classes"`

	// LayersIndex indicates the Spring-Boot-Layers-Index of a Spring Boot application.
	LayersIndex string `mapstructure:"layers-index" properties:"Spring-Boot-Layers-Index,default=" toml:"layers-index"`

	// Lib indicates the Spring-Boot-Lib of a Spring Boot application.
	Lib string `mapstructure:"lib" properties:"Spring-Boot-Lib,default=" toml:"lib"`

	// MainClass indicates the Main-Class of a Spring Boot application.
	MainClass string `mapstructure:"main-class" properties:"Main-Class,default=" toml:"main-class"`

	// Version indicates the Spring-Boot-Version of a Spring Boot application.
	Version string `mapstructure:"version" properties:"Spring-Boot-Version,default=" toml:"version"`
}

func (m Metadata) Identity() (string, string) {
	return "Spring Boot", m.Version
}

// NewMetadata creates a new Metadata returning false if Spring-Boot-Version is not defined.
func NewMetadata(application application.Application, logger logger.Logger) (Metadata, bool, error) {
	md := Metadata{}

	m, err := manifest.NewManifest(application, logger)
	if err != nil {
		return Metadata{}, false, err
	}

	if err := m.Decode(&md); err != nil {
		return Metadata{}, false, err
	}

	if md.Version == "" {
		return Metadata{}, false, nil
	}

	return md, true, nil
}
