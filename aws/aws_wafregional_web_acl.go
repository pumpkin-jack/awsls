// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
)

func ListWafregionalWebAcl(client *Client) ([]Resource, error) {
	req := client.wafregionalconn.ListWebACLsRequest(&wafregional.ListWebACLsInput{})

	var result []Resource

	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	if len(resp.WebACLs) > 0 {
		for _, r := range resp.WebACLs {

			result = append(result, Resource{
				Type:   "aws_wafregional_web_acl",
				ID:     *r.WebACLId,
				Region: client.wafregionalconn.Config.Region,
			})
		}
	}

	return result, nil
}
