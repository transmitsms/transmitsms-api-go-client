/*
 * transmitsms_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package lists_pkg

import "transmitsms_lib/configuration_pkg"

/*
 * Interface for the LISTS_IMPL
 */
type LISTS interface {
    AddToList (string, string, string, *string, *string, *string) (interface{}, error)

    AddFieldToList (string, string, string, *string, map[string]interface{}) (interface{}, error)

    AddList (string, string, *string, map[string]interface{}) (interface{}, error)

    OptoutListMember (string, string, string) (interface{}, error)

    DeleteFromList (int64, string, string) (interface{}, error)

    GetList (string, string, *string, *string, *string) (interface{}, error)

    GetLists (string, *string, *string) (interface{}, error)

    EditListMember (string, string, string, *string, *string, *string, map[string]interface{}) (interface{}, error)

    GetContact (string, string, string) (interface{}, error)

    AddContactsBulk (string, string, string, *string, *string, map[string]interface{}) (interface{}, error)

    RemoveList (string, string) (interface{}, error)
}

/*
 * Factory for the LISTS interaface returning LISTS_IMPL
 */
func NewLISTS(config configuration_pkg.CONFIGURATION) *LISTS_IMPL {
    client := new(LISTS_IMPL)
    client.config = config
    return client
}