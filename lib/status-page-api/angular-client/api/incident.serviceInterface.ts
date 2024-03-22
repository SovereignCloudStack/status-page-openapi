/**
 * SCS Status Page API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.1
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { HttpHeaders }                                       from '@angular/common/http';

import { Observable }                                        from 'rxjs';

import { GetIncident200Response } from '../model/models';
import { GetIncidentUpdate200Response } from '../model/models';
import { GetIncidentUpdates200Response } from '../model/models';
import { GetIncidents200Response } from '../model/models';
import { IdField } from '../model/models';
import { Incident } from '../model/models';
import { IncidentUpdate } from '../model/models';
import { Order } from '../model/models';


import { Configuration }                                     from '../configuration';



export interface IncidentServiceInterface {
    defaultHeaders: HttpHeaders;
    configuration: Configuration;

    /**
     * Create a new incident.
     * 
     * @param incident Send an incident.
     */
    createIncident(incident?: Incident, extraHttpRequestParams?: any): Observable<IdField>;

    /**
     * Create a new update to a specific incident.
     * 
     * @param incidentId Incident ID is required in path.
     * @param incidentUpdate Send an incident update.
     */
    createIncidentUpdate(incidentId: string, incidentUpdate?: IncidentUpdate, extraHttpRequestParams?: any): Observable<Order>;

    /**
     * Delete an incident.
     * 
     * @param incidentId Incident ID is required in path.
     */
    deleteIncident(incidentId: string, extraHttpRequestParams?: any): Observable<{}>;

    /**
     * Delete a specific update from a specific incident
     * 
     * @param incidentId Incident ID is required in path.
     * @param updateOrder Incident update order (inremental) is required in path.
     */
    deleteIncidentUpdate(incidentId: string, updateOrder: number, extraHttpRequestParams?: any): Observable<{}>;

    /**
     * Get a specific incident by id.
     * 
     * @param incidentId Incident ID is required in path.
     */
    getIncident(incidentId: string, extraHttpRequestParams?: any): Observable<GetIncident200Response>;

    /**
     * Get a specific update from a specific incident.
     * 
     * @param incidentId Incident ID is required in path.
     * @param updateOrder Incident update order (inremental) is required in path.
     */
    getIncidentUpdate(incidentId: string, updateOrder: number, extraHttpRequestParams?: any): Observable<GetIncidentUpdate200Response>;

    /**
     * Get a list of updates from a specific incident.
     * 
     * @param incidentId Incident ID is required in path.
     */
    getIncidentUpdates(incidentId: string, extraHttpRequestParams?: any): Observable<GetIncidentUpdates200Response>;

    /**
     * Get a list of incidents between two points in time.
     * 
     * @param start Start of time frame to query for (RFC3339).
     * @param end End of time frame to query for (RFC3339).
     */
    getIncidents(start: string, end: string, extraHttpRequestParams?: any): Observable<GetIncidents200Response>;

    /**
     * Update an incident.
     * 
     * @param incidentId Incident ID is required in path.
     * @param incident Send an incident.
     */
    updateIncident(incidentId: string, incident?: Incident, extraHttpRequestParams?: any): Observable<{}>;

    /**
     * Update a specific update from a specific incident.
     * 
     * @param incidentId Incident ID is required in path.
     * @param updateOrder Incident update order (inremental) is required in path.
     * @param incidentUpdate Send an incident update.
     */
    updateIncidentUpdate(incidentId: string, updateOrder: number, incidentUpdate?: IncidentUpdate, extraHttpRequestParams?: any): Observable<{}>;

}
