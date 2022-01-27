/*
 * loadbalancer
 *
 * <br/>https://ncloud.apigw.ntruss.com/loadbalancer/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package loadbalancer

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/url"
	"reflect"
	"strings"
	"bytes"

	"golang.org/x/net/context"
)

// Linger please
var (
	_ context.Context
)

type V2ApiService service


/* V2ApiService 
 @param addLoadBalancerSslCertificateRequest addLoadBalancerSslCertificateRequest
 @return *AddLoadBalancerSslCertificateResponse*/
func (a *V2ApiService) AddLoadBalancerSslCertificate(addLoadBalancerSslCertificateRequest *AddLoadBalancerSslCertificateRequest) (*AddLoadBalancerSslCertificateResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AddLoadBalancerSslCertificateResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/addLoadBalancerSslCertificate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/x-www-form-urlencoded",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = addLoadBalancerSslCertificateRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return &successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return &successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	if localVarHttpResponse.StatusCode >= 300 || (localVarHttpResponse.StatusCode < 300 && !strings.HasPrefix(string(bodyBytes), `{`)) {
		return &successPayload, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if !strings.Contains(string(bodyBytes), `{"error"`) && strings.HasPrefix(string(bodyBytes), `{`) {
		if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
			return &successPayload, err
		}
	}


	return &successPayload, err
}

/* V2ApiService 
 @param addServerInstancesToLoadBalancerRequest addServerInstancesToLoadBalancerRequest
 @return *AddServerInstancesToLoadBalancerResponse*/
func (a *V2ApiService) AddServerInstancesToLoadBalancer(addServerInstancesToLoadBalancerRequest *AddServerInstancesToLoadBalancerRequest) (*AddServerInstancesToLoadBalancerResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AddServerInstancesToLoadBalancerResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/addServerInstancesToLoadBalancer"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/x-www-form-urlencoded",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = addServerInstancesToLoadBalancerRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return &successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return &successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	if localVarHttpResponse.StatusCode >= 300 || (localVarHttpResponse.StatusCode < 300 && !strings.HasPrefix(string(bodyBytes), `{`)) {
		return &successPayload, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if !strings.Contains(string(bodyBytes), `{"error"`) && strings.HasPrefix(string(bodyBytes), `{`) {
		if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
			return &successPayload, err
		}
	}


	return &successPayload, err
}

/* V2ApiService 
 @param changeLoadBalancedServerInstancesRequest changeLoadBalancedServerInstancesRequest
 @return *ChangeLoadBalancedServerInstancesResponse*/
func (a *V2ApiService) ChangeLoadBalancedServerInstances(changeLoadBalancedServerInstancesRequest *ChangeLoadBalancedServerInstancesRequest) (*ChangeLoadBalancedServerInstancesResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ChangeLoadBalancedServerInstancesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/changeLoadBalancedServerInstances"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/x-www-form-urlencoded",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = changeLoadBalancedServerInstancesRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return &successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return &successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	if localVarHttpResponse.StatusCode >= 300 || (localVarHttpResponse.StatusCode < 300 && !strings.HasPrefix(string(bodyBytes), `{`)) {
		return &successPayload, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if !strings.Contains(string(bodyBytes), `{"error"`) && strings.HasPrefix(string(bodyBytes), `{`) {
		if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
			return &successPayload, err
		}
	}


	return &successPayload, err
}

/* V2ApiService 
 @param changeLoadBalancerInstanceConfigurationRequest changeLoadBalancerInstanceConfigurationRequest
 @return *ChangeLoadBalancerInstanceConfigurationResponse*/
func (a *V2ApiService) ChangeLoadBalancerInstanceConfiguration(changeLoadBalancerInstanceConfigurationRequest *ChangeLoadBalancerInstanceConfigurationRequest) (*ChangeLoadBalancerInstanceConfigurationResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ChangeLoadBalancerInstanceConfigurationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/changeLoadBalancerInstanceConfiguration"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/x-www-form-urlencoded",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = changeLoadBalancerInstanceConfigurationRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return &successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return &successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	if localVarHttpResponse.StatusCode >= 300 || (localVarHttpResponse.StatusCode < 300 && !strings.HasPrefix(string(bodyBytes), `{`)) {
		return &successPayload, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if !strings.Contains(string(bodyBytes), `{"error"`) && strings.HasPrefix(string(bodyBytes), `{`) {
		if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
			return &successPayload, err
		}
	}


	return &successPayload, err
}

