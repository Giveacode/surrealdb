// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package item

import (
	"github.com/abcum/surreal/sql"
	"github.com/abcum/surreal/util/data"
)

func (this *Doc) Blaze(ast *sql.SelectStatement) (res interface{}) {

	doc := data.New()

	for _, v := range ast.Expr {
		if _, ok := v.Expr.(*sql.All); ok {
			doc = this.current
			break
		}
	}

	for _, v := range ast.Expr {
		switch e := v.Expr.(type) {
		default:
			doc.Set(e, v.Alias)
		case bool, int64, float64, string:
			doc.Set(e, v.Alias)
		case []interface{}, map[string]interface{}:
			doc.Set(e, v.Alias)
		case *sql.Null:
			doc.Set(nil, v.Alias)
		case *sql.Ident:
			doc.Set(this.current.Get(e.ID).Data(), v.Alias)
		case *sql.All:
			break
		}
	}

	return doc.Data()

}
