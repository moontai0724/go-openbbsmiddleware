package schema

import (
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/mock_http"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
)

func TestUpdateArticleSummaries(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer Article_c.Drop()

	ret := mock_http.LoadGeneralArticles(nil)

	updateNanoTS := types.NowNanoTS()

	articleSummaries0 := make([]*ArticleSummary, len(ret.Articles))
	for idx, each_b := range ret.Articles {
		articleSummaries0[idx] = NewArticleSummary(each_b, updateNanoTS)
	}

	query0 := &ArticleQuery{BBoardID: bbs.BBoardID("1_test1"), ArticleID: bbs.ArticleID("19bWBI4Zokcool")}
	articleSummary0 := &ArticleSummary{
		BBoardID:   bbs.BBoardID("1_test1"),
		ArticleID:  bbs.ArticleID("19bWBI4Zokcool"),
		IsDeleted:  false,
		CreateTime: types.NanoTS(1234567890000000000),
		MTime:      types.NanoTS(1234567889000000000),

		Recommend:    8,
		Owner:        bbs.UUserID("okcool"),
		Title:        "然後呢？～",
		Money:        3,
		Class:        "問題",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS,
	}

	articleSummary1 := &ArticleSummary{
		BBoardID:   bbs.BBoardID("1_test1"),
		ArticleID:  bbs.ArticleID("1VrooM21teemo"),
		IsDeleted:  false,
		CreateTime: types.NanoTS(1607937174000000000),
		MTime:      types.NanoTS(1607937100000000000),

		Recommend:    3,
		Owner:        bbs.UUserID("teemo"),
		Title:        "再來呢？～",
		Money:        12,
		Class:        "問題",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS,
	}

	articleSummary2 := &ArticleSummary{
		BBoardID:   bbs.BBoardID("1_test1"),
		ArticleID:  bbs.ArticleID("1VrooM21teem2"),
		IsDeleted:  false,
		CreateTime: types.NanoTS(1607937174000000000),
		MTime:      types.NanoTS(1607937100000000000),

		Recommend:    4,
		Owner:        bbs.UUserID("teem2"),
		Title:        "再來呢2？～",
		Money:        15,
		Class:        "問題",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS,
	}

	updateNanoTS1 := types.NowNanoTS()

	articleSummary3 := &ArticleSummary{
		BBoardID:   bbs.BBoardID("1_test1"),
		ArticleID:  bbs.ArticleID("1VrooM21teemo"),
		IsDeleted:  false,
		CreateTime: types.NanoTS(1607937174000000000),
		MTime:      types.NanoTS(1607937100000000000),

		Recommend:    12,
		Owner:        bbs.UUserID("teemo"),
		Title:        "再來呢3？～",
		Money:        20,
		Class:        "問題",
		Filemode:     0,
		UpdateNanoTS: updateNanoTS1,
	}

	articleSummaries1 := []*ArticleSummary{articleSummary1, articleSummary2}

	query1 := &ArticleQuery{BBoardID: bbs.BBoardID("1_test1"), ArticleID: bbs.ArticleID("1VrooM21teemo")}

	query2 := &ArticleQuery{BBoardID: bbs.BBoardID("1_test1"), ArticleID: bbs.ArticleID("1VrooM21teem2")}

	articleSummaries2 := []*ArticleSummary{articleSummary2, articleSummary3}

	type args struct {
		a            []*ArticleSummary
		updateNanoTS types.NanoTS
	}
	tests := []struct {
		name     string
		args     args
		query    *ArticleQuery
		expected *ArticleSummary
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{a: articleSummaries0, updateNanoTS: updateNanoTS},
			query:    query0,
			expected: articleSummary0,
		},
		{
			args:     args{a: articleSummaries0, updateNanoTS: updateNanoTS},
			query:    query1,
			expected: articleSummary1,
		},
		{
			args:     args{a: articleSummaries1, updateNanoTS: updateNanoTS},
			query:    query2,
			expected: articleSummary2,
		},
		{
			args:     args{a: articleSummaries2, updateNanoTS: updateNanoTS1},
			query:    query1,
			expected: articleSummary3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateArticleSummaries(tt.args.a, tt.args.updateNanoTS); (err != nil) != tt.wantErr {
				t.Errorf("UpdateArticleSummaries() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, _ := GetArticleSummary(tt.query.BBoardID, tt.query.ArticleID)
			testutil.TDeepEqual(t, "got", got, tt.expected)
		})
	}
}
