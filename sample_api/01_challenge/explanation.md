/*
Golangのポイント
1. スコープについて
　Golangにはクラスがありません。
　Javaだったら、privateやpublicをクラス名やメソッド名の先頭につけることで
　アクセス範囲を決められたと思います。

　Golangでは、クラスで関数にアクセスする範囲を決められないので、
　関数名のイニシャルを大文字にするか小文字にするかで
　アクセスできる範囲を変更します。
　イニシャルを小文字にすると、他のディレクトリから
　関数に直接アクセスすることができなくなります。

　例えば、handler/hello.goでは、func Hello(){}という関数を作成し、
　router/router.goでHello関数を呼び出していますが、
　もし func Hello(){} を func hello(){}とすると、
　router/router.goから直接hello関数を呼び出すことができなくなります。

2．エラーハンドリングについて
　Golangには例外エラーをキャッチする機能がありません。
　try-exceptionで自動的に例外エラーを補足できないため、
　エラーが発生する箇所はif文を利用してエラーを補足します。

　例えば、文字列をint型に直す際に、文字列が数値ではない場合の
　エラーを補足したい場合は、以下のようにします。

	intSample := "1"

	value1, err := strconv.Atoi(intSample) // intSampleを数値に変換
	if err != nil { // エラーを補足（intSampleは数値のため、エラーなし）
		// 何らかのエラー処理
	}

	strSample := "Sample"

	value2, err := strconv.Atoi(strSample) // strSampleを数値に変換
	if err != nil { // エラーを補足 （strSampleは数値のため、エラー発生）
		// 何らかのエラー処理
	}
*/