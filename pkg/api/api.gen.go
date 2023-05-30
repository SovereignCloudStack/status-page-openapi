// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
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
	// AffectedBy A list of impacts for a component or incident.
	// In case of an incident the IDs refer to a component and vice versa.
	AffectedBy *ImpactList `json:"affectedBy,omitempty"`

	// DisplayName Short and describing name.
	DisplayName *DisplayName `json:"displayName,omitempty"`

	// Id Identification for objects. UUID preferred.
	Id *Id `json:"id,omitempty"`

	// Labels Labels are free text key value pairs for components.
	Labels *Labels `json:"labels,omitempty"`
}

// Date Point in time. Nullable for open end timeframes.
type Date = time.Time

// Description A longer text with detailed information.
type Description = string

// DisplayName Short and describing name.
type DisplayName = string

// Id Identification for objects. UUID preferred.
type Id = string

// Impact An impact connects a component and an incident with an impact type.
type Impact struct {
	// Reference Identification for objects. UUID preferred.
	Reference *Id         `json:"reference,omitempty"`
	Type      *ImpactType `json:"type,omitempty"`
}

// ImpactList A list of impacts for a component or incident.
// In case of an incident the IDs refer to a component and vice versa.
type ImpactList = []Impact

// ImpactType defines model for ImpactType.
type ImpactType struct {
	// Description A longer text with detailed information.
	Description *Description `json:"description,omitempty"`

	// DisplayName Short and describing name.
	DisplayName *DisplayName `json:"displayName,omitempty"`

	// Id Identification for objects. UUID preferred.
	Id *Id `json:"id,omitempty"`
}

// Incident defines model for Incident.
type Incident struct {
	// Affects A list of impacts for a component or incident.
	// In case of an incident the IDs refer to a component and vice versa.
	Affects *ImpactList `json:"affects,omitempty"`

	// BeganAt Point in time. Nullable for open end timeframes.
	BeganAt *Date `json:"beganAt"`

	// Description A longer text with detailed information.
	Description *Description `json:"description,omitempty"`

	// DisplayName Short and describing name.
	DisplayName *DisplayName `json:"displayName,omitempty"`

	// EndedAt Point in time. Nullable for open end timeframes.
	EndedAt *Date `json:"endedAt"`

	// Id Identification for objects. UUID preferred.
	Id *Id `json:"id,omitempty"`

	// Phase To reference a phase, its generation and order is needed.
	Phase *PhaseReference `json:"phase,omitempty"`

	// Updates Positive and incrementing number for ordering and identfication of e.g. sub resources.
	Updates *Incremental `json:"updates,omitempty"`
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

	// Order Positive and incrementing number for ordering and identfication of e.g. sub resources.
	Order *Incremental `json:"order,omitempty"`
}

// Incremental Positive and incrementing number for ordering and identfication of e.g. sub resources.
type Incremental = int

// Labels Labels are free text key value pairs for components.
type Labels map[string]string

// Phase A single phase is just its name.
// It can be referenced by its generation and order.
// See: #/components/schemas/PhaseReference
type Phase = string

// PhaseList Phase resources are always handled as a list.
type PhaseList struct {
	// Generation Positive and incrementing number for ordering and identfication of e.g. sub resources.
	Generation *Incremental `json:"generation,omitempty"`
	Phases     []Phase      `json:"phases"`
}

// PhaseReference To reference a phase, its generation and order is needed.
type PhaseReference struct {
	// DisplayName Short and describing name.
	DisplayName *DisplayName `json:"displayName,omitempty"`

	// Generation Positive and incrementing number for ordering and identfication of e.g. sub resources.
	Generation *Incremental `json:"generation,omitempty"`

	// Order Positive and incrementing number for ordering and identfication of e.g. sub resources.
	Order *Incremental `json:"order,omitempty"`
}

// ComponentIdPathParameter Identification for objects. UUID preferred.
type ComponentIdPathParameter = Id

// ImpactTypeIdPathParameter Identification for objects. UUID preferred.
type ImpactTypeIdPathParameter = Id

// IncidentIdPathParameter Identification for objects. UUID preferred.
type IncidentIdPathParameter = Id

// IncidentUpdateIdPathParameter Positive and incrementing number for ordering and identfication of e.g. sub resources.
type IncidentUpdateIdPathParameter = Incremental

// IdResponse Identification for objects. UUID preferred.
type IdResponse = Id

// IncrementalResponse Positive and incrementing number for ordering and identfication of e.g. sub resources.
type IncrementalResponse = Incremental

// SuccessResponse defines model for SuccessResponse.
type SuccessResponse = bool

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

// GetIncidentsParams defines parameters for GetIncidents.
type GetIncidentsParams struct {
	// Start Start of time frame to query for (RFC3339).
	Start time.Time `form:"start" json:"start"`

	// End End of time frame to query for (RFC3339).
	End time.Time `form:"end" json:"end"`
}

// GetPhaseListParams defines parameters for GetPhaseList.
type GetPhaseListParams struct {
	Generation *Incremental `form:"generation,omitempty" json:"generation,omitempty"`
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
	// (DELETE /incidents/{incidentId}/updates/{updateId})
	DeleteIncidentUpdate(ctx echo.Context, incidentId IncidentIdPathParameter, updateId IncidentUpdateIdPathParameter) error
	// Get a specific update from a specific incident.
	// (GET /incidents/{incidentId}/updates/{updateId})
	GetIncidentUpdate(ctx echo.Context, incidentId IncidentIdPathParameter, updateId IncidentUpdateIdPathParameter) error
	// Update a specific update from a specific incident.
	// (PATCH /incidents/{incidentId}/updates/{updateId})
	UpdateIncidentUpdate(ctx echo.Context, incidentId IncidentIdPathParameter, updateId IncidentUpdateIdPathParameter) error
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

	// ------------- Path parameter "updateId" -------------
	var updateId IncidentUpdateIdPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "updateId", runtime.ParamLocationPath, ctx.Param("updateId"), &updateId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter updateId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteIncidentUpdate(ctx, incidentId, updateId)
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

	// ------------- Path parameter "updateId" -------------
	var updateId IncidentUpdateIdPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "updateId", runtime.ParamLocationPath, ctx.Param("updateId"), &updateId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter updateId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetIncidentUpdate(ctx, incidentId, updateId)
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

	// ------------- Path parameter "updateId" -------------
	var updateId IncidentUpdateIdPathParameter

	err = runtime.BindStyledParameterWithLocation("simple", false, "updateId", runtime.ParamLocationPath, ctx.Param("updateId"), &updateId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter updateId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateIncidentUpdate(ctx, incidentId, updateId)
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
	router.DELETE(baseURL+"/incidents/:incidentId/updates/:updateId", wrapper.DeleteIncidentUpdate)
	router.GET(baseURL+"/incidents/:incidentId/updates/:updateId", wrapper.GetIncidentUpdate)
	router.PATCH(baseURL+"/incidents/:incidentId/updates/:updateId", wrapper.UpdateIncidentUpdate)
	router.GET(baseURL+"/phases", wrapper.GetPhaseList)
	router.POST(baseURL+"/phases", wrapper.CreatePhaseList)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"iCD+FwAA///uLJ3G/i0AAA==",
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
