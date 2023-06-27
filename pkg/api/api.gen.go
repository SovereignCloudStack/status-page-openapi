// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// Component defines model for Component.
type Component struct {
	// ActivelyAffectedBy A list of impacts for an component.
	// Impacts reference incidents.
	ActivelyAffectedBy *ImpactIncidentList `json:"activelyAffectedBy,omitempty"`

	// DisplayName Short and describing name.
	DisplayName *DisplayName `json:"displayName,omitempty"`

	// Labels Labels are free text key value pairs for components.
	Labels *Labels `json:"labels,omitempty"`
}

// ComponentResponseData defines model for ComponentResponseData.
type ComponentResponseData struct {
	// ActivelyAffectedBy A list of impacts for an component.
	// Impacts reference incidents.
	ActivelyAffectedBy *ImpactIncidentList `json:"activelyAffectedBy,omitempty"`

	// DisplayName Short and describing name.
	DisplayName *DisplayName `json:"displayName,omitempty"`

	// Id Identification for objects. UUID preferred.
	Id Id `json:"id"`

	// Labels Labels are free text key value pairs for components.
	Labels *Labels `json:"labels,omitempty"`
}

// Date Point in time. Nullable for open end timeframes.
type Date = time.Time

// Description A longer text with detailed information.
type Description = string

// DisplayName Short and describing name.
type DisplayName = string

// Generation Incremental as generation field for responses.
type Generation struct {
	// Generation Positive and incrementing number for ordering and identfication of e.g. sub resources.
	Generation Incremental `json:"generation"`
}

// Id Identification for objects. UUID preferred.
type Id = string

// IdField Id field for responses.
type IdField struct {
	// Id Identification for objects. UUID preferred.
	Id Id `json:"id"`
}

// Impact An impact connects a component and an incident with an impact type.
type Impact struct {
	// Reference Identification for objects. UUID preferred.
	Reference *Id `json:"reference,omitempty"`

	// Type Identification for objects. UUID preferred.
	Type *Id `json:"type,omitempty"`
}

// ImpactComponentList A list of impacts for an incident.
// Impacts reference components.
type ImpactComponentList = []Impact

// ImpactIncidentList A list of impacts for an component.
// Impacts reference incidents.
type ImpactIncidentList = []Impact

// ImpactType defines model for ImpactType.
type ImpactType struct {
	// Description A longer text with detailed information.
	Description *Description `json:"description,omitempty"`

	// DisplayName Short and describing name.
	DisplayName *DisplayName `json:"displayName,omitempty"`
}

// ImpactTypeResponseData defines model for ImpactTypeResponseData.
type ImpactTypeResponseData struct {
	// Description A longer text with detailed information.
	Description *Description `json:"description,omitempty"`

	// DisplayName Short and describing name.
	DisplayName *DisplayName `json:"displayName,omitempty"`

	// Id Identification for objects. UUID preferred.
	Id Id `json:"id"`
}

// Incident defines model for Incident.
type Incident struct {
	// Affects A list of impacts for an incident.
	// Impacts reference components.
	Affects *ImpactComponentList `json:"affects,omitempty"`

	// BeganAt Point in time. Nullable for open end timeframes.
	BeganAt *Date `json:"beganAt"`

	// Description A longer text with detailed information.
	Description *Description `json:"description,omitempty"`

	// DisplayName Short and describing name.
	DisplayName *DisplayName `json:"displayName,omitempty"`

	// EndedAt Point in time. Nullable for open end timeframes.
	EndedAt *Date `json:"endedAt"`

	// Phase To reference a phase, its generation and order is needed.
	Phase *PhaseReference `json:"phase,omitempty"`

	// Updates List of Incrementals for referencing subresources.
	Updates *IncrementalList `json:"updates,omitempty"`
}

