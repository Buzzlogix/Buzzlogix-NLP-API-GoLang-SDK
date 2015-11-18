/*
 * buzzlogixtextanalysisapi
 *
 * This file was automatically generated for Buzzlogix by APIMATIC BETA v2.0 on 11/18/2015
 */

package keywords


/*
 * Interface for the KEYWORDS_IMPL
 */
type KEYWORDS interface {
    CreateReturnEnglishKeywords (string) (interface{}, error)

    CreateReturnEnglishKeywordsForm (string, string) (interface{}, error)
}

/*
 * Factory for the KEYWORDS interaface returning KEYWORDS_IMPL
 */
func NewKEYWORDS() KEYWORDS {
    return &KEYWORDS_IMPL{}
}