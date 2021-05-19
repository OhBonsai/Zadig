/*
Copyright 2021 The KodeRover Authors.

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

package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/koderover/zadig/lib/microservice/aslan/core/workflow/service/workflow"
	internalhandler "github.com/koderover/zadig/lib/microservice/aslan/internal/handler"
)

func ListServiceWorkflows(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JsonResponse(c, ctx) }()

	productName := c.Param("productName")
	envName := c.Param("envName")
	serviceName := c.Param("serviceName")
	serviceType := c.Param("serviceType")

	ctx.Resp, ctx.Err = workflow.ListServiceWorkflows(productName, envName, serviceName, serviceType, ctx.Logger)
}