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
import {
    EntCoolroom,
    EntCoolroomFromJSON,
    EntCoolroomFromJSONTyped,
    EntCoolroomToJSON,
    EntPatient,
    EntPatientFromJSON,
    EntPatientFromJSONTyped,
    EntPatientToJSON,
    EntRelative,
    EntRelativeFromJSON,
    EntRelativeFromJSONTyped,
    EntRelativeToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDeceasedReceiveEdges
 */
export interface EntDeceasedReceiveEdges {
    /**
     * 
     * @type {EntCoolroom}
     * @memberof EntDeceasedReceiveEdges
     */
    coolroom?: EntCoolroom;
    /**
     * 
     * @type {EntUser}
     * @memberof EntDeceasedReceiveEdges
     */
    doctor?: EntUser;
    /**
     * 
     * @type {EntPatient}
     * @memberof EntDeceasedReceiveEdges
     */
    patient?: EntPatient;
    /**
     * 
     * @type {EntRelative}
     * @memberof EntDeceasedReceiveEdges
     */
    relative?: EntRelative;
}

export function EntDeceasedReceiveEdgesFromJSON(json: any): EntDeceasedReceiveEdges {
    return EntDeceasedReceiveEdgesFromJSONTyped(json, false);
}

export function EntDeceasedReceiveEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDeceasedReceiveEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'coolroom': !exists(json, 'Coolroom') ? undefined : EntCoolroomFromJSON(json['Coolroom']),
        'doctor': !exists(json, 'Doctor') ? undefined : EntUserFromJSON(json['Doctor']),
        'patient': !exists(json, 'Patient') ? undefined : EntPatientFromJSON(json['Patient']),
        'relative': !exists(json, 'Relative') ? undefined : EntRelativeFromJSON(json['Relative']),
    };
}

export function EntDeceasedReceiveEdgesToJSON(value?: EntDeceasedReceiveEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'coolroom': EntCoolroomToJSON(value.coolroom),
        'doctor': EntUserToJSON(value.doctor),
        'patient': EntPatientToJSON(value.patient),
        'relative': EntRelativeToJSON(value.relative),
    };
}