// IncidentResponseData defines model for IncidentResponseData.
type IncidentResponseData struct {
	// Affects A list of impacts for an incident.
	// Impacts reference components.
	Affects *ImpactComponentList `json:"affects,omitempty"`

	// BeganAt Point in time. Nullable for open end timeframes.
	BeganAt *Date `json:"beganAt"`

	// Description A longer text with detailed information.
	Description *Description `json:"description,omitempty"`

	// DisplayName Short and describing name.
	DisplayName *DisplayName `json:"displayName,omitempty"`

	// EndedAt Point in time. Nullable for open end timeframes.
	EndedAt *Date `json:"endedAt"`

	// Id Identification for objects. UUID preferred.
	Id Id `json:"id"`

	// Phase To reference a phase, its generation and order is needed.
	Phase *PhaseReference `json:"phase,omitempty"`

	// Updates List of Incrementals for referencing subresources.
	Updates *IncrementalList `json:"updates,omitempty"`
}

// IncidentUpdate An update is a sub resource to an incident.
// It's identified by the incident ID and its own order.
// Updates happen in a given order.
type IncidentUpdate struct {
	// CreatedAt Point in time. Nullable for open end timeframes.
	CreatedAt *Date `json:"createdAt"`

	// Description A longer text with detailed information.
	Description *Description `json:"description,omitempty"`

	// DisplayName Short and describing name.
	DisplayName *DisplayName `json:"displayName,omitempty"`
}

// IncidentUpdateResponseData defines model for IncidentUpdateResponseData.
type IncidentUpdateResponseData struct {
	// CreatedAt Point in time. Nullable for open end timeframes.
	CreatedAt *Date `json:"createdAt"`

	// Description A longer text with detailed information.
	Description *Description `json:"description,omitempty"`

	// DisplayName Short and describing name.
	DisplayName *DisplayName `json:"displayName,omitempty"`

	// Order Positive and incrementing number for ordering and identfication of e.g. sub resources.
	Order Incremental `json:"order"`
}

// Incremental Positive and incrementing number for ordering and identfication of e.g. sub resources.
type Incremental = int

// IncrementalList List of Incrementals for referencing subresources.
type IncrementalList = []Incremental

// Labels Labels are free text key value pairs for components.
type Labels map[string]string

// Order Incremental as order field for responses.
type Order struct {
	// Order Positive and incrementing number for ordering and identfication of e.g. sub resources.
	Order Incremental `json:"order"`
}

// Phase A single phase is just its name.
// It can be referenced by its generation and order.
// See: #/components/schemas/PhaseReference
type Phase = string

// PhaseList Phase resources are always handled as a list.
type PhaseList struct {
	Phases []Phase `json:"phases"`
}

// PhaseListResponseData defines model for PhaseListResponseData.
type PhaseListResponseData struct {
	// Generation Positive and incrementing number for ordering and identfication of e.g. sub resources.
	Generation Incremental `json:"generation"`
	Phases     []Phase     `json:"phases"`
}

// PhaseReference defines model for PhaseReference.
type PhaseReference struct {
	// Generation Positive and incrementing number for ordering and identfication of e.g. sub resources.
	Generation Incremental `json:"generation"`

	// Order Positive and incrementing number for ordering and identfication of e.g. sub resources.
	Order Incremental `json:"order"`
}

// ComponentIdPathParameter Identification for objects. UUID preferred.
type ComponentIdPathParameter = Id

// GenerationQueryParameter Positive and incrementing number for ordering and identfication of e.g. sub resources.
type GenerationQueryParameter = Incremental

// ImpactTypeIdPathParameter Identification for objects. UUID preferred.
type ImpactTypeIdPathParameter = Id

// IncidentIdPathParameter Identification for objects. UUID preferred.
type IncidentIdPathParameter = Id

// IncidentUpdateOrderPathParameter Positive and incrementing number for ordering and identfication of e.g. sub resources.
type IncidentUpdateOrderPathParameter = Incremental

// ComponentListResponse defines model for ComponentListResponse.
type ComponentListResponse struct {
	Data []ComponentResponseData `json:"data"`
}

// ComponentResponse defines model for ComponentResponse.
type ComponentResponse struct {
	Data ComponentResponseData `json:"data"`
}

// GenerationResponse Incremental as generation field for responses.
type GenerationResponse = Generation

// IdResponse Id field for responses.
type IdResponse = IdField

