/*
 * transmitsms_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package TransmitSMSClient

import(
	"transmitsms_lib/configuration_pkg"
	"transmitsms_lib/sms_pkg"
	"transmitsms_lib/numbers_pkg"
	"transmitsms_lib/emailsms_pkg"
	"transmitsms_lib/resellers_pkg"
	"transmitsms_lib/account_pkg"
	"transmitsms_lib/keywords_pkg"
	"transmitsms_lib/lists_pkg"
)
/*
 * Client structure as interface implementation
 */
type TRANSMITSMS_IMPL struct {
     sms sms_pkg.SMS
     numbers numbers_pkg.NUMBERS
     emailsms emailsms_pkg.EMAILSMS
     resellers resellers_pkg.RESELLERS
     account account_pkg.ACCOUNT
     keywords keywords_pkg.KEYWORDS
     lists lists_pkg.LISTS
     config  configuration_pkg.CONFIGURATION
}

/**
     * Access to Configuration
     * @return Returns the Configuration instance
*/
func (me *TRANSMITSMS_IMPL) Configuration() configuration_pkg.CONFIGURATION {
    return me.config
}
/**
     * Access to SMS controller
     * @return Returns the SMS() instance
*/
func (me *TRANSMITSMS_IMPL) SMS() sms_pkg.SMS {
    if(me.sms) == nil {
        me.sms = sms_pkg.NewSMS(me.config)
    }
    return me.sms
}
/**
     * Access to Numbers controller
     * @return Returns the Numbers() instance
*/
func (me *TRANSMITSMS_IMPL) Numbers() numbers_pkg.NUMBERS {
    if(me.numbers) == nil {
        me.numbers = numbers_pkg.NewNUMBERS(me.config)
    }
    return me.numbers
}
/**
     * Access to EmailSMS controller
     * @return Returns the EmailSMS() instance
*/
func (me *TRANSMITSMS_IMPL) EmailSMS() emailsms_pkg.EMAILSMS {
    if(me.emailsms) == nil {
        me.emailsms = emailsms_pkg.NewEMAILSMS(me.config)
    }
    return me.emailsms
}
/**
     * Access to Resellers controller
     * @return Returns the Resellers() instance
*/
func (me *TRANSMITSMS_IMPL) Resellers() resellers_pkg.RESELLERS {
    if(me.resellers) == nil {
        me.resellers = resellers_pkg.NewRESELLERS(me.config)
    }
    return me.resellers
}
/**
     * Access to Account controller
     * @return Returns the Account() instance
*/
func (me *TRANSMITSMS_IMPL) Account() account_pkg.ACCOUNT {
    if(me.account) == nil {
        me.account = account_pkg.NewACCOUNT(me.config)
    }
    return me.account
}
/**
     * Access to Keywords controller
     * @return Returns the Keywords() instance
*/
func (me *TRANSMITSMS_IMPL) Keywords() keywords_pkg.KEYWORDS {
    if(me.keywords) == nil {
        me.keywords = keywords_pkg.NewKEYWORDS(me.config)
    }
    return me.keywords
}
/**
     * Access to Lists controller
     * @return Returns the Lists() instance
*/
func (me *TRANSMITSMS_IMPL) Lists() lists_pkg.LISTS {
    if(me.lists) == nil {
        me.lists = lists_pkg.NewLISTS(me.config)
    }
    return me.lists
}

