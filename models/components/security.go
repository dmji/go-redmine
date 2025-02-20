// Code generated by Speakeasy (https://speakeasy.com).

package components

type Security struct {
	BasicAuth         *SchemeBasicAuth `security:"scheme,type=http,subtype=basic"`
	APIKeyAuth        *string          `security:"scheme,type=apiKey,subtype=header,name=X-Redmine-API-Key,env=redmine_api_key_auth"`
	APIKeyInQueryAuth *string          `security:"scheme,type=apiKey,subtype=query,name=key,env=redmine_api_key_in_query_auth"`
}

func (o *Security) GetBasicAuth() *SchemeBasicAuth {
	if o == nil {
		return nil
	}
	return o.BasicAuth
}

func (o *Security) GetAPIKeyAuth() *string {
	if o == nil {
		return nil
	}
	return o.APIKeyAuth
}

func (o *Security) GetAPIKeyInQueryAuth() *string {
	if o == nil {
		return nil
	}
	return o.APIKeyInQueryAuth
}