// ImpactTypeListResponse defines model for ImpactTypeListResponse.
type ImpactTypeListResponse struct {
	Data []ImpactTypeResponseData `json:"data"`
}

// ImpactTypeResponse defines model for ImpactTypeResponse.
type ImpactTypeResponse struct {
	Data ImpactTypeResponseData `json:"data"`
}

// IncidentListResponse defines model for IncidentListResponse.
type IncidentListResponse struct {
	Data []IncidentResponseData `json:"data"`
}

// IncidentResponse defines model for IncidentResponse.
type IncidentResponse struct {
	Data IncidentResponseData `json:"data"`
}

// IncidentUpdateListResponse defines model for IncidentUpdateListResponse.
type IncidentUpdateListResponse struct {
	Data []IncidentUpdateResponseData `json:"data"`
}

// IncidentUpdateResponse defines model for IncidentUpdateResponse.
type IncidentUpdateResponse struct {
	Data IncidentUpdateResponseData `json:"data"`
}

// OrderResponse Incremental as order field for responses.
type OrderResponse = Order

// PhaseListResponse defines model for PhaseListResponse.
type PhaseListResponse struct {
	Data PhaseListResponseData `json:"data"`
}

// ComponentRequest defines model for ComponentRequest.
type ComponentRequest = Component

// ImpactTypeRequest defines model for ImpactTypeRequest.
type ImpactTypeRequest = ImpactType

// IncidentRequest defines model for IncidentRequest.
type IncidentRequest = Incident

// IncidentUpdateRequest An update is a sub resource to an incident.
// It's identified by the incident ID and its own order.
// Updates happen in a given order.
type IncidentUpdateRequest = IncidentUpdate

// PhaseListRequest Phase resources are always handled as a list.
type PhaseListRequest = PhaseList

// GetIncidentsParams defines parameters for GetIncidents.
type GetIncidentsParams struct {
	// Start Start of time frame to query for (RFC3339).
	Start time.Time `form:"start" json:"start"`

	// End End of time frame to query for (RFC3339).
	End time.Time `form:"end" json:"end"`
}

// GetPhaseListParams defines parameters for GetPhaseList.
type GetPhaseListParams struct {
	// Generation Optional generation in query. E.g. ?generation=7
	Generation *GenerationQueryParameter `form:"generation,omitempty" json:"generation,omitempty"`
}

// CreateComponentJSONRequestBody defines body for CreateComponent for application/json ContentType.
type CreateComponentJSONRequestBody = Component

// UpdateComponentJSONRequestBody defines body for UpdateComponent for application/json ContentType.
type UpdateComponentJSONRequestBody = Component

// CreateImpactTypeJSONRequestBody defines body for CreateImpactType for application/json ContentType.
type CreateImpactTypeJSONRequestBody = ImpactType

// UpdateImpactTypeJSONRequestBody defines body for UpdateImpactType for application/json ContentType.
type UpdateImpactTypeJSONRequestBody = ImpactType

// CreateIncidentJSONRequestBody defines body for CreateIncident for application/json ContentType.
type CreateIncidentJSONRequestBody = Incident

// UpdateIncidentJSONRequestBody defines body for UpdateIncident for application/json ContentType.
type UpdateIncidentJSONRequestBody = Incident

// CreateIncidentUpdateJSONRequestBody defines body for CreateIncidentUpdate for application/json ContentType.
type CreateIncidentUpdateJSONRequestBody = IncidentUpdate

// UpdateIncidentUpdateJSONRequestBody defines body for UpdateIncidentUpdate for application/json ContentType.
type UpdateIncidentUpdateJSONRequestBody = IncidentUpdate

