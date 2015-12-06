/*
 * buzzlogixtextanalysisapi
 *
 * This file was automatically generated for buzzlogix by APIMATIC BETA v2.0 on 12/06/2015
 */

package twittersentiment


/*
 * Interface for the TWITTERSENTIMENT_IMPL
 */
type TWITTERSENTIMENT interface {
    CreateReturnEnglishTwitterSentimentPlaintext (string) (interface{}, error)

    CreateReturnEnglishTwitterSentimentMultipartForm ([]byte) (interface{}, error)

    CreateReturnEnglishTwitterSentimentEncodedForm (string) (interface{}, error)
}

/*
 * Factory for the TWITTERSENTIMENT interaface returning TWITTERSENTIMENT_IMPL
 */
func NewTWITTERSENTIMENT() TWITTERSENTIMENT {
    return &TWITTERSENTIMENT_IMPL{}
}