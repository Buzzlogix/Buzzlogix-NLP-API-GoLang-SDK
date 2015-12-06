/*
 * buzzlogixtextanalysisapi
 *
 * This file was automatically generated for buzzlogix by APIMATIC BETA v2.0 on 12/06/2015
 */

package objectivity


/*
 * Interface for the OBJECTIVITY_IMPL
 */
type OBJECTIVITY interface {
    CreateReturnEnglishObjectivityPlaintext (string) (interface{}, error)

    CreateReturnEnglishObjectivityMultipartForm ([]byte) (interface{}, error)

    CreateReturnEnglishObjectivityEncodedForm (string) (interface{}, error)
}

/*
 * Factory for the OBJECTIVITY interaface returning OBJECTIVITY_IMPL
 */
func NewOBJECTIVITY() OBJECTIVITY {
    return &OBJECTIVITY_IMPL{}
}