/* V2ApiService 
 @param createLoadBalancerInstanceRequest createLoadBalancerInstanceRequest
 @return *CreateLoadBalancerInstanceResponse*/
func (a *V2ApiService) CreateLoadBalancerInstance(createLoadBalancerInstanceRequest *CreateLoadBalancerInstanceRequest) (*CreateLoadBalancerInstanceResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  CreateLoadBalancerInstanceResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/createLoadBalancerInstance"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/x-www-form-urlencoded",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = createLoadBalancerInstanceRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return &successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return &successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	if localVarHttpResponse.StatusCode >= 300 || (localVarHttpResponse.StatusCode < 300 && !strings.HasPrefix(string(bodyBytes), `{`)) {
		return &successPayload, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if !strings.Contains(string(bodyBytes), `{"error"`) && strings.HasPrefix(string(bodyBytes), `{`) {
		if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
			return &successPayload, err
		}
	}


	return &successPayload, err
}

/* V2ApiService 
 @param deleteLoadBalancerInstancesRequest deleteLoadBalancerInstancesRequest
 @return *DeleteLoadBalancerInstancesResponse*/
func (a *V2ApiService) DeleteLoadBalancerInstances(deleteLoadBalancerInstancesRequest *DeleteLoadBalancerInstancesRequest) (*DeleteLoadBalancerInstancesResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DeleteLoadBalancerInstancesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/deleteLoadBalancerInstances"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/x-www-form-urlencoded",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = deleteLoadBalancerInstancesRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return &successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return &successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	if localVarHttpResponse.StatusCode >= 300 || (localVarHttpResponse.StatusCode < 300 && !strings.HasPrefix(string(bodyBytes), `{`)) {
		return &successPayload, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if !strings.Contains(string(bodyBytes), `{"error"`) && strings.HasPrefix(string(bodyBytes), `{`) {
		if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
			return &successPayload, err
		}
	}


	return &successPayload, err
}

/* V2ApiService 
 @param deleteLoadBalancerSslCertificateRequest deleteLoadBalancerSslCertificateRequest
 @return *DeleteLoadBalancerSslCertificateResponse*/
func (a *V2ApiService) DeleteLoadBalancerSslCertificate(deleteLoadBalancerSslCertificateRequest *DeleteLoadBalancerSslCertificateRequest) (*DeleteLoadBalancerSslCertificateResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DeleteLoadBalancerSslCertificateResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/deleteLoadBalancerSslCertificate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/x-www-form-urlencoded",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = deleteLoadBalancerSslCertificateRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return &successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return &successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	if localVarHttpResponse.StatusCode >= 300 || (localVarHttpResponse.StatusCode < 300 && !strings.HasPrefix(string(bodyBytes), `{`)) {
		return &successPayload, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if !strings.Contains(string(bodyBytes), `{"error"`) && strings.HasPrefix(string(bodyBytes), `{`) {
		if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
			return &successPayload, err
		}
	}


	return &successPayload, err
}

/* V2ApiService 
 @param deleteServerInstancesFromLoadBalancerRequest deleteServerInstancesFromLoadBalancerRequest
 @return *DeleteServerInstancesFromLoadBalancerResponse*/
func (a *V2ApiService) DeleteServerInstancesFromLoadBalancer(deleteServerInstancesFromLoadBalancerRequest *DeleteServerInstancesFromLoadBalancerRequest) (*DeleteServerInstancesFromLoadBalancerResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DeleteServerInstancesFromLoadBalancerResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/deleteServerInstancesFromLoadBalancer"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/x-www-form-urlencoded",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = deleteServerInstancesFromLoadBalancerRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return &successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return &successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	if localVarHttpResponse.StatusCode >= 300 || (localVarHttpResponse.StatusCode < 300 && !strings.HasPrefix(string(bodyBytes), `{`)) {
		return &successPayload, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if !strings.Contains(string(bodyBytes), `{"error"`) && strings.HasPrefix(string(bodyBytes), `{`) {
		if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
			return &successPayload, err
		}
	}


	return &successPayload, err
}

/* V2ApiService 
 @param getLoadBalancedServerInstanceListRequest getLoadBalancedServerInstanceListRequest
 @return *GetLoadBalancedServerInstanceListResponse*/
