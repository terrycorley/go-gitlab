// Copyright 2014 The go-gitlab AUTHORS. All rights reserved.
// Copyright 2014 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gitlab

// import (
// 	"bytes"
// )

// // MarkdownOptions specifies optional parameters to the Markdown method.
// type MarkdownOptions struct {
// 	// Mode identifies the rendering mode.  Possible values are:
// 	//   markdown - render a document as plain Markdown, just like
// 	//   README files are rendered.
// 	//
// 	//   gfm - to render a document as user-content, e.g. like user
// 	//   comments or issues are rendered. In GFM mode, hard line breaks are
// 	//   always taken into account, and issue and user mentions are linked
// 	//   accordingly.
// 	//
// 	// Default is "markdown".
// 	Mode string

// 	// Context identifies the repository context.  Only taken into account
// 	// when rendering as "gfm".
// 	Context string
// }

// type markdownRequest struct {
// 	Text    *string `json:"text,omitempty"`
// 	Mode    *string `json:"mode,omitempty"`
// 	Context *string `json:"context,omitempty"`
// }

// // Markdown renders an arbitrary Markdown document.
// //
// // GitHub API docs: https://developer.github.com/v3/markdown/
// func (c *Client) Markdown(text string, opt *MarkdownOptions) (string, *Response, error) {
// 	request := &markdownRequest{Text: String(text)}
// 	if opt != nil {
// 		if opt.Mode != "" {
// 			request.Mode = String(opt.Mode)
// 		}
// 		if opt.Context != "" {
// 			request.Context = String(opt.Context)
// 		}
// 	}

// 	req, err := c.NewRequest("POST", "/markdown", request)
// 	if err != nil {
// 		return "", nil, err
// 	}

// 	buf := new(bytes.Buffer)
// 	resp, err := c.Do(req, buf)
// 	if err != nil {
// 		return "", resp, err
// 	}

// 	return buf.String(), resp, nil
// }

// // ListEmojis returns the emojis available to use on GitHub.
// //
// // GitHub API docs: https://developer.github.com/v3/emojis/
// func (c *Client) ListEmojis() (map[string]string, *Response, error) {
// 	req, err := c.NewRequest("GET", "/emojis", nil)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	var emoji map[string]string
// 	resp, err := c.Do(req, &emoji)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return emoji, resp, nil
// }

// // APIMeta represents metadata about the GitHub API.
// type APIMeta struct {
// 	// An Array of IP addresses in CIDR format specifying the addresses
// 	// that incoming service hooks will originate from on GitHub.com.
// 	Hooks []string `json:"hooks,omitempty"`

// 	// An Array of IP addresses in CIDR format specifying the Git servers
// 	// for GitHub.com.
// 	Git []string `json:"git,omitempty"`

// 	// Whether authentication with username and password is supported.
// 	// (GitHub Enterprise instances using CAS or OAuth for authentication
// 	// will return false. Features like Basic Authentication with a
// 	// username and password, sudo mode, and two-factor authentication are
// 	// not supported on these servers.)
// 	VerifiablePasswordAuthentication *bool `json:"verifiable_password_authentication,omitempty"`
// }

// // APIMeta returns information about GitHub.com, the service. Or, if you access
// // this endpoint on your organization’s GitHub Enterprise installation, this
// // endpoint provides information about that installation.
// //
// // GitHub API docs: https://developer.github.com/v3/meta/
// func (c *Client) APIMeta() (*APIMeta, *Response, error) {
// 	req, err := c.NewRequest("GET", "/meta", nil)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	meta := new(APIMeta)
// 	resp, err := c.Do(req, meta)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return meta, resp, nil
// }