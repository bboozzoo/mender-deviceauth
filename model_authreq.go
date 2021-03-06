// Copyright 2016 Mender Software AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package main

import (
	"errors"
	"time"
)

// note: fields with underscores need the 'bson' decorator
// otherwise the underscore will be removed upon write to mongo
type AuthReq struct {
	IdData      string    `json:"id_data" bson:"id_data"`
	TenantToken string    `json:"tenant_token" bson:"tenant_token"`
	PubKey      string    `json:"pubkey"`
	DeviceId    string    `json:"device_id" bson:"device_id"`
	Timestamp   time.Time `json:"ts" bson:"ts"`
	Status      string    `json:"status"`
	SeqNo       uint64    `json:"seq_no" bson:"seq_no"`
}

func (r *AuthReq) Validate() error {
	if r.IdData == "" {
		return errors.New("id_data must be provided")
	}

	if r.PubKey == "" {
		return errors.New("pubkey must be provided")
	}

	if r.SeqNo == 0 {
		return errors.New("seq_no must be provided")
	}

	// not checking tenant token for now - TODO
	return nil
}
