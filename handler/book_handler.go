package handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/yyh-gl/go-api-server-by-ddd/domain/model"
	"github.com/yyh-gl/go-api-server-by-ddd/usecase"
	"net/http"
)

func BookIndex(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	var books []*model.Book
	var err error
	books, err = usecase.IBookUsecase(usecase.BookUsecase{}).GetAll()
	if err != nil {
		// TODO: エラーレスポンスを返す
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(books); err != nil {
		// TODO: エラーレスポンスを返す
		panic(err)
	}
}
