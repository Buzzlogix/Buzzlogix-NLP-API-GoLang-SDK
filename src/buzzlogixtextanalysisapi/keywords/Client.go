/*
 * buzzlogixtextanalysisapi
 *
 * This file was automatically generated for buzzlogix by APIMATIC BETA v2.0 on 12/06/2015
 */
package keywords

import(
    "github.com/apimatic/unirest-go"
    "buzzlogixtextanalysisapi"
    "buzzlogixtextanalysisapi/apihelper"

)
/*
 * Client structure as interface implementation
 */
type KEYWORDS_IMPL struct { }

/**
 * The text should be provided as text/plain in the body
 * @param    string        body     parameter: Required
 * @return	Returns the interface{} response from the API call
 */
func (me *KEYWORDS_IMPL) CreateReturnEnglishKeywordsTextPlain (
            body string) (interface{}, error) {
    //the base uri for api requests
    queryBuilder := buzzlogixtextanalysisapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/keywords"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "X-Mashape-Key" : buzzlogixtextanalysisapi.Config.XMashapeKey,
    }

    //prepare API request
    request := unirest.Post(queryBuilder, headers, body)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper.NewAPIError("HTTP Response Not OK" , response.Code, response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }
    
    //returning the response
    var retVal interface{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * The text should be provided as multipart/form-data with the key 'text'. Files can be uploaded.
 * @param    []byte        text     parameter: Required
 * @return	Returns the interface{} response from the API call
 */
func (me *KEYWORDS_IMPL) CreateReturnEnglishKeywordsMultipartFormData (
            text []byte) (interface{}, error) {
    //the base uri for api requests
    queryBuilder := buzzlogixtextanalysisapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/keywords"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "X-Mashape-Key" : buzzlogixtextanalysisapi.Config.XMashapeKey,
    }

    //form parameters
    parameters := map[string]interface{} {

        "text" : text,

    }


    //prepare API request
    request := unirest.Post(queryBuilder, headers, parameters)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper.NewAPIError("HTTP Response Not OK" , response.Code, response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }
    
    //returning the response
    var retVal interface{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Supply the text being classified using key 'text' in a form. 
 * @param    string        text     parameter: Required
 * @return	Returns the interface{} response from the API call
 */
func (me *KEYWORDS_IMPL) CreateReturnEnglishKeywordsXWwwFormUrlencoded (
            text string) (interface{}, error) {
    //the base uri for api requests
    queryBuilder := buzzlogixtextanalysisapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/keywords"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "X-Mashape-Key" : buzzlogixtextanalysisapi.Config.XMashapeKey,
    }

    //form parameters
    parameters := map[string]interface{} {

        "text" : text,

    }


    //prepare API request
    request := unirest.Post(queryBuilder, headers, parameters)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper.NewAPIError("HTTP Response Not OK" , response.Code, response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }
    
    //returning the response
    var retVal interface{}
    err = json.Unmarshal(response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

