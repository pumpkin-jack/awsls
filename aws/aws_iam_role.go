// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func ListIamRole(client *Client) ([]Resource, error) {
	req := client.iamconn.ListRolesRequest(&iam.ListRolesInput{})

	var result []Resource

	p := iam.NewListRolesPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Roles {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t := *r.CreateDate
			result = append(result, Resource{
				Type:      "aws_iam_role",
				ID:        *r.RoleName,
				Region:    client.iamconn.Config.Region,
				Tags:      tags,
				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
