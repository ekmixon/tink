// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

package testutil

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	tinkpb "github.com/google/tink/go/proto/tink_go_proto"
)

func tinkRootPath() (string, error) {
	root := os.Getenv("TEST_SRCDIR")
	if root == "" {
		return "", errors.New("TEST_SRCDIR not found")
	}
	path := root + "/google3/third_party/tink"
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return path, nil
	}
	path2 := root + "/tink_base"
	if _, err := os.Stat(path2); !os.IsNotExist(err) {
		return path2, nil
	}
	return "", errors.New("Tink root path not found")
}

// KeyTemplateProto reads a KeyTemplate from tink/testdata/templates.
func KeyTemplateProto(dir string, name string) (*tinkpb.KeyTemplate, error) {
	root, err := tinkRootPath()
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(root + "/testdata/templates/" + dir + "/" + name)
	if err != nil {
		return nil, err
	}
	template := &tinkpb.KeyTemplate{}
	err = proto.UnmarshalText(string(data), template)
	if err != nil {
		return nil, err
	}
	return template, nil
}
