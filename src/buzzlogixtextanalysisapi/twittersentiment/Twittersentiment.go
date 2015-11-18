/*
 * buzzlogixtextanalysisapi
 *
 * This file was automatically generated for Buzzlogix by APIMATIC BETA v2.0 on 11/18/2015
 */

package twittersentiment


/*
 * Interface for the TWITTERSENTIMENT_IMPL
 */
type TWITTERSENTIMENT interface {
    CreateReturnEnglishTwitterSentiment (string) (interface{}, error)

    CreateReturnEnglishTwitterSentimentForm (string) (interface{}, error)
}

/*
 * Factory for the TWITTERSENTIMENT interaface returning TWITTERSENTIMENT_IMPL
 */
func NewTWITTERSENTIMENT() TWITTERSENTIMENT {
    return &TWITTERSENTIMENT_IMPL{}
}