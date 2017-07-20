// Copyright 2017 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fileadapter

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/casbin/casbin/model"
	"github.com/casbin/casbin/persist"
	"github.com/casbin/casbin/util"
)

// Adapter represents the file adapter for policy storage.
// It can load policy from file or save policy to file.
type Adapter struct {
	filePath string
}

// NewAdapter is the constructor for Adapter.
func NewAdapter(filePath string) *Adapter {
	a := Adapter{}
	a.filePath = filePath
	return &a
}

// LoadPolicy loads policy from file.
func (a *Adapter) LoadPolicy(model model.Model) error {
	if a.filePath == "" {
		return errors.New("Invalid file path, file path cannot be empty")
	}

	err := a.loadPolicyFile(model, persist.LoadPolicyLine)
	return err
}

// SavePolicy saves policy to file.
func (a *Adapter) SavePolicy(model model.Model) error {
	if a.filePath == "" {
		return errors.New("Invalid file path, file path cannot be empty")
	}

	var tmp bytes.Buffer

	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			tmp.WriteString(ptype + ", ")
			tmp.WriteString(util.ArrayToString(rule))
			tmp.WriteString("\n")
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			tmp.WriteString(ptype + ", ")
			tmp.WriteString(util.ArrayToString(rule))
			tmp.WriteString("\n")
		}
	}

	err := a.savePolicyFile(strings.TrimRight(tmp.String(), "\n"))
	return err
}

func (a *Adapter) loadPolicyFile(model model.Model, handler func(string, model.Model)) error {
	f, err := os.Open(a.filePath)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		handler(line, model)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func (a *Adapter) savePolicyFile(text string) error {
	f, err := os.Create(a.filePath)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	w.WriteString(text)
	w.Flush()
	f.Close()
	return nil
}

func (a *Adapter) AddPolicy(sec string, ptype string, policy []string) error {
	return errors.New("not implemented")
}

func (a *Adapter) RemovePolicy(sec string, ptype string, policy []string) error {
	return errors.New("not implemented")
}

func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return errors.New("not implemented")
}