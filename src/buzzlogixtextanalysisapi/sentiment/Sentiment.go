/*
 * buzzlogixtextanalysisapi
 *
 * This file was automatically generated for buzzlogix by APIMATIC BETA v2.0 on 12/06/2015
 */

package sentiment


/*
 * Interface for the SENTIMENT_IMPL
 */
type SENTIMENT interface {
    CreateReturnEnglishGeneralSentimentPlaintext (string) (interface{}, error)

    CreateReturnEnglishGeneralSentimentMultipartForm ([]byte) (interface{}, error)

    CreateReturnEnglishGeneralSentimentEncodedForm (string) (interface{}, error)
}

/*
 * Factory for the SENTIMENT interaface returning SENTIMENT_IMPL
 */
func NewSENTIMENT() SENTIMENT {
    return &SENTIMENT_IMPL{}
}