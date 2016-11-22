//  Copyright (c) 2014 Couchbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package query

import (
	"encoding/json"

	"github.com/blevesearch/bleve/index"
	"github.com/blevesearch/bleve/mapping"
	"github.com/blevesearch/bleve/search"
	"github.com/blevesearch/bleve/search/searcher"
)

type MatchNoneQuery struct {
	BoostVal *Boost `json:"boost,omitempty"`
}

// NewMatchNoneQuery creates a Query which will not
// match any documents in the index.
func NewMatchNoneQuery() *MatchNoneQuery {
	return &MatchNoneQuery{}
}

func (q *MatchNoneQuery) SetBoost(b float64) {
	boost := Boost(b)
	q.BoostVal = &boost
}

func (q *MatchNoneQuery) Boost() float64{
	if q.BoostVal != nil {
		return q.BoostVal.Value()
	}
	return 0
}

func (q *MatchNoneQuery) Searcher(i index.IndexReader, m mapping.IndexMapping, explain bool) (search.Searcher, error) {
	return searcher.NewMatchNoneSearcher(i)
}

func (q *MatchNoneQuery) MarshalJSON() ([]byte, error) {
	tmp := map[string]interface{}{
		"boost":      q.BoostVal,
		"match_none": map[string]interface{}{},
	}
	return json.Marshal(tmp)
}
