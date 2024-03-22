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

import { Generation } from '../model/models';
import { GetPhaseList200Response } from '../model/models';
import { PhaseList } from '../model/models';


import { Configuration }                                     from '../configuration';



export interface PhaseServiceInterface {
    defaultHeaders: HttpHeaders;
    configuration: Configuration;

    /**
     * Create a new generation of the phase list.
     * 
     * @param phaseList Send a new list of phases.
     */
    createPhaseList(phaseList?: PhaseList, extraHttpRequestParams?: any): Observable<Generation>;

    /**
     * Get the current generation list of phases.
     * 
     * @param generation Optional generation in query. E.g. ?generation&#x3D;7
     */
    getPhaseList(generation?: number, extraHttpRequestParams?: any): Observable<GetPhaseList200Response>;

}