func (a *V2ApiService) GetLoadBalancedServerInstanceList(getLoadBalancedServerInstanceListRequest *GetLoadBalancedServerInstanceListRequest) (*GetLoadBalancedServerInstanceListResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  GetLoadBalancedServerInstanceListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/getLoadBalancedServerInstanceList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/x-www-form-urlencoded",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = getLoadBalancedServerInstanceListRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return &successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return &successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	if localVarHttpResponse.StatusCode >= 300 || (localVarHttpResponse.StatusCode < 300 && !strings.HasPrefix(string(bodyBytes), `{`)) {
		return &successPayload, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if !strings.Contains(string(bodyBytes), `{"error"`) && strings.HasPrefix(string(bodyBytes), `{`) {
		if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
			return &successPayload, err
		}
	}


	return &successPayload, err
}

/* V2ApiService 
 @param getLoadBalancerInstanceListRequest getLoadBalancerInstanceListRequest
 @return *GetLoadBalancerInstanceListResponse*/
func (a *V2ApiService) GetLoadBalancerInstanceList(getLoadBalancerInstanceListRequest *GetLoadBalancerInstanceListRequest) (*GetLoadBalancerInstanceListResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  GetLoadBalancerInstanceListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/getLoadBalancerInstanceList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/x-www-form-urlencoded",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = getLoadBalancerInstanceListRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return &successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return &successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	if localVarHttpResponse.StatusCode >= 300 || (localVarHttpResponse.StatusCode < 300 && !strings.HasPrefix(string(bodyBytes), `{`)) {
		return &successPayload, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if !strings.Contains(string(bodyBytes), `{"error"`) && strings.HasPrefix(string(bodyBytes), `{`) {
		if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
			return &successPayload, err
		}
	}


	return &successPayload, err
}

/* V2ApiService 
 @param getLoadBalancerSslCertificateListRequest getLoadBalancerSslCertificateListRequest
 @return *GetLoadBalancerSslCertificateListResponse*/
func (a *V2ApiService) GetLoadBalancerSslCertificateList(getLoadBalancerSslCertificateListRequest *GetLoadBalancerSslCertificateListRequest) (*GetLoadBalancerSslCertificateListResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  GetLoadBalancerSslCertificateListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/getLoadBalancerSslCertificateList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/x-www-form-urlencoded",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = getLoadBalancerSslCertificateListRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return &successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return &successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	if localVarHttpResponse.StatusCode >= 300 || (localVarHttpResponse.StatusCode < 300 && !strings.HasPrefix(string(bodyBytes), `{`)) {
		return &successPayload, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if !strings.Contains(string(bodyBytes), `{"error"`) && strings.HasPrefix(string(bodyBytes), `{`) {
		if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
			return &successPayload, err
		}
	}


	return &successPayload, err
}

/* V2ApiService 
 @param getLoadBalancerTargetServerInstanceListRequest getLoadBalancerTargetServerInstanceListRequest
 @return *GetLoadBalancerTargetServerInstanceListResponse*/
func (a *V2ApiService) GetLoadBalancerTargetServerInstanceList(getLoadBalancerTargetServerInstanceListRequest *GetLoadBalancerTargetServerInstanceListRequest) (*GetLoadBalancerTargetServerInstanceListResponse, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  GetLoadBalancerTargetServerInstanceListResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/getLoadBalancerTargetServerInstanceList"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/x-www-form-urlencoded",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = getLoadBalancerTargetServerInstanceListRequest
	v := reflect.ValueOf(localVarPostBody).Elem().FieldByName("UserData")
	if v.IsValid() && v.CanAddr() {
		ptr := v.Addr().Interface().(**string)
		if *ptr != nil {
			**ptr = base64.StdEncoding.EncodeToString([]byte(**ptr))
		}
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return &successPayload, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return &successPayload, err
	}
	defer localVarHttpResponse.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)

	if localVarHttpResponse.StatusCode >= 300 || (localVarHttpResponse.StatusCode < 300 && !strings.HasPrefix(string(bodyBytes), `{`)) {
		return &successPayload, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if !strings.Contains(string(bodyBytes), `{"error"`) && strings.HasPrefix(string(bodyBytes), `{`) {
		if err = json.Unmarshal(bodyBytes[bytes.IndexAny(bytes.Trim(bodyBytes, "{"), "{"):len(bodyBytes)-1], &successPayload); err != nil {
			return &successPayload, err
		}
	}


	return &successPayload, err
}

