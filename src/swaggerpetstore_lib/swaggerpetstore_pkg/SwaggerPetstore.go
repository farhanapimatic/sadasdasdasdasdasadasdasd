/*
 * swaggerpetstore_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */

package SwaggerPetstoreClient

import(
	"swaggerpetstore_lib/configuration_pkg"
	"swaggerpetstore_lib/pet_pkg"
	"swaggerpetstore_lib/store_pkg"
	"swaggerpetstore_lib/user_pkg"
)


/*
 * Interface for the SWAGGERPETSTORE_IMPL
 */
type SWAGGERPETSTORE interface {
     Pet()                   pet_pkg.PET
     Store()                 store_pkg.STORE
     User()                  user_pkg.USER
     Configuration()         configuration_pkg.CONFIGURATION
}

/*
 * Factory for the SWAGGERPETSTORE interaface returning SWAGGERPETSTORE_IMPL
 */
func NewSWAGGERPETSTORE() SWAGGERPETSTORE {
    swaggerPetstoreClient := new(SWAGGERPETSTORE_IMPL)
    swaggerPetstoreClient.config = configuration_pkg.NewCONFIGURATION()
    return swaggerPetstoreClient
}
