/*
 * Copyright 2018 the original author or authors.
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package shell

import (
	"github.com/projectriff/riff-cli/pkg/options"
	"path/filepath"
	"github.com/projectriff/riff-cli/pkg/initializers/core"
)

var shellFunctionDockerfileTemplate = `
FROM projectriff/shell-function-invoker:{{.RiffVersion}}
ARG FUNCTION_URI="/{{.ArtifactBase}}"
ADD {{.Artifact}} /
ENV FUNCTION_URI $FUNCTION_URI
`

func generateShellFunctionDockerFile(opts options.InitOptions) (string, error) {
	dockerFileTokens := core.DockerFileTokens{
		Artifact:     opts.Artifact,
		ArtifactBase: filepath.Base(opts.Artifact),
		RiffVersion:  opts.RiffVersion,
	}
	return core.GenerateFunctionDockerFileContents(shellFunctionDockerfileTemplate, "docker-shell", dockerFileTokens)
}