// CreatePhaseListJSONRequestBody defines body for CreatePhaseList for application/json ContentType.
type CreatePhaseListJSONRequestBody = PhaseList

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get a list of components.
	// (GET /components)
	GetComponents(ctx echo.Context) error
	// Create a new component.
	// (POST /components)
	CreateComponent(ctx echo.Context) error
	// Delete a component.
	// (DELETE /components/{componentId})
	DeleteComponent(ctx echo.Context, componentId ComponentIdPathParameter) error
	// Get a specific component by id.
	// (GET /components/{componentId})
	GetComponent(ctx echo.Context, componentId ComponentIdPathParameter) error
	// Update a component.
	// (PATCH /components/{componentId})
	UpdateComponent(ctx echo.Context, componentId ComponentIdPathParameter) error
	// Get a list of impact types.
	// (GET /impacttypes)
	GetImpactTypes(ctx echo.Context) error
	// Create a new impact type.
	// (POST /impacttypes)
	CreateImpactType(ctx echo.Context) error
	// Delete an impact type.
	// (DELETE /impacttypes/{impactTypeId})
	DeleteImpactType(ctx echo.Context, impactTypeId ImpactTypeIdPathParameter) error
	// Get a specific impact type by id.
	// (GET /impacttypes/{impactTypeId})
	GetImpactType(ctx echo.Context, impactTypeId ImpactTypeIdPathParameter) error
	// Update a specific impact type.
	// (PATCH /impacttypes/{impactTypeId})
	UpdateImpactType(ctx echo.Context, impactTypeId ImpactTypeIdPathParameter) error
	// Get a list of incidents between two points in time.
	// (GET /incidents)
	GetIncidents(ctx echo.Context, params GetIncidentsParams) error
	// Create a new incident.
	// (POST /incidents)
	CreateIncident(ctx echo.Context) error
	// Delete an incident.
	// (DELETE /incidents/{incidentId})
	DeleteIncident(ctx echo.Context, incidentId IncidentIdPathParameter) error
	// Get a specific incident by id.
	// (GET /incidents/{incidentId})
	GetIncident(ctx echo.Context, incidentId IncidentIdPathParameter) error
	// Update an incident.
	// (PATCH /incidents/{incidentId})
	UpdateIncident(ctx echo.Context, incidentId IncidentIdPathParameter) error
	// Get a list of updates from a specific incident.
	// (GET /incidents/{incidentId}/updates)
	GetIncidentUpdates(ctx echo.Context, incidentId IncidentIdPathParameter) error
	// Create a new update to a specific incident.
	// (POST /incidents/{incidentId}/updates)
	CreateIncidentUpdate(ctx echo.Context, incidentId IncidentIdPathParameter) error
	// Delete a specific update from a specific incident
	// (DELETE /incidents/{incidentId}/updates/{updateOrder})
	DeleteIncidentUpdate(ctx echo.Context, incidentId IncidentIdPathParameter, updateOrder IncidentUpdateOrderPathParameter) error
	// Get a specific update from a specific incident.
	// (GET /incidents/{incidentId}/updates/{updateOrder})
	GetIncidentUpdate(ctx echo.Context, incidentId IncidentIdPathParameter, updateOrder IncidentUpdateOrderPathParameter) error
	// Update a specific update from a specific incident.
	// (PATCH /incidents/{incidentId}/updates/{updateOrder})
	UpdateIncidentUpdate(ctx echo.Context, incidentId IncidentIdPathParameter, updateOrder IncidentUpdateOrderPathParameter) error
	// Get the current generation list of phases.
	// (GET /phases)
	GetPhaseList(ctx echo.Context, params GetPhaseListParams) error
	// Create a new generation of the phase list.
	// (POST /phases)
	CreatePhaseList(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetComponents converts echo context to params.
func (w *ServerInterfaceWrapper) GetComponents(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComponents(ctx)
	return err
}

// CreateComponent converts echo context to params.
func (w *ServerInterfaceWrapper) CreateComponent(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateComponent(ctx)
	return err
}

// DeleteComponent converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteComponent(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "componentId" -------------
	var componentId ComponentIdPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "componentId", runtime.ParamLocationPath, ctx.Param("componentId"), &componentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter componentId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteComponent(ctx, componentId)
	return err
}

// GetComponent converts echo context to params.
func (w *ServerInterfaceWrapper) GetComponent(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "componentId" -------------
	var componentId ComponentIdPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "componentId", runtime.ParamLocationPath, ctx.Param("componentId"), &componentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter componentId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComponent(ctx, componentId)
	return err
}

