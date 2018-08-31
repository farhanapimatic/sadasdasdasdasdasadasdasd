/*
 * swaggerpetstore_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */
package models_pkg

import(
    "encoding/json"
)

/**
 * Type definition for OAuthScopeEnum enum
 */
type OAuthScopeEnum int

/**
 * Value collection for OAuthScopeEnum enum
 */
const (
    OAuthScope_WRITEPETS OAuthScopeEnum = 1 + iota
    OAuthScope_READPETS
)

func (r OAuthScopeEnum) MarshalJSON() ([]byte, error) { 
    s := OAuthScopeEnumToValue(r)
    return json.Marshal(s) 
} 

func (r *OAuthScopeEnum) UnmarshalJSON(data []byte) error { 
    var s string 
    json.Unmarshal(data, &s)
    v :=  OAuthScopeEnumFromValue(s)
    *r = v 
    return nil 
 } 


/**
 * Converts OAuthScopeEnum to its string representation
 */
func OAuthScopeEnumToValue(oAuthScopeEnum OAuthScopeEnum) string {
    switch oAuthScopeEnum {
        case OAuthScope_WRITEPETS:
    		return "writepets"		
        case OAuthScope_READPETS:
    		return "readpets"		
        default:
        	return "writepets"
    }
}

/**
 * Converts OAuthScopeEnum Array to its string Array representation
*/
func OAuthScopeEnumArrayToValue(oAuthScopeEnum []OAuthScopeEnum) []string {
    convArray := make([]string,len( oAuthScopeEnum))
    for i:=0; i<len(oAuthScopeEnum);i++ {
        convArray[i] = OAuthScopeEnumToValue(oAuthScopeEnum[i])
    }
    return convArray
}


/**
 * Converts given value to its enum representation
 */
func OAuthScopeEnumFromValue(value string) OAuthScopeEnum {
    switch value {
        case "writepets":
            return OAuthScope_WRITEPETS
        case "readpets":
            return OAuthScope_READPETS
        default:
            return OAuthScope_WRITEPETS
    }
}
