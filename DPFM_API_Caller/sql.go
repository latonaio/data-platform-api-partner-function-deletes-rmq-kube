package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-partner-function-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-partner-function-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) PartnerFunction(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.PartnerFunction {

	where := strings.Join([]string{
		fmt.Sprintf("WHERE partnerFunction.PartnerFunction = \"%s\ ", input.PartnerFunction.PartnerFunction),
	}, "")

	rows, err := c.db.Query(
		`SELECT 
    	partnerFunction.PartnerFunction
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_partnerFunction_partnerFunction_data as partnerFunction 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPartnerFunction(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
