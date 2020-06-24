/*
 * transmitsms_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package sms_pkg

import "transmitsms_lib/configuration_pkg"

/*
 * Interface for the SMS_IMPL
 */
type SMS interface {
    SendSMS (string, string, *int64, *string, *string, *int64, *string, *string, *string, *string, *bool, *string, *string) (interface{}, error)

    FormatNumber (string, string, string) (interface{}, error)

    GetSMS (string, string) (interface{}, error)

    GetSMSStats (string, string) (interface{}, error)

    GetSMSResponses (string, *string, *string, *string, *string, *string, *string, *string, *string) (interface{}, error)

    GetUserSMSResponses (string, *string, *string, *string, *string, *string, *string) (interface{}, error)

    GetLinkHits (string, string) (interface{}, error)

    GetSMSSent (string, string, *string, *string, *string, *string) (interface{}, error)

    GetSMSDeliveryStatus (string, string, string) (interface{}, error)

    CancelSMS (string, string) (interface{}, error)
}

/*
 * Factory for the SMS interaface returning SMS_IMPL
 */
func NewSMS(config configuration_pkg.CONFIGURATION) *SMS_IMPL {
    client := new(SMS_IMPL)
    client.config = config
    return client
}