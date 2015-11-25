/*
 * buzzlogixtextanalysisapi
 *
 * This file was automatically generated for buzzlogix by APIMATIC BETA v2.0 on 11/25/2015
 */

package objectivity


/*
 * Interface for the OBJECTIVITY_IMPL
 */
type OBJECTIVITY interface {
    CreateReturnEnglishObjectivity (string) (interface{}, error)

    CreateReturnEnglishObjectivityForm (string) (interface{}, error)
}

/*
 * Factory for the OBJECTIVITY interaface returning OBJECTIVITY_IMPL
 */
func NewOBJECTIVITY() OBJECTIVITY {
    return &OBJECTIVITY_IMPL{}
}