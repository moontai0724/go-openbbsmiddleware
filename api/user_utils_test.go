package api

import (
	"fmt"
	"testing"

	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
)

func Test_deserializeEmailToken(t *testing.T) {
	setupTest()
	defer teardownTest()

	jwt0, _ := pttbbsapi.CreateEmailToken("SYSOP", "", "test@ptt.test", pttbbsapi.CONTEXT_CHANGE_EMAIL)

	content0 := fmt.Sprintf("SYSOP, http://localhost:3457/user/SYSOP/changeemail?%v=%v", types.EMAIL_TOKEN_NAME, jwt0)

	type args struct {
		userID          bbs.UUserID
		token           string
		urlTemplate     string
		contentTemplate string
	}
	tests := []struct {
		name            string
		args            args
		expectedContent string
	}{
		// TODO: Add test cases.
		{
			args:            args{userID: "SYSOP", token: jwt0, urlTemplate: CHANGE_EMAIL_R, contentTemplate: types.EMAILTOKEN_TEMPLATE_CONTENT},
			expectedContent: content0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotContent := deserializeEmailToken(tt.args.userID, tt.args.token, tt.args.urlTemplate, tt.args.contentTemplate); gotContent != tt.expectedContent {
				t.Errorf("deserializeEmailToken() = %v, want %v", gotContent, tt.expectedContent)
			}
		})
	}
}
