package main //20 variant

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Triangle struct {
	Side1 *int `json:"side1"`
	Side2 *int `json:"side2"`
	Side3 *int `json:"side3"`
}

type Result struct {
	Result string `json:"result"`
}

func triangleHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Данный метод не поддерживается"))
		return
	}
	var triangle Triangle
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&triangle)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if strings.HasPrefix(err.Error(), "json: cannot unmarshal number") {
			w.Write([]byte("Число вышло за пределы разрядной сетки или не является целым числом!"))
		} else if strings.HasPrefix(err.Error(), "json: cannot unmarshal string") {
			w.Write([]byte("Введите число!"))
		} else {
			w.Write([]byte("Ошибка запроса! " + err.Error()))
		}
		return
	}
	if triangle.Side1 == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("1 сторона не указана!"))
		return
	}
	if triangle.Side2 == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("2 cторона не указана!"))
		return
	}
	if triangle.Side3 == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("3 cторона не указана!"))
		return
	}
	if int(*triangle.Side1) <= 0 || int(*triangle.Side2) <= 0 || int(*triangle.Side3) <= 0 {
		w.WriteHeader(400)
		w.Write([]byte("Все стороны должны быть положительны!"))
		return
	}

	var res Result
	if *triangle.Side1+*triangle.Side2 > *triangle.Side3 && *triangle.Side1+*triangle.Side3 > *triangle.Side2 && *triangle.Side3+*triangle.Side2 > *triangle.Side1 {
		res.Result = "Треугольник существует"
	} else {
		res.Result = "Треугольник не существует"
	}
	w.Header().Set("Content-Type", "application/json") //в формате json
	w.WriteHeader(http.StatusOK)
	respByte, _ := json.Marshal(res)
	w.Write(respByte)

}

func main() {
	http.HandleFunc("/triangle", triangleHandle)

	fmt.Println("Запуск сервера:")
	err := http.ListenAndServe("127.0.0.1:8081", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера!: ", err)
	}
}
