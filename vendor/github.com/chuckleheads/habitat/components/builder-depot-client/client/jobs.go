package client

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/chuckleheads/habitat/components/builder-depot-client/types"
	pkg "github.com/chuckleheads/habitat/components/core/ident"
)

// ScheduleJob schedules a new job with builder for a given package
func (c *Client) ScheduleJob(ident pkg.Ident, group bool) error {
	res, err := c.req.SetPathParams(map[string]string{
		"origin": ident.Origin,
		"name":   ident.Name,
	}).
		SetQueryParam("package_only", fmt.Sprintf("%t", group)).
		Post("depot/pkgs/schedule/{origin}/{name}")

	if err != nil {
		return err
	}
	if res.StatusCode() != 200 {
		return errors.New(res.Status())
	}
	return nil
}

func (c *Client) FetchRdeps(ident pkg.Ident) ([]string, error) {
	res, err := c.req.SetPathParams(map[string]string{
		"ident": ident.String(),
	}).Get("rdeps/{ident}")

	if err != nil {
		return nil, err
	}
	var rdeps types.RDepsResp
	err = json.Unmarshal(res.Body(), &rdeps)
	if err != nil {
		return nil, err
	}
	return rdeps.RDeps, nil
}
