/*
 * Minio Cloud Storage, (C) 2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package metadata

import (
	"testing"

	. "github.com/minio/minio/internal/gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestCache(c *C) {
	cache := NewCache()
	data := []byte("Hello, world!")
	ok := cache.Set("filename", data)

	c.Assert(ok, Equals, true)
	storedata := cache.Get("filename")

	c.Assert(ok, Equals, true)
	c.Assert(data, DeepEquals, storedata)

	cache.Delete("filename")

	ok = cache.Exists("filename")
	c.Assert(ok, Equals, false)
}
