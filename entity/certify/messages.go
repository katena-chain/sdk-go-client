/**
 * Copyright (c) 2018, TransChain.
 *
 * This source code is licensed under the Apache 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package certify

import (
    "encoding/json"
    "errors"
    "fmt"

    "github.com/katena-chain/sdk-go-client/utils"
)

var certificateFactories = map[string]func() Certificate{
    CertificateV1Type: func() Certificate { return &CertificateV1{} },
}

var secretFactories = map[string]func() Secret{
    SecretV1Type: func() Secret { return &SecretV1{} },
}

// MsgCreateCertificate is a wrapper to indicate that a create certificate action should be applied in a transaction.
// It should implement the Message interface.
type MsgCreateCertificate struct {
    Certificate Certificate `json:"certificate"`
}

// GetType returns the type value (Message interface requirement).
func (mcc MsgCreateCertificate) GetType() string {
    return MsgCreateCertificateType
}

// MarshalJSON converts a MsgCreateCertificate into a jsonMsgCreateCertificate and wraps
// it in a JSONWrapper to indicate which message action and message value should be
// unmarshaled back.
func (mcc MsgCreateCertificate) MarshalJSON() ([]byte, error) {
    type jsonAlias MsgCreateCertificate
    value, _ := json.Marshal((*jsonAlias)(&mcc))
    return json.Marshal(&utils.JSONWrapper{
        Type:  mcc.GetType(),
        Value: value,
    })
}

// UnmarshalJSON converts an already unwrapped json value to a MsgCreateCertificate.
// It handles the Certificate interface conversion as well.
func (mcc *MsgCreateCertificate) UnmarshalJSON(data []byte) error {
    var jsonMsg struct {
        Certificate *utils.JSONWrapper `json:"certificate"`
    }
    if err := json.Unmarshal(data, &jsonMsg); err != nil {
        return err
    }
    if certificateFactory, ok := certificateFactories[jsonMsg.Certificate.Type]; ok {
        certificate := certificateFactory()
        if err := json.Unmarshal(jsonMsg.Certificate.Value, certificate); err != nil {
            return err
        }
        mcc.Certificate = certificate
    } else {
        return errors.New(fmt.Sprintf("unknown certificate type: %s", jsonMsg.Certificate.Type))
    }
    return nil
}

// MsgCreateSecret is a wrapper to indicate that a create secret action should be applied in a transaction.
// It should implement the Message interface.
type MsgCreateSecret struct {
    Secret Secret `json:"secret"`
}

// GetType returns the type value (Message interface requirement).
func (mcc MsgCreateSecret) GetType() string {
    return MsgCreateSecretType
}

// MarshalJSON converts a MsgCreateSecret into a jsonMsgCreateSecret and wraps
// it in a JSONWrapper to indicate which message action and message value should be
// unmarshaled back.
func (mcc MsgCreateSecret) MarshalJSON() ([]byte, error) {
    type jsonAlias MsgCreateSecret
    value, _ := json.Marshal((*jsonAlias)(&mcc))
    return json.Marshal(&utils.JSONWrapper{
        Type:  mcc.GetType(),
        Value: value,
    })
}

// UnmarshalJSON converts an already unwrapped json value to a MsgCreateSecret.
// It handles the Secret interface conversion as well.
func (mcc *MsgCreateSecret) UnmarshalJSON(data []byte) error {
    var jsonMsg struct {
        Secret *utils.JSONWrapper `json:"secret"`
    }
    if err := json.Unmarshal(data, &jsonMsg); err != nil {
        return err
    }
    if secretFactory, ok := secretFactories[jsonMsg.Secret.Type]; ok {
        secret := secretFactory()
        if err := json.Unmarshal(jsonMsg.Secret.Value, secret); err != nil {
            return err
        }
        mcc.Secret = secret
    } else {
        return errors.New(fmt.Sprintf("unknown secret type: %s", jsonMsg.Secret.Type))
    }
    return nil
}
