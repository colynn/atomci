/*
Copyright 2021 The AtomCI Group Authors.

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

package api

import (
	"github.com/go-atomci/atomci/internal/core/pipelinemgr"
	"github.com/go-atomci/atomci/internal/middleware/log"
)

// TaskTemplateController ...
type TaskTemplateController struct {
	BaseController
}

// GetTaskTmpls ..
func (p *TaskTemplateController) GetTaskTmpls() {
	pm := pipelinemgr.NewPipelineManager()
	rsp, err := pm.GetTaskTmpls()
	if err != nil {
		p.HandleInternalServerError(err.Error())
		log.Log.Error("get all task templates occur error: %s", err.Error())
		return
	}
	p.Data["json"] = NewResult(true, rsp, "")
	p.ServeJSON()
}

// GetTaskTmplsByPagination ..
func (p *TaskTemplateController) GetTaskTmplsByPagination() {
	filterQuery := p.GetFilterQuery()
	pm := pipelinemgr.NewPipelineManager()
	rsp, err := pm.GetTaskTmplsByPagination(filterQuery)
	if err != nil {
		p.HandleInternalServerError(err.Error())
		log.Log.Error("get all task templates occur error: %s", err.Error())
		return
	}
	p.Data["json"] = NewResult(true, rsp, "")
	p.ServeJSON()
}

// UpdateTaskTmpl ..
func (p *TaskTemplateController) UpdateTaskTmpl() {
	stepID, _ := p.GetInt64FromPath(":step_id")
	request := pipelinemgr.TaskTmplReq{}
	p.DecodeJSONReq(&request)
	pm := pipelinemgr.NewPipelineManager()
	err := pm.UpdateTaskTmpl(&request, stepID)
	if err != nil {
		p.HandleInternalServerError(err.Error())
		log.Log.Error("update flow step occur error: %s", err.Error())
		return
	}
	p.Data["json"] = NewResult(true, nil, "")
	p.ServeJSON()
}

// CreateTaskTmpl ..
func (p *TaskTemplateController) CreateTaskTmpl() {
	request := pipelinemgr.TaskTmplReq{}
	creator := p.User
	p.DecodeJSONReq(&request)
	pm := pipelinemgr.NewPipelineManager()
	err := pm.CreateTaskTmpl(&request, creator)
	if err != nil {
		p.HandleInternalServerError(err.Error())
		log.Log.Error("create flow step occur error: %s", err.Error())
		return
	}
	p.Data["json"] = NewResult(true, nil, "")
	p.ServeJSON()
}

// DeleteTaskTmpl ..
func (p *TaskTemplateController) DeleteTaskTmpl() {
	stepID, _ := p.GetInt64FromPath(":step_id")
	pm := pipelinemgr.NewPipelineManager()
	err := pm.DeleteTaskTmpl(stepID)
	if err != nil {
		p.HandleInternalServerError(err.Error())
		log.Log.Error("delete flow step occur error: %s", err.Error())
		return
	}
	p.Data["json"] = NewResult(true, nil, "")
	p.ServeJSON()
}
