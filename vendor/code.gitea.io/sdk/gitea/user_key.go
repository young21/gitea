// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gitea

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// PublicKey publickey is a user key to push code to repository
type PublicKey struct {
	ID      int64     `json:"id"`
	Key     string    `json:"key"`
	URL     string    `json:"url,omitempty"`
	Title   string    `json:"title,omitempty"`
	// swagger:strfmt date-time
	Created time.Time `json:"created_at,omitempty"`
}

// ListPublicKeys list all the public keys of the user
func (c *Client) ListPublicKeys(user string) ([]*PublicKey, error) {
	keys := make([]*PublicKey, 0, 10)
	return keys, c.getParsedResponse("GET", fmt.Sprintf("/users/%s/keys", user), nil, nil, &keys)
}

// ListMyPublicKeys list all the public keys of current user
func (c *Client) ListMyPublicKeys() ([]*PublicKey, error) {
	keys := make([]*PublicKey, 0, 10)
	return keys, c.getParsedResponse("GET", "/user/keys", nil, nil, &keys)
}

// GetPublicKey get current user's public key by key id
func (c *Client) GetPublicKey(keyID int64) (*PublicKey, error) {
	key := new(PublicKey)
	return key, c.getParsedResponse("GET", fmt.Sprintf("/user/keys/%d", keyID), nil, nil, &key)
}

// CreatePublicKey create public key with options
func (c *Client) CreatePublicKey(opt CreateKeyOption) (*PublicKey, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	key := new(PublicKey)
	return key, c.getParsedResponse("POST", "/user/keys", jsonHeader, bytes.NewReader(body), key)
}

// DeletePublicKey delete public key with key id
func (c *Client) DeletePublicKey(keyID int64) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("/user/keys/%d", keyID), nil, nil)
	return err
}
