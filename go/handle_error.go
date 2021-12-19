package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// 関数は関数名(引数 型) (戻り値の型) のように書く
func isOdd(num int) (bool, error) {
	if num == 0 {
		return false, errors.New("Can't specify zero")
	}
	return num%2 == 1, nil
}

func main() {
	fmt.Println("Please enter the number")

	var value string
	fmt.Scan(&value) // 定義した変数のポインタを渡すことで、後続処理では変数に値が設定された状態になる

	// 入力された文字列を数値に変換する
	num, err := strconv.Atoi(value)
	if err != nil {
		// エラーを出力してここで終了
		fmt.Println(err.Error())
		os.Exit(1)
	}

	result, err := isOdd(num)
	// 呼び出し側は err が nil かどうかを判定する必要がある
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(fmt.Sprintf("IsOdd? %t", result))
}
