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
    EntBorrowEdges,
    EntBorrowEdgesFromJSON,
    EntBorrowEdgesFromJSONTyped,
    EntBorrowEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBorrow
 */
export interface EntBorrow {
    /**
     * 
     * @type {EntBorrowEdges}
     * @memberof EntBorrow
     */
    edges?: EntBorrowEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntBorrow
     */
    id?: number;
    /**
     * Type holds the value of the "type" field.
     * @type {string}
     * @memberof EntBorrow
     */
    type?: string;
}

export function EntBorrowFromJSON(json: any): EntBorrow {
    return EntBorrowFromJSONTyped(json, false);
}

export function EntBorrowFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBorrow {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntBorrowEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'type': !exists(json, 'type') ? undefined : json['type'],
    };
}

export function EntBorrowToJSON(value?: EntBorrow | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntBorrowEdgesToJSON(value.edges),
        'id': value.id,
        'type': value.type,
    };
}


