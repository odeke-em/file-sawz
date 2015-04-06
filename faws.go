//   Copyright 2015 Emmanuel Odeke. All Rights Reserved.
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package faws

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/awslabs/aws-sdk-go/aws"
)

const (
	AWSAccessId  = "AWSAccessKeyId"
	AWSSecretKey = "AWSSecretKey"
)

type credentials struct {
	secretKey string `json:"aws_secret"`
	// accessId: Your public key
	accessId string `json:"aws_access_id"`
}

func FileCredentials(p string) (*credentials, error) {
	var data []byte
	var err error

	if data, err = ioutil.ReadFile(p); err != nil {
		return nil, err
	}

	var creds credentials
	err = json.Unmarshal(data, &creds)

	return &creds, err
}

func (cred *Credentials) Serialize(outpath string) error {
	var data []byte
	if data, err = json.Marshall(cred); err != nil {
		return err
	}
	return ioutil.WriteFile(p, data, 0600)
}

func EnvCredentials() *aws.Credentials {
	return &aws.Credentials{
		SecretAccessKey: os.Getenv(AWSSecretKey),
		AccessKeyID:     os.Getenv(AWSAccessId),
	}
}
