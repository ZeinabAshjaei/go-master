package api

import (
	"bytes"
	"fmt"
	mockdb "github.com/ZeinabAshjaei/go-master/db/mock"
	db "github.com/ZeinabAshjaei/go-master/db/sqlc"
	"github.com/ZeinabAshjaei/go-master/utils"
	"github.com/goccy/go-json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAccountAPI(t *testing.T) {
	account := getRandomAccount()

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockStore := mockdb.NewMockStore(controller)

	mockStore.EXPECT().
		GetAccount(gomock.Any(), account.ID).
		Times(1).
		Return(account, nil)

	url := fmt.Sprintf("/accounts/%d", account.ID)
	server := NewServer(mockStore)
	recorder := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
	checkResponseBodyMatch(t, recorder.Body, account)
}

func getRandomAccount() db.Account {
	return db.Account{
		ID:       utils.RandomInt(1, 1000),
		Owner:    utils.RandomOwner(),
		Currency: utils.RandomCurrency(),
		Balance:  utils.RandomInt(100, 10000),
	}
}

func checkResponseBodyMatch(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var getAccount db.Account
	err = json.Unmarshal(data, &getAccount)
	require.NoError(t, err)
	require.Equal(t, getAccount, account)
}
