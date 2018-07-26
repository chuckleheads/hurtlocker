package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/chuckleheads/habitat/components/builder-depot-client/types"
	pkg "github.com/chuckleheads/habitat/components/core/ident"
)

func (c *Client) SearchPackage(searchTerm string) []pkg.Ident {
	encodedTerm := url.QueryEscape(searchTerm)
	test, err := c.req.SetPathParams(map[string]string{
		"term": encodedTerm,
	}).Get("depot/pkgs/search/{term}")

	if err != nil {
		fmt.Println(err)
	}
	var res types.PagedSearchResp
	err = json.Unmarshal(test.Body(), &res)
	if err != nil {
		fmt.Println(err)
	}
	return res.Data
}
