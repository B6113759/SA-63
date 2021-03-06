/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ControllersDeceasedReceive
 */
export interface ControllersDeceasedReceive {
    /**
     * 
     * @type {number}
     * @memberof ControllersDeceasedReceive
     */
    coolroomID?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersDeceasedReceive
     */
    deathtime?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersDeceasedReceive
     */
    patientID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersDeceasedReceive
     */
    relativeID?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersDeceasedReceive
     */
    userID?: number;
}

export function ControllersDeceasedReceiveFromJSON(json: any): ControllersDeceasedReceive {
    return ControllersDeceasedReceiveFromJSONTyped(json, false);
}

export function ControllersDeceasedReceiveFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersDeceasedReceive {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'coolroomID': !exists(json, 'coolroomID') ? undefined : json['coolroomID'],
        'deathtime': !exists(json, 'deathtime') ? undefined : json['deathtime'],
        'patientID': !exists(json, 'patientID') ? undefined : json['patientID'],
        'relativeID': !exists(json, 'relativeID') ? undefined : json['relativeID'],
        'userID': !exists(json, 'userID') ? undefined : json['userID'],
    };
}

export function ControllersDeceasedReceiveToJSON(value?: ControllersDeceasedReceive | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'coolroomID': value.coolroomID,
        'deathtime': value.deathtime,
        'patientID': value.patientID,
        'relativeID': value.relativeID,
        'userID': value.userID,
    };
}


