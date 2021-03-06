/*
Copyright 2018 Google LLC

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

package commands

import (
	"github.com/docker/docker/builder/dockerfile/instructions"
	"github.com/google/go-containerregistry/v1"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type DockerCommand interface {
	// ExecuteCommand is responsible for:
	// 	1. Making required changes to the filesystem (ex. copying files for ADD/COPY or setting ENV variables)
	//  2. Updating metadata fields in the config
	// It should not change the config history.
	ExecuteCommand(*v1.Config) error
	// The config history has a "created by" field, should return information about the command
	CreatedBy() string
	// A list of files to snapshot, empty for metadata commands or nil if we don't know
	FilesToSnapshot() []string
}

func GetCommand(cmd instructions.Command, buildcontext string) (DockerCommand, error) {
	switch c := cmd.(type) {
	case *instructions.RunCommand:
		return &RunCommand{cmd: c}, nil
	case *instructions.CopyCommand:
		return &CopyCommand{cmd: c, buildcontext: buildcontext}, nil
	case *instructions.ExposeCommand:
		return &ExposeCommand{cmd: c}, nil
	case *instructions.EnvCommand:
		return &EnvCommand{cmd: c}, nil
	case *instructions.WorkdirCommand:
		return &WorkdirCommand{cmd: c}, nil
	case *instructions.AddCommand:
		return &AddCommand{cmd: c, buildcontext: buildcontext}, nil
	case *instructions.CmdCommand:
		return &CmdCommand{cmd: c}, nil
	case *instructions.EntrypointCommand:
		return &EntrypointCommand{cmd: c}, nil
	case *instructions.LabelCommand:
		return &LabelCommand{cmd: c}, nil
	case *instructions.UserCommand:
		return &UserCommand{cmd: c}, nil
	case *instructions.OnbuildCommand:
		return &OnBuildCommand{cmd: c}, nil
	case *instructions.VolumeCommand:
		return &VolumeCommand{cmd: c}, nil
	case *instructions.MaintainerCommand:
		logrus.Warnf("%s is deprecated, skipping", cmd.Name())
		return nil, nil
	}
	return nil, errors.Errorf("%s is not a supported command", cmd.Name())
}
