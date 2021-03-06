/*******************************************************************************
 * Copyright © 2017-2018 VMware, Inc. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 * @author: Huaqiao Zhang, <huaqiaoz@vmware.com>
 *******************************************************************************/

package domain

import (
	_ "encoding/json"
	"gopkg.in/mgo.v2/bson"
)

type Gateway struct {
	Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Address     string        `json:"address"`
	Created     int64         `bson:"created" json:"created"`
	Modified    int64         `bson:"modified" json:"modified"`
}
