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
    EntStadiumEdges,
    EntStadiumEdgesFromJSON,
    EntStadiumEdgesFromJSONTyped,
    EntStadiumEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntStadium
 */
export interface EntStadium {
    /**
     * 
     * @type {EntStadiumEdges}
     * @memberof EntStadium
     */
    edges?: EntStadiumEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntStadium
     */
    id?: number;
    /**
     * Namestadium holds the value of the "namestadium" field.
     * @type {string}
     * @memberof EntStadium
     */
    namestadium?: string;
}

export function EntStadiumFromJSON(json: any): EntStadium {
    return EntStadiumFromJSONTyped(json, false);
}

export function EntStadiumFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntStadium {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntStadiumEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'namestadium': !exists(json, 'namestadium') ? undefined : json['namestadium'],
    };
}

export function EntStadiumToJSON(value?: EntStadium | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntStadiumEdgesToJSON(value.edges),
        'id': value.id,
        'namestadium': value.namestadium,
    };
}


