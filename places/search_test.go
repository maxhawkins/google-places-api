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
				Type:   Bar,
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

func TestTextSearchValidate(t *testing.T) {
	for _, test := range []struct {
		Name string
		Call TextSearchCall
		Want error
	}{
		{
			Name: "Missing search criteria",
			Call: TextSearchCall{},
			Want: errEmptyQuery,
		},
		{
			Name: "With search criteria",
			Call: TextSearchCall{
				queryStr: "コナミコマンド",
			},
			Want: nil,
		},
		{
			Name: "Missing radius",
			Call: TextSearchCall{
				queryStr: "foo",
				lat:      0.0,
				lng:      0.1,
			},
			Want: errMissingRadius,
		},
		{
			Name: "Incorrect radius",
			Call: TextSearchCall{
				queryStr: "foo",
				lat:      27.988056,
				lng:      86.925278,
				Radius:   (maximumRadius + 1),
			},
			Want: errRadiusIsTooGreat,
		},
		{
			Name: "With correct location and radius",
			Call: TextSearchCall{
				queryStr: "foo",
				lat:      27.988056,
				lng:      86.925278,
				Radius:   maximumRadius,
			},
			Want: nil,
		},
	} {
		got := test.Call.validate()
		if got != test.Want {
			t.Errorf("TextSearchCall{%v}.query() = %#v, want %#v",
				test.Name, got, test.Want)
		}
	}
}
