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

package feed

import (
	"context"

	"github.com/logikoisto/logicrec/stub/feed"
)

var feedServer = &ServerMockImpl{}

// ServerMockImpl impl
type ServerMockImpl struct {
}

// GetFeed 具体实现
func (*ServerMockImpl) GetFeed(context.Context, *feed.FeedRequest) (*feed.FeedResponse, error) {
	return nil, nil
}

// TODO
// 1  实现一个 feed的rpc
// 2  写一个单元测试
// 3  完善log 和 err 处理机制
