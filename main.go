/*
 * Copyright 2018. Akamai Technologies, Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"os"
	"fmt"
	"bytes"
	"log"
	"encoding/json"
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/papi-v1"
        "github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"

) 

type Input struct  {
	edgerc string
	section string
	PropertyId string
	GroupId string
}

type Output struct {
	Rules string
}

func main() {

	var record Input

	err := json.NewDecoder(os.Stdin).Decode(&record)
	if err != nil {
		log.Fatal(err)
	}

        config, err := edgegrid.Init(record.edgerc, record.section)
        if err != nil {
		log.Fatal(err)
        }

	papi.Init(config)

	// Get Property
        property := papi.NewProperty(papi.NewProperties())
        property.PropertyID = record.PropertyId
        property.GroupID = record.GroupId
        err = property.GetProperty()

	if err != nil {
		log.Fatal(err)
	}

	// Get Property Rules
	rules, err := property.GetRules();

	if err != nil {
		log.Fatal(err)
	}

	jsonBody, err := json.Marshal(rules)
	buf := bytes.NewBufferString("")
	buf.Write(jsonBody)

	var output Output
	output.Rules = buf.String()

	jsonBody, err = json.Marshal(output)
	buf = bytes.NewBufferString("")
	buf.Write(jsonBody)

	fmt.Print(buf.String())
	
}
