/**
 * SCS Status Page API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.1.2
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { HttpHeaders }                                       from '@angular/common/http';

import { Observable }                                        from 'rxjs';

import { Component } from '../model/models';
import { GetComponent200Response } from '../model/models';
import { GetComponents200Response } from '../model/models';
import { IdField } from '../model/models';


import { Configuration }                                     from '../configuration';



export interface ComponentServiceInterface {
    defaultHeaders: HttpHeaders;
    configuration: Configuration;

    /**
     * Create a new component.
     * 
     * @param component Send a component.
     */
    createComponent(component?: Component, extraHttpRequestParams?: any): Observable<IdField>;

    /**
     * Delete a component.
     * 
     * @param componentId Component ID is required in path.
     */
    deleteComponent(componentId: string, extraHttpRequestParams?: any): Observable<{}>;

    /**
     * Get a specific component by id.
     * 
     * @param componentId Component ID is required in path.
     * @param at An optional at time stamp, used as reference time for the component.
     */
    getComponent(componentId: string, at?: string, extraHttpRequestParams?: any): Observable<GetComponent200Response>;

    /**
     * Get a list of components.
     * 
     * @param at An optional at time stamp, used as reference time for the component.
     */
    getComponents(at?: string, extraHttpRequestParams?: any): Observable<GetComponents200Response>;

    /**
     * Update a component.
     * 
     * @param componentId Component ID is required in path.
     * @param component Send a component.
     */
    updateComponent(componentId: string, component?: Component, extraHttpRequestParams?: any): Observable<{}>;

}
