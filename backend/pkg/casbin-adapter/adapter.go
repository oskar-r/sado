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

package mysqladapter

import (
	"log"
	"my-archive/backend/data"
	"strings"

	"github.com/casbin/casbin/model"
	//_ "github.com/go-sql-driver/mysql" // This is for MySQL initialization.
)

// Adapter represents the MySQL adapter for policy storage.
type Adapter struct {
	db *data.PSQLRepo
}

// NewAdapter is the constructor for Adapter.
func NewPsqlAdapter(driverName string, dataSourceName string) *Adapter {
	r, err := data.NewPSQL(dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return &Adapter{
		db: r,
	}
}

func (a *Adapter) Close() {
	a.db.Close()
}

func (a *Adapter) Ping() {
	a.db.Ping()
}

func (a *Adapter) close() {
	a.db.Close()
}

func loadPolicyLine(line string, model model.Model) {
	if line == "" {
		return
	}

	tokens := strings.Split(line, ", ")

	key := tokens[0]
	sec := key[:1]
	model[sec][key].Policy = append(model[sec][key].Policy, tokens[1:])
}

// LoadPolicy loads policy from database.
func (a *Adapter) LoadPolicy(model model.Model) error {
	err := a.db.CreatePolicyDB()
	if err != nil {
		return err
	}

	lines, err := a.db.LoadPolicies("")
	if err != nil {
		return err
	}
	for _, line := range lines {
		loadPolicyLine(line, model)
	}
	return err
}

func policyLine(ptype string, rule []string) []interface{} {

	var p []interface{}

	if ptype != "" {
		p = append(p, ptype)
	} else {
		return p
	}
	for i, v := range rule {
		if i > 2 && ptype != "p" {
			break
		}
		p = append(p, v)
	}

	return p
}

// SavePolicy saves policy to database.
func (a *Adapter) SavePolicy(model model.Model) error {
	return a.db.SaveAllPolicy("", model)
}

func (a *Adapter) AddPolicy(sec string, ptype string, policy []string) error {

	rule := policyLine(ptype, policy)
	_, err := a.db.AddPolicy(rule)
	return err
}

func (a *Adapter) RemovePolicy(sec string, ptype string, policy []string) error {
	rule := policyLine(ptype, policy)
	_, err := a.db.DeletePolicy(rule)
	return err
}

func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	var s []string
	for _, v := range fieldValues {
		s = append(s, v)
	}
	rule := policyLine(ptype, s)
	_, err := a.db.DeletePolicy(rule)
	return err
}
