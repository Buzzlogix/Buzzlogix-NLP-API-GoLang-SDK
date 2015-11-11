/*
 * buzzlogixtextanalysisapi
 *
 * This file was automatically generated for Buzzlogix by APIMATIC BETA v2.0 on 11/09/2015
 */
package objectivity

import(
    "github.com/apimatic/unirest-go"
    "buzzlogixtextanalysisapi"
    "buzzlogixtextanalysisapi/apihelper"

)
/*
 * Client structure as interface implementation
 */
type OBJECTIVITY_IMPL struct { }

/**
 * Use this endpoint to retrieve whether the provided text is subjective or objective.
 * @param    string        body     parameter: Required
 * @return	Returns the  response from the API call
 */
func (me *OBJECTIVITY_IMPL) CreateReturnObjectivity (
            body string) (error) {
    //the base uri for api requests
    queryBuilder := buzzlogixtextanalysisapi.BASEURI;
        
    //prepare query string for API call
    queryBuilder = queryBuilder + "/objectivity"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    queryBuilder, err = apihelper.CleanUrl(queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //prepare API request
    request := unirest.Post(queryBuilder, headers, body)

    //Custom Authentication to be added by the developers for authorization
    apihelper.AppendCustomAuthParams(request)
    //and invoke the API call request to fetch the response
    response, err := unirest.AsString(request);
    if err != nil {
        //error in API invocation
        return err
    }

    //error handling using HTTP status codes
    if (response.Code == 401) {
        err = apihelper.NewAPIError("No API Key found in headers, body or querystring", response.Code, response.RawBody)
    } else if (response.Code == 403) {
        err = apihelper.NewAPIError("Invalid authentication credentials", response.Code, response.RawBody)
    } else if (response.Code == 429) {
        err = apihelper.NewAPIError("API rate limit exceeded", response.Code, response.RawBody)
    } else if (response.Code < 200) || (response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", response.Code, response.RawBody)
        }
    if(err != nil) {
        //error detected in status code validation
        return err
    }
    
    //returning the response
    return nil
}

