/**
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 *
 * OpenAPI Generator version: 4.2.1-SNAPSHOT
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.OpenApiPetstore) {
      root.OpenApiPetstore = {};
    }
    root.OpenApiPetstore.ArrayOfNumberOnly = factory(root.OpenApiPetstore.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';



  /**
   * The ArrayOfNumberOnly model module.
   * @module model/ArrayOfNumberOnly
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>ArrayOfNumberOnly</code>.
   * @alias module:model/ArrayOfNumberOnly
   * @class
   */
  var exports = function() {
    var _this = this;

  };

  /**
   * Constructs a <code>ArrayOfNumberOnly</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/ArrayOfNumberOnly} obj Optional instance to populate.
   * @return {module:model/ArrayOfNumberOnly} The populated <code>ArrayOfNumberOnly</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('ArrayNumber')) {
        obj['ArrayNumber'] = ApiClient.convertToType(data['ArrayNumber'], ['Number']);
      }
    }
    return obj;
  }

  /**
   * @member {Array.<Number>} ArrayNumber
   */
  exports.prototype['ArrayNumber'] = undefined;



  return exports;
}));


