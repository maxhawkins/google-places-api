package places

import "testing"

var dummyService = &Service{}

func TestNearbyCallValidate(t *testing.T) {
	for _, test := range []struct {
		Name string
		Call NearbyCall
		Want error
	}{
		{
			Name: "empty call",
			Call: NearbyCall{},
			Want: errInvalidByProminence,
		},
		{
			Name: "valid with default ranking",
			Call: NearbyCall{
				Radius: 5,
			},
			Want: nil,
		},
		{
			Name: "valid prominence call",
			Call: NearbyCall{
				Radius: 5,
				RankBy: RankByProminence,
			},
			Want: nil,
		},
		{
			Name: "invalid prominence call",
			Call: NearbyCall{
				RankBy: RankByProminence,
			},
			Want: errInvalidByProminence,
		},
		{
			Name: "invalid distance call",
			Call: NearbyCall{
				RankBy: RankByDistance,
			},
			Want: errInvalidByDistance,
		},
		{
			Name: "distance call with types",
			Call: NearbyCall{
				RankBy: RankByDistance,
				Types:  []FeatureType{Bar},
			},
			Want: nil,
		},
		{
			Name: "distance call with name",
			Call: NearbyCall{
				RankBy: RankByDistance,
				Name:   "name",
			},
			Want: nil,
		},
		{
			Name: "distance call with keyword",
			Call: NearbyCall{
				RankBy:  RankByDistance,
				Keyword: "keyword",
			},
			Want: nil,
		},
	} {
		got := test.Call.validate()
		if got != test.Want {
			t.Errorf("NearbyCall{%v}.query() = %#v, want %#v",
				test.Name, got, test.Want)
		}
	}
}
