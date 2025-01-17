package store

import (
	"context"
	"os"
	"samsamoohooh-api/internal/application/domain"
	"samsamoohooh-api/internal/application/presenter"
	"samsamoohooh-api/internal/infra/validator"
	"samsamoohooh-api/pkg/prettier"
	"samsamoohooh-api/pkg/storetest"
	"samsamoohooh-api/pkg/testutil"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := storetest.SetUp(storetest.DefaultCtx)
	if err != nil {
		panic(err)
	}

	exitCode := m.Run()

	err = storetest.Shutdwon(storetest.DefaultCtx)
	if err != nil {
		panic(err)
	}

	os.Exit(exitCode)
}

func TestCreateUser(t *testing.T) {
	db := storetest.GetMysql(t)

	type args struct {
		ctx    context.Context
		params *presenter.CreateUserParams
	}

	tests := []struct {
		name string
		args
		expect    *domain.User
		expectErr bool
	}{
		{
			name: "[성공] 의도적으로 사용자를 생성한 경우",
			args: args{
				ctx: context.Background(),
				params: &presenter.CreateUserParams{
					Nickname:   "홍길동",
					Resolution: lo.ToPtr("저는 독서를 정말로 좋아합니다!"),
					Provider:   domain.Provider("GOOGLE"),
				},
			},
			expect: &domain.User{
				ID:         1,
				Nickname:   "홍길동",
				Resolution: lo.ToPtr("저는 독서를 정말로 좋아합니다!"),
				Provider:   "GOOGLE",
			},
		},
		{
			name: "[실패] not null 제약조건 위배",
			args: args{
				ctx: context.Background(),
				params: &presenter.CreateUserParams{
					Nickname: "",
					Provider: "",
				},
			},
			expect:    nil,
			expectErr: true,
		},

		{
			name: "[실패] 필수 값을 넣지 않은 경우",
			args: args{
				ctx: context.Background(),
				params: &presenter.CreateUserParams{
					Nickname:   "",
					Resolution: lo.ToPtr("안녕!"),
					Provider:   "",
				},
			},
			expect:    nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userStore := NewUserStore(db, validator.NewValidator())

			createdUser, err := userStore.CreateUser(tt.ctx, tt.params)
			ignoredCreatedUser := testutil.IgnoreFields(createdUser, "CreatedAt", "UpdatedAt")

			assert.Equalf(t, tt.expectErr, err != nil, "\nerr: %v", err)
			assert.Equalf(t, tt.expect, ignoredCreatedUser, "\ntt.expected: %+v\nres: %+v", prettier.Pretty(tt.expect), prettier.Pretty(ignoredCreatedUser))
		})
	}
}

func TestFindUser(t *testing.T) {
	db := storetest.GetMysql(t)

	// [사전 준비] 조회에 사용할 사용자들 생성
	var user = &domain.User{
		Nickname: "July",
		Provider: "APPLE",
	}

	err := db.
		Create(user).
		Error
	assert.NoError(t, err)

	// 테스트 케이스들
	type args struct {
		ctx    context.Context
		params *presenter.FoundUserParams
	}

	tests := []struct {
		name      string
		expectErr bool
		expect    *domain.User
		args
	}{
		{
			name: "[성공] id가 1인 사용자를 조회하는 경우",
			args: args{
				ctx: context.Background(),
				params: &presenter.FoundUserParams{
					UserID: 1,
				},
			},
			expect: user,
		},
		{
			name: "[실패] id가 0인 사용자를 조회하는 경우",
			args: args{
				ctx: context.Background(),
				params: &presenter.FoundUserParams{
					UserID: 0,
				},
			},
			expect:    nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userStore := NewUserStore(db, validator.NewValidator())

			res, err := userStore.FindUser(tt.ctx, tt.params)

			assert.Equalf(t, tt.expectErr, err != nil, "\ntt.expectErr = %+v\nerr = %+v", tt.expectErr, err)
			assert.Equalf(t, tt.expect, res, "\ntt.expected: %+v\nres: %+v", prettier.Pretty(tt.expect), prettier.Pretty(res))

		})
	}
}
