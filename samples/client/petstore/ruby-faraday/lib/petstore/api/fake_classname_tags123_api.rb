=begin
#OpenAPI Petstore

#This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

The version of the OpenAPI document: 1.0.0

Generated by: https://openapi-generator.tech
OpenAPI Generator version: 4.2.1-SNAPSHOT

=end

require 'cgi'

module Petstore
  class FakeClassnameTags123Api
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end
    # To test class name in snake case
    # To test class name in snake case
    # @param body [Client] client model
    # @param [Hash] opts the optional parameters
    # @return [Client]
    def test_classname(body, opts = {})
      data, _status_code, _headers = test_classname_with_http_info(body, opts)
      data
    end

    # To test class name in snake case
    # To test class name in snake case
    # @param body [Client] client model
    # @param [Hash] opts the optional parameters
    # @return [Array<(Client, Integer, Hash)>] Client data, response status code and response headers
    def test_classname_with_http_info(body, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: FakeClassnameTags123Api.test_classname ...'
      end
      # verify the required parameter 'body' is set
      if @api_client.config.client_side_validation && body.nil?
        fail ArgumentError, "Missing the required parameter 'body' when calling FakeClassnameTags123Api.test_classname"
      end
      # resource path
      local_var_path = '/fake_classname_test'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json'])
      # HTTP header 'Content-Type'
      header_params['Content-Type'] = @api_client.select_header_content_type(['application/json'])

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:body] || @api_client.object_to_http_body(body) 

      # return_type
      return_type = opts[:return_type] || 'Client' 

      # auth_names
      auth_names = opts[:auth_names] || ['api_key_query']

      new_options = opts.merge(
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:PATCH, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: FakeClassnameTags123Api#test_classname\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end
  end
end
