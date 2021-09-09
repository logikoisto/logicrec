// Copyright 2021 logicrec Project Authors
//
// Licensed under the Apache License, Version 2.0 (the "License")
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"github.com/logikoisto/logicrec/base/conf"
	"github.com/zhenghaoz/gorse/base"
)

// FeedServer is feed server
type FeedServer struct {
	conf *conf.FeedConf
}

// NewFeedServer is feed service
func NewFeedServer(cfg *conf.FeedConf) *FeedServer {
	return &FeedServer{conf: cfg}
}

// Run is run
func (f *FeedServer) Run() {
	base.Logger().Info("hello logicrec")
}

// TODO 写算法逻辑啦
