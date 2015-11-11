/*
 * buzzlogixtextanalysisapi
 *
 * This file was automatically generated for Buzzlogix by APIMATIC BETA v2.0 on 11/09/2015
 */

package sentiment


/*
 * Interface for the SENTIMENT_IMPL
 */
type SENTIMENT interface {
    CreateReturnGeneralSentiment (string) (error)
}

/*
 * Factory for the SENTIMENT interaface returning SENTIMENT_IMPL
 */
func NewSENTIMENT() SENTIMENT {
    return &SENTIMENT_IMPL{}
}