// UpdateComponent converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateComponent(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "componentId" -------------
	var componentId ComponentIdPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "componentId", runtime.ParamLocationPath, ctx.Param("componentId"), &componentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter componentId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateComponent(ctx, componentId)
	return err
}

// GetImpactTypes converts echo context to params.
func (w *ServerInterfaceWrapper) GetImpactTypes(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetImpactTypes(ctx)
	return err
}

// CreateImpactType converts echo context to params.
func (w *ServerInterfaceWrapper) CreateImpactType(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateImpactType(ctx)
	return err
}

// DeleteImpactType converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteImpactType(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "impactTypeId" -------------
	var impactTypeId ImpactTypeIdPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "impactTypeId", runtime.ParamLocationPath, ctx.Param("impactTypeId"), &impactTypeId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter impactTypeId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteImpactType(ctx, impactTypeId)
	return err
}

// GetImpactType converts echo context to params.
func (w *ServerInterfaceWrapper) GetImpactType(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "impactTypeId" -------------
	var impactTypeId ImpactTypeIdPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "impactTypeId", runtime.ParamLocationPath, ctx.Param("impactTypeId"), &impactTypeId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter impactTypeId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetImpactType(ctx, impactTypeId)
	return err
}

// UpdateImpactType converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateImpactType(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "impactTypeId" -------------
	var impactTypeId ImpactTypeIdPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "impactTypeId", runtime.ParamLocationPath, ctx.Param("impactTypeId"), &impactTypeId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter impactTypeId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateImpactType(ctx, impactTypeId)
	return err
}

// GetIncidents converts echo context to params.
func (w *ServerInterfaceWrapper) GetIncidents(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetIncidentsParams
	// ------------- Required query parameter "start" -------------

	err = runtime.BindQueryParameter("form", true, true, "start", ctx.QueryParams(), &params.Start)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter start: %s", err))
	}

	// ------------- Required query parameter "end" -------------

	err = runtime.BindQueryParameter("form", true, true, "end", ctx.QueryParams(), &params.End)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter end: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetIncidents(ctx, params)
	return err
}

// CreateIncident converts echo context to params.
func (w *ServerInterfaceWrapper) CreateIncident(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateIncident(ctx)
	return err
}

// DeleteIncident converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteIncident(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "incidentId" -------------
	var incidentId IncidentIdPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "incidentId", runtime.ParamLocationPath, ctx.Param("incidentId"), &incidentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter incidentId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteIncident(ctx, incidentId)
	return err
}

// GetIncident converts echo context to params.
func (w *ServerInterfaceWrapper) GetIncident(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "incidentId" -------------
	var incidentId IncidentIdPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "incidentId", runtime.ParamLocationPath, ctx.Param("incidentId"), &incidentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter incidentId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetIncident(ctx, incidentId)
	return err
}

// UpdateIncident converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateIncident(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "incidentId" -------------
	var incidentId IncidentIdPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "incidentId", runtime.ParamLocationPath, ctx.Param("incidentId"), &incidentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter incidentId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateIncident(ctx, incidentId)
	return err
}

// GetIncidentUpdates converts echo context to params.
func (w *ServerInterfaceWrapper) GetIncidentUpdates(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "incidentId" -------------
	var incidentId IncidentIdPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "incidentId", runtime.ParamLocationPath, ctx.Param("incidentId"), &incidentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter incidentId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetIncidentUpdates(ctx, incidentId)
	return err
}

// CreateIncidentUpdate converts echo context to params.
func (w *ServerInterfaceWrapper) CreateIncidentUpdate(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "incidentId" -------------
	var incidentId IncidentIdPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "incidentId", runtime.ParamLocationPath, ctx.Param("incidentId"), &incidentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter incidentId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateIncidentUpdate(ctx, incidentId)
	return err
}

