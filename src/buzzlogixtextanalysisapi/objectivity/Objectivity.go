/*
 * buzzlogixtextanalysisapi
 *
 * This file was automatically generated for Buzzlogix by APIMATIC BETA v2.0 on 11/09/2015
 */

package objectivity


/*
 * Interface for the OBJECTIVITY_IMPL
 */
type OBJECTIVITY interface {
    CreateReturnObjectivity (string) (error)
}

/*
 * Factory for the OBJECTIVITY interaface returning OBJECTIVITY_IMPL
 */
func NewOBJECTIVITY() OBJECTIVITY {
    return &OBJECTIVITY_IMPL{}
}