/*
 *  Licensed to the Apache Software Foundation (ASF) under one
 *  or more contributor license agreements.  See the NOTICE file
 *  distributed with this work for additional information
 *  regarding copyright ownership.  The ASF licenses this file
 *  to you under the Apache License, Version 2.0 (the
 *  "License"); you may not use this file except in compliance
 *  with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing,
 *  software distributed under the License is distributed on an
 *   * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 *  KIND, either express or implied.  See the License for the
 *  specific language governing permissions and limitations
 *  under the License.
 */

package synapse

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"time"

	"github.com/apache/synapse-go/internal/app/adapters/mediation"
	"github.com/apache/synapse-go/internal/pkg/core/artifacts"
	"github.com/apache/synapse-go/internal/pkg/core/deployers"
	"github.com/apache/synapse-go/internal/pkg/core/utils"
)

func Run(ctx context.Context) error {

	start := time.Now()
	PrintWelcomeMessage()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(ctx)
	ctx = context.WithValue(ctx, utils.WaitGroupKey, &wg)

	// TODO: Load configuration
	// TODO: Create loggers and add to context

	defer cancel()

	// create instace variables

	// Adding config context to the GO context
	conCtx := artifacts.GetConfigContext()

	ctx = context.WithValue(ctx, utils.ConfigContextKey, conCtx)

	mediationEngine := mediation.NewMediationEngine()

	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path: ", err)
	}
	binDir := filepath.Dir(exePath)

	artifactsPath := filepath.Join(binDir, "..", "artifacts")
	deployer := deployers.NewDeployer(artifactsPath, mediationEngine)
	err = deployer.Deploy(ctx)
	if err != nil {
		fmt.Println("Error deploying artifacts: ", err)
	}

	elapsed := time.Since(start)
	fmt.Println("Server started in : ", elapsed)

	<-ctx.Done()
	wg.Wait()
	return nil
}

func PrintWelcomeMessage() {
	colors := []string{
		"\033[31m", // Red
		"\033[33m", // Yellow
		"\033[32m", // Green
		"\033[36m", // Cyan
		"\033[34m", // Blue
		"\033[35m", // Magenta
	}

	// ANSI code to reset color to default
	reset := "\033[0m"

	art := []string{
		"",
		"      _/_/_/                                                             ",
		"   _/        _/    _/  _/_/_/      _/_/_/  _/_/_/      _/_/_/    _/_/    ",
		"    _/_/    _/    _/  _/    _/  _/    _/  _/    _/  _/_/      _/_/_/_/   ",
		"       _/  _/    _/  _/    _/  _/    _/  _/    _/      _/_/  _/          ",
		"_/_/_/      _/_/_/  _/    _/    _/_/_/  _/_/_/    _/_/_/      _/_/_/     ",
		"               _/                      _/                                ",
		"          _/_/                        _/                                 ",
	}
	// Iterate over each line of the ASCII art
	for _, line := range art {
		// Iterate over each character in the line
		for i, char := range line {
			// Select color based on character position to create a gradient
			color := colors[i%len(colors)]
			// Print the colored character without adding a newline
			fmt.Printf("%s%c", color, char)
		}
		// Reset color at the end of each line and add a newline
		fmt.Println(reset)
	}
}
