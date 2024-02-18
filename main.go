package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SMAは単純移動平均線を表す構造体
type SMA struct {
	period int       // 移動平均の期間
	values []float64 // 移動平均の値の配列
}

// NewSMAは新しいSMAを作成する関数
func NewSMA(period int) *SMA {
	return &SMA{
		period: period,
		values: make([]float64, 0, period),
	}
}

// Addは新しい値を追加するメソッド
func (s *SMA) Add(value float64) {
	// 配列が満杯の場合は先頭の要素を削除する
	if len(s.values) == s.period {
		s.values = s.values[1:]
	}
	// 配列の末尾に値を追加する
	s.values = append(s.values, value)
}

// Getは現在の移動平均の値を取得するメソッド
func (s *SMA) Get() float64 {
	// 配列が空の場合は0を返す
	if len(s.values) == 0 {
		return 0
	}
	// 配列の要素の合計を求める
	sum := 0.0
	for _, v := range s.values {
		sum += v
	}
	// 合計を要素数で割って平均を求める
	return sum / float64(len(s.values))
}

// Strategyはゴールデンクロス・デッドクロス戦略を表すインタフェース
type Strategy interface {
	// Updateは新しい価格を受け取って戦略を更新するメソッド
	Update(price float64)
	// Signalは現在の売買シグナルを返すメソッド
	Signal() string
}

// SMAStrategyはSMAを使った戦略を表す構造体
type SMAStrategy struct {
	shortSMA *SMA   // 短期移動平均線
	longSMA  *SMA   // 長期移動平均線
	signal   string // 売買シグナル
}

// NewSMAStrategyは新しいSMAStrategyを作成する関数
func NewSMAStrategy(shortPeriod, longPeriod int) *SMAStrategy {
	return &SMAStrategy{
		shortSMA: NewSMA(shortPeriod),
		longSMA:  NewSMA(longPeriod),
		signal:   "wait", // 初期値はwait
	}
}

// Updateは新しい価格を受け取って戦略を更新するメソッド
func (s *SMAStrategy) Update(price float64) {
	// SMAに値を追加する
	s.shortSMA.Add(price)
	s.longSMA.Add(price)
	// SMAの値を取得する
	shortMA := s.shortSMA.Get()
	longMA := s.longSMA.Get()
	// クロスの状態を判定する
	if shortMA > longMA {
		// 短期線が長期線より上ならゴールデンクロス
		s.signal = "buy"
	} else if shortMA < longMA {
		// 短期線が長期線より下ならデッドクロス
		s.signal = "sell"
	} else {
		// 短期線と長期線が重なっている場合は待機
		s.signal = "wait"
	}
}

// Signalは現在の売買シグナルを返すメソッド
func (s *SMAStrategy) Signal() string {
	return s.signal
}

func main() {
	// 乱数のシードを設定する
	rand.Seed(time.Now().UnixNano())
	// SMAの期間を設定する
	shortPeriod := 5
	longPeriod := 20
	// SMA戦略を作成する
	strategy := NewSMAStrategy(shortPeriod, longPeriod)
	// 50回の価格変動をシミュレーションする
	price := 100.0 // 初期価格
	for i := 0; i < 50; i++ {
		// 価格を前回から±10%の範囲で変動させる
		price = price * (1 + rand.Float64()*0.2 - 0.1)
		// 戦略を更新する
		strategy.Update(price)
		// 価格とシグナルを表示する
		fmt.Printf("price: %.2f, signal: %s\n", price, strategy.Signal())
	}
}
