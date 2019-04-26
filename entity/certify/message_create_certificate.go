/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package certify

import (
    "github.com/katena-chain/sdk-go-client/utils"
)

// MsgCreateCertificate is wrapper to indicate that a create certificate action should be applied in a transaction.
// It should implement the Message interface.
type MsgCreateCertificate struct {
    Certificate Certificate `json:"certificate"`
}

// MsgCreateCertificate constructor.
func NewMsgCreateCertificate(certificate Certificate) *MsgCreateCertificate {
    return &MsgCreateCertificate{
        Certificate: certificate,
    }
}

// jsonMsgCreateCertificate wraps a MsgCreateCertificate for its json marshaled value.
type jsonMsgCreateCertificate MsgCreateCertificate

// ToTypedObject indicates which message action and message value should be signed (Message interface requirement).
func (mcc *MsgCreateCertificate) ToTypedObject() *utils.JSONWrapper {
    return &utils.JSONWrapper{
        Type:  mcc.GetType(),
        Value: (*jsonMsgCreateCertificate)(mcc),
    }
}

// GetType returns the type value (Message interface requirement).
func (mcc *MsgCreateCertificate) GetType() string {
    return msgCreateCertificateType
}
