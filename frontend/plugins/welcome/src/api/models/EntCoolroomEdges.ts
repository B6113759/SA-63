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
    EntCoolroomType,
    EntCoolroomTypeFromJSON,
    EntCoolroomTypeFromJSONTyped,
    EntCoolroomTypeToJSON,
    EntDeceasedReceive,
    EntDeceasedReceiveFromJSON,
    EntDeceasedReceiveFromJSONTyped,
    EntDeceasedReceiveToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntCoolroomEdges
 */
export interface EntCoolroomEdges {
    /**
     * 
     * @type {EntCoolroomType}
     * @memberof EntCoolroomEdges
     */
    coolroomtype?: EntCoolroomType;
    /**
     * Deceasedreceives holds the value of the deceasedreceives edge.
     * @type {Array<EntDeceasedReceive>}
     * @memberof EntCoolroomEdges
     */
    deceasedreceives?: Array<EntDeceasedReceive>;
}

export function EntCoolroomEdgesFromJSON(json: any): EntCoolroomEdges {
    return EntCoolroomEdgesFromJSONTyped(json, false);
}

export function EntCoolroomEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntCoolroomEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'coolroomtype': !exists(json, 'Coolroomtype') ? undefined : EntCoolroomTypeFromJSON(json['Coolroomtype']),
        'deceasedreceives': !exists(json, 'Deceasedreceives') ? undefined : ((json['Deceasedreceives'] as Array<any>).map(EntDeceasedReceiveFromJSON)),
    };
}

export function EntCoolroomEdgesToJSON(value?: EntCoolroomEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'coolroomtype': EntCoolroomTypeToJSON(value.coolroomtype),
        'deceasedreceives': value.deceasedreceives === undefined ? undefined : ((value.deceasedreceives as Array<any>).map(EntDeceasedReceiveToJSON)),
    };
}

