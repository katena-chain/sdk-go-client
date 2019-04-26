/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package utils

import (
    "net/url"
    "path"
)

// GetUri joins the base path and paths array and adds the query values to return a new url.
func GetUri(basePath string, paths []string, queryValues map[string]string) (*url.URL, error) {
    uri, err := url.Parse(basePath)
    if err != nil {
        return nil, err
    }
    uri.Path = path.Join(append([]string{uri.Path}, paths...)...)
    uriQuery := uri.Query()
    for index, value := range queryValues {
        uriQuery.Set(index, value)
    }
    uri.RawQuery = uriQuery.Encode()
    return uri, nil
}
