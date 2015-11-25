/*
 * buzzlogixtextanalysisapi
 *
 * This file was automatically generated for buzzlogix by APIMATIC BETA v2.0 on 11/25/2015
 */

package keywords


/*
 * Interface for the KEYWORDS_IMPL
 */
type KEYWORDS interface {
    CreateReturnEnglishKeywords (string) (interface{}, error)

    CreateReturnEnglishKeywordsForm (string) (interface{}, error)
}

/*
 * Factory for the KEYWORDS interaface returning KEYWORDS_IMPL
 */
func NewKEYWORDS() KEYWORDS {
    return &KEYWORDS_IMPL{}
}