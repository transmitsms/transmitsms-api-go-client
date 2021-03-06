/*
 * transmitsms_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package keywords_pkg


import(
	"encoding/json"
	"github.com/apimatic/unirest-go"
	"transmitsms_lib/apihelper_pkg"
	"transmitsms_lib/configuration_pkg"
)
/*
 * Client structure as interface implementation
 */
type KEYWORDS_IMPL struct {
     config configuration_pkg.CONFIGURATION
}

/**
 * Get a list of existing keywords.
 * @param    string         format     parameter: Required
 * @param    *string        number     parameter: Optional
 * @param    *string        page       parameter: Optional
 * @param    *string        max        parameter: Optional
 * @return	Returns the interface{} response from the API call
 */
func (me *KEYWORDS_IMPL) GetKeywords (
            format string,
            number *string,
            page *string,
            max *string) (interface{}, error) {
    //the endpoint path uri
    _pathUrl := "/get-keywords.{format}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "format" : format,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "number" : number,
        "page" : page,
        "max" : max,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
    }

    //prepare API request
    _request := unirest.GetWithAuth(_queryBuilder, headers, me.config.Username(), me.config.Password())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal interface{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Add a keyword to an existing virtual number.
 * @param    string         keyword             parameter: Required
 * @param    string         number              parameter: Required
 * @param    string         format              parameter: Required
 * @param    *string        reference           parameter: Optional
 * @param    *string        listId              parameter: Optional
 * @param    *string        welcomeMessage      parameter: Optional
 * @param    *string        membersMessage      parameter: Optional
 * @param    *string        activate            parameter: Optional
 * @param    *string        forwardUrl          parameter: Optional
 * @param    *string        forwardEmail        parameter: Optional
 * @param    *string        forwardSms          parameter: Optional
 * @return	Returns the interface{} response from the API call
 */
func (me *KEYWORDS_IMPL) AddKeyword (
            keyword string,
            number string,
            format string,
            reference *string,
            listId *string,
            welcomeMessage *string,
            membersMessage *string,
            activate *string,
            forwardUrl *string,
            forwardEmail *string,
            forwardSms *string) (interface{}, error) {
    //the endpoint path uri
    _pathUrl := "/add-keyword.{format}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "format" : format,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "keyword" : keyword,
        "number" : number,
        "reference" : reference,
        "list_id" : listId,
        "welcome_message" : welcomeMessage,
        "members_message" : membersMessage,
        "activate" : activate,
        "forward_url" : forwardUrl,
        "forward_email" : forwardEmail,
        "forward_sms" : forwardSms,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
    }

    //prepare API request
    _request := unirest.GetWithAuth(_queryBuilder, headers, me.config.Username(), me.config.Password())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal interface{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Edit an existing keyword.
 * @param    string         keyword             parameter: Required
 * @param    string         number              parameter: Required
 * @param    string         format              parameter: Required
 * @param    *string        reference           parameter: Optional
 * @param    *string        status              parameter: Optional
 * @param    *string        listId              parameter: Optional
 * @param    *string        welcomeMessage      parameter: Optional
 * @param    *string        membersMessage      parameter: Optional
 * @param    *string        activate            parameter: Optional
 * @param    *string        forwardUrl          parameter: Optional
 * @param    *string        forwardEmail        parameter: Optional
 * @param    *string        forwardSms          parameter: Optional
 * @return	Returns the interface{} response from the API call
 */
func (me *KEYWORDS_IMPL) EditKeyword (
            keyword string,
            number string,
            format string,
            reference *string,
            status *string,
            listId *string,
            welcomeMessage *string,
            membersMessage *string,
            activate *string,
            forwardUrl *string,
            forwardEmail *string,
            forwardSms *string) (interface{}, error) {
    //the endpoint path uri
    _pathUrl := "/edit-keyword.{format}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper_pkg.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "format" : format,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration_pkg.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "keyword" : keyword,
        "number" : number,
        "reference" : reference,
        "status" : status,
        "list_id" : listId,
        "welcome_message" : welcomeMessage,
        "members_message" : membersMessage,
        "activate" : activate,
        "forward_url" : forwardUrl,
        "forward_email" : forwardEmail,
        "forward_sms" : forwardSms,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
    }

    //prepare API request
    _request := unirest.GetWithAuth(_queryBuilder, headers, me.config.Username(), me.config.Password())
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,false);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal interface{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

