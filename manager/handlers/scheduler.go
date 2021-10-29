/*
 *     Copyright 2020 The Dragonfly Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package handlers

import (
	"net/http"

	"d7y.io/dragonfly/v2/manager/types"
	"github.com/gin-gonic/gin"
)

// @Summary Create Scheduler
// @Description create by json config
// @Tags Scheduler
// @Accept json
// @Produce json
// @Param Scheduler body types.CreateSchedulerRequest true "Scheduler"
// @Success 200 {object} model.Scheduler
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /schedulers [post]
func (h *Handlers) CreateScheduler(ctx *gin.Context) {
	var json types.CreateSchedulerRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	scheduler, err := h.service.CreateScheduler(ctx.Request.Context(), json)
	if err != nil {
		ctx.Error(err) // nolint: errcheck
		return
	}

	ctx.JSON(http.StatusOK, scheduler)
}

// @Summary Destroy Scheduler
// @Description Destroy by id
// @Tags Scheduler
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /schedulers/{id} [delete]
func (h *Handlers) DestroyScheduler(ctx *gin.Context) {
	var params types.SchedulerParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	if err := h.service.DestroyScheduler(ctx.Request.Context(), params.ID); err != nil {
		ctx.Error(err) // nolint: errcheck
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary Update Scheduler
// @Description Update by json config
// @Tags Scheduler
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param Scheduler body types.UpdateSchedulerRequest true "Scheduler"
// @Success 200 {object} model.Scheduler
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /schedulers/{id} [patch]
func (h *Handlers) UpdateScheduler(ctx *gin.Context) {
	var params types.SchedulerParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.Error(err) // nolint: errcheck
		return
	}

	var json types.UpdateSchedulerRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.Error(err) // nolint: errcheck
		return
	}

	scheduler, err := h.service.UpdateScheduler(ctx.Request.Context(), params.ID, json)
	if err != nil {
		ctx.Error(err) // nolint: errcheck
		return
	}

	ctx.JSON(http.StatusOK, scheduler)
}

// @Summary Get Scheduler
// @Description Get Scheduler by id
// @Tags Scheduler
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} model.Scheduler
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /schedulers/{id} [get]
func (h *Handlers) GetScheduler(ctx *gin.Context) {
	var params types.SchedulerParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	scheduler, err := h.service.GetScheduler(ctx.Request.Context(), params.ID)
	if err != nil {
		ctx.Error(err) // nolint: errcheck
		return
	}

	ctx.JSON(http.StatusOK, scheduler)
}

// @Summary Get Schedulers
// @Description Get Schedulers
// @Tags Scheduler
// @Accept json
// @Produce json
// @Param page query int true "current page" default(0)
// @Param per_page query int true "return max item count, default 10, max 50" default(10) minimum(2) maximum(50)
// @Success 200 {object} []model.Scheduler
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /schedulers [get]
func (h *Handlers) GetSchedulers(ctx *gin.Context) {
	var query types.GetSchedulersQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	h.setPaginationDefault(&query.Page, &query.PerPage)
	schedulers, count, err := h.service.GetSchedulers(ctx.Request.Context(), query)
	if err != nil {
		ctx.Error(err) // nolint: errcheck
		return
	}

	h.setPaginationLinkHeader(ctx, query.Page, query.PerPage, int(count))
	ctx.JSON(http.StatusOK, schedulers)
}
