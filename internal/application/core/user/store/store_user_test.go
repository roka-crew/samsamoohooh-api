package store

import (
	"context"
	"os"
	"samsamoohooh-api/internal/application/domain"
	"samsamoohooh-api/internal/application/presenter"
	"samsamoohooh-api/internal/infra/storetest"
	"samsamoohooh-api/pkg/prettier"
	"samsamoohooh-api/pkg/testutil"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	storetest.SetUp()

	exitCode := m.Run()

	storetest.Shutdwon()

	os.Exit(exitCode)
}

func TestCreateUser(t *testing.T) {
	db := storetest.GetMysql(t)

	type args struct {
		ctx    context.Context
		params *presenter.CreateUserParams
	}

	tests := []struct {
		name      string
		expectErr bool
		expected  *domain.User
		args
	}{
		{
			name:      "[성공] 의도적으로 사용자를 생성한 경우",
			expectErr: false,
			expected: &domain.User{
				ID:         1,
				Nickname:   "홍길동",
				Resolution: lo.ToPtr("저는 독서를 정말로 좋아합니다!"),
				Provider:   "GOOGLE",
			},
			args: args{
				ctx: context.Background(),
				params: &presenter.CreateUserParams{
					Nickname:   "홍길동",
					Resolution: lo.ToPtr("저는 독서를 정말로 좋아합니다!"),
					Provider:   domain.Provider("GOOGLE"),
				},
			},
		},

		{
			name:      "[실패] 필수 값을 넣지 않은 경우",
			expectErr: true,
			expected:  nil,
			args: args{
				ctx: context.Background(),
				params: &presenter.CreateUserParams{
					Nickname:   "",
					Resolution: lo.ToPtr("안녕!"),
					Provider:   "",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userStore := NewUserStore(db, storetest.GetValidator())

			createdUser, err := userStore.CreateUser(tt.ctx, tt.params)
			ignoredCreatedUser := testutil.IgnoreFields(createdUser, "CreatedAt", "UpdatedAt")

			assert.Equalf(t, tt.expectErr, err != nil, "\nerr: %v", err)
			assert.Equalf(t, tt.expected, ignoredCreatedUser, "\ntt.expected: %+v\nres: %+v", prettier.Pretty(tt.expected), prettier.Pretty(ignoredCreatedUser))
		})
	}
}
