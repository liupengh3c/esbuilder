package esbuilder

// For more details, see
// https://www.elastic.co/guide/en/elasticsearch/reference/7.10/query-dsl-match-query.html

type matchQuery struct {
	name                string
	text                interface{}
	operator            string // or / and
	analyzer            string
	boost               *float64
	fuzziness           string
	prefixLength        *int
	maxExpansions       *int
	minimumShouldMatch  string
	fuzzyRewrite        string
	lenient             *bool
	fuzzyTranspositions *bool
	zeroTermsQuery      string
	cutoffFrequency     *float64
	// queryName           string
}

// NewmatchQuery creates and initializes a new matchQuery.
func NewMatchQuery(name string, text interface{}) *matchQuery {
	return &matchQuery{name: name, text: text}
}

// Operator sets the operator to use when using a boolean query.
// Can be "AND" or "OR" (default).
func (q *matchQuery) Operator(operator string) *matchQuery {
	q.operator = operator
	return q
}
func (q *matchQuery) Analyzer(analyzer string) *matchQuery {
	q.analyzer = analyzer
	return q
}

// Fuzziness sets the fuzziness when evaluated to a fuzzy query type.
// Defaults to "AUTO".
func (q *matchQuery) Fuzziness(fuzziness string) *matchQuery {
	q.fuzziness = fuzziness
	return q
}

// PrefixLength sets the length of a length of common (non-fuzzy)
// prefix for fuzzy match queries. It must be non-negative.
func (q *matchQuery) PrefixLength(prefixLength int) *matchQuery {
	q.prefixLength = &prefixLength
	return q
}

// MaxExpansions is used with fuzzy or prefix type queries. It specifies
// the number of term expansions to use. It defaults to unbounded so that
// its recommended to set it to a reasonable value for faster execution.
func (q *matchQuery) MaxExpansions(maxExpansions int) *matchQuery {
	q.maxExpansions = &maxExpansions
	return q
}

// CutoffFrequency can be a value in [0..1] (or an absolute number >=1).
// It represents the maximum treshold of a terms document frequency to be
// considered a low frequency term.
func (q *matchQuery) CutoffFrequency(cutoff float64) *matchQuery {
	q.cutoffFrequency = &cutoff
	return q
}

// MinimumShouldMatch sets the optional minimumShouldMatch value to
// apply to the query.
func (q *matchQuery) MinimumShouldMatch(minimumShouldMatch string) *matchQuery {
	q.minimumShouldMatch = minimumShouldMatch
	return q
}

// FuzzyRewrite sets the fuzzy_rewrite parameter controlling how the
// fuzzy query will get rewritten.
func (q *matchQuery) FuzzyRewrite(fuzzyRewrite string) *matchQuery {
	q.fuzzyRewrite = fuzzyRewrite
	return q
}

// FuzzyTranspositions sets whether transpositions are supported in
// fuzzy queries.
//
// The default metric used by fuzzy queries to determine a match is
// the Damerau-Levenshtein distance formula which supports transpositions.
// Setting transposition to false will
// * switch to classic Levenshtein distance.
// * If not set, Damerau-Levenshtein distance metric will be used.
func (q *matchQuery) FuzzyTranspositions(fuzzyTranspositions bool) *matchQuery {
	q.fuzzyTranspositions = &fuzzyTranspositions
	return q
}

// Lenient specifies whether format based failures will be ignored.
func (q *matchQuery) Lenient(lenient bool) *matchQuery {
	q.lenient = &lenient
	return q
}

// ZeroTermsQuery can be "all" or "none".
func (q *matchQuery) ZeroTermsQuery(zeroTermsQuery string) *matchQuery {
	q.zeroTermsQuery = zeroTermsQuery
	return q
}

// Boost sets the boost to apply to this query.
func (q *matchQuery) Boost(boost float64) *matchQuery {
	q.boost = &boost
	return q
}

// Source returns JSON for the function score query.
func (q *matchQuery) Build() (interface{}, error) {
	// {"match":{"name":{"query":"value","type":"boolean/phrase"}}}
	source := make(map[string]interface{})

	match := make(map[string]interface{})
	source["match"] = match

	query := make(map[string]interface{})
	match[q.name] = query

	query["query"] = q.text

	if q.operator != "" {
		query["operator"] = q.operator
	}
	if q.analyzer != "" {
		query["analyzer"] = q.analyzer
	}
	if q.fuzziness != "" {
		query["fuzziness"] = q.fuzziness
	}
	if q.prefixLength != nil {
		query["prefix_length"] = *q.prefixLength
	}
	if q.maxExpansions != nil {
		query["max_expansions"] = *q.maxExpansions
	}
	if q.minimumShouldMatch != "" {
		query["minimum_should_match"] = q.minimumShouldMatch
	}
	if q.fuzzyRewrite != "" {
		query["fuzzy_rewrite"] = q.fuzzyRewrite
	}
	if q.lenient != nil {
		query["lenient"] = *q.lenient
	}
	if q.fuzzyTranspositions != nil {
		query["fuzzy_transpositions"] = *q.fuzzyTranspositions
	}
	if q.zeroTermsQuery != "" {
		query["zero_terms_query"] = q.zeroTermsQuery
	}
	if q.cutoffFrequency != nil {
		query["cutoff_frequency"] = *q.cutoffFrequency
	}
	if q.boost != nil {
		query["boost"] = *q.boost
	}

	return source, nil
}