// DeleteIncidentUpdate converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteIncidentUpdate(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "incidentId" -------------
	var incidentId IncidentIdPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "incidentId", runtime.ParamLocationPath, ctx.Param("incidentId"), &incidentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter incidentId: %s", err))
	}

	// ------------- Path parameter "updateOrder" -------------
	var updateOrder IncidentUpdateOrderPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "updateOrder", runtime.ParamLocationPath, ctx.Param("updateOrder"), &updateOrder)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter updateOrder: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteIncidentUpdate(ctx, incidentId, updateOrder)
	return err
}

// GetIncidentUpdate converts echo context to params.
func (w *ServerInterfaceWrapper) GetIncidentUpdate(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "incidentId" -------------
	var incidentId IncidentIdPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "incidentId", runtime.ParamLocationPath, ctx.Param("incidentId"), &incidentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter incidentId: %s", err))
	}

	// ------------- Path parameter "updateOrder" -------------
	var updateOrder IncidentUpdateOrderPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "updateOrder", runtime.ParamLocationPath, ctx.Param("updateOrder"), &updateOrder)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter updateOrder: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetIncidentUpdate(ctx, incidentId, updateOrder)
	return err
}

// UpdateIncidentUpdate converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateIncidentUpdate(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "incidentId" -------------
	var incidentId IncidentIdPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "incidentId", runtime.ParamLocationPath, ctx.Param("incidentId"), &incidentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter incidentId: %s", err))
	}

	// ------------- Path parameter "updateOrder" -------------
	var updateOrder IncidentUpdateOrderPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "updateOrder", runtime.ParamLocationPath, ctx.Param("updateOrder"), &updateOrder)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter updateOrder: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateIncidentUpdate(ctx, incidentId, updateOrder)
	return err
}

// GetPhaseList converts echo context to params.
func (w *ServerInterfaceWrapper) GetPhaseList(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPhaseListParams
	// ------------- Optional query parameter "generation" -------------

	err = runtime.BindQueryParameter("form", true, false, "generation", ctx.QueryParams(), &params.Generation)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter generation: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPhaseList(ctx, params)
	return err
}

// CreatePhaseList converts echo context to params.
func (w *ServerInterfaceWrapper) CreatePhaseList(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreatePhaseList(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/components", wrapper.GetComponents)
	router.POST(baseURL+"/components", wrapper.CreateComponent)
	router.DELETE(baseURL+"/components/:componentId", wrapper.DeleteComponent)
	router.GET(baseURL+"/components/:componentId", wrapper.GetComponent)
	router.PATCH(baseURL+"/components/:componentId", wrapper.UpdateComponent)
	router.GET(baseURL+"/impacttypes", wrapper.GetImpactTypes)
	router.POST(baseURL+"/impacttypes", wrapper.CreateImpactType)
	router.DELETE(baseURL+"/impacttypes/:impactTypeId", wrapper.DeleteImpactType)
	router.GET(baseURL+"/impacttypes/:impactTypeId", wrapper.GetImpactType)
	router.PATCH(baseURL+"/impacttypes/:impactTypeId", wrapper.UpdateImpactType)
	router.GET(baseURL+"/incidents", wrapper.GetIncidents)
	router.POST(baseURL+"/incidents", wrapper.CreateIncident)
	router.DELETE(baseURL+"/incidents/:incidentId", wrapper.DeleteIncident)
	router.GET(baseURL+"/incidents/:incidentId", wrapper.GetIncident)
	router.PATCH(baseURL+"/incidents/:incidentId", wrapper.UpdateIncident)
	router.GET(baseURL+"/incidents/:incidentId/updates", wrapper.GetIncidentUpdates)
	router.POST(baseURL+"/incidents/:incidentId/updates", wrapper.CreateIncidentUpdate)
	router.DELETE(baseURL+"/incidents/:incidentId/updates/:updateOrder", wrapper.DeleteIncidentUpdate)
	router.GET(baseURL+"/incidents/:incidentId/updates/:updateOrder", wrapper.GetIncidentUpdate)
	router.PATCH(baseURL+"/incidents/:incidentId/updates/:updateOrder", wrapper.UpdateIncidentUpdate)
	router.GET(baseURL+"/phases", wrapper.GetPhaseList)
	router.POST(baseURL+"/phases", wrapper.CreatePhaseList)